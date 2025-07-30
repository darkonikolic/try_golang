# Pravila za Golang projekat - DDD + Clean Architecture

## **🎯 GLAVNI PRINCIPI**

### **DDD (Domain-Driven Design) - UVEK**
- **Entities**: UVEK sa ID i lifecycle (User, Product, Order)
- **Value Objects**: UVEK za validation (Email, UserID, Money)
- **Domain Services**: UVEK za business logic (UserService, OrderService)
- **Repository Pattern**: UVEK za data access (UserRepository, OrderRepository)
- **Domain Events**: UVEK za komunikaciju između agregata

### **Clean Architecture - UVEK**
- **Domain**: UVEK (entities, value objects, domain services, repository interfaces)
- **Application**: UVEK (orchestration, DTOs, handlers)
- **Infrastructure**: UVEK (database, external services, repository implementations)
- **UI**: UVEK (HTTP handlers, request/response DTOs)

## **🏗️ ARHITEKTURA - JASNE GRANICE**

### **Fizička struktura**
```
cmd/
  api/
    main.go              # Entry point - UVEK
internal/
  domain/
    user/
      entity/            # User entity - UVEK
      repository/        # UserRepository interface - UVEK
      service/           # UserService - UVEK
      errors/            # Domain errors - UVEK
    order/
      entity/            # Order entity - UVEK
      repository/        # OrderRepository interface - UVEK
      service/           # OrderService - UVEK
      errors/            # Domain errors - UVEK
  application/
    handler/             # HTTP handlers - UVEK
    dto/                 # Request/Response DTOs - UVEK
  infrastructure/
    database/            # Database connection - UVEK
    repository/          # Repository implementations - UVEK
    external/            # External services - UVEK
pkg/
  logger/                # Shared utilities - UVEK
  config/                # Configuration - UVEK
```

### **Dependency Flow - UVEK**
```
UI → Application → Domain ← Infrastructure
```

### **Zabranjene zavisnosti - NIKAD**
- ❌ Domain → Application
- ❌ Domain → Infrastructure  
- ❌ Domain → UI
- ❌ Application → Infrastructure
- ❌ UI → Infrastructure

## **💻 TEHNIČKA PRAVILA**

### **Docker-first pristup - UVEK**
- Sve komande se izvršavaju kroz Docker kontejnere
- Nema lokalnog pokretanja Go-a, go mod-a ili drugih alata
- Makefile komande koriste `docker-compose exec app`
- Multi-stage Docker build za production

### **Go standardi - UVEK**
- Sav kod mora biti u skladu sa Go coding standards
- `gofmt` za formatiranje, `golint` za linting
- `go vet` za static analysis
- Jedan package po direktorijumu, jasna struktura

### **TDD - Test Driven Development - UVEK**
- **UVEK** prvo piši test koji pada (RED)
- **UVEK** implementiraj minimalno da test prođe (GREEN)
- **UVEK** refaktoriši kod (REFACTOR)
- **UVEK** za sve: entities, value objects, services, handlers
- **UVEK** test coverage > 80%
- **UVEK** `go test` pre svakog commit-a

## **🔧 IMPLEMENTACIJA PRAVILA**

### **Domain Layer - UVEK**
```go
// UVEK - Entity sa ID
type User struct {
    ID        UserID
    Email     Email
    Name      string
    CreatedAt time.Time
    UpdatedAt time.Time
}

// UVEK - Value Object sa validation
type Email string
func (e Email) Validate() error { /* validation logic */ }

// UVEK - Repository interface
type UserRepository interface {
    Save(user *User) error
    FindByID(id UserID) (*User, error)
    FindByEmail(email Email) (*User, error)
}

// UVEK - Domain service
type UserService struct {
    repo UserRepository
}
```

### **Application Layer - UVEK**
```go
// UVEK - Handler sa DTOs
type CreateUserHandler struct {
    userService *domain.UserService
}

func (h *CreateUserHandler) Handle(req CreateUserRequest) (*CreateUserResponse, error) {
    // UVEK - orchestration logic
}
```

### **Infrastructure Layer - UVEK**
```go
// UVEK - Repository implementation
type PostgresUserRepository struct {
    db *sql.DB
}

func (r *PostgresUserRepository) Save(user *domain.User) error {
    // UVEK - database logic
}
```

## **🚫 ZABRANJENA PRAVILA**

### **NIKAD u Domain layer:**
- ❌ `interface{}` ili `any` za domenske podatke
- ❌ Framework-specific kod (Gin, GORM)
- ❌ Database connection
- ❌ HTTP handlers
- ❌ External service calls

### **NIKAD u Application layer:**
- ❌ Direct database access
- ❌ Framework-specific logic
- ❌ Business logic (to ide u Domain)

### **NIKAD u Infrastructure layer:**
- ❌ Business logic
- ❌ Domain entities modification
- ❌ HTTP handlers

