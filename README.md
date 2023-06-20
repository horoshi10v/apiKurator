# apiKurator
## Go REST API Backend with GORM, MySQL, and OAuth Google

This repository contains a Go-based REST API backend application that utilizes the GORM library for interacting with a MySQL database and OAuth Google for authentication. It provides a starting point for building a robust backend server for your web or mobile applications.

## Prerequisites

Before running this application, ensure that you have the following prerequisites installed on your system:

- Go (version 1.16 or higher)
- MySQL database
- Google OAuth credentials

## Installation

1. Clone this repository to your local machine:

   ```bash
   git clone https://github.com/your-username/your-repo.git
   ```

2. Change into the project directory:

   ```bash
   cd your-repo
   ```

3. Install the necessary dependencies using the following command:

   ```bash
   go mod download
   ```

4. Set up the MySQL database by creating a new database and updating the configuration in the `config/config.go` file to match your database settings:

   ```go
   package config

   const (
   	DBUsername = "your-username"
   	DBPassword = "your-password"
   	DBHost     = "localhost"
   	DBPort     = "3306"
   	DBName     = "your-database"
   )
   ```

5. Obtain the Google OAuth credentials by creating a new project in the Google Developer Console and downloading the `credentials.json` file. Rename it to `client_secret.json` and place it in the project's root directory.

6. Start the application by running the following command:

   ```bash
   go run main.go
   ```

7. The API server will be up and running on `http://localhost:8080`.

## API Endpoints

The following API endpoints are available:

- **POST** `/api/auth/google/login`: Initiates the Google OAuth login process.
- **GET** `/api/auth/google/callback`: Callback endpoint for handling the OAuth authorization code.
- **GET** `/api/auth/google/logout`: Revokes the Google OAuth access token and logs out the user.
- **GET** `/api/users`: Retrieves a list of all users.
- **POST** `/api/users`: Creates a new user.
- **GET** `/api/users/{id}`: Retrieves a user by ID.
- **PUT** `/api/users/{id}`: Updates a user by ID.
- **DELETE** `/api/users/{id}`: Deletes a user by ID.

## Libraries Used

The following libraries were used in this project:

- [Gin](https://github.com/gin-gonic/gin): HTTP web framework
- [GORM](https://gorm.io/): ORM library for interacting with the database
- [MySQL](https://github.com/go-sql-driver/mysql): MySQL driver for Go
- [Google OAuth](https://github.com/googleapis/google-api-go-client): Google OAuth client library for Go

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please feel free to open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).
