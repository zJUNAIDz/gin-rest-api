# Gin REST API

## Overview

This is a RESTful API built with the [Gin Web Framework](https://gin-gonic.com/) in Go. It is designed to manage users, events, and event registrations, offering endpoints for creating, retrieving, updating, and deleting resources. Authentication is implemented using JWT.

---

## Features

- **User Authentication**:

  - User signup and login using JWT-based authentication.
  - Secure password hashing with bcrypt.

- **Event Management**:

  - Create, update, delete, and view events.
  - Users can register for and cancel event participation.

- **Database Integration**:

  - Uses SQLite for persistent data storage.

- **Cross-Platform Executables**:
  - Precompiled executables available for Windows and Linux systems.

---

## Installation and Setup

### Prerequisites

- Go 1.20 or higher installed on your system (if building from source).

### Download Executables

Precompiled executables for Windows and Linux are available in the [Releases](https://github.com/zjunaidz/gin-rest-api/releases) section of the repository.

#### For Windows

1. Download the `.exe` file from the Releases.
2. Double-click the file to run the server.

#### For Linux

1. Download the Linux executable.
2. Grant execute permissions:

   ```bash
   chmod +x gin-rest-api
   ```

3. Run the executable:

   ```bash
   ./gin-rest-api
   ```

### Build from Source

#### Clone the Repository

```bash
git clone https://github.com/zjunaidz/gin-rest-api.git
cd gin-rest-api
```

#### Environment Configuration

1.Create a `.env` file in the project root directory by copying `.example.env`:

```bash
cp .example.env .env
```

2.Replace the placeholder values in `.env` with your configuration.

```env
  JWT_SECRET="your-jwt-secret"
```

#### Install Dependencies

```bash
go mod tidy
```

#### Run the Application

```bash
go run main.go
```

The application will run on `http://localhost:8080` by default.

---

## API Endpoints

### Public Endpoints

#### User Authentication

- **1. Signup**

- **URL**: `/signup`
- **Method**: `POST`
- **Body**:

  ```json
  {
    "email": "user@example.com",
    "password": "your_password"
  }
  ```

- **Response**:

  ```json
  {
    "message": "User created successfully.",
    "user": {
      "id": 1,
      "email": "user@example.com"
    }
  }
  ```

- **2. Login**

- **URL**: `/login`
- **Method**: `POST`
- **Body**:

  ```json
  {
    "email": "user@example.com",
    "password": "your_password"
  }
  ```

- **Response**:

  ```json
  {
    "message": "Login successful.",
    "token": "your_jwt_token"
  }
  ```

#### Fetch Events

- **3. Get All Events**

- **URL**: `/events`
- **Method**: `GET`
- **Response**:

  ```json
  [
    {
      "id": 1,
      "name": "Event 1",
      "description": "Description for Event 1",
      "location": "Location 1",
      "datetime": "2023-01-01T12:00:00Z",
      "user_id": 1
    }
  ]
  ```

- **4. Get Event by ID**

- **URL**: `/events/:id`
- **Method**: `GET`

---

### Protected Endpoints (Authentication Required)

Add the JWT token in the `Authorization` header:

```http
Authorization: Bearer your_jwt_token
```

#### Event Management

- **1. Create Event**

- **URL**: `/events/create`
- **Method**: `POST`
- **Body**:

  ```json
  {
    "name": "Event Name",
    "description": "Event Description",
    "location": "Event Location"
  }
  ```

- **2. Update Event**

- **URL**: `/events/update/:id`
- **Method**: `PUT`

- **3. Delete Event**

- **URL**: `/events/delete/:id`
- **Method**: `DELETE`

#### Event Registrations

- **4. Register for Event**

- **URL**: `/events/register/:id`
- **Method**: `POST`

- **5. Cancel Registration**

- **URL**: `/events/cancel/:id`
- **Method**: `DELETE`

---

## Project Structure

```plaintext
.
├── db             # Database initialization and migrations
├── middlewares    # Authentication middleware
├── models         # Business logic for users and events
├── routes         # API route handlers
├── utils          # Utility functions (JWT, password hashing)
├── main.go        # Application entry point
```

---

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a new branch: `git checkout -b feature-name`.
3. Make your changes and commit them.
4. Push to your fork and create a pull request.

---

## License

This project is licensed under the [MIT License](LICENSE).
