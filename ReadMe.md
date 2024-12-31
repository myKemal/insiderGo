# MongoAPI Project

This project is a Go-based RESTful API application with MongoDB integration. It includes Swagger documentation for easy API exploration and Docker support for containerized deployment.

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

2. **Create a .env File**
   Create a `.env` file in the project root with the following content:
   ```env
   MONGODB_URI=mongodb+srv://<username>:<password>@cluster.mongodb.net/<database>?retryWrites=true&w=majority
   PORT=8080
   ```
   Replace `<username>`, `<password>`, and `<database>` with your MongoDB Atlas credentials.

3. **Build and Run the Docker Container**
   ```bash
   docker build -t mongoapi .
   docker run -p 8080:8080 --env-file .env mongoapi
   ```

4. **Access the Application**
    - API Base URL: `http://localhost:8080`
    - Swagger Documentation: `http://localhost:8080/swagger/`

## Swagger Documentation

Swagger is integrated into the project for exploring and testing API endpoints.

### Access Swagger UI

1. Start the application as described above.
2. Open your browser and navigate to `http://localhost:8080/swagger/`.
3. Use the Swagger UI to test and explore the API endpoints.

### Example Endpoints

- **GET /unsent-messages**: Fetch all messages with a sending status of `not_sent`.

## Project Structure

```plaintext
mongoApi/
├── app/
│   ├── app.go
│   ├── handler/
│   │   ├── handler.go
│   ├── repo/
│   │   ├── repo.go
│   ├── router/
│   │   ├── router.go
│   ├── daos/
│   │   ├── message.go
│   ├── dtos/
│   │   ├── message.go
├── main.go
├── Dockerfile
├── go.mod
└── README.md
```

## Stopping the Application
To stop the running container:
```bash
docker stop $(docker ps -q --filter ancestor=mongoapi)
```

