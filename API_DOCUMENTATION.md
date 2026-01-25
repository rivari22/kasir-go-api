# Kasir API Documentation

## Overview

This API provides category management endpoints for a point-of-sale (kasir) system. All endpoints return JSON responses and use standard HTTP status codes.

**Base URL**: `http://localhost:8080`

**Note**: This API uses in-memory storage. Data will be lost when the server restarts.

---

## Data Models

### Category
```json
{
  "id": 1,
  "name": "Category Name",
  "description": "Category Description"
}
```

| Field | Type | Description |
|-------|------|-------------|
| id | integer | Auto-generated unique identifier |
| name | string | Category name (required) |
| description | string | Category description (optional) |

### Generic Response
All API responses follow this structure:
```json
{
  "message": "string",
  "data": any
}
```

---

## Endpoints

### Categories

#### Get All Categories
```http
GET /categories
```

**Response** (200 OK):
```json
{
  "message": "success get categories",
  "data": [
    {
      "id": 1,
      "name": "Electronics",
      "description": "Electronic devices and accessories"
    },
    {
      "id": 2,
      "name": "Groceries",
      "description": "Food and household items"
    }
  ]
}
```

---

#### Create Category
```http
POST /categories
```

**Request Body**:
```json
{
  "name": "Category Name",
  "description": "Category Description"
}
```

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| name | string | Yes | Category name |
| description | string | No | Category description |

**Response** (200 OK):
```json
{
  "message": "success create new category",
  "data": "Category Name"
}
```

**Error Response** (400 Bad Request):
```json
{
  "message": "name is required"
}
```

---

#### Get Category by ID
```http
GET /categories/{id}
```

**Path Parameters**:
| Parameter | Type | Description |
|-----------|------|-------------|
| id | integer | Category ID |

**Response** (200 OK):
```json
{
  "message": "success get category by id",
  "data": {
    "id": 1,
    "name": "Electronics",
    "description": "Electronic devices and accessories"
  }
}
```

**Error Response** (404 Not Found):
```json
{
  "message": "category not found"
}
```

---

#### Update Category
```http
PUT /categories/{id}
```

**Path Parameters**:
| Parameter | Type | Description |
|-----------|------|-------------|
| id | integer | Category ID |

**Request Body**:
```json
{
  "name": "Updated Category Name",
  "description": "Updated Category Description"
}
```

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| name | string | Yes | Category name |
| description | string | No | Category description |

**Response** (200 OK):
```json
{
  "message": "success update category by id",
  "data": 1
}
```

**Error Responses**:

- 400 Bad Request - Name is required:
```json
{
  "message": "name is required"
}
```

- 404 Not Found - Category doesn't exist:
```json
{
  "message": "category not found"
}
```

---

#### Delete Category
```http
DELETE /categories/{id}
```

**Path Parameters**:
| Parameter | Type | Description |
|-----------|------|-------------|
| id | integer | Category ID |

**Response** (200 OK):
```json
{
  "message": "success delete category by id",
  "data": 0
}
```

**Error Response** (404 Not Found):
```json
{
  "message": "category not found"
}
```

---

## HTTP Status Codes

| Code | Description |
|------|-------------|
| 200 | Success |
| 400 | Bad Request - Invalid input or missing required fields |
| 404 | Not Found - Resource doesn't exist |
| 405 | Method Not Allowed - Invalid HTTP method for endpoint |

---

## CORS Configuration

The API is configured with the following CORS headers:

| Header | Value |
|--------|-------|
| Access-Control-Allow-Origin | `*` (all origins) |
| Access-Control-Allow-Methods | `GET, POST, PUT, DELETE, OPTIONS` |
| Access-Control-Allow-Headers | `Content-Type, Authorization` |

**Note**: The CORS configuration allows all origins for development. For production, this should be restricted to specific allowed origins.

---

## Technical Details

### Server Configuration
- **Port**: 8080
- **Storage**: In-memory (data is lost on server restart)
- **Concurrency**: Thread-safe operations using `sync.RWMutex`

### ID Generation
- Category IDs are auto-incremented based on the highest existing ID
- First category created will have ID 1

### Error Handling
All errors return a JSON response with a descriptive message:
```json
{
  "message": "Error description"
}
```

---

## Example Usage

### cURL Examples

**Create a category:**
```bash
curl -X POST http://localhost:8080/categories \
  -H "Content-Type: application/json" \
  -d '{"name": "Electronics", "description": "Electronic devices"}'
```

**Get all categories:**
```bash
curl http://localhost:8080/categories
```

**Get category by ID:**
```bash
curl http://localhost:8080/categories/1
```

**Update a category:**
```bash
curl -X PUT http://localhost:8080/categories/1 \
  -H "Content-Type: application/json" \
  -d '{"name": "Updated Electronics", "description": "Updated description"}'
```

**Delete a category:**
```bash
curl -X DELETE http://localhost:8080/categories/1
```
