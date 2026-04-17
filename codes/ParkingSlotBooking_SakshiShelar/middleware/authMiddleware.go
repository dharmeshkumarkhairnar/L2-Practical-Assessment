package middleware

import (
	"fmt"
	"log"
	"net/http"
	"practical-assessment/constant"
	"practical-assessment/model"
	"practical-assessment/utils"
	"practical-assessment/utils/redis"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()

		redisClient, err := redis.GetRedisClient()
		if redisClient == nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, model.ErrorMessageResponse{
				Message: model.ErrorMessage{
					Key:      constant.Redis,
					ErrorMsg: err.Error(),
				},
				ErrorMessage: constant.UserUnauthorized,
			})
			ctx.Abort()
			return
		} 
		
		authHeader := ctx.GetHeader(constant.Authorization)
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, model.ErrorMessageResponse{
				Message: model.ErrorMessage{
					Key:      constant.Authorization,
					ErrorMsg: constant.HeaderMissing,
				},
				ErrorMessage: constant.UserUnauthorized,
			})
			ctx.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, constant.Bearer)
		ctx.Set(constant.Token, tokenString)

		redisKey := tokenString
		ctx.Set(constant.RedisKey, redisKey)

		existsInRedis, err := redisClient.Exists(ctx.Request.Context(), redisKey).Result()
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorMessageResponse{
				Message: model.ErrorMessage{
					Key:      constant.Redis,
					ErrorMsg: err.Error(),
				},
				ErrorMessage: constant.InternalServer,
			})
			ctx.Abort()
			log.Fatalf("$$$$$$$$$$$$$$$ Redis error: %v", err)
			return
		}

		if existsInRedis == 0 {
			ctx.IndentedJSON(http.StatusUnauthorized,  model.ErrorMessageResponse{
				Message: model.ErrorMessage{
					Key:      "generic",
					ErrorMsg: constant.UserAlreadyLoggedOut,
				},
				ErrorMessage: constant.UserUnauthorized,
			})
			ctx.Abort()
			return
		}

		email, err := utils.GetEmailFromToken(tokenString)
		fmt.Println(email)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, model.ErrorMessageResponse{
				Message: model.ErrorMessage{
					Key:      constant.Header,
					ErrorMsg: err.Error(),
				},
				ErrorMessage: constant.UserUnauthorized,
			})
			ctx.Abort()
			return
		}

		userId, _ := redisClient.Get(ctx, tokenString).Int()
		ctx.Set("user_id", userId)

		tokenExpiry, err := utils.GetExpiryFromToken(tokenString)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, model.ErrorMessageResponse{
				Message: model.ErrorMessage{
					Key:      constant.Header,
					ErrorMsg: err.Error(),
				},
				ErrorMessage: constant.UserUnauthorized,
			})
			ctx.Abort()
			return
		}

		ctx.Set(constant.Email, email)
		ctx.Set(constant.Expiry, tokenExpiry)

		ctx.Next()
		log.Printf("Completed in: %v", time.Since(start))
	}

}
