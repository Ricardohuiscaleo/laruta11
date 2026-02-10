# La Ruta 11 - Project Structure

## Monorepo Organization

This is a monorepo containing 3 independent applications deployed separately:

```
laruta11/
├── app/              # Customer menu app (app.laruta11.cl)
├── caja/             # Admin/cashier system (caja.laruta11.cl)
├── landing/          # Public landing page (laruta11.cl)
├── docs/             # Project documentation
└── legacy-apis/      # Deprecated PHP APIs (not tracked in git)
```

## Application Structure

### /app - Customer Menu Application
```
app/
├── src/
│   ├── components/   # React components (MenuApp, ProductCard, Cart, etc.)
│   ├── pages/        # Astro pages (routes)
│   ├── layouts/      # Page layouts
│   ├── hooks/        # React hooks
│   ├── icons/        # Icon components
│   ├── utils/        # Utility functions
│   └── mock/         # Mock data for development
├── api/              # PHP backend APIs
│   ├── orders/       # Order management endpoints
│   ├── users/        # User authentication and management
│   ├── coupons/      # Discount code system
│   ├── notifications/# Push notification system
│   ├── tuu/          # TUU payment integration
│   ├── tracker/      # Analytics tracking
│   └── *.php         # Core API endpoints
├── public/           # Static assets
├── sql/              # Database schema files
└── game-isolated/    # Embedded mini-games
```

### /caja - Admin/Cashier System
```
caja/
├── api/              # PHP backend APIs (being migrated to Go)
├── api-go/           # Go API (NEW - Dockerfile deploy)
│   ├── main.go
│   ├── handlers.go
│   ├── Dockerfile
│   ├── go.mod
│   └── go.sum
├── src/
│   ├── components/   # React components (ComprasApp, InventarioApp, etc.)
│   ├── pages/        # Admin dashboard pages
│   └── layouts/      # Admin layouts
├── docs/             # Technical documentation
├── sql/              # Database migrations
└── nixpacks.toml     # Astro frontend deploy config
```

### /landing - Public Website
```
landing/
├── api-go/           # Go API for S3 (Dockerfile deploy)
│   ├── main.go
│   ├── Dockerfile
│   ├── go.mod
│   └── go.sum
├── src/
│   ├── components/   # Landing page components
│   ├── pages/        # Public pages
│   └── layouts/      # Landing layouts
└── vendor/           # PHP dependencies (Composer)
```

## Core Components

### Frontend Architecture
- **Framework**: Astro with React integration
- **UI Components**: React functional components with hooks
- **Styling**: Tailwind CSS
- **State Management**: React hooks (useState, useEffect)
- **PWA**: Service workers for offline capability

### Backend Architecture
- **API Layer**: PHP REST APIs
- **Database**: MySQL with direct PDO connections
- **File Storage**: AWS S3 for product images
- **Authentication**: Session-based with JWT tokens
- **Payment Integration**: TUU payment gateway

### Database Schema
- **productos**: Product catalog with pricing and availability
- **ingredientes**: Ingredient inventory with stock levels
- **recetas**: Recipe definitions linking products to ingredients
- **ventas**: Sales transactions and order history
- **usuarios**: User accounts and authentication
- **combos**: Combo deal definitions
- **notifications**: Push notification queue
- **wallet_transactions**: Cashback and loyalty points

## Architectural Patterns

### API Design
- RESTful endpoints with JSON responses
- Consistent error handling with status codes
- CORS enabled for cross-origin requests
- Session validation middleware

### Data Flow
1. Frontend makes API request
2. PHP endpoint validates session/auth
3. Database query executed via PDO
4. Response formatted as JSON
5. Frontend updates UI state

### Inventory Management
- **Recipe-based**: Products linked to ingredients via recipes
- **Automatic calculation**: Stock levels computed from ingredient availability
- **Purchase recommendations**: AI-powered suggestions based on sales patterns
- **Audit trail**: All inventory changes logged

### Deployment Strategy
- Each app deployed independently on Easypanel
- **Astro frontends**: Nixpacks with `nixpacks.toml`
- **Go APIs**: Dockerfile (NOT Nixpacks - causes go.sum issues)
- Shared database across applications
- Environment-specific configuration via .env files
- Static assets served from public directories

### Go API Deployment (CRITICAL)
**Always use Dockerfile for Go APIs:**
```dockerfile
FROM golang:1.21-alpine
WORKDIR /app
COPY go.mod .
RUN go mod download && go mod tidy
COPY . .
RUN go build -o api .
EXPOSE 3001
CMD ["./api"]
```

**Steps:**
1. Generate `go.sum` locally: `go mod tidy`
2. Commit `go.mod` and `go.sum`
3. Easypanel: Build Method = Dockerfile, Build Path = `path/to/api-go`
4. Set environment variables
5. Deploy