### **NIKAD globalno:**
- ❌ Global variables
- ❌ Panic u production kodu
- ❌ Direktno pokretanje alata van Docker-a

## **📋 PROCES**

### **Code Review - UVEK**
- Svaka izmena se review-uje
- Provera arhitekture i Go standarda
- Fokus na funkcionalnost

### **Refaktorisanje - UVEK**
- Kada je potrebno, ne unapred
- Fokus na čist kod
- Pojednostavljivanje kompleksnosti

## **🎓 MENTORSTVO I KRITIČKI PRISTUP**

### **Kritički pristup umesto potvrđivanja - UVEK**
- **NIKAD** samo "da, u pravu si" - nema vrednosti za učenje
- **UVEK** analiziraj zahtev/ideju sa argumentima
- **UVEK** objasni zašto je nešto dobro/loše
- **UVEK** predloži best practice sa trade-offs
- **UVEK** daj konkretne primere i alternativne pristupe

### **Kako odgovarati na zahteve:**
- **Kada je u pravu**: "Da, u pravu si + dodatni kontekst + best practice"
- **Kada nije u pravu**: "Ne, jer... + argumenti + alternativni pristup"
- **Kada je kompleksno**: "Kompleksno je... + trade-offs + preporučeni pristup"

### **Nikad samo potvrđivanje:**
- ❌ "Da, u pravu si"
- ❌ "To je dobra ideja"
- ❌ "Može tako"
- ✅ "Da, ali...", "Ne, jer...", "Kompleksno je..."
- ✅ Argumenti + best practices + trade-offs

### **Kritička komunikacija - obavezna:**
- **AKO JE LOŠE - MORAŠ DA KAŽEŠ DA JE LOŠE**
- **AKO JE RIZIČNO - MORAŠ DA UPOZORIŠ**
- **AKO JE NEOPTIMALNO - MORAŠ DA PREDLOŽIŠ BOLJE**
- **NIKAD** ne ulepšavaj stvarnost
- **NIKAD** ne skrivaj probleme
- **UVEK** budi direktan i kritičan

## **📚 DOKUMENTACIJA**

### **Glavni dokument prioritet:**
- **PRIORITET 1**: `docs/days/day_X.md` - trenutni dan
- **PRIORITET 2**: `docs/progressive-learning-plan.md` - glavni plan
- **PRIORITET 3**: `docs/project-scope.md` - enterprise scope

### **Pravilo za usklađivanje:**
- **UVEK** prvo ažuriraj trenutni dan
- **UVEK** zatim ažuriraj sve ostale dokumente da se slažu
- **UVEK** obriši ili refaktoriši konfliktne delove
- **NIKAD** nemoj imati konfliktne informacije

## **🔧 MAKEFILE KOMANDE**

```bash
# Osnovne komande
make dev          # Development server
make dev-stop     # Stop development
make bash         # Ulazak u kontejner

# Development
make test         # Testovi
make lint         # Linting
make fmt          # Formatiranje
make check        # Sve provere

# Database
make db-up        # Start database
make db-down      # Stop database

# Build
make build        # Build aplikacije
make clean        # Clean sve
```

## **💬 KOMUNIKACIJA**

### **Pravilo za komunikaciju sa mentorom (AI asistentom)**
- **Ako korisnikov zahtev nije potpuno jasan, asistent MORA da traži pojašnjenje pre nego što odgovori.**
- **NIKADA ne nagađaj ili daj generički odgovor ako nisi siguran šta korisnik traži.**
- **Direktno pitaj: 'Možeš li precizirati šta tačno želiš?' ili 'Nisam siguran da li sam razumeo, možeš li pojasniti?'**
- **Cilj: Uvek dati tačan, koncizan i kontekstualno relevantan odgovor.**

### **Pravilo za predupređivanje frustracije**
- **AKO korisnik postavlja pitanje koje može biti interpretirano na više načina, asistent MORA da pita za pojašnjenje.**
- **AKO korisnik koristi nejasne termine ili reference, asistent MORA da traži preciziranje.**
- **AKO korisnik traži "fix" bez konteksta, asistent MORA da pita "Šta tačno treba da popravim?"**
- **AKO korisnik postavlja pitanje koje može biti tehnički ili konceptualno, asistent MORA da pita "Da li te zanima tehnička implementacija ili konceptualno objašnjenje?"**
- **Cilj: Sprečiti frustraciju pre nego što nastane.**

### **Terminologija ispravke (obavezno)**
- **OBVEZNO** ispraviti pogrešnu terminologiju (PHP/Java → Go)
- **OBVEZNO** ispraviti: "referenca" → "pointer", "exception" → "error"
- **OBVEZNO** ispraviti: "class" → "struct", "inheritance" → "composition"
- **GREŠKE U KUCANJU** su OK (typos)
- **TERMINOLOGIJA** nije OK - mora biti Go-specific 