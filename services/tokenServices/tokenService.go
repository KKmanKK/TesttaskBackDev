package tokenServices

import (
	"context"
	"fmt"
	"os"
	"restjwtgo/initializers"
	"restjwtgo/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Генерация access токена
func SetAccessToken(id string) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": id,
		"exp": time.Now().Add(time.Second * 20).Unix(),
	})

	tokenString, err := accessToken.SignedString([]byte(os.Getenv("ACCESS_TOKEN_KEY")))

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// Генерация resfresh токена
func SetRefreshToken(id string) (string, error) {
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := refreshToken.SignedString([]byte(os.Getenv("REFRESH_TOKEN_KEY")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// Сохранение access токена в базу данных
func SaveToken(user_id string, refreshToken string) {
	coll := initializers.GetDbCollection("token")
	token := models.Token{}
	err := coll.FindOne(context.Background(), bson.M{"user_id": user_id}).Decode(&token)
	if err == nil {
		token.RefreshToken = refreshToken
		_, err = coll.UpdateOne(context.Background(), bson.M{"user_id": user_id}, token)
		if err != nil {
			fmt.Println("Udpate")
		}
		return
	}
	newtoken := models.Token{ID: primitive.NewObjectID(), User_id: user_id, RefreshToken: refreshToken}
	coll.InsertOne(context.Background(), newtoken)
	return
}

// Валидация access токена
func ValidateAccessToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("ACCESS_TOKEN_KEY")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// Валидация refresh токена
func ValidateRefreshToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("REFRESH_TOKEN_KEY")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// Нахождение refresh токена по его значению
func FindToken(refreshToken string) (interface{}, error) {
	coll := initializers.GetDbCollection("token")
	token := models.Token{}
	err := coll.FindOne(context.Background(), bson.M{"refreshToken": refreshToken}).Decode(&token)

	return token, err
}
