# User Management System

## Overview

The User Management System is a web application that allows administrators to manage users. It includes features such as user authentication, role-based access control, and user data management.

## Technologies
- Golang
- PostgreSQL
- Gin
- GORM
- JWT
- Bcrypt


## Features

- User Authentication
- Role-Based Access Control
- User Data Management
- Admin Dashboard

## Installation

### Prerequisites

- Go 1.18 or higher
- PostgreSQL

### Environment Variables

Create a `.env` file in the root directory and add the following environment variables:

    DATABASE_HOST=your_database_host 
    DATABASE_USER=your_database_user 
    DATABASE_PASSWORD=your_database_password 
    DATABASE_NAME=your_database_name 
    DATABASE_PORT=your_database_port

### Steps

1. Clone the repository:
   ```sh
   git clone https://github.com/dev-saiful/umanagement.git
   cd umanagement

2. Install dependencies:
    ```sh
    go mod tidy

3. Initialize the database:
    ```sh
    go run main.go

## Usage
Running the Application
To start the application, run:
```sh
go run main.go
```

API Endpoints

Authentication

* Login
    * ```POST /api/login```
    * Request Body: ```{ "email": "user@example.com", "password": "password" }```
    * Response: ```{ "token": "jwt_token" }```


User Management

* Get All Users

    * ```GET /api/users```
    * Headers: ```{ "Authorization": "Bearer jwt_token" }```
    * Response: ```[ { "id": 1, "email": "user@example.com", "username": "username" } ]```

* Get User by ID

    * ```GET /api/users/:id```
    * Headers: ```{ "Authorization": "Bearer jwt_token" }```
    * Response: ```{ "id": 1, "email": "user@example.com", "username": "username" }```

* Get Admin by Email

    * ```GET /api/admin/:email```
    * Headers: ```{ "Authorization": "Bearer jwt_token" }```
    Response: ```{ "id": 1, "email": "admin@example.com", "username": "admin", "role": "admin" }```


## Contributing
1. Fork the repository

2. Create a new branch ```(git checkout -b feature-branch)```

3. Make your changes
4. Commit your changes ```(git commit -m 'Add new feature')```
5. Push to the branch ```(git push origin feature-branch)```
6. Create a new Pull Request


## License
This project is licensed under the MIT License.

```sh
This README provides an overview of the project, installation instructions, usage details, and contribution guidelines. Adjust the content as needed to fit your specific project requirements.