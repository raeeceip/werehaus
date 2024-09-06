# Items

Manage inventory items in the warehouse.

## Get Items

Retrieve a list of items.

**URL**: `/api/items`

**Method**: `GET`

**Query Parameters**:
- `page` (optional): Page number for pagination
- `limit` (optional): Number of items per page
- `search` (optional): Search term to filter items

**Headers**: 
- `Authorization: Bearer [token]`

### Success Response

**Code**: `200 OK`

**Content example**

```json
[
    {
        "id": 1,
        "name": "Laptop",
        "description": "High-performance laptop",
        "quantity": 50
    },
    {
        "id": 2,
        "name": "Mouse",
        "description": "Wireless mouse",
        "quantity": 100
    }
]
```

## Create Item

Create a new item in the inventory.

**URL**: `/api/items`

**Method**: `POST`

**Headers**: 
- `Authorization: Bearer [token]`
- `Content-Type: application/json`

**Data constraints**

```json
{
    "name": "[item name]",
    "description": "[item description]",
    "quantity": "[integer]"
}
```

**Data example**

```json
{
    "name": "Keyboard",
    "description": "Mechanical keyboard",
    "quantity": 30
}
```

### Success Response

**Code**: `201 CREATED`

**Content example**

```json
{
    "id": 3,
    "name": "Keyboard",
    "description": "Mechanical keyboard",
    "quantity": 30
}
```

## Update Item

Update an existing item in the inventory.

**URL**: `/api/items/:id`

**Method**: `PUT`

**Headers**: 
- `Authorization: Bearer [token]`
- `Content-Type: application/json`

**Data constraints**

```json
{
    "name": "[item name]",
    "description": "[item description]",
    "quantity": "[integer]"
}
```

**Data example**

```json
{
    "name": "Keyboard",
    "description": "Ergonomic mechanical keyboard",
    "quantity": 25
}
```

### Success Response

**Code**: `200 OK`

**Content example**

```json
{
    "id": 3,
    "name": "Keyboard",
    "description": "Ergonomic mechanical keyboard",
    "quantity": 25
}
```

## Delete Item

Delete an item from the inventory.

**URL**: `/api/items/:id`

**Method**: `DELETE`

**Headers**: 
- `Authorization: Bearer [token]`

### Success Response

**Code**: `200 OK`

**Content example**

```json
{
    "message": "Item successfully deleted"
}
```