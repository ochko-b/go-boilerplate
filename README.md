# Go Boilerplate

This is a boilerplate project for building RESTful APIs with Go, utilizing PostgreSQL for data persistence and JWT for authentication. It provides a structured foundation for scalable web applications, complete with Docker support and SQL code generation.

## Tech Stack

- **Go**: The core programming language.
- **Fiber**: A fast and flexible web framework for Go (replacing the standard `net/http`, as seen in `main.go`).
- **sqlc**: Generates type-safe Go code from SQL queries.
- **PostgreSQL**: The relational database (confirmed via `docker-compose.yml`).
- **JWT**: Authentication via JSON Web Tokens (implemented in `internal/utils/jwt.go`).
- **godotenv**: Loads environment variables from a `.env` file (used in `main.go`).
- **Middleware**: Custom middleware for authentication, CORS, and logging (in `internal/middleware/`).
- **Docker**: Containerization support with `Dockerfile` and `docker-compose.yml`.

## File Structure

- **`cmd/server/`**: Application entry point and routing setup.

  - `main.go`: Initializes the configuration, database connection, services, handlers, and starts the Fiber server.
  - `routes/`:
    - `auth.go`: Defines authentication routes (e.g., login, register).
    - `setup.go`: Configures the Fiber app with routes and middleware.
    - `user.go`: Defines user-related routes (e.g., user profile, update).

- **`internal/`**: Core application logic (private to the project).

  - `config/`: Configuration management (`config.go`).
  - `database/`: Database setup and migrations.
    - `connection.go`: Establishes the database connection.
    - `migrations/`: SQL migration files (e.g., `0001_initial_schema.up.sql`).
  - `handlers/`: HTTP request handlers (`auth.go`, `health.go`, `user.go`).
  - `middleware/`: Request processing utilities (`auth.go`, `cors.go`, `logger.go`).
  - `models/`: Data structures (`auth.go`, `users.go`).
  - `repository/`: Data access layer (`repository.go`).
  - `services/`: Business logic (`auth.go`, `user.go`).
  - `utils/`: Helper functions (`jwt.go`, `password.go`, `response.go`, `uuid.go`).

- **`pkg/`**: Reusable, public packages.

  - `validator/`: Validation utilities (`validator.go`).

- **`generated/`**: Auto-generated code.

  - `sqlc/`: SQLc-generated files (`db.go`, `models.go`, `querier.go`, `users.sql.go`).

- **`sql/`**: SQL definitions.

  - `queries/`: SQL query files (`users.sql`).
  - `schema.sql`: Database schema definition.

- **`bin/`**: Compiled binaries (e.g., `main`).

- **`Dockerfile`**: Docker configuration for building the app image.

- **`docker-compose.yml`**: Multi-container setup for the app and PostgreSQL database.

- **`Makefile`**: Build, run, test, and automation scripts.

- **`go.mod` & `go.sum`**: Go module dependencies.

- **`sqlc.yaml`**: Configuration for sqlc code generation.

## Setup and Installation

1. **Install Go**:

   - Download and install from [golang.org](https://golang.org/dl/).

2. **Set Up PostgreSQL**:

   - Use Docker Compose (from `docker-compose.yml`) to run PostgreSQL: `docker-compose up -d`.
   - Alternatively, install PostgreSQL locally and create a database: `createdb goappdb`.

3. **Load Environment Variables**:

   - The project uses `godotenv` (in `main.go`) to load variables from a `.env` file.
   - Create a `.env` file in the root directory with:
     ```bash
     DB_HOST=localhost  # or 'postgres' if using Docker Compose
     DB_PORT=5432
     DB_USER=postgres
     DB_PASSWORD=password
     DB_NAME=goappdb
     JWT_SECRET=your-super-secret-jwt-key
     PORT=3000
     ```
   - These are loaded via `config.Load()` in `main.go`.

4. **Run Database Migrations**:

   - Apply migrations with `make migrate-up` (from `Makefile`).
   - This uses `golang-migrate` to run SQL migrations from `internal/database/migrations/`.

5. **Build the Application**:

   - Run `make build` (from `Makefile`) to compile the binary into `bin/main`.

6. **Run the Application**:
   - Execute `make run` (from `Makefile`) or `./bin/main` to start the server.
   - If using Docker Compose, run `docker-compose up` to start both the app and database.

## Development

### Adding a New Route

- Create a handler in `internal/handlers/` (e.g., `newfeature.go`).
- Register the route in `cmd/server/routes/setup.go` or the relevant routes file.

### Adding a New Model

- Define the struct in `internal/models/` (e.g., `newmodel.go`).
- Add a migration in `internal/database/migrations/` (e.g., `0002_add_new_table.up.sql`).
- Apply the migration with `make migrate-up`.

### Modifying SQL Queries

- Update or add queries in `sql/queries/` (e.g., `newquery.sql`).
- Regenerate SQLc code with `make sqlc-generate` (from `Makefile`).

### Running Tests

- Run tests with `make test` (from `Makefile`, executes `go test -v ./...`).
- Add tests in relevant directories (e.g., `internal/services/user_test.go`).

### Development Setup

- Use `docker-compose up -d` (from `docker-compose.yml`) to start the database.
- Apply migrations and generate code with `make dev-setup` if defined in `Makefile`.

## Deployment

- **Docker Compose**:

  - Build and run the app and database with `docker-compose up --build` (from `docker-compose.yml`).
  - This manages both the app and PostgreSQL containers.

- **Manual Deployment**:
  - Build the binary with `make build`.
  - Deploy `bin/main` to your server.
  - Ensure PostgreSQL is running and environment variables are set.
