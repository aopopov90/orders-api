
## Building and executing
```bash
# initialize a project
go mod init github.com/aopopov90/orders-api

# run 
go run main.go

# build
go build -o api main.go

# run executable
./api 
```

## Adding dependencies

Add using the package manager.
```bash
go get -u github.com/go-chi/chi/v5
```

`go.mod` will have a new entry: `require github.com/go-chi/chi/v5 v5.0.12 // indirect `

Another approach is to add dependency under the import statement first and then run `go mod tidy`.

## Testing

Start redis:
```bash
docker run --name=redis-devel --publish=6379:6379 --hostname=redis --restart=on-failure --detach redis:latest
```
Create an order:
```bash
curl -X POST -d '{"customer_id":"'$(uuidgen)'","line_items":[{"item_id":"'$(uuidgen)'","quantity":5,"price":1999}]}' localhost:9000/orders
```

Checker orders using `redis-cli`:
```bash
# connect to redis container
docker exec -it 517096bd2474 sh

# start the cli
redis-cli

# check the id is in the orders set
 SMEMBERS orders
```