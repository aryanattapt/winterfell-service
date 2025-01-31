# Winterfell Project Service

This is a Go-based project. Follow the steps below to clone the repository and run the project locally.

## Prerequisites

Before running the project, ensure you have the following installed:

- **Go** (version 1.23.5) – [Install Go](https://golang.org/dl/)
- **Git** – [Install Git](https://git-scm.com/)
- **Docker** – [Install Docker](https://www.docker.com/get-started)
  
You need Docker to run database as a container. Database, we will use:
- **Mongo DB 8**
- **Elastic Search 8**
- **Postgre SQL 17**

## Steps to Run the Project

### 1. Clone the Repository

Start by cloning the repository to your local machine. Run the following command in your terminal:

```bash
git clone https://github.com/aryanattapt/wintefell-service.git
```

### 2. Install Dependency

To Install All Dependency please Run the following command in your terminal:

```bash
go mod tidy
```

### 3. Run Project

To run the project via HTTP Network, please Run the following command in your terminal:

```bash
cd cmd/app/
go run main.go
```

Open Browser And Type http://localhost:8000