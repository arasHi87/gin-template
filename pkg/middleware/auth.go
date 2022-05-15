package middleware

import (
	"net/http"
	"strings"

	"github.com/arashi87/gin-template/pkg/common"
	"github.com/arashi87/gin-template/pkg/model"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		// get jwt token from Authorization field in header
		token := ctx.Request.Header.Get("Authorization")

		//  check if token is blank
		if token == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"msg": "token miss"})
			ctx.Abort()
			return
		}

		// split token in format
		parts := strings.SplitN(token, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			ctx.JSON(http.StatusBadRequest, gin.H{"msg": "token format error, it should start with Bearer"})
			ctx.Abort()
			return
		}

		// parse token
		claims, err := common.ParseToken(parts[1])
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"msg": "invalid token"})
			ctx.Abort()
			return
		}

		// set user model in context
		var user model.UserModel
		common.DB.First(&user, claims.ID)
		ctx.Set("user", user)
		ctx.Next()
	}
}
