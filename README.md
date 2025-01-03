# Task-Scheduler
# Task Scheduler

A simple full-stack Task Manager application where users can create, read, update, and delete tasks. 
The application consists of a **Golang backend** and a **VueJS frontend**.

---

## Features

- **Task Management**: Add, view, update, delete Tasks.
- **Database Integration**: Storage using PostgreSQL.
- **Dockerized Deployment**: Run the app using Docker Compose.

---

## Prerequisites

Ensure the following are installed on your system:

- **Go** (Golang): [Download and Install Go](https://golang.org/dl/)
- **Node.js and npm**: [Download and Install Node.js](https://nodejs.org/)
- **PostgreSQL**: Installed locally or in Docker.(https://hub.docker.com/_/postgres )
- **Docker** (Optional for containerized setup): [Get Docker](https://www.docker.com/)

---

## Installation

### 1. Clone the Repository
```bash
git clone https://github.com/username/task-scheduler.git
cd task-scheduler
```
### 2. Run backend and frontend
```bash
cd backend
go run main.go

cd frontend 
cd task-scheduler
npm run serve
```
The frontend will run on  "http://localhost:5173"