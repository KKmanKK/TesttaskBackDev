package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func ValidateRerfresh(cnt *gin.Context) {
	var accessToken string
	// var err error

	authHeader := cnt.GetHeader("Authorization")
	words := strings.Split(authHeader, " ")
	for i, val := range words {
		if i == 1 {
			accessToken = val
		}
	}
	if accessToken == "" {
		cnt.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// userAcces, err := tokenServices.ValidateAccessToken(accessToken)
	// if err != nil {
	// 	cnt.AbortWithStatus(http.StatusUnauthorized)
	// 	return
	// }

	// if claims, ok := userAcces.Claims.(jwt.MapClaims); ok && userAcces.Valid {

	// 	if float64(time.Now().Unix()) > claims["exp"].(float64) {
	// 		cnt.AbortWithStatus(http.StatusUnauthorized)
	// 	}
	// 	user := models.User{}

	// 	initializers.DB.First(&user, "id = ?", claims["sub"])

	// 	if user.ID == 0 {
	// 		cnt.AbortWithStatus(http.StatusUnauthorized)
	// 	}
	// 	cnt.Set("user", user)
	// 	cnt.Next()
	// }

}
