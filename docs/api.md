# API docs

The following endpoints are available in the API:

**1. Get Vending Machine Endpoint**
   
This endpoint is used to get the details of a vending machine.

**HTTP Method:** GET

**URL:** `/api/vending-machine/{id}`

**Request URL Parameters:**

`id`: The ID of the vending machine to retrieve.

**Response Body:**

```json
{
  "id": 1,
  "Status": "idle",
  "Inventory": {
    "Cola": 9,
    "Coffee": 20
  },
  "Coins": 1
}
```

**2. Insert Coin Endpoint**

This endpoint is used to insert a coin into the vending machine.


**HTTP Method:** POST

**URL:** `/api/v1/vending-machine/insert-coin`

**Response Body:**

```json
{
  "machine_id": 1
}
```

**Response Body:**

```json
{
  "id": 1,
  "message": "Coin inserted successfully",
  "status": "product_selecting",
  "cola": 10,
  "coffee": 20
}
```

**3. Select Product Endpoint**

This endpoint is used to select a product from the vending machine.

**HTTP Method:** POST

**URL:** `/api/v1/vending-machine/select-product`

**Response Body:**

```json
{
  "id": 1,
  "product": "cola"
}
```

**Response Body:**

```json
{
  "id": 1,
  "message": "Enjoy your colaaa :))))))",
  "status": "idle"
}
```