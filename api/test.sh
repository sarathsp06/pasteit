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
Get "a6e0667c-9704-4f64-a2db-c18aab4bbd73"