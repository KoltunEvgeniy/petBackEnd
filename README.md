Backend service written in Go for managing appointments, schedules, and availability between clients and masters.

---

##  Tech Stack
- Go
- Gin — HTTP framework
- PostgreSQL
- SQLX
- JWT
- Database migrations

---

## Project Structure

meawby/
├── cmd/ # Application entrypoints
│ └── api/ # HTTP server
├── internal/
│ ├── handler/ # HTTP handlers
│ ├── service/ # Business logic
│ ├── repository/ # Database access layer
│ ├── model/ # Domain models
│ └── middleware/ # Auth, logging
├── migrations/ # Database migrations
├── pkg/ # Shared utilities
├── go.mod
└── README.md


---

## Run Locally

### 1. Install dependencies
```bash
go mod download
2. Run database migrations
bash
go run cmd/migrate/main.go up
3. Start the server
bash
go run cmd/api/main.go
Server will start on:

arduino
http://localhost:8080
 Authentication
JWT-based authentication

Protected endpoints require access token

Token is passed via HTTP header:

makefile
Authorization: Bearer <token>
Main Features
User authentication and authorization

Client and master roles

Fixed time-slot scheduling

Availability management

Appointment creation and listing

PostgreSQL transactions

Repository pattern

 Testing
bash
go test ./...
🛠 Development Notes
Fixed slot duration

No overlapping appointments

Clear separation of layers

Business logic independent from transport layer