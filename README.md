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

## Learning Goals

* Build backend APIs with Go
* Learn concurrency and goroutines
* Develop REST APIs using Gin
* Work with PostgreSQL databases
* Use GORM ORM effectively
* Build production-ready backend services

---

Learning backend development one day at a time. 🚀
