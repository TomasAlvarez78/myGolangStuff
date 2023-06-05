

## Add Pauli
```sh
    curl http://localhost:8080/friends \
        --include \
        --header "Content-Type: application/json" \
        --request "POST" \
        --data '{"id": 8,"personName": "Pauwuli","age": 22,"career": "Software Engineering"}'
```


## Get friends
```sh
    curl http://localhost:8080/friends \
        --header "Content-Type: application/json" \
        --request "GET"
```
