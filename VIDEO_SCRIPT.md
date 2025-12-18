# Video Presentation Script: Microservices CI/CD Pipeline

**Duration**: ~5 Minutes
**Speaker**: DevOps Team Lead

---

## 0:00 - 0:45: Introduction
**Visual**: Slide showing the 5 Microservices (Auth, User, Product, Order, Payment) and the high-level goal (Independent Deployment).
**Audio**: "Hello, I am the lead DevOps engineer. Today I will present our robust CI/CD pipeline design for a distributed ecosystem of five microservices. Our goal was to ensure that each service—Auth, User, Product, Order, and Payment—can be developed, tested, and deployed entirely independently while maintaining system integrity."

## 0:45 - 1:30: Architecture & Version Control
**Visual**: Show the file structure in VS Code (Monorepo structure). Zoom in on `services/` folder.
**Audio**: "We adopted a Monorepo strategy. This allows us to keep all our services in one place for better visibility and shared tooling. However, strict independence is enforced via our CI/CD triggers. As you can see, we have a `services` directory containing all 5 services, each with its own Dockerfile and package configuration."

## 1:30 - 2:30: Build & Test Automation (The CI Phase)
**Visual**: Open `.github/workflows/user-service-ci.yml`. Highlight the `paths` trigger and the `pytest` / `flake8` steps.
**Audio**: "The core of our independence lies in GitHub Actions Path Filtering. Look at this workflow for the User Service. It *only* triggers when files in `services/user-service` are changed.
Crucially, we don't just build; we verify. We have implemented a polyglot testing strategy:
*   **Mocha** for our Node.js services.
*   **Pytest and Flake8** for Python.
*   **Go Test** for Golang.
The pipeline fails immediately if any test fails or if code quality drops, ensuring only clean code reaches the build stage."

## 2:30 - 3:30: Containerization & Registry
**Visual**: Show the `Dockerfile` in `services/user-service` and then the `docker build` step in the YAML.
**Audio**: "Once testing passes, we move to containerization. We use Docker to package every service. In the CI pipeline, we build a unique image tagged with the Git Commit SHA. This ensures traceability—we know exactly which code is running in any container. These images are then pushed to our central Container Registry."

## 3:30 - 4:15: Continuous Deployment (The CD Phase)
**Visual**: Diagram of the flow: Staging -> Integration Test -> Production.
**Audio**: "For deployment, we follow a rigorous path. The new image is first deployed to a Staging environment. Here, we run automated integration tests to ensure the service plays nicely with others. Only after these tests pass do we promote the image to Production. This prevents a broken microservice from taking down the whole system."

## 4:15 - 5:00: Local Development & Live Demo
**Visual**: Show `docker-compose.yml`, run `docker-compose up`, and then **open the browser to localhost:3000, 5002, etc.** to show the Status Pages.
**Audio**: "Finally, for developer experience, we provide a `docker-compose` file that spins up the entire ecosystem locally.
As you can see here, I can access every service independently in the browser. Each service renders its own status page showing its version and technology stack. This proves that our ports are correctly mapped and the services are healthy.
In summary, this pipeline delivers modularity, high stability, and complete automation, allowing our team to ship features faster and safer."

---
**End of Presentation**
