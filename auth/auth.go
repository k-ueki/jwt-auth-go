package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"time"
)

func GetTokenHandler(c echo.Context)(string,error){
	// set jwt.header
	token:=jwt.New(jwt.SigningMethodHS256)

	// set jwt.claims
	claims:=token.Claims.(jwt.MapClaims)
	claims["admin"] = true
	claims["name"] = "test-user"
	claims["yotei"] = "go to hogehoge"
	// issuedAt: 発行日時
	claims["iat"] = time.Now()
	// expirationTime: 失効日時
	claims["exp"] = time.Now().Add(time.Hour*1).Unix()
	// subject: 用途
	claims["sub"] = "test-use"

	// secret: SIGNEDKEY
	//tokenString,err:=token.SignedString([]byte(os.Getenv("SIGNEDKEY")))
	tokenString,err:=token.SignedString([]byte("secret"))
	if err != nil {
		return "",err
	}
	return tokenString,nil
}
