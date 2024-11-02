# vendingMachine
Create CRUD API for Vending Machine Case

SETUP: 

1. Initialize the Go Project: First, make sure to initialize a Go module.
```
go mod init vending-machine
```

2. Install Mux: We'll use Gorilla Mux for routing.
```
go get -u github.com/gorilla/mux
```


Explanation

Struct: 
  Item represents an item in the vending machine with Name and Price fields.

Endpoints:
  GET /items: Retrieves all items.
  GET /items/{name}: Retrieves a single item by its name.
  POST /items: Creates a new item after checking that the price denomination is valid.
  PUT /items/{name}: Updates an item’s details if it exists and if the new price is valid.
  DELETE /items/{name}: Deletes an item by name.

Validation: 
  isValidDenomination ensures that the price is either 2000 or 5000, as per the requirements.


Running the API

1. Save the code to a file, for example, main.go.
2. Run the API:
   ```
   go run main.go
   ```
3. Test the endpoints using curl or a tool like Postman.


Testing the API

For testing, you can use curl in the terminal or a tool like Postman.

Assume the API is running on http://localhost:8000.

1. Get All Items

Request:
```
curl -X GET http://localhost:8000/items
```

Response:
```json
[
    { "name": "Aqua", "price": 2000 },
    { "name": "Sosro", "price": 5000 },
    { "name": "Cola", "price": 7000 },
    { "name": "Milo", "price": 9000 },
    { "name": "Coffee", "price": 12000 }
]
```

2. Get a Single Item by Name

Request:
```
curl -X GET http://localhost:8000/items/Sosro
```

Response:
```json
{ "name": "Sosro", "price": 5000 }
```

If the item does not exist:
Response:
```json
"Item not found"
```

3. Create a New Item
When creating a new item, make sure the price is either 2000 or 5000, as only these denominations are accepted.

Request:
```
curl -X POST http://localhost:8000/items \
-H "Content-Type: application/json" \
-d '{"name": "Sprite", "price": 5000}'
```

Response:
```json
{ "name": "Sprite", "price": 5000 }
```

4. Update an Existing Item
Updating an item requires specifying the name in the URL and providing a new price with the correct denomination.

Request:
```
curl -X PUT http://localhost:8000/items/Cola \
-H "Content-Type: application/json" \
-d '{"name": "Cola", "price": 5000}'
```

Response:
```json
{ "name": "Cola", "price": 5000 }
```

5. Delete an Item

Request:
```
curl -X DELETE http://localhost:8000/items/Milo
```

Response:
```json
"Item deleted"
```

If the item does not exist:

Response:
```json
"Item not found"
```
