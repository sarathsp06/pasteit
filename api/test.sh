function Create {
curl --request "POST" \
     --location "http://localhost:8081/twirp/pasteit.Paster/CreatePaste" \
     --header "Content-Type:application/json" \
     --data '{"headers": {"version" : "1.0.0"} ,"content":"This is sparta","ttl":3600}' \
     --verbose
}


function Get() {
data="{\"id\": \"$1\"}"
curl --request "POST" \
     --location "http://localhost:8081/twirp/pasteit.Paster/GetPaste" \
     --header "Content-Type:application/json" \
     --data "${data}" \
     --verbose    
}

#Create
Get "8e968bf5-741b-4ba5-812c-206064d022a3"