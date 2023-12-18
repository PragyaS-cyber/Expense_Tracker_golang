# Expense Tracking System - GoFr Golang

A simple HTTP (REST) API for managing expenses of a user. The API supports CRUD operations and integrates with a database for data persistence.

## Table of Contents

- [Features](#features)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Running the API](#running-the-api)
- [API Endpoints](#api-endpoints)
- [Usage Examples](#usage-examples)
- [Testing](#testing)
- [Contributing](#contributing)

## Features

- Create, Read, Update, and Delete tasks.
- User authentication for personalized task management.
- Manage task categories for better organization.

## Getting Started

### Prerequisites

- Go programming language installed
- PostgreSQL database (can be run using Docker)

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/PragyaS-cyber/Expense_Tracker_golang.git
   cd expense_tracker
   ```

2. Install dependencies:

	```bash
	go mod download
	```



1. Set up PostgreSQL database:
	```bash
	docker run -d -p 5432:5432 --name todo-db -e POSTGRES_USER=your_username -e POSTGRES_PASSWORD=your_password -e POSTGRES_DB=todo_db postgres
	```

2. Update database connection details:
	```bash
	// Update connection details based on Dockerized PostgreSQL instance
	db, err = gorm.Open("postgres", "host=localhost user=your_username dbname=todo_db sslmode=disable password=your_password")
	```

### Run the API:
	```bash
	go run main.go
	```

   


![Alt text](image.png)