package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/sarathsp06/pasteit/rpc/server"
)

// Create global To-Do counter
var addr string = "https://pasteit-cb6fdypqlq-an.a.run.app"

func readContent() []byte {
	fi, err := os.Stdin.Stat()
	if err != nil {
		return nil
	}
	if fi.Size() > 0 {
		r := bufio.NewReader(os.Stdin)
		data, _ := io.ReadAll(r)
		return data
	}
	return nil
}

func main() {
	// Define our interaction flags.
	action := flag.String("command", "list", "Command to interact with pasteit. create, get, list")
	title := flag.String("title", "Title", "Title for the paste, valid for create")
	id := flag.String("id", "", "The paste id")
	flag.Parse()

	// Create the generated Twirp Protobuf Client for our Todo service
	addr = "http://0.0.0.0:8080"
	client := server.NewPasterProtobufClient(addr, &http.Client{Timeout: time.Second})

	switch *action {
	case "create":
		content := readContent()
		if len(content) == 0 {
			log.Fatal("Content should be available ")
		}
		paste, err := client.CreatePaste(context.Background(), &server.CreatePasteRequest{
			Title:   *title,
			Content: string(content),
		})
		if err != nil {
			fmt.Println("could not paste:", err)
		} else {
			fmt.Printf("Created paste id:%s", paste.GetUuid())
		}
	case "get":
		if len(*id) == 0 {
			log.Fatal("id should be available ")
		}
		paste, err := client.GetPaste(context.Background(), &server.GetPasteRequest{
			Id: *id,
		})
		if err != nil {
			fmt.Println("could get paste:", err)
		} else {
			fmt.Printf("Got paste content:\n\n%s", paste.GetContent())
		}
	default:
		fmt.Println("invalid command, please try again.")
		flag.PrintDefaults()
	}
}
