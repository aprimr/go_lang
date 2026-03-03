# Expense Tracker API

A REST API built with Go and PostgreSQL for tracking personal expenses.

## Tech Stack
- Go (net/http)
- PostgreSQL (Neon)
- pgx driver

## Features
- CRUD operations for expenses
- Pagination support
- Expense summary by category

## Endpoints
GET    /expenses?page=1&limit=10
POST   /expenses
GET    /expenses/:id
PUT    /expenses/:id
DELETE /expenses/:id
GET    /expenses/summary

## Setup
1. Clone the repo
2. Copy .env.example to .env and fill in values
3. Run: go run main.go