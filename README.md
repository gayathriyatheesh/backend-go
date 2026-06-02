# Go Backend Learning Journey

This repository contains my Go backend development learning journey, including Go basics, concurrency, goroutines, channels, and backend API development using the Gin framework.

---

# Week 1

## Day 1 — Go Basics

Topics covered:
- Hello World
- Variables
- Arithmetic operations
- Loops
- Functions
- Packages

Programs created:
- hello.go
- math.go
- loop.go
- function.go
- packages.go

---

## Day 2 — Goroutines & Concurrency

Topics covered:
- Goroutines
- WaitGroups
- Concurrent execution
- Channels
- Sender & Receiver pattern
- Synchronization

Programs created:
- goroutines.go
- channels.go

Key learning:
- Go executes goroutines concurrently
- Channels help goroutines communicate safely
- WaitGroups wait for all goroutines to finish

---

## Day 3 — Gin Backend API

Built a backend API using Go + Gin framework.

Topics covered:
- HTTP methods
- Routing
- Handlers
- JSON APIs
- REST API basics
- Models and handlers separation
- API testing using Postman

Project:
- StreamCart Backend API

Endpoints:
- GET /health
- GET /products
- GET /products/:id
- POST /products

Files:
- main.go
- models.go
- handlers.go

Tech used:
- Go
- Gin Framework
- Postman
- Git & GitHub

---

# Running the Project

## Install Gin

```bash
go get -u github.com/gin-gonic/gin
```

## Run API

```bash
go run .
```

Server runs on:

```bash
http://localhost:8080
```

---

# Goals

- Learn backend engineering with Go
- Build scalable APIs
- Learn concurrency deeply
- Prepare for real backend projects
- Build production-level systems
