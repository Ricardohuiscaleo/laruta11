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
		if (req.User == os.Getenv("CAJA_USER_CAJERA") && req.Pass == os.Getenv("CAJA_PASSWORD_CAJERA")) || (req.User == os.Getenv("CAJA_USER_ADMIN") && req.Pass == os.Getenv("CAJA_PASSWORD_ADMIN")) {
			valid, role = true, "caja"
		}
	case "inventario":
		if req.User == os.Getenv("INVENTARIO_USER") && req.Pass == os.Getenv("INVENTARIO_PASSWORD") {
			valid, role = true, "inventario"
		}
	case "comandas":
		if req.User == "comandas" && req.Pass == os.Getenv("CAJA_PASSWORD_CAJERA") {
			valid, role = true, "comandas"
		}
	case "admin":
		if (req.User == "admin" && req.Pass == os.Getenv("ADMIN_PASSWORD_ADMIN")) || (req.User == "ricardo" && req.Pass == os.Getenv("ADMIN_PASSWORD_RICARDO")) || (req.User == "manager" && req.Pass == os.Getenv("ADMIN_PASSWORD_MANAGER")) || (req.User == "ruta11" && req.Pass == os.Getenv("ADMIN_PASSWORD_RUTA11")) {
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
