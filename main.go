package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/k-ueki/jwt-auth-go/auth"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type User struct{
	ID string `json:"id"`
	Name string `json:"name"`
}

func login(c echo.Context)error{
	token,err:=auth.GetTokenHandler(c)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK,map[string]string{
		"token":token,
	})
}

func restricted(c echo.Context)error{
	user := c.Get("user").(*jwt.Token)
  	claims := user.Claims.(jwt.MapClaims)
  	name := claims["name"].(string)
	yotei := claims["yotei"].(string)
  	return c.String(http.StatusOK, "You are " +name +"\n and your next yotei is "+yotei)
}

func main() {
	e:=echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/login",login)

	r:=e.Group("/restricted")
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("",restricted)

	e.Logger.Fatal(e.Start(":8080"))
}
