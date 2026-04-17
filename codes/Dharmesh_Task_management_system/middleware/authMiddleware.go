package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
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
		log.Printf("Request:%s", ctx.Request.Method)

		defer func() {
			log.Printf("Request completed in %d ms", time.Since(start).Microseconds())
		}()

		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" {
			errorMsg := model.ErrorMessage{
				Key:     "header",
				Message: "authorization header is missing",
			}
			ctx.AbortWithStatusJSON(http.StatusBadRequest, &model.ErrorAPIResponse{
				ErrorMsg: errorMsg,
				Message:  "login failed",
			})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := utils.ParseToken(tokenString)

		if err != nil {
			errorMsg := model.ErrorMessage{
				Key:     "token",
				Message: err.Error(),
			}
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, &model.ErrorAPIResponse{
				ErrorMsg: errorMsg,
				Message:  "login failed",
			})
			return
		}

		claims, err := utils.VerifyToken(token)

		if err != nil {
			errorMsg := model.ErrorMessage{
				Key:     "token",
				Message: err.Error(),
			}
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, &model.ErrorAPIResponse{
				ErrorMsg: errorMsg,
				Message:  "login failed",
			})
			return
		}

		userID := claims["iss"].(float64)

		redisClient := redis.GetRedisClient()
		cacheKey := fmt.Sprintf("ACTIVE_TOKEN_%d", int64(userID))
		c := context.Background()

		exits, err := redisClient.Exists(c, cacheKey).Result()

		if err != nil {
			errorMsg := model.ErrorMessage{
				Key:     "redis",
				Message: "redis connection error",
			}
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, &model.ErrorAPIResponse{
				ErrorMsg: errorMsg,
				Message:  "login failed",
			})
			return
		}

		if exits == 0 {
			errorMsg := model.ErrorMessage{
				Key:     "token",
				Message: "token is invalid",
			}
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, &model.ErrorAPIResponse{
				ErrorMsg: errorMsg,
				Message:  "login failed",
			})
			return
		}

		ctx.Set("userID", int64(userID))
		ctx.Set("token", tokenString)

		ctx.Next()

	}
}
