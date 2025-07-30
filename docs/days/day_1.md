# **📅 Dan 1: Docker-First Go Development Setup**

## **🎯 Cilj Dana**
Postaviti Docker development environment za Go aplikaciju i kreirati čistu projektnu strukturu koja će podržati Clean Architecture - sve unutar Docker containera.

## **📋 Šta Radimo Danas**

### **1. Docker Development Environment (2 sata)** ✅
- [x] **Dockerfile za Go development**: Multi-stage build
- [x] **docker-compose.yml**: Development environment setup
- [x] **Volume mounting**: Hot reload za development
- [x] **Docker development workflow**: Real-world scenario

### **2. Go Module Setup u Docker (2 sata)** ✅
- [x] **Go modul unutar containera**: `go mod init helloworld`
- [x] **Dependency management**: `go mod tidy` u Docker
- [ ] **Go tools u container**: `go fmt`, `go vet`, `go test`
- [ ] **Container-aware development**: Docker kao development platform

### **3. Project Structure u Docker (2 sata)** ✅
- [x] **Clean Architecture folder structure**:
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

### **4. Development Workflow (1 sat)** ✅
- [x] **Docker Compose commands**: `docker-compose up`, `docker-compose down`
- [x] **Hot reload setup**: Air/Realize za development
- [ ] **Debug u Docker**: Delve debugger u container
- [x] **Development vs production**: Environment separation

## **🏭 Production vs Development Setup**

### **Development Environment (`docker-compose.yml`)**
```yaml
# Development - sve u jednom fajlu
services:
  app:
    build:
      target: development  # Hot reload, debug tools
    volumes:
      - .:/app            # Live code changes
    environment:
      - GO_ENV=development
  postgres:
    image: postgres:15-alpine  # Lokalna baza za development
```

### **Production Environment (`docker-compose.prod.yml`)**
```yaml
# Production - samo aplikacija
services:
  app:
    build:
      target: production   # Optimizovan build
    environment:
      - GO_ENV=production
      # Database connection preko env vars
      # - DATABASE_URL=postgres://user:pass@host:port/db
    healthcheck:
      test: ["CMD", "wget", "--spider", "http://localhost:8080/health"]
    deploy:
      resources:
        limits:
          memory: 512M
          cpus: '0.5'
```

### **Ključne Razlike**

| Aspekt | Development | Production |
|--------|-------------|------------|
| **Database** | Lokalna PostgreSQL | Cloud database (AWS RDS, GCP) |
| **Hot Reload** | ✅ Air tool | ❌ Optimizovan build |
| **Volume Mounts** | ✅ Live code changes | ❌ Static binary |
| **Resource Limits** | ❌ Unlimited | ✅ Memory/CPU limits |
| **Health Checks** | ❌ Basic | ✅ Comprehensive |
| **Environment** | Development tools | Production optimized |

### **Zašto Production Nema Database?**

1. **Security**: Database credentials ne treba da budu u Docker Compose
2. **Scalability**: Database se skalira odvojeno od aplikacije
3. **Backup**: Cloud databases imaju automatske backup-ove
4. **Monitoring**: Cloud databases imaju built-in monitoring
5. **High Availability**: Cloud databases imaju failover

### **Production Deployment Workflow**

```bash
# 1. Build production image
make prod-build

# 2. Set environment variables
export DATABASE_URL="postgres://user:pass@rds.amazonaws.com:5432/db"

# 3. Deploy to production
make prod

# 4. Monitor
make prod-logs
```

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
- [x] Kreirati multi-stage Dockerfile za Go
- [x] Postaviti Docker Compose development environment
- [x] Organizovati kod po Clean Architecture principima
- [x] Implementirati hot reload u Docker

### **Konceptualne Veštine**
- [x] Razumeti Docker-first development
- [x] Razumeti container vs host development
- [x] Razumeti Clean Architecture u containers
- [x] Razumeti development vs production builds

### **Praktične Veštine**
- [x] Koristiti Docker CLI komande
- [x] Raditi sa Docker Compose
- [x] Debug u container environment
- [x] Organizovati kod za container deployment

## **🎯 Deliverables**

### **Docker Files**
- [x] `Dockerfile` (development + production stages)
- [x] `docker-compose.yml` (development environment)
- [x] `.dockerignore` (optimize builds)
- [x] `docker-compose.prod.yml` (production setup)

### **Go Application**
- [x] `go.mod` i `go.sum` (unutar Docker)
- [x] Osnovna projektna struktura
- [x] `main.go` sa "Hello World"
- [x] Hot reload setup sa Air

### **Dokumentacija**
- [x] README.md sa Docker setup instrukcijama
- [x] Development workflow dokumentacija
- [x] Docker commands cheatsheet
- [x] Environment setup guide

### **Verifikacija**
- [x] `docker-compose up` radi
- [x] Hot reload funkcioniše
- [x] `docker build` uspešan
- [x] Kod je formatiran sa `go fmt`

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

**🔄 Dan 1 - U toku...** 🐳 