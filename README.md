# Backend Go Learning Journey 🚀

This repository contains my journey of learning backend development with Go, Gin, PostgreSQL, and GORM.

## Week 1 Progress

### Day 1 - Go Basics

* Variables, data types, functions
* Structs and interfaces
* Error handling
* Packages and modules

### Day 2 - Concurrency

* Goroutines
* WaitGroups
* Channels
* Concurrent programming patterns

### Day 3 - Gin Framework & REST APIs

Built my first REST API using Gin.

#### Features

* Health check endpoint
* Product endpoints
* JSON responses
* Route handling

#### Endpoints

* `GET /health`
* `GET /products`
* `GET /products/:id`
* `POST /products`

### Day 4 - PostgreSQL & GORM Integration

Connected the API to a PostgreSQL database using GORM.

#### Technologies Used

* Go
* Gin Framework
* PostgreSQL
* GORM ORM

#### Features Implemented

* PostgreSQL database setup
* Database connection using GORM
* AutoMigrate for automatic table creation
* Product model persistence
* Create products through API
* Retrieve products from database

#### Database Models

##### Product

* ID
* Name
* Price
* Category

##### User

* ID
* Name
* Email

#### API Endpoints

##### GET /products

Returns all products stored in PostgreSQL.

##### POST /products

Creates a new product and saves it to PostgreSQL.

Example Request:

```json
{
  "name": "Laptop",
  "price": 999.99,
  "category": "electronics"
}
```

## Project Structure

```text
week1/
├── day1-basics/
├── day2-concurrency/
├── day3-gin/
└── day4-database/
```

## Setup

### Clone Repository

```bash
git clone <repository-url>
cd gayathri-backend-go
```

### Install Dependencies

```bash
go mod tidy
```

### PostgreSQL Configuration

Create a `.env` file:

```env
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=streamcart
DB_PORT=5432
```

### Run Application

```bash
go run .
```

### Day 5 - JWT Aunthenication & Register User

POST /register

Creates a new user account.

Example Request:
{
"name": "John",
"email": "[john@test.com](mailto:john@test.com)",
"password": "password123"
}

### Login User

POST /login

Returns a JWT token for authenticated access.

Example Request:
{
"email": "[john@test.com](mailto:john@test.com)",
"password": "password123"
}

### Protected Routes

The following routes require a JWT token:

* GET /products
* POST /products

Include the token in the request header:

Authorization: Bearer YOUR_TOKEN_HERE


## Learning Goals

- Learn Go fundamentals and syntax
- Understand goroutines, channels, and concurrency
- Build REST APIs using Gin
- Understand HTTP methods (GET, POST, PUT, DELETE)
- Work with PostgreSQL databases
- Use GORM for database operations (CRUD)
- Design and use Go structs as database models
- Understand authentication vs authorization
- Secure passwords using bcrypt hashing
- Implement JWT-based authentication
- Create and validate JWT tokens
- Protect API routes using Gin middleware
- Build backend services following REST principles
- Learn Git and GitHub workflow for version control

---

Learning backend development one day at a time. 🚀
