# DevOps Microservices CI/CD Project

This project contains the source code, CI/CD pipeline configuration, and design documentation for a distributed microservices application.

## 📂 Project Structure

```
.
├── .github/workflows/       # CI/CD Pipelines (GitHub Actions)
│   ├── auth-service-ci.yml
│   ├── user-service-ci.yml
│   └── product-service-ci.yml
├── services/                # Source code for Microservices
│   ├── auth-service/        # Node.js
│   ├── user-service/        # Python/Flask
│   ├── product-service/     # Go
│   ├── order-service/       # Node.js
│   └── payment-service/     # Python/Flask
├── DESIGN.md                # Detailed Pipeline Architecture & Design
├── VIDEO_SCRIPT.md          # Script for the video presentation
└── docker-compose.yml       # Local development setup
```

## 🚀 Getting Started

1.  **Read the Design**: Start with [DESIGN.md](DESIGN.md) to understand the architecture and pipeline strategy.
2.  **Run Locally**:
    ```bash
    docker-compose up --build
    ```
    This will build and start all 5 microservices locally.
3.  **Check Workflows**: Navigate to `.github/workflows` to see the YAML configurations for independent CI/CD.

## 🔌 Service Endpoints
*   **Auth Service**: http://localhost:3000
*   **User Service**: http://localhost:5002
*   **Product Service**: http://localhost:8081
*   **Order Service**: http://localhost:3001
*   **Payment Service**: http://localhost:5001

## 🛠 Features
*   **Monorepo Structure**: Centralized management with independent deployments.
*   **Path-Based Triggers**: Modifying one service only triggers that service's pipeline.
*   **Polyglot Support**: Pipelines for Node.js, Python, and Go.
*   **Containerization**: Dockerfiles for all services.
