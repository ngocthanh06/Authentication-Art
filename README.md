# Authentication Service

## Overview
This project provides a secure authentication service for your application. It includes features such as user registration, login, password reset, and token-based authentication.

## Getting Started

### Prerequisites
* Go 1.20
* Docker
* Docker Compose
* PostgreSQL

### Installation
1. Clone the repository:
   ```bash
   git clone [git@github.com:ngocthanh06/Authentication-Art.git](git@github.com:ngocthanh06/Authentication-Art.git)
   ```
2. Copy .env
   ```bash 
   cp .env.example .env
   ```
 
3. Build source:
   ```bash
   make build

4. Generate jwt key:
   ```bash
   make init-jwt-key
   ```

5. Run migrate:
   ```bash
   make migrate-up
   ```
6. Access link: localhost:8080/register/
```curl
   curl --location 'localhost:8080/v1/register' \
   --header 'Content-Type: application/x-www-form-urlencoded' \
   --data-urlencode 'title=user' \
   --data-urlencode 'first_name=user' \
   --data-urlencode 'email=user@gmail.com' \
   --data-urlencode 'last_name=user' \
   --data-urlencode 'password=password'
```

## Structure

```markdown
   ## Authentication structure
   
   |-- app # This directory commonly houses the main application logic
      |-- authentication # main application
   |-- cmd # command
   |-- docker
   |-- internal
      |-- common # Reusable code components like utility functions or data structures.
      |-- config # Configuration files for the application
      |-- database # Code related to database interaction, including models or database access layer.
      |-- handlers #  Functions responsible for handling incoming requests and generating responses. (Similar to controllers in Laravel)
      |-- middleware # Middleware functions that intercept requests before reaching handlers, adding functionality like authentication or logging.
      |-- models # Data structures representing entities in the application, often mapped to database tables.
      |-- providers # Code that registers services with the application's dependency injection container.
      |-- repositories # Abstractions responsible often interacting with models.
      |-- routes # Defines the routing rules for the application, mapping URLs to handlers
      |-- services # Business logic encapsulated in reusable services used by handlers or other parts of the application.
      |-- utils # Utility functions specific to the project's needs.
   |-- makefile # This file might contain build instructions for the project, automating tasks like compiling code or running tests.
   |-- readme.me
   |-- go.mod # declares the project's dependencies and their versions.
   |-- go.sum # contains cryptographic checksums for downloaded dependencies, ensuring their integrity.
   |-- .air.toml
   |-- .env.example
   |-- .gitignore
```
### How to use
Add new file: `./internal/repository/your_repository.go`
```go
   package repositories
   
   import (
      "gorm.io/gorm"
   )
   
   func NewAuthRepository(dtb *gorm.DB) *DbStorage {
      return &DbStorage {
         db: dtb,
      }
   }
```
Add new file: `./internal/services/your_service.go`
```go
   package services
   
   type YourService struct {
      yourRepository *repositories.DbStorage
      userRepository *repositories.DbStorage
   }
   
   func NewAuthService(yourRepository *repositories.DbStorage) *YourService {
      return &YourService {
         yourRepository: yourRepository,
      }
   }

```
Open file `./internal/providers/provider.go`
```go
   var YourServ *services.YourService
   
   func ConfigSetupProviders() {
        ...
       // repository
       YourRepo := repositories.NewYourRepository(database.DB)
       YourServ = services.NewYourService(YourRepo)
       ...
   }
```
# Still Development
