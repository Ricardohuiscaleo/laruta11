#!/bin/bash
# Script para reorganizar api-go en carpetas

cd "$(dirname "$0")"

# Crear estructura de carpetas
mkdir -p auth compras inventory quality catalog orders shared

# Mover auth.go a carpeta auth/
cat > auth/handler.go << 'EOF'
package auth

import (
	"database/sql"
	"os"
	"github.com/gin-gonic/gin"
)

type Handler struct{ DB *sql.DB }

func (h *Handler) Login(c *gin.Context) {
	var req struct {
		User string `json:"user"`
		Pass string `json:"pass"`
		Type string `json:"type"`
	}
	c.BindJSON(&req)
	valid, role := false, ""
	switch req.Type {
	case "caja":
		if (req.User == "ruta11caja" && req.Pass == os.Getenv("CAJA_USER_CAJERA")) || (req.User == "admin123" && req.Pass == os.Getenv("CAJA_USER_ADMIN")) {
			valid, role = true, "caja"
		}
	case "inventario":
		if req.User == os.Getenv("INVENTARIO_USER") && req.Pass == os.Getenv("INVENTARIO_PASSWORD") {
			valid, role = true, "inventario"
		}
	case "comandas":
		if req.User == "comandas" && req.Pass == os.Getenv("CAJA_USER_CAJERA") {
			valid, role = true, "comandas"
		}
	case "admin":
		if (req.User == "admin" && req.Pass == os.Getenv("ADMIN_USER_ADMIN")) || (req.User == "ricardo" && req.Pass == os.Getenv("ADMIN_USER_RICARDO")) || (req.User == "manager" && req.Pass == os.Getenv("ADMIN_USER_MANAGER")) || (req.User == "ruta11" && req.Pass == os.Getenv("ADMIN_USER_RUTA11")) {
			valid, role = true, "admin"
		}
	}
	if valid {
		c.JSON(200, gin.H{"success": true, "role": role, "user": req.User})
	} else {
		c.JSON(401, gin.H{"success": false, "error": "Credenciales inválidas"})
	}
}

func (h *Handler) Check(c *gin.Context) {
	if user := c.GetHeader("X-User"); user != "" {
		c.JSON(200, gin.H{"success": true, "authenticated": true})
	} else {
		c.JSON(401, gin.H{"success": false, "authenticated": false})
	}
}

func (h *Handler) Logout(c *gin.Context) {
	c.JSON(200, gin.H{"success": true})
}
EOF

# Mover compras.go a carpeta compras/
mv compras.go compras/handler.go 2>/dev/null || echo "compras.go ya movido"

# Crear shared/db.go
cat > shared/db.go << 'EOF'
package shared

import "database/sql"

type Server struct{ DB *sql.DB }
EOF

# Actualizar main.go
cat > main.go << 'EOF'
package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"api-go/auth"
	"api-go/compras"
	"api-go/shared"
)

func main() {
	dsn := os.Getenv("APP_DB_USER") + ":" + os.Getenv("APP_DB_PASS") + "@tcp(" + os.Getenv("APP_DB_HOST") + ")/" + os.Getenv("APP_DB_NAME") + "?parseTime=true"
	db, _ := sql.Open("mysql", dsn)
	defer db.Close()
	db.SetMaxOpenConns(25)

	authH := &auth.Handler{DB: db}
	comprasH := &compras.Handler{DB: db}

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,X-User")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	r.GET("/api/health", func(c *gin.Context) { c.JSON(200, gin.H{"status": "ok"}) })

	// Auth
	r.POST("/api/auth/login", authH.Login)
	r.GET("/api/auth/check", authH.Check)
	r.POST("/api/auth/logout", authH.Logout)

	// Compras
	r.GET("/api/compras", comprasH.GetCompras)
	r.GET("/api/compras/items", comprasH.GetItems)
	r.GET("/api/compras/proveedores", comprasH.GetProveedores)
	r.GET("/api/compras/saldo", comprasH.GetSaldo)
	r.GET("/api/compras/historial-saldo", comprasH.GetHistorialSaldo)
	r.GET("/api/compras/precio-historico", comprasH.GetPrecioHistorico)
	r.POST("/api/compras", comprasH.Registrar)
	r.DELETE("/api/compras/:id", comprasH.Delete)
	r.POST("/api/compras/:id/respaldo", comprasH.UploadRespaldo)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3002"
	}
	log.Printf("🚀 API on :%s", port)
	r.Run(":" + port)
}
EOF

echo "✅ Estructura de carpetas creada"
echo "📁 auth/, compras/, shared/"
echo "⚠️  Ejecuta: go mod tidy"
