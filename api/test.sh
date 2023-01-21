BASE_URL=https://pasteit-cb6fdypqlq-an.a.run.app/
#local 
#BASE_URL=http://localhost:8081
function Create {
curl --request "POST" \
     --location "${BASE_URL}/twirp/pasteit.Paster/CreatePaste" \
     --header "Content-Type:application/json" \
     --data '{"headers": {"version" : "1.0.0"} ,"content":"This is sparta","ttl":3600}' 
}


function Get() {
data="{\"id\": \"$1\"}"
curl --request "POST" \
     --location "${BASE_URL}/twirp/pasteit.Paster/GetPaste" \
     --header "Content-Type:application/json" \
     --data "${data}" \
     --verbose    
}

#Create
Create 
Get 7b04551b-8053-4f01-b58a-568e621be962