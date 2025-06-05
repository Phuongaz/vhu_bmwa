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

### Authentication

- `POST /api/auth/register` — Register new user  
- `POST /api/auth/login` — User login and token 
- `POST /api/auth/logout` — User logout (requires authentication)
- `GET /api/profile` — Get user profile (requires authentication)

- `GET /api/oauth2/authorize` — Google login
- `GET /api/oauth2/callback` — Google callback

### Product Management (Admin only)

- `GET /api/products` — Get all products  
- `POST /api/products` — Create a new product  
- `PUT /api/products/:id` — Update a product  
- `DELETE /api/products/:id` — Delete a product  

---

## Logging System

- `access.log` — Logs all HTTP requests  
- `security.log` — Logs login attempts, rate limiting, and suspicious input  
- `error.log` — Logs backend errors with stack traces  

---

## Postman Collection

[REST API SECURE.postman_collection.json](REST%20API%20SECURE.postman_collection.json)
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
