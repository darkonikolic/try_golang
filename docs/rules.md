# Pravila za Golang projekat

## Tehnička pravila

### Docker-first pristup
- Sve komande se izvršavaju kroz Docker kontejnere
- Nema lokalnog pokretanja Go-a, go mod-a ili drugih alata
- Makefile komande koriste `docker-compose exec app`
- Multi-stage Docker build za production

### Go standardi
- Sav kod mora biti u skladu sa Go coding standards
- `gofmt` za formatiranje, `golint` za linting
- `go vet` za static analysis
- Jedan package po direktorijumu, jasna struktura

### Clean Architecture
- Jasna separacija slojeva: UI → Application → Domain ← Infrastructure
- Domain ne zna za Application/Infrastructure
- Infrastructure implementira Domain interfejse

## Arhitektura

### DDD + Hexagonal + CQRS
- **Domain**: Modeli, Value Objects, Servisi, Eventi, Repository interfejsi
- **Application**: Handleri, Komande, DTO-ovi, orkestracija domena
- **Infrastructure**: Repository implementacije, eksterni servisi, adapteri
- **UI**: HTTP handlers (Gin), request/response DTO-ovi

### Fizička separacija slojeva
```
cmd/
  api/
    main.go              # Entry point za API
internal/
  domain/
    model/               # Entiteti
    valueobject/         # Value Objects
    service/             # Domain servisi
    repository/           # Repository interfejsi
    event/               # Domain eventi
    exception/           # Domain exceptioni
  application/
    handler/             # Handleri, application interfejsi
    command/             # Komande
    dto/                 # Application DTO-ovi
  infrastructure/
    persistence/         # Repository implementacije
    messaging/           # Event adapteri
    database/            # Database konfiguracija
  ui/
    http/
      handler/           # HTTP handlers (Gin)
      middleware/        # Gin middleware
      dto/               # Request/Response DTO-ovi
pkg/
  logger/                # Shared logging
  config/                # Configuration management
  database/              # Database utilities
```

### CQRS (Command Query Responsibility Segregation)
- Razdvojeni read/write interfejsi
- `UserReadRepository` za čitanje
- `UserWriteRepository` za pisanje

## Proces

### **TDD - Test Driven Development**
- **UVEK** prvo piši test koji pada (RED)
- **UVEK** implementiraj minimalno da test prođe (GREEN)
- **UVEK** refaktoriši kod (REFACTOR)
- **UVEK** za sve: entities, value objects, services, handlers
- **UVEK** test coverage > 80%
- **UVEK** `go test` pre svakog commit-a

### Code Review
- Svaka izmena se review-uje
- Provera arhitekture i Go standarda
- Fokus na funkcionalnost

### Refaktorisanje
- Kada je potrebno, ne unapred
- Fokus na čist kod
- Pojednostavljivanje kompleksnosti

## Pragmatičnost

### KISS princip (Keep It Simple, Stupid)
- Jednostavnost pre kompleksnosti
- Fokus na funkcionalnost, ne na proces
- Ne over-engineer-uj za mali projekat

### **DDD i Clean Architecture za učenje**
- **DA** Value Objects - za učenje Go structs i methods
- **DA** Domain Entities - za učenje Go pointers i receivers
- **DA** Repository Pattern - za učenje Go interfaces
- **DA** Clean Architecture - za učenje Go package organization
- **DA** Error Handling - za učenje Go error patterns
- **NE** Over-engineer - fokus na učenje, ne na kompleksnost

### Fokus na rezultate
- Radi šta treba da se uradi
- Ne komplikuj nepotrebno
- Refaktorisi kada je potrebno

### **Pragmatična Clean Architecture**
- **Application layer**: SAMO ako je potreban (kompleksna orkestracija)
- **Domain service direktno u UI**: Ako je jednostavno
- **Interfejsi fleksibilno**: Koristi gde treba, ne forsiraj
- **Anticorruption layer**: SAMO ako je potreban (kompleksni external API)
- **Infrastructure layer**: UVEK (database, external services)
- **UI layer**: UVEK (HTTP handlers)

### **Docker-first sa pragmatičnim pristupom**
- **Development**: Local + Docker opcije
- **Production**: UVEK Docker
- **Firme zahtevaju**: Docker environment
- **Optimizacija**: Brz development workflow

## Mentorstvo i kritički pristup

### **Kritički pristup umesto potvrđivanja**
- **NIKAD** samo "da, u pravu si" - nema vrednosti za učenje
- **UVEK** analiziraj zahtev/ideju sa argumentima
- **UVEK** objasni zašto je nešto dobro/loše
- **UVEK** predloži best practice sa trade-offs
- **UVEK** daj konkretne primere i alternativne pristupe

