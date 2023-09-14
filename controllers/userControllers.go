package controllers

import (
	"net/http"
	"restjwtgo/services/userServices"

	"github.com/gin-gonic/gin"
)

func SingUp(cnt *gin.Context) {

	var body struct {
		Email    string
		Password string
	}
	if cnt.Bind(&body) != nil {
		cnt.JSON(http.StatusBadRequest, gin.H{
			"error": "falid to read body",
		})
		return
	}

	userData, err := userServices.CreateUser(body.Email, body.Password)
	if err != nil {
		cnt.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	cnt.SetSameSite(http.SameSiteLaxMode)
	cnt.SetCookie("Authorization", userData.RefreshToken, 3600*24, "", "", false, true)

	cnt.JSON(http.StatusOK, gin.H{
		"user": userData,
	})

}

func Refresh(cnt *gin.Context) {
	refreshToken, err := cnt.Cookie("Authorization")
	if err != nil {
		cnt.AbortWithStatus(http.StatusUnauthorized)
	}
	var userData, erro = userServices.Refresh(refreshToken)
	if erro != nil {
		cnt.JSON(http.StatusUnauthorized, gin.H{
			"errors": erro.Error(),
		})
		return
	}
	cnt.SetSameSite(http.SameSiteLaxMode)
	cnt.SetCookie("Authorization", userData.RefreshToken, 3600*24, "", "", false, true)
	cnt.JSON(http.StatusOK, gin.H{
		"user": userData,
	})
}
