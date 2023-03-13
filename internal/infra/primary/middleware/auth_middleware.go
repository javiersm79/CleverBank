package middleware

import (
	"cleverbank/internal/infra/config/properties"
	"fmt"
	"github.com/golang-jwt/jwt"
)

func IsAuthorized(userToken string) (bool, error) {

	/*if r.Header["Token"] == nil {
		var err Error
		err = SetError(err, "No Token Found")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}*/

	var mySigningKey = []byte(properties.GetApplicationProperties().SecretKeyJWT)

	_, err := jwt.Parse(userToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return false, fmt.Errorf("There was an error in parsing token.")
		}
		return mySigningKey, nil
	})

	if err != nil {
		return false, fmt.Errorf("Your Token has been expired.")
	}

	return true, nil

}
