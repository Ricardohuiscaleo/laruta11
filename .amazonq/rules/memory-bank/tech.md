# La Ruta 11 - Technology Stack

## Programming Languages

### Frontend
- **JavaScript/JSX**: React components and client-side logic
- **TypeScript**: Type definitions (env.d.ts)
- **HTML/CSS**: Markup and styling via Tailwind

### Backend
- **Go**: Primary backend language for APIs (replacing PHP)
- **PHP 7.4+**: Legacy backend (being migrated to Go)
- **SQL**: Database queries and schema definitions

## Frameworks & Libraries

### Frontend Stack
- **Astro 4.11.5**: Static site generator with partial hydration
- **React 18.3.1**: UI component library
- **React DOM 18.3.1**: React rendering
- **Tailwind CSS 3.4.17**: Utility-first CSS framework
- **@astrojs/react**: Astro-React integration
- **@astrojs/tailwind**: Astro-Tailwind integration
- **@astrojs/node**: Node.js adapter for SSR

### UI Components & Icons
- **lucide-react 0.400.0**: Icon library
- **react-icons 5.5.0**: Additional icon set
- **recharts 3.1.2**: Data visualization charts

### Backend Stack
- **Go 1.21+**: Primary API language
  - `github.com/gin-gonic/gin`: HTTP framework
  - `github.com/go-sql-driver/mysql`: MySQL driver
  - `github.com/aws/aws-sdk-go-v2`: AWS S3 integration
- **PHP 7.4+** (Legacy): Being migrated to Go
- **MySQL**: Relational database via PDO/sql.DB
- **AWS S3**: Image storage with official SDK

## Build System

### Package Management
- **npm**: Node.js package manager
- **package.json**: Dependency definitions in each app
- **Composer**: PHP dependency manager (landing app)

### Build Commands
```bash
# Development
npm run dev        # Start dev server with hot reload
npm run start      # Alias for dev

# Production
npm run build      # Build for production
npm run preview    # Preview production build
```

### Build Configuration
- **astro.config.mjs**: Astro configuration per app
- **tailwind.config.mjs**: Tailwind customization
- **nixpacks.toml**: Deployment configuration for Easypanel

## Database

### Technology
- **MySQL 9.6**: Relational database (migrated from Hostinger to VPS)
- **Host**: `websites_mysql-laruta11` (internal) or `76.13.126.63:3306` (external)
- **Database**: `laruta11`
- **User**: `laruta11_user`
- **Connection**: Direct MySQL connections via Go `database/sql` or PHP PDO

### Database Management Tools
- **Beekeeper Studio** (recommended): Modern, free, native M1 support
  - Download: https://www.beekeeperstudio.io/
  - Fast, clean UI, multiple tabs, dark mode
- **TablePlus** (premium): Fastest option, $89 lifetime
  - Download: https://tableplus.com/
  - Native M1, best performance
- **Adminer** (web): Lightweight phpMyAdmin alternative
  - Single PHP file, can deploy on VPS
- **phpMyAdmin** (legacy): Avoid, slow and outdated

### Key Tables
- productos, ingredientes, recetas
- ventas, usuarios, orders
- combos, notifications, wallet_transactions
- tuu_orders, reviews, analytics

### Connection Management
- **config.php**: Database credentials and connection setup
- **load-env.php**: Environment variable loader
- Direct PDO connections in API endpoints

## External Services

### Payment Processing
- **TUU**: Chilean payment gateway integration
- Webhook callbacks for payment confirmation
- Transfer payment verification

### Cloud Storage
- **AWS S3**: Product image storage
- **S3Manager.php**: Custom S3 upload/download wrapper

### AI/Analytics
- **Google Gemini API**: AI-powered business analysis
- Custom analytics tracking system

### Live Streaming
- **YouTube Live API**: Contest streaming integration
- Real-time viewer tracking

## Development Tools

### Version Control
- **Git**: Source control
- **.gitignore**: Excludes node_modules, .env, build artifacts

### Environment Management
- **.env**: Environment variables (not committed)
- **.env.example**: Template for required variables
- **load-env.php**: PHP environment loader

### Deployment
- **Easypanel**: Hosting platform (VPS)
- **Astro apps**: Nixpacks with `nixpacks.toml`
- **Go APIs**: Dockerfile (NOT Nixpacks)
- **Manual rebuild required**: Auto-deploy unreliable

### Go API Deployment Process
1. Generate `go.sum` locally: `cd api-go && go mod tidy`
2. Commit: `git add -A && git commit -m "..." && git push`
3. Easypanel → Service → **Rebuild** (manual)
4. Hard refresh browser: `Cmd+Shift+R`
5. Test: `/api/health`

### Active Services
- `landing-r11`: Astro (laruta11.cl)
- `app-r11`: Astro (app.laruta11.cl)
- `caja-r11`: Astro (caja.laruta11.cl)
- `api-go-landing-r11`: Go API S3 (Dockerfile)
- `api-go-caja-r11`: Go API caja (Dockerfile)

## API Architecture

### REST Endpoints
- JSON request/response format
- Session-based authentication
- CORS headers for cross-origin requests

### File Structure
```
api-go/              # Go APIs (NEW)
├── main.go          # Server setup
├── handlers.go      # Endpoint handlers
├── Dockerfile       # Build config
├── go.mod
└── go.sum

api/                 # PHP APIs (Legacy)
├── orders/          # Order management
├── users/           # Authentication
├── coupons/         # Discounts
├── notifications/   # Push notifications
└── *.php            # Core endpoints
```

### Common Patterns
- Database connection via config.php
- Session validation in protected endpoints
- JSON response with status codes
- Error handling with try-catch blocks

## Progressive Web App (PWA)

### Features
- **manifest.json**: App metadata and icons
- **sw.js**: Service worker for offline support
- **Cache-first strategy**: Fast loading
- **Install prompt**: Add to home screen

### Audio Assets
- notificacion.mp3, comanda.mp3, cupon.mp3
- Sound effects for user interactions

## Security

### Authentication
- Session-based auth with PHP sessions
- JWT tokens for API authentication
- Password hashing (bcrypt)

### Data Protection
- Prepared statements (PDO) prevent SQL injection
- Input validation and sanitization
- CORS configuration for API access
- Environment variables for sensitive data
