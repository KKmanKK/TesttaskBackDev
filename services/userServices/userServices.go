package userServices

import (
	"context"
	"errors"
	"restjwtgo/initializers"
	"restjwtgo/models"
	"restjwtgo/services/tokenServices"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

// Создаем структуру для удобного возврата данных из функции
type Data struct {
	User         models.User
	AccessToken  string
	RefreshToken string
}

// Функция для создания пользователя
func CreateUser(email string, pass string) (Data, error) {
	var data Data = Data{}

	hash, err := bcrypt.GenerateFromPassword([]byte(pass), 5)

	if err != nil {
		return data, errors.New("falid to hash password")
	}
	coll := initializers.GetDbCollection("users")
	data.User.Email = email
	data.User.Password = string(hash)
	data.User.ID = primitive.NewObjectID()

	_, err = coll.InsertOne(context.Background(), data.User)

	if err != nil {
		return data, errors.New("falid to create user")
	}
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"email": 1},
		Options: options.Index().SetUnique(true),
	}
	_, err = coll.Indexes().CreateOne(context.Background(), indexModel)

	accessToken, err := tokenServices.SetAccessToken(data.User.ID.Hex())
	if err != nil {
		return data, errors.New("Failed create accessToken")
	}

	refreshToken, err := tokenServices.SetRefreshToken(data.User.ID.Hex())
	tokenServices.SaveToken(data.User.ID.Hex(), refreshToken)

	if err != nil {
		return data, errors.New("Failed create refreshToken")
	}

	data.AccessToken = accessToken
	data.RefreshToken = refreshToken

	return data, nil
}

// Функция refresh для обновления access и refresh токенов
func Refresh(refreshToken string) (Data, error) {
	coll := initializers.GetDbCollection("users")
	var data Data = Data{}
	token, err := tokenServices.ValidateRefreshToken(refreshToken)
	if err != nil {
		return data, errors.New("User is unauthorized")
	}

	if _, err := tokenServices.FindToken(refreshToken); err != nil {
		return data, errors.New("User is unauthorized")
	}
	var userId string
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId = claims["sub"].(string)
	}
	objId, _ := primitive.ObjectIDFromHex(userId)
	result := coll.FindOne(context.Background(), bson.M{"_id": objId}).Decode(&data.User)
	if result != nil {

		return data, result
	}
	accessToken, err := tokenServices.SetAccessToken(data.User.ID.Hex())
	if err != nil {
		return data, errors.New("Failed create accessToken")
	}

	refrToken, err := tokenServices.SetRefreshToken(data.User.ID.Hex())
	tokenServices.SaveToken(data.User.ID.Hex(), refrToken)
	if err != nil {
		return data, errors.New("Failed create refreshToken")
	}
	data.AccessToken = accessToken
	data.RefreshToken = refrToken
	return data, nil
}
