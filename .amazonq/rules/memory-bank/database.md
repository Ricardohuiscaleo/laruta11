# Database Migration Guide

## Current Setup

**Database**: MySQL 9.6 on VPS (Easypanel)
**Previous**: MySQL on Hostinger
**Migration Date**: February 2026

## Connection Details

### Internal (from APIs in VPS)
```
Host: websites_mysql-laruta11
Port: 3306
Database: laruta11
User: laruta11_user
Password: CCoonn22kk11@
```

### External (from local machine)
```
Host: 76.13.126.63
Port: 3306
Database: laruta11
User: laruta11_user
Password: CCoonn22kk11@
Root Password: b38027bdba3d91da9453
```

## Database Statistics

- **Products**: 226
- **Users**: 67
- **Orders (tuu_orders)**: 922
- **Total Tables**: 70

## Migration Process (for reference)

1. Export from Hostinger via phpMyAdmin
2. Fix collation: `sed -i 's/utf8mb4_uca1400_ai_ci/utf8mb4_unicode_ci/g' backup.sql`
3. Import to VPS: `docker exec -i <container> mysql -u root -p<pass> laruta11 < backup.sql`
4. Update API environment variables

## Recommended Tools

### Beekeeper Studio (Free, M1 Native)
- Best for daily use
- Fast, modern UI
- Download: https://www.beekeeperstudio.io/

### TablePlus (Premium, M1 Native)
- Fastest option
- $89 lifetime license
- Download: https://tableplus.com/

### Adminer (Web-based)
- Lightweight phpMyAdmin alternative
- Can deploy on VPS for remote access

## Future Migration: MySQL → PostgreSQL

**When to consider:**
- 1000+ orders/day
- Complex analytics queries slow
- Need better JSON support
- Need full-text search

**Current status:** MySQL is sufficient for current scale
