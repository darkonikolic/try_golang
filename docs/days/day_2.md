# **📅 Dan 2: User Entity i Domain Layer**

## **🎯 Cilj Dana**
Implementirati User Entity sa Go structs i methods, postaviti Domain layer, i konektovati aplikaciju na PostgreSQL database - sve unutar Docker environment-a.

## **📋 Šta Radimo Danas**

### **1. User Entity Design (2 sata)** ✅
- [x] **Go structs za User**: ID, Email, Name, CreatedAt, UpdatedAt
- [x] **Value Objects**: Email validation, UserID type
- [x] **Methods na User struct**: Validate(), Update(), IsActive()
- [x] **TDD pristup**: Prvo test, pa implementacija

### **2. Domain Layer Setup (2 sata)**
- [ ] **Domain folder structure**: `internal/domain/user/`
- [ ] **Repository interface**: `UserRepository`
- [ ] **Domain services**: `UserService`
- [ ] **Domain exceptions**: `UserNotFoundError`, `InvalidEmailError`

### **3. Database Connection (2 sata)**
- [ ] **PostgreSQL setup**: Konekcija u Docker
- [ ] **Database schema**: Users tabela
- [ ] **Connection pooling**: U Docker environment
- [ ] **Environment variables**: Database config

### **4. Infrastructure Layer (1 sat)**
- [ ] **Repository implementation**: PostgreSQL UserRepository
- [ ] **Database migrations**: SQL scripts
- [ ] **Connection management**: Open/Close, transactions

## **🎓 Koncepti koje Učimo**

### **Go Structs i Methods**
- **Struct definition**: `type User struct`
- **Method receivers**: Pointer vs Value receivers
- **Embedding**: Composition over inheritance
- **Validation**: Custom validation methods

### **Domain-Driven Design u Go**
- **Entities**: User sa ID i lifecycle
- **Value Objects**: Email, UserID kao types
- **Repository Pattern**: Interface za data access
- **Domain Services**: Business logic

### **Database u Docker**
- **PostgreSQL**: Relational database
- **Connection pooling**: Efficient connections
- **Migrations**: Schema versioning
- **Transactions**: ACID properties

### **TDD u Go**
- **Test structure**: `*_test.go` files
- **Table-driven tests**: Multiple test cases
- **Mocking**: Interface-based testing
- **Test coverage**: Go test tools

## **🔧 Alati koje Koristimo**

### **Go Tools**
- **Go structs**: Custom types and methods
- **Go interfaces**: Repository pattern
- **Go testing**: `go test`, `go test -v`
- **Go vet**: Static analysis

### **Database Tools**
- **PostgreSQL**: Relational database
- **SQL**: Schema creation, queries
- **Connection pooling**: Efficient DB access
- **Migrations**: Schema versioning

### **Docker Tools**
- **Docker Compose**: Multi-service setup
- **Volume mounting**: Database persistence
- **Environment variables**: Configuration
- **Health checks**: Database monitoring

## **📚 Learning Objectives**

### **Tehničke Veštine**
- [ ] Kreirati Go structs sa methods
- [ ] Implementirati Repository pattern
- [ ] Konektovati na PostgreSQL u Docker
- [ ] Pisati TDD tests za User entity

### **Konceptualne Veštine**
- [ ] Razumeti DDD u Go
- [ ] Razumeti Value Objects vs Entities
- [ ] Razumeti Repository pattern
- [ ] Razumeti database transactions

### **Praktične Veštine**
- [ ] Koristiti Go testing framework
- [ ] Raditi sa PostgreSQL u Docker
- [ ] Implementirati proper error handling
- [ ] Organizovati kod po DDD principima

## **🎯 Deliverables**

### **Go Code**
- [ ] `internal/domain/user/user.go` (User struct)
- [ ] `internal/domain/user/user_repository.go` (Repository interface)
- [ ] `internal/domain/user/service.go` (Domain service)
- [ ] `internal/domain/user/errors.go` (Domain exceptions)

### **Database**
- [ ] PostgreSQL connection u Docker
- [ ] Users table schema
- [ ] Database migrations
- [ ] Connection pooling setup

### **Tests**
- [ ] `internal/domain/user/user_test.go`
- [ ] `internal/domain/user/user_repository_test.go`
- [ ] `internal/domain/user/service_test.go`
- [ ] Test coverage > 80%

### **Infrastructure**
- [ ] `internal/infrastructure/persistence/user_repository.go`
- [ ] `internal/infrastructure/database/connection.go`
- [ ] `scripts/db/migrations/001_create_users.sql`

## **🚀 Next Steps**

### **Sutra (Dan 3)**
- HTTP handlers za User CRUD
- Gin framework integration
- Request/Response DTOs
- API endpoints testing

### **Ove Nedelje**
- **Dan 4**: User CRUD API
- **Dan 5**: Authentication/Authorization
- **Dan 6-7**: Advanced features

## **💡 Tips & Tricks**

### **Go Best Practices**
- **Pointer receivers**: Za methods koji menjaju struct
- **Value receivers**: Za methods koji samo čitaju
- **Interface segregation**: Small, focused interfaces
- **Error handling**: Return errors, don't panic

### **DDD in Go**
- **Entities**: Have identity, can change
- **Value Objects**: Immutable, no identity
- **Repository**: Interface for data access
- **Domain Services**: Business logic

### **Database Best Practices**
- **Connection pooling**: Reuse connections
- **Transactions**: ACID properties
- **Migrations**: Version control for schema
- **Indexes**: Performance optimization

### **TDD Workflow**
- **Red**: Write failing test
- **Green**: Make test pass
- **Refactor**: Clean up code
- **Repeat**: For each feature

---

**🔄 Dan 2 - Spreman za početak!** 🐳 