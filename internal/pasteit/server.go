package pasteit

import (
	"context"
	"fmt"
	"time"

	store "github.com/sarathsp06/pasteit/internal/store"
	pb "github.com/sarathsp06/pasteit/rpc/server"
	uuid "github.com/satori/go.uuid"
	"github.com/twitchtv/twirp"
)

var validationError error

type server struct {
	maxContent int
	minContent int
	store      store.Store
}

// NewServer creates an instance of
func NewServer(minContent int, maxContent int, store store.Store) pb.Paster {
	return &server{
		maxContent: maxContent,
		minContent: minContent,
		store:      store,
	}
}

// CreatePaste creates a paste
func (s *server) CreatePaste(ctx context.Context, pr *pb.CreatePasteRequest) (*pb.CreatePasteResponse, error) {
	if len(pr.Content) > s.maxContent || len(pr.Content) < s.minContent {
		return nil, twirp.InvalidArgumentError("Content", "Invalid content size").
			WithMeta("min", fmt.Sprintf("%d", s.minContent)).
			WithMeta("max", fmt.Sprintf("%d", s.maxContent))
	}

	if pr.Ttl == 0 {
		pr.Ttl = 60 * 60
	}

	paste := store.Paste{
		ID:      uuid.NewV4(),
		Body:    pr.Content,
		Headers: pr.Headers,
		Expiry:  time.Now().Add(time.Second * time.Duration(pr.Ttl)),
	}

	paste, err := s.store.SavePaste(ctx, paste)
	if err := castError(err); err != nil {
		return nil, err
	}
	return &pb.CreatePasteResponse{
		Uuid:      paste.ID.String(),
		CreatedAt: uint64(paste.CreatedAt.Unix()),
	}, err
}

// GetPaste  gets paste given id
func (s *server) GetPaste(ctx context.Context, pr *pb.GetPasteRequest) (*pb.GetPasteResponse, error) {
	id, err := uuid.FromString(pr.Id)
	if err != nil {
		return nil, twirp.InvalidArgumentError("ID", "Invalid ID, not a uuid").WithMeta("id", pr.Id)
	}

	paste, err := s.store.GetPaste(ctx, id)
	if err := castError(err); err != nil {
		return nil, err
	}
	return &pb.GetPasteResponse{
		Title:   paste.Title,
		Content: paste.Body,
		Headers: paste.Headers,
	}, nil
}
