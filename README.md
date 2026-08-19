# ForgeFlow
![Go](https://img.shields.io/badge/Go-1.24-00ADD8?style=for-the-badge&logo=go)
![Gin](https://img.shields.io/badge/Gin-Web_Framework-00A86B?style=for-the-badge)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-16-336791?style=for-the-badge&logo=postgresql)
![Docker](https://img.shields.io/badge/Docker-Containerized-2496ED?style=for-the-badge&logo=docker)
![Swagger](https://img.shields.io/badge/Swagger-OpenAPI-85EA2D?style=for-the-badge&logo=swagger)
![JWT](https://img.shields.io/badge/JWT-Authentication-000000?style=for-the-badge&logo=jsonwebtokens)

Production-ready backend engineering project built with Go, Gin, PostgreSQL, JWT Authentication, Docker, Swagger, and modern cloud-native engineering practices.

## Tech Stack

- Go
- Gin
- PostgreSQL
- GORM
- JWT Authentication
- Docker
- Swagger (OpenAPI)

## Current Features

- RESTful APIs
- JWT Authentication
- Protected Routes
- CRUD Operations
- Server-side Pagination
- Centralized Validation
- Structured Logging
- Request ID Middleware
- Dependency Injection
- Repository Pattern
- Unit Testing
- Swagger Documentation

---

##  Project Progress

###  Completed

- [x] REST API Development
- [x] PostgreSQL Integration
- [x] GORM ORM
- [x] JWT Authentication
- [x] Protected Routes
- [x] CRUD Operations
- [x] Docker & Docker Compose
- [x] Swagger/OpenAPI Documentation
- [x] Environment Configuration
- [x] Health, Readiness & Version Endpoints
- [x] Request Validation
- [x] Structured Logging
- [x] Request ID Middleware
- [x] Repository Pattern
- [x] Dependency Injection
- [x] Authentication Unit Tests
- [x] Server-side Pagination
- [x] Reusable Query Builder

### In Progress

- [ ] Search
- [ ] Filtering
- [ ] Dynamic Sorting
- [ ] Project Unit Tests
- [ ] Redis Caching
- [ ] Rate Limiting
- [ ] Background Jobs

###  Planned

- [ ] Prometheus Metrics
- [ ] Grafana Dashboards
- [ ] GitHub Actions CI/CD
- [ ] Kubernetes Deployment
- [ ] Helm Charts
- [ ] Terraform Infrastructure
- [ ] AWS Deployment

---

#  Architecture

```text
                 Client
                    │
                    ▼
              Gin Router
                    │
      ┌─────────────┴─────────────┐
      ▼                           ▼
 Authentication             Project APIs
      │                           │
      ▼                           ▼
   Auth Service             Project Service
      │                           │
      └─────────────┬─────────────┘
                    ▼
            Repository Layer
                    │
                    ▼
           PostgreSQL + GORM
```

---

#  Project Structure

```text
forgeflow/

├── app/
│   ├── cmd/
│   ├── internal/
│   │   ├── handlers/
│   │   ├── middleware/
│   │   ├── models/
│   │   ├── repository/
│   │   ├── routes/
│   │   └── service/
│   └── pkg/
│       ├── dto/
│       ├── jwt/
│       ├── response/
│       └── validation/
│
├── docs/
├── deploy/
├── monitoring/
├── infrastructure/
├── docker-compose.yml
└── README.md
```

---

#  API Features

- JWT Authentication
- User Registration & Login
- Protected Routes
- Project CRUD
- Pagination
- Validation
- Swagger Documentation
- Structured Logging
- Request ID Middleware
- Health Endpoint
- Readiness Endpoint
- Version Endpoint

---

# Coming Soon

-  Search
- Filtering
- Dynamic Sorting
- Redis Cache
- Prometheus Metrics
- Grafana Dashboards
- More Unit Tests
- GitHub Actions CI/CD
- Kubernetes
- Helm
- Terraform
- AWS Deployment
