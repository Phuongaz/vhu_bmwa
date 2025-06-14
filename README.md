# VHU_BMWA — Secure RESTful API Example

## Overview

**VHU_BMWA** is a full-stack project that demonstrates how to build a secure RESTful API with authentication, logging, and containerized deployment. This project was developed as part of the Web & Application Security course at VHU.

---

## Objectives

- Develop a RESTful API using Golang and Gin framework
- Implement JWT-based authentication and role-based access control
- Add structured logging (access, security, and error logs)
- Apply basic security measures: bcrypt, rate limiting, input validation
- Deploy the system using Docker and Nginx
- Implement OAuth2.0 authentication with Google
---

## Technology Stack
- **Frontend**: Vue.js 3, Vite, Tailwind CSS, Pinia
- **Backend**: Golang (Gin)
- **Database**: MySQL
- **Authentication**: JWT
- **Logging**: Logrus
- **Containerization**: Docker & Docker Compose
- **Proxy**: Nginx
- **OAuth2.0**: Google
---

## Security Features

- JWT token authentication
- Password hashing with bcrypt
- Rate limiting (10 requests/second)
- Input validation
- Role-based access for admin endpoints
- OAuth2.0 authentication with Google
---

## API Endpoints

### V1
This version is included more vulnerabilities.
- No JWT token authentication (Just store username/password in localStorage)
- Authorize with header X-Role (take it from localStorage)
- No rate limiting
- No input validation
- Raw SQL query

##### Authentication

- `POST /api/v1/auth/register` — Register new user  
- `POST /api/v1/auth/login` — User login
- `POST /api/v1/auth/logout` — User logout (requires authentication)
- `GET /api/v1/profile` — Get user profile (requires authentication)

### Product Management (Admin only)
- `GET /api/v1/products` — Get all products  
- `POST /api/v1/products` — Create a new product  
- `PUT /api/v1/products/:id` — Update a product  
- `DELETE /api/v1/products/:id` — Delete a product  

### V2
This version is included more security features.
- JWT token authentication
- Rate limiting (10 requests/second)
- Input validation
- Role-based access for admin endpoints
- OAuth2.0 authentication with Google
- Prepared statement
- Use ORM (Gorm)
- Store JWT token in cookies

##### Authentication

- `POST /api/v2/auth/register` — Register new user  
- `POST /api/v2/auth/login` — User login
- `POST /api/v2/auth/logout` — User logout (requires authentication)
- `GET /api/v2/profile` — Get user profile (requires authentication)

### Product Management (Admin only)
- `GET /api/v2/products` — Get all products  
- `POST /api/v2/products` — Create a new product  
- `PUT /api/v2/products/:id` — Update a product  
- `DELETE /api/v2/products/:id` — Delete a product  
---

## Logging System

- `access.log` — Logs all HTTP requests  
- `security.log` — Logs login attempts, rate limiting, and suspicious input  
- `error.log` — Logs backend errors with stack traces  

---
## Deployment

### Requirements

- Docker  
- Docker Compose

### Steps

```bash
git clone https://github.com/phuongaz/VHU_BMWA.git
cd VHU_BMWA
docker-compose up --build -d
---
```
