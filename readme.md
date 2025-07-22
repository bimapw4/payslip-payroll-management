# Payslips & Payroll Management API

> A project **Golang Fiber** REST API with **PostgreSQL**, built with modular architecture including Authentication, Payroll processing, Employee payslips, Attendance tracking, Overtime, and other.


### Features
> 
| Module             | Description                                        |
| ------------------ | -------------------------------------------------- |
| **Payroll**        | Create & run payroll, lock the payroll period      |
| **Payslip**        | Generate employee payslip and summary for admin    |
| **Attendance**     | Track check-in/check-out with validation           |
| **Overtime**       | Manage overtime (max 3 hours per day)              |
| **Reimbursement**  | Input & update reimbursement with user-level guard |
| **Audit Log**      | Log all HTTP requests & responses (middleware)     |
| **Authentication** | JWT-based login with bcrypt password hashing       |

### Project Structure
<pre>
├── bootstrap/                # App initialization (DB, DI, Migrations)
│   ├── db.go
│   ├── migrate.go
│   └── providers.go
│
├── internal/                 
│   ├── business/             # Business logic / usecases
│   ├── common/               # Common helpers (JWT, context, bcrypt, etc.)
│   ├── consts/                # Global constants
│   ├── entity/               # Request payloads / DTO (input layer)
│   ├── handlers/             # HTTP handlers (Fiber endpoints)
│   ├── middleware/           # Middleware (Audit log, Auth guard)
│   ├── presentations/        # DB models & API response structures
│   ├── provider/             # Dependency injection & service registry
│   ├── repositories/         # Data access layer (SQLX + PostgreSQL)
│   ├── response/             # API response wrapper (success / error)
│   ├── routes/                # HTTP route definitions (Fiber)
│   └── migrations/           # SQL migration scripts
│
├── pkg/                      
│   ├── databasex/            # Additional DB helper functions
│   └── meta/                 # Pagination, metadata utilities
│
├── .env                      # Environment variables
├── .env.example              # Sample environment file
├── docker-compose.yml        # Docker service setup
├── dockerfile                 # Dockerfile for app build
├── go.mod                     # Go modules
├── go.sum                     
├── main.go                   # Application entry point
└── readme.md                 
</pre>

### Database Schema
Main tables:
* users
* payroll
* attendance
* overtime
* reimbursement
* payslip_summary
* audit_log

### Authentication
* Login with username & password
* Passwords are hashed using bcrypt
* JWT access tokens are issued for API access

### Sample Endpoints
| Method | Endpoint               | Description                |
| ------ | ---------------------- | -------------------------- |
| POST   | `/login`               | Login & get JWT token      |
| POST   | `/attendance`          | Check-in / Check-out       |
| POST   | `/overtime`            | Submit overtime            |
| POST   | `/payroll`             | Create payroll             |
| PUT    | `/payroll/running/:id` | Run payroll process        |
| GET    | `/payroll/generate/payslips/:id` | Generate payslip for user  |
| GET    | `/payroll/summary/payslip/:id` | Get List payslip summary report for admin |

this is the postman collection


### Testing
✅ Unit Test
* Business layer unit tested
* Using gomock + testify

### Run Tests
Run all tests:
```
go test ./internal/... -v
```

### Technology Stack
* Golang (1.21+)

* Fiber (HTTP Framework)

* SQLX + PostgreSQL

* JWT v4

* Gomock + Testify (Testing)

* Docker / Docker Compose

### 📄 API Logging
All API requests and responses are logged to the audit log table using auditlog middleware.

### Env Example
```
APP_NAME = Payroll Payslip
PORT = 8083

DB_HOST = 
DB_USER = 
DB_PASSWORD = 
DB_NAME = 
DB_PORT = 

JWT_SECRET_KEY = 
JWT_LIFESPAN = 
```