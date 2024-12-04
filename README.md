# service-todolist

A Todo List service implemented using Go with **Gin Gonic**, **GORM**, and **PostgreSQL**, following **Clean Architecture** principles.

---

## **Features**

- Manage todos with Create, Read, Update, and Delete operations.
- Search todos by title, due date, priority, and completion status.
- API documentation available via Swagger.

---

## **Getting Started**

### **1. Prerequisites**

Ensure you have the following tools installed:

- [Docker](https://www.docker.com/)
- [Go](https://golang.org/) (version 1.19 or later)

---

### **2. Setup**

#### **Option 1: Using `make`**

Run the following command to set up the database and perform migrations:

```bash
make setup
```

#### **Option 2: Using `Manual Setup`**

1. Start the database container:

```bash
docker compose -f docker-compose.db.yml up -d
```

2. Run the database migration:

```bash
go run database/migration/migration.go
```
