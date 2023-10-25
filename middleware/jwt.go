// middleware/jwt.go

package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type CustomClaims struct {
	jwt.StandardClaims
	FkCompany  int  `json:"fk_company"`
	FkCustomer int  `json:"fk_customer"`
	UserId     uint `json:"UserId"`
}

func ValidateTokenMiddleware() gin.HandlerFunc {
	jwtKey := os.Getenv("JWT_KEY")
	return func(c *gin.Context) {

		if c.Request.Method == "OPTIONS" {
			c.Next()
			return
		}
		if strings.HasPrefix(c.Request.URL.Path, "/swagger") {
			c.Next()
			return
		}
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Falta el encabezado de autorización"})
			c.Abort()
			return
		}

		tokenStr := strings.Replace(authHeader, "Bearer ", "", -1)
		token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		})

		// Comprobar error
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token no válido"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			// Almacenar FkCompany y FkCustomer en el contexto
			c.Set("FkCompany", claims.FkCompany)
			c.Set("FkCustomer", claims.FkCustomer)
			c.Set("UserId", claims.UserId)

		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token no válido"})
			c.Abort()
			return
		}

		c.Next()
	}
}
