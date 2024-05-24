# My Project

This project is a web application built using [Golang][docker].

## Prerequisites

Before running the project, make sure you have the following installed:

- [Git](https://git-scm.com/downloads)
- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Getting Started

1. Clone the repository or pull the latest changes:

```bash
git clone https://github.com/PhuPhuoc/nextbean_center_api.git
docker compose up --build
```

2. Run the database migrations to set up your database schema:
```bash
migrate -path db/migrations -database "mysql://{username}:{password}@tcp({host}:{port})/{database}" up
```