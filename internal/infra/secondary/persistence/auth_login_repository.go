package persistence

import (
	"cleverbank/internal/core/domain/auth"
	"cleverbank/internal/infra/config/properties"
	"cleverbank/internal/infra/secondary/persistence/dto"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthLoginRepository struct {
	Dbconex *gorm.DB
}

func (alr AuthLoginRepository) Login(req auth.LoginReq) (auth.LoginResponse, error) {
	var authUser dto.User

	if err := alr.Dbconex.Where("email = 	?", req.Email).First(&authUser).Error; err != nil {
		return auth.LoginResponse{}, fmt.Errorf("user not found: %s", req.Email)
	}
	check := CheckPasswordHash(req.Password, authUser.Password)

	if !check {
		return auth.LoginResponse{}, fmt.Errorf("user not found: %s", req.Email)
	}

	validToken, err := GenerateJWT(authUser.Email)
	if err != nil {
		return auth.LoginResponse{}, fmt.Errorf("failed to generate token")
	}

	return auth.LoginResponse{
		Email:       req.Email,
		TokenString: validToken,
	}, nil
}
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateJWT(email string) (string, error) {
	properties.GetApplicationProperties()
	var mySigningKey = []byte(properties.GetApplicationProperties().SecretKeyJWT)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Minute * 10).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Errorf("Something went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}