### **Kako odgovarati na zahteve:**
- **Kada je u pravu**: "Da, u pravu si + dodatni kontekst + best practice"
- **Kada nije u pravu**: "Ne, jer... + argumenti + alternativni pristup"
- **Kada je kompleksno**: "Kompleksno je... + trade-offs + preporučeni pristup"

### **Mentorstvo principi:**
- **Kritički pristup**: Ne samo potvrđivanje, već argumentovana analiza
- **Argumentovano**: Sa konkretnim razlozima i best practices
- **Korisno**: Za učenje i razvoj, ne samo za potvrđivanje
- **Enterprise mindset**: Senior level pristup sa trade-offs

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

### **Odgovornost za odluke:**
- **Konačna odluka je na tebi** - ti odlučuješ šta ćeš implementirati
- **Moja obaveza** - da budem kritičan i direktan
- **Ako je loše** - moram da kažem da je loše
- **Ako je rizično** - moram da upozorim
- **Ako je neoptimalno** - moram da predložim bolje

## Dokumentacija i konzistentnost

### **Glavni dokument prioritet:**
- **PRIORITET 1**: `docs/progressive-learning-plan.md` - glavni plan (6 meseci)
- **PRIORITET 2**: `docs/step-X-*.md` - detaljni koraci
- **PRIORITET 3**: `docs/mentorship-plan.md` - overview
- **PRIORITET 4**: `docs/project-scope.md` - enterprise scope

### **Pravilo za usklađivanje:**
- **UVEK** prvo ažuriraj `daily-steps-plan.md`
- **UVEK** zatim ažuriraj sve ostale dokumente da se slažu
- **UVEK** obriši ili refaktoriši konfliktne delove
- **NIKAD** nemoj imati konfliktne informacije

### **Proces održavanja:**
- **Pre svake promene**: Proveri da li se slaže sa glavnim planom
- **Nakon svake promene**: Ažuriraj sve povezane dokumente
- **Redovno**: Review i cleanup konfliktnih delova

### **Cleanup pravila:**
- **Obriši duplikate**: Ako postoje dva dokumenta za istu stvar
- **Pojednostavi**: Ako je dokument previše kompleksan
- **Uskladi**: Ako se dokumenti ne slažu
- **Prioritizuj**: Glavni plan je `daily-steps-plan.md`

## Makefile komande

```bash
# Osnovne komande
make up          # Pokretanje
make down        # Gašenje
make bash        # Ulazak u kontejner

# Development
make go-mod-tidy     # go mod tidy
make test            # Testovi
make test-coverage   # Testovi sa coverage
make lint            # Linting
make fmt             # Formatiranje

# Database
make db-create       # Kreiraj bazu
make db-destroy      # Obriši bazu
make db-recreate     # Rekreiraj bazu
make db-migrate      # Migracije

# Build
make build           # Build aplikacije
make build-docker    # Build Docker image
make run             # Pokretanje aplikacije

# Kubernetes
make k8s-apply       # Apply Kubernetes manifests
make k8s-delete      # Delete Kubernetes resources
make k8s-logs        # View logs
```

## Zabranjena pravila

- ❌ `interface{}` ili `any` za domenske podatke
- ❌ Mešanje logike između slojeva
- ❌ Framework-specific kod u Domain sloju
- ❌ Direktno pokretanje alata van Docker-a
- ❌ Global variables
- ❌ Panic u production kodu 

## Organizacija koda i dokumentacije

### **Steps.md je smernica**
- [x] `docs/steps.md` je samo smernica za implementaciju
- [x] Kod se piše u source fajlovima, ne u dokumentaciji
- [x] Nakon završetka, steps.md se briše i kreira novi
- [x] Fokus na implementaciju, ne na dokumentaciju

### **Scripts i test fajlovi**
- [x] **`scripts/`** - svi pomoćni scriptovi
- [x] **`scripts/test/`** - test fajlovi za eksperimente
- [x] **`scripts/debug/`** - debug fajlovi
- [x] **`scripts/temp/`** - privremeni fajlovi

### **Dokumentacija organizacija**
- [x] **`docs/check_later.md`** - beleske za kasnije proveru
- [x] **`docs/experience/`** - iskustva i beleske
- [x] **`docs/learned/`** - naučene lekcije
- [x] **`docs/troubleshooting/`** - rešeni problemi

### **Beleske za kasnije**
- [x] Sve što kažeš "zabeleži za kasnije" ide u `docs/check_later.md`
- [x] Organizuj po kategorijama (Go, DDD, Architecture, etc.)
- [x] Dodaj datum i kontekst

### **Empty notes pravilo**
- [x] `docs/empty-notes.md` je za tvoje lične beleske
- [x] Ako sadrži duplikate sa drugim dokumentima - obriši
- [x] Ako ima bitne informacije - prebaci u odgovarajući dokument
- [x] Fokus na lične beleske, ne na duplikate 