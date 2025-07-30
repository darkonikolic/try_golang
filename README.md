# Try Golang - Go Project with DDD Architecture

A Go project demonstrating Domain-Driven Design (DDD) architecture with Docker Compose setup for development and production.

## 🏗️ Project Structure

```
try_golang/
├── cmd/api/           # Application entry point
├── internal/          # Private application code
│   ├── application/   # Application services
│   ├── domain/        # Domain models and business logic
│   ├── infrastructure/# External interfaces (DB, APIs)
│   └── interfaces/    # HTTP handlers, CLI
├── pkg/              # Public packages
├── docs/             # Documentation
├── scripts/          # Build and deployment scripts
├── test/             # Integration tests
├── Dockerfile        # Multi-stage Docker build
├── docker-compose.yml # Development environment
├── Makefile          # Build automation
└── .air.toml         # Hot reload configuration
```

## 🚀 Quick Start

### Prerequisites

- Docker & Docker Compose
- Make (optional, for convenience)

### Development

```bash
# Start development server with hot reload
make dev

# Start in background
make dev-d

# Stop development server
make dev-stop

# View logs
make logs
```

### Production

```bash
# Build production image
make build-prod

# Run production container
docker run -p 8080:8080 helloworld_app
```

## 📋 Available Commands

### Development Commands

| Command | Description |
|---------|-------------|
| `make dev` | Start development server with hot reload |
| `make dev-d` | Start development server in background |
| `make dev-stop` | Stop development server |
| `make dev-restart` | Restart development server |
| `make logs` | Show development logs |
| `make logs-app` | Show application logs only |

### Testing Commands

| Command | Description |
|---------|-------------|
| `make test` | Run tests in Docker Compose |
| `make test-verbose` | Run tests with verbose output |
| `make test-coverage` | Run tests with coverage report |

### Build Commands

| Command | Description |
|---------|-------------|
| `make build` | Build all Docker Compose services |
| `make build-app` | Build application service only |
| `make build-prod` | Build production image |
| `make build-multi` | Build multi-platform binaries |

### Database Commands

| Command | Description |
|---------|-------------|
| `make db-up` | Start database only |
| `make db-down` | Stop database only |
| `make db-logs` | Show database logs |
| `make db-reset` | Reset database (remove volume) |

### Code Quality Commands

| Command | Description |
|---------|-------------|
| `make fmt` | Format Go code |
| `make vet` | Vet Go code |
| `make lint` | Run linter |
| `make check` | Run all code quality checks |

### Dependency Commands

| Command | Description |
|---------|-------------|
| `make deps` | Download dependencies |
| `make deps-update` | Update dependencies |
| `make deps-vendor` | Vendor dependencies |

### Utility Commands

| Command | Description |
|---------|-------------|
| `make shell` | Open shell in application container |
| `make bash` | Open bash shell in application container |
| `make shell-db` | Open shell in database container |
| `make ps` | Show running containers |
| `make status` | Show status of all services |
| `make install-tools` | Install development tools |

### Cleanup Commands

| Command | Description |
|---------|-------------|
| `make clean` | Clean build artifacts and Docker Compose resources |
| `make clean-compose` | Clean Docker Compose resources |
| `make clean-images` | Clean Docker images |
| `make clean-all` | Clean everything (containers, images, volumes) |

## 🐳 Docker Architecture

### Multi-Stage Dockerfile

The project uses a multi-stage Dockerfile with three stages:

1. **Development Stage** (`golang:1.24-alpine`)
   - Includes Air for hot reload
   - Development tools and dependencies
   - Source code mounted via volume

2. **Builder Stage** (`golang:1.24-alpine`)
   - Compiles the application
   - Creates statically linked binary
   - Optimized for Linux deployment

3. **Production Stage** (`alpine:latest`)
   - Minimal runtime image (~10MB)
   - Contains only the compiled binary
   - Non-root user for security

### Docker Compose Services

- **app**: Go application with hot reload
- **postgres**: PostgreSQL database (optional)

## 🔧 Configuration

### Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | `8080` | Application port |
| `GO_ENV` | `development` | Go environment |

### Database Configuration

| Variable | Default | Description |
|----------|---------|-------------|
| `POSTGRES_DB` | `helloworld` | Database name |
| `POSTGRES_USER` | `postgres` | Database user |
| `POSTGRES_PASSWORD` | `postgres` | Database password |

## 📊 API Endpoints

### Health Check
```
GET /health
```
Returns application status and version information.

### Root Endpoint
```
GET /
```
Returns welcome message and API documentation.

### API v1
```
GET /api/v1/
```
Returns available API endpoints.

## 🛠️ Development Workflow

### 1. Start Development Environment
```bash
make dev
```

### 2. Make Code Changes
Edit files in your IDE - changes are automatically detected and the server restarts.

### 3. Run Tests
```bash
make test
```

### 4. Check Code Quality
```bash
make check
```

### 5. Build for Production
```bash
make build-prod
```

## 🔍 Troubleshooting

### Port Already in Use
```bash
# Stop all containers
make clean-containers

# Or stop specific service
make dev-stop
```

### Database Issues
```bash
# Reset database
make db-reset

# Check database logs
make db-logs
```

### Build Issues
```bash
# Clean everything and rebuild
make clean-all
make build
```

### Hot Reload Not Working
```bash
# Restart development server
make dev-restart

# Check Air configuration
cat .air.toml
```

## 📚 Learning Resources

- [Go Documentation](https://golang.org/doc/)
- [Domain-Driven Design](https://martinfowler.com/bliki/DomainDrivenDesign.html)
- [Docker Compose](https://docs.docker.com/compose/)
- [Air - Live Reload](https://github.com/air-verse/air)

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Run tests: `make test`
5. Check code quality: `make check`
6. Submit a pull request

## 📄 License

This project is licensed under the MIT License.

---

**Happy Coding! 🚀** 