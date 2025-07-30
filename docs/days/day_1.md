# **📅 Dan 1: Docker-First Go Development Setup**

## **🎯 Cilj Dana**
Postaviti Docker development environment za Go aplikaciju i kreirati čistu projektnu strukturu koja će podržati Clean Architecture - sve unutar Docker containera.

## **📋 Šta Radimo Danas**

### **1. Docker Development Environment (2 sata)**
- **Dockerfile za Go development**: Multi-stage build
- **docker-compose.yml**: Development environment setup
- **Volume mounting**: Hot reload za development
- **Docker development workflow**: Real-world scenario

### **2. Go Module Setup u Docker (2 sata)**
- **Go modul unutar containera**: `go mod init helloworld`
- **Dependency management**: `go mod tidy` u Docker
- **Go tools u container**: `go fmt`, `go vet`, `go test`
- **Container-aware development**: Docker kao development platform

### **3. Project Structure u Docker (2 sata)**
- **Clean Architecture folder structure**:
  ```
  helloworld/
  ├── cmd/           # Application entry points
  ├── internal/      # Private application code
  │   ├── domain/    # Business entities and rules
  │   ├── application/ # Use cases and application services
  │   ├── infrastructure/ # External concerns (DB, HTTP)
  │   └── interfaces/ # UI layer (HTTP handlers)
  ├── pkg/           # Public libraries
  ├── docs/          # Documentation
  ├── scripts/       # Build and deployment scripts
  ├── test/          # Integration tests
  ├── Dockerfile     # Multi-stage build
  └── docker-compose.yml # Development environment
  ```

### **4. Development Workflow (1 sat)**
- **Docker Compose commands**: `docker-compose up`, `docker-compose down`
- **Hot reload setup**: Air/Realize za development
- **Debug u Docker**: Delve debugger u container
- **Development vs production**: Environment separation

## **🎓 Koncepti koje Učimo**

### **Docker Development**
- **Multi-stage Dockerfile**: Development vs production builds
- **Volume mounting**: Code changes without rebuild
- **Docker Compose**: Multi-service development
- **Container networking**: Service communication

### **Go Module System u Docker**
- **go.mod**: Module definition in container
- **go.sum**: Dependency verification
- **Module path**: Import path structure
- **Version management**: Semantic versioning

### **Project Organization**
- **Clean Architecture**: Layer separation in containers
- **Dependency direction**: Inward dependencies only
- **Package naming**: Go conventions
- **Docker-aware structure**: Container-friendly organization

### **Development Workflow**
- **Docker-first development**: Container as primary environment
- **Hot reload**: Real-time code changes
- **Debug in container**: Remote debugging
- **Production-like development**: Real-world scenario

## **🔧 Alati koje Koristimo**

### **Docker Tools**
- **Docker CLI**: `docker build`, `docker run`, `docker-compose`
- **Multi-stage builds**: Development and production images
- **Volume mounting**: Code synchronization
- **Container networking**: Service communication

### **Go Tools u Docker**
- **Go compiler**: `go build`, `go run` in container
- **Go modules**: `go mod tidy`, `go mod download`
- **Go tools**: `go fmt`, `go vet`, `go test`
- **Hot reload**: Air/Realize for development

### **Development Tools**
- **Docker Compose**: Multi-service orchestration
- **Volume mounts**: Code changes without rebuild
- **Environment variables**: Configuration management
- **Health checks**: Container monitoring

## **📚 Learning Objectives**

### **Tehničke Veštine**
- [ ] Kreirati multi-stage Dockerfile za Go
- [ ] Postaviti Docker Compose development environment
- [ ] Organizovati kod po Clean Architecture principima
- [ ] Implementirati hot reload u Docker

### **Konceptualne Veštine**
- [ ] Razumeti Docker-first development
- [ ] Razumeti container vs host development
- [ ] Razumeti Clean Architecture u containers
- [ ] Razumeti development vs production builds

### **Praktične Veštine**
- [ ] Koristiti Docker CLI komande
- [ ] Raditi sa Docker Compose
- [ ] Debug u container environment
- [ ] Organizovati kod za container deployment

## **🎯 Deliverables**

### **Docker Files**
- [ ] `Dockerfile` (development + production stages)
- [ ] `docker-compose.yml` (development environment)
- [ ] `.dockerignore` (optimize builds)
- [ ] `docker-compose.prod.yml` (production setup)

### **Go Application**
- [ ] `go.mod` i `go.sum` (unutar Docker)
- [ ] Osnovna projektna struktura
- [ ] `main.go` sa "Hello World"
- [ ] Hot reload setup sa Air

### **Dokumentacija**
- [ ] README.md sa Docker setup instrukcijama
- [ ] Development workflow dokumentacija
- [ ] Docker commands cheatsheet
- [ ] Environment setup guide

### **Verifikacija**
- [ ] `docker-compose up` radi
- [ ] Hot reload funkcioniše
- [ ] `docker build` uspešan
- [ ] Kod je formatiran sa `go fmt`

## **🚀 Next Steps**

### **Sutra (Dan 2)**
- User Entity kreiranje u Docker
- Go structs i methods u container
- Domain layer setup
- Database connection u Docker

### **Ove Nedelje**
- **Dan 3-4**: User Entity + Go Structs (Docker)
- **Dan 5-7**: HTTP Handlers + Gin Framework (Docker)

## **💡 Tips & Tricks**

### **Docker Best Practices**
- **Multi-stage builds**: Separate development and production
- **Volume mounting**: Code changes without rebuild
- **Non-root user**: Security best practice
- **Health checks**: Container monitoring
- **Environment variables**: Configuration management

### **Go in Docker Best Practices**
- **Alpine Linux**: Smaller base images
- **Go modules**: Proper dependency management
- **Build caching**: Optimize build times
- **Static linking**: Self-contained binaries

### **Development Workflow**
- **Docker Compose**: Multi-service development
- **Hot reload**: Real-time development
- **Debug in container**: Remote debugging
- **Production-like**: Real-world scenario

### **Clean Architecture in Docker**
- **Layer separation**: Clear boundaries in containers
- **Dependency injection**: Container-aware DI
- **Interface segregation**: Small, focused interfaces
- **Single responsibility**: Each container has one purpose

---

**Da li si spreman da počnemo sa Docker-First Dan 1?** 🐳 