package util

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func ValidateToken (ctx *gin.Context ,token string ) error{
	claims := &Claims{}
	jwtSignedKey := []byte("secret_key")
	tokenParse, err := jwt.ParseWithClaims(token, claims, 
	func(t *jwt.Token) (interface{}, error){
		return jwtSignedKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			ctx.JSON(http.StatusUnauthorized, err)
			return err
		}
		ctx.JSON(http.StatusBadRequest, err)
		return err
	}

	if !tokenParse.Valid {
		ctx.JSON(http.StatusUnauthorized, "Invalid Token")
			return err
	}

	ctx.Next()
	return nil
}

func VerifyBaererToken (ctx *gin.Context ) error{
	authorizationHeaderKey := ctx.GetHeader("authorization")
	fields := strings.Fields(authorizationHeaderKey)
	tokenToValidate := fields[1]
	errorValidate := ValidateToken(ctx, tokenToValidate)
	
	if errorValidate != nil {
		return errorValidate
	}

	return nil
	
}