# StreamCart Backend API

A Go backend for an e-commerce + social feed platform.

## Building with Go + Gin framework

### Setup

```bash
go get -u github.com/gin-gonic/gin
go run .
```

### API Endpoints

- GET /health - Health check
- GET /products - List all products
- GET /products/:id - Get product by ID
- POST /products - Create new product