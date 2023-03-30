package middleware

import (
	config2 "Skyriders/config"
	"Skyriders/model"
	"Skyriders/utils"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authorize(obj string, act string, enforcer *casbin.Enforcer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sub, existed := ctx.Get("currentUser")
		if !existed {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "user has not logged in yet"})
			return
		}

		err := enforcer.LoadPolicy()
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "failed to load policy"})
			return
		}

		userObj, ok := sub.(*model.User)
		if !ok {
			ctx.AbortWithStatus(http.StatusInternalServerError)
		}

		ok, err = enforcer.Enforce(fmt.Sprint(userObj.ID.Hex()), obj, act)

		if err != nil {
			ctx.AbortWithStatusJSON(500, gin.H{"message": "error occurred while authorizing user"})
			return
		}

		if !ok {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "you are not authorized"})
			return
		}

		ctx.Next()
	}
}

func Anonymous() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accessToken, err := ctx.Cookie("access_token")
		if err != nil {
			ctx.Next()
		}

		config, _ := config2.LoadConfig(".")
		_, err = utils.ValidateToken(accessToken, config.AccessTokenPublicKey)
		if err == nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "You are already logged in"})
			return
		}
		ctx.Next()
	}
}
