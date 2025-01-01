# InsiderGo Project

This project is a Go-based  API application with MongoDB,Redis/InMemory integration.
It includes Swagger documentation for easy API exploration and Docker support for containerized deployment.
Also if you need to use the cross compilation option for Go use  can be examined makefile

## Features
- MongoDB integration for storing and fetching messages.
- RESTful API using Gorilla Mux.
- Swagger documentation for API endpoints.
- Docker support for containerized deployment.

## Prerequisites
- Docker installed on your system.

## Running the Project

1. **Clone the Repository**
   ```bash
   git clone https://github.com/myKemal/insiderGo.git
   cd insiderGo
   ```

2**Build and Run Methods**
   ```bash
   make compose-up
   ```
   ```bash
   make help
   ```

3**Access the Application**
    - API Base URL: `http://localhost:8090`
    - Swagger Documentation: `http://localhost:8090/swagger/`

   
## Project Structure

```plaintext
insiderGo/
├── app/                # Core application code and components
│   ├── config/         # Configuration files (e.g., reading environment variables)
│   ├── daos/           # Data Access Objects: structures representing database entities
│   ├── docs/           # Auto-generated Swagger documentation files
│   ├── dtos/           # Data Transfer Objects: structures for API input/output
│   ├── handler/        # HTTP handler functions: business logic for API endpoints       
│   ├── initialize/     # Initialization logic: setting up DB, cache, and other services         
│   ├── model/          # Core models used across the application
│   ├── repository/     # Repository layer for database and cache interactions
│   ├── router/         # HTTP router configuration (e.g., defining API routes)
│   ├── services/       # Business logic and reusable services
│   ├── app.go
├── main.go
├── docker-compose.override.yml     # Override file for Docker Compose (e.g.,choose temp memory as inMemory)
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── Makefile                        # Automation of common tasks (e.g., build, run, test)
└── README.md
```

### Future Improvements

The following enhancements are planned or could be considered in the future:

1. **Improved Performance:**
    - When the unsent message list in the instance decrease below a certain limit, unsent messages can be retrieved from the DB.
    - The period duration and the number of messages sent can be updated dynamically according to data such as the reading speed in the DB regarding the sending of 2 messages in 2 minutes, the statistical distribution of unsent data over time, etc.

2. **Monitoring and Logging:**
    - Integrate OpenTelemetry or Prometheus for distributed tracing and metric collection.
    - Add structured logging with configurable levels (debug, info, warning, error).
   
3. **CI/CD Integration:**
- Automate build, test, and deployment processes with GitHub Actions or GitLab CI/CD.


