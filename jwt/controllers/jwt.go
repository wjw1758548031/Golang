package controllers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/kataras/iris"
	"time"
)

type user struct {
}

//测试用
func ReqJwt(ctx iris.Context) {
	user := make(map[string]interface{})
	user["id"] = "123456"
	user["name"] = "wjw"
	//CreateToken(user,1)
	fmt.Println("进入")
}

var secretKey string = "WkIgHSGLjCO4Z8GjjoAwRUK3yGwkgZyG"

/*
*创建正式token
 */
func CreateToken(ctx iris.Context) {
	datas := make(map[string]interface{})
	expires := 1
	datas["uid"] = int64(11)
	datas["name"] = int64(11)
	if expires == 0 {
		expires = 1
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	for k, v := range datas {
		claims[k] = v
	}
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(expires) * 24).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		panic("TOKEN_ERROR")
	}
	fmt.Println("tokenString:", tokenString)
	return
}

//获取检查token
func VerifyToken(ctx iris.Context) {
	token, err := request.ParseFromRequest(ctx.Request(), request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})
	if token == nil || err != nil {
		fmt.Println("PLEASE_LOGIN")
	}

	tokenData := token.Claims.(jwt.MapClaims)
	if int64(tokenData["exp"].(float64)) < time.Now().Unix() {
		fmt.Println("TOKEN_HAS_EPIRED")
	}
	users := make(map[string]interface{}, 0)
	users["name"] = tokenData["name"]
	users["uid"] = tokenData["uid"]
	ctx.Values().Set("users", users)
	fmt.Println("users:", users)
	ctx.Next()
	return
}

//获取检查token
func QueryToken(ctx iris.Context) {
	users := ctx.Values().Get("users")
	fmt.Println("users", users)
	uid := ctx.Values().Get("users").(map[string]interface{})["uid"]
	name := ctx.Values().Get("users").(map[string]interface{})["name"]
	fmt.Println("uid:", uid)
	fmt.Println("name:", name)
}
