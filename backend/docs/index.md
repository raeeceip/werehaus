# Warehouse Management API

Welcome to the Warehouse Management API documentation. This API allows you to manage items, locations, and issues in a warehouse system.

## Getting Started

To get started with the API, you'll need to:

1. [Authenticate](#authentication) with the API
2. Make requests to the various endpoints

## Base URL

All API requests should be made to:

```
http://localhost:3000/api
```

## Authentication

To use the API, you need to authenticate. See the [Authentication](api/auth.md) section for details.

## API Sections

- [Items](api/items.md): Manage inventory items
- [Locations](api/locations.md): Manage warehouse locations
- [Issues](api/issues.md): Handle item transfers and issues
- [Reports](api/reports.md): Generate various reports

## Error Handling

The API uses conventional HTTP response codes to indicate the success or failure of an API request. In general:

- 2xx range indicate success
- 4xx range indicate an error that failed given the information provided (e.g., a required parameter was omitted, etc.)
- 5xx range indicate an error with our servers