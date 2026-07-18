# Dockerized Go Services

This repository contains a collection of **Dockerized Go (Golang) web services** developed as part of learning **Docker, Go, and Microservices**. Each service is implemented as an independent Go application and containerized using a custom Dockerfile.

## Technologies Used

- Go (Golang)
- Docker
- HTTP Web Services
- Git & GitHub

## Project Structure

```
Docker_services/
│
├── Calculator/
├── Greeting/
├── HelloWorld/
├── Square/
├── Temperature/
├── Weather/
├── Factorial/
├── EvenOdd/
├── RandomNumber/
└── CurrentTime/
```

Each service contains:

```
ServiceName/
├── Dockerfile
├── go.mod
├── README.md
└── *.go
```

## Services Included

| Service | Description | Port |
|----------|-------------|------|
| Calculator | Performs basic arithmetic operations | 8080 |
| Greeting | Returns a personalized greeting | 8081 |
| Hello World | Displays a Hello World message | 8082 |
| Square | Calculates the square of a number | 8083 |
| Temperature | Converts Celsius to Fahrenheit | 8084 |
| Weather | Displays sample weather information | 8085 |
| Factorial | Calculates the factorial of a number | 8086 |
| Even/Odd | Checks whether a number is even or odd | 8087 |
| Random Number | Generates a random number | 8088 |
| Current Time | Displays the current system date and time | 8089 |

## Running a Service

Navigate to any service folder.

Example:

```bash
cd Greeting
```

Build the Docker image:

```bash
docker build -t greeting-service .
```

Run the container:

```bash
docker run -d -p 8081:8080 --name greeting-container greeting-service
```

Open in your browser:

```
http://localhost:8081/greet?name=Grace
```

## Learning Outcomes

- Understanding Docker fundamentals
- Creating custom Dockerfiles
- Building Docker images
- Running Docker containers
- Developing RESTful web services in Go
- Managing projects using Git and GitHub
- Organizing multiple microservices in a single repository

## Author

**Grace Paul**

B.Tech Computer Science and Engineering

Christ College of Engineering, Irinjalakuda
