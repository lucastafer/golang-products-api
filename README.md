# Go Products API Project

This was my first simple project using Golang, developed with the goal of getting to know the language better. I built a CRUD API from scratch, based on the tutorial by [GoLabTutorials](https://www.youtube.com/@GoLabTutoriais) on YouTube. At the end of the video, I was challenged to implement two additional routes, which I successfully completed.

The project serves as a foundation for creating a RESTful API using the Gin framework, connecting to a PostgreSQL database for CRUD (Create, Read, Update, Delete) operations on products. The application runs in a Docker container, with Docker Compose.

## Technologies

- **Go (Golang)** as programming language;
- **Gin** as Web Framework web for Go;
- **PostgreSQL** as relational database;
- **Docker** for containerizing the application, ensuring consistent execution across environments.

## Pre-requisites

Make sure you have the following tools installed:

- [Go](https://golang.org/dl/)
- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)

## How to run project

### 1. Clone repo

```bash
git clone https://github.com/lucastafer/golang-products-api.git
cd golang-products-api
```

### 2. Build and run the containers

`docker-compose up --build`

### 3. Access API via Postman or similar

Once the containers are up, the API will be available at `http://localhost:8000`.
