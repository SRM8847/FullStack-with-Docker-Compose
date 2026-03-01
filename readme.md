# 🐳 FullStack Application with Docker Compose

A fully containerized full-stack web application built with a **Go** backend, **JavaScript/HTML/CSS** frontend, **PostgreSQL** database, and **Nginx** as a reverse proxy — all orchestrated using **Docker Compose**.

---

## 📋 Table of Contents

- [Overview](#overview)
- [Tech Stack](#tech-stack)
- [Project Structure](#project-structure)
- [Architecture](#architecture)
- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
  - [1. Clone the Repository](#1-clone-the-repository)
  - [2. Configure Environment Variables](#2-configure-environment-variables)
  - [3. Build and Run](#3-build-and-run)
  - [4. Access the Application](#4-access-the-application)
- [Services](#services)
- [Environment Variables](#environment-variables)
- [Useful Docker Commands](#useful-docker-commands)
- [Development Tips](#development-tips)
- [Troubleshooting](#troubleshooting)
- [Contributing](#contributing)
- [License](#license)

---

## 🧭 Overview

This project demonstrates how to containerize and orchestrate a complete full-stack web application using Docker and Docker Compose. It is ideal as a learning template or a starting point for production-grade containerized applications.

With a single command (`docker compose up -d --build`), the entire stack — frontend, backend, database, and reverse proxy — is spun up and ready to use.

---

## 🛠 Tech Stack

| Layer         | Technology                        |
|---------------|-----------------------------------|
| **Frontend**  | HTML, CSS, JavaScript             |
| **Backend**   | Go (Golang)                       |
| **Database**  | PostgreSQL                        |
| **Proxy**     | Nginx (Reverse Proxy)             |
| **Container** | Docker & Docker Compose           |

---

## 📁 Project Structure

```
FullStack-with-Docker-Compose/
├── backend/               # Go REST API server
│   ├── Dockerfile         # Multi-stage build for Go app
│   ├── main.go            # Entry point
│   └── ...
├── frontend/              # Static frontend (HTML/CSS/JS)
│   ├── Dockerfile         # Nginx-served static files
│   ├── index.html
│   ├── style.css
│   └── ...
├── db/                    # Database init scripts
│   └── init.sql           # SQL schema / seed data
├── nginx/                 # Nginx configuration
│   └── nginx.conf         # Reverse proxy routing rules
├── docker-compose.yaml    # Docker Compose service definitions
├── .gitignore             # Ignores .env and other artifacts
└── README.md              # Project documentation
```

---

## 🏗 Architecture

```
                        ┌─────────────────────────────────────┐
                        │           Docker Network             │
                        │                                      │
  Browser  ──────────▶  │  Nginx (Port 80)                     │
                        │    │                                 │
                        │    ├──▶  /api/*  ──▶  Go Backend     │
                        │    │                    │            │
                        │    └──▶  /*  ──────▶  Frontend       │
                        │                     (Static Files)   │
                        │         Go Backend ──▶ PostgreSQL    │
                        │                                      │
                        └─────────────────────────────────────┘
```

- **Nginx** acts as a reverse proxy, routing `/api/*` requests to the Go backend and all other routes to the frontend.
- **Go Backend** handles API logic and communicates with the PostgreSQL database.
- **PostgreSQL** persists data and is initialized using scripts in the `db/` directory.
- All services communicate over an isolated Docker internal network.

---

## ✅ Prerequisites

Make sure you have the following installed on your machine:

- [Docker](https://docs.docker.com/get-docker/) (v20.10+)
- [Docker Compose](https://docs.docker.com/compose/install/) (v2.0+ — included in Docker Desktop)

To verify your installations:

```bash
docker --version
docker compose version
```

---

## 🚀 Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/SRM8847/FullStack-with-Docker-Compose.git
cd FullStack-with-Docker-Compose
```

### 2. Configure Environment Variables

Create a `.env` file in the **root of the project folder**. This file holds your PostgreSQL credentials:

```bash
touch .env
```

Add the following content to `.env`:

```env
POSTGRES_USER=appuser
POSTGRES_PASSWORD=changeme
POSTGRES_DB=appdb
```

> ⚠️ **Security Note:** Never commit your `.env` file to version control. It is already included in `.gitignore`. Feel free to change the values above to your liking.

### 3. Build and Run

Build all Docker images and start all services in detached mode with a single command:

```bash
docker compose up -d --build
```

This will:
- Build the Go backend image
- Build the frontend image
- Pull the official PostgreSQL image
- Configure and start Nginx as a reverse proxy
- Connect all services over a shared Docker network

### 4. Access the Application

Once all containers are running, open your browser and navigate to:

```
http://localhost
```

The Nginx proxy will serve the frontend at the root URL and forward API calls accordingly.

---

## 📦 Services

The `docker-compose.yaml` defines the following services:

| Service      | Description                                      | Internal Port |
|--------------|--------------------------------------------------|---------------|
| `frontend`   | Static HTML/CSS/JS served via a web server       | 80            |
| `backend`    | Go REST API server                               | 8080          |
| `db`         | PostgreSQL database                              | 5432          |
| `nginx`      | Reverse proxy routing traffic to frontend/backend| 80 (exposed)  |

---

## 🔐 Environment Variables

The following environment variables are used by the application and should be defined in your `.env` file:

| Variable            | Description                         | Example Value |
|---------------------|-------------------------------------|---------------|
| `POSTGRES_USER`     | PostgreSQL username                 | `appuser`     |
| `POSTGRES_PASSWORD` | PostgreSQL password                 | `changeme`    |
| `POSTGRES_DB`       | Name of the PostgreSQL database     | `appdb`       |

---

## 🧰 Useful Docker Commands

```bash
# Start all services (detached, rebuild images)
docker compose up -d --build

# Start all services without rebuilding
docker compose up -d

# Stop all running containers
docker compose stop

# Stop and remove containers, networks, and volumes
docker compose down

# Remove containers AND volumes (⚠️ deletes DB data)
docker compose down -v

# View logs for all services
docker compose logs -f

# View logs for a specific service (e.g., backend)
docker compose logs -f backend

# List all running containers
docker compose ps

# Execute a command inside a running container
docker compose exec backend sh

# Rebuild a single service
docker compose up -d --build backend
```

---

## 💡 Development Tips

**Rebuilding after code changes:**
If you make changes to the Go backend or frontend source files, rebuild the affected service:
```bash
docker compose up -d --build backend
# or
docker compose up -d --build frontend
```

**Connecting to the PostgreSQL database directly:**
```bash
docker compose exec db psql -U appuser -d appdb
```

**Checking container health:**
```bash
docker compose ps
```
All services should show a status of `running`.

**Inspecting the Docker network:**
```bash
docker network ls
docker network inspect <network_name>
```

---

## 🔧 Troubleshooting

**Problem: Port 80 is already in use**
> Another service (like Apache or a local Nginx) might be using port 80.
```bash
# Find and stop the conflicting process
sudo lsof -i :80
sudo kill -9 <PID>
```
Alternatively, change the exposed port in `docker-compose.yaml`:
```yaml
ports:
  - "8080:80"   # change 8080 to any free port
```

**Problem: Database connection refused**
> The backend may start before PostgreSQL is fully ready. Add a `depends_on` with a health check in `docker-compose.yaml`, or wait a few seconds and restart the backend:
```bash
docker compose restart backend
```

**Problem: Changes not reflected after rebuild**
> Docker may use cached layers. Force a fresh build:
```bash
docker compose build --no-cache
docker compose up -d
```

**Problem: `.env` values not being picked up**
> Ensure the `.env` file is in the **root directory** of the project (same level as `docker-compose.yaml`) and does not have any extra spaces or quotes around values.

---

## 🤝 Contributing

Contributions are welcome! To contribute:

1. Fork the repository
2. Create a new branch: `git checkout -b feature/your-feature-name`
3. Make your changes and commit: `git commit -m "Add your feature"`
4. Push to the branch: `git push origin feature/your-feature-name`
5. Open a Pull Request

Please make sure your code is clean, well-commented, and follows existing conventions.

---

## 📄 License

This project is open source. Feel free to use it as a reference or starting point for your own containerized applications.

---

> Made with ❤️ using Go, Docker, PostgreSQL, and Nginx.
