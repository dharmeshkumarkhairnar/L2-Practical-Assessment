package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"practical-assessment/constant"
	"practical-assessment/model"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

func AuthMiddleware(redisClient *redis.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, model.ErrorAPIResponse{
				Message: model.ErrorMessage{Key: "token", ErrorMessage: "authorization header missing"},
				Error:   "unauthorized",
			})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		if tokenString == authHeader {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, model.ErrorAPIResponse{
				Message: model.ErrorMessage{Key: "token", ErrorMessage: "invalid header format"},
				Error:   "invalid token",
			})
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if token.Method != jwt.SigningMethodHS256 {
				return nil, errors.New("Unexpected signing method")
			}
			return []byte(constant.AccessSecretKey), nil
		})
		if err != nil {
			fmt.Println("token parsing error:", err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, model.ErrorAPIResponse{
				Message: model.ErrorMessage{Key: "token", ErrorMessage: "invalid token"},
				Error:   err.Error(),
			})
			return
		}
		ctx.Set("token", tokenString)

		//extracting claims for verification
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, model.ErrorAPIResponse{
				Message: model.ErrorMessage{
					Key:          "token",
					ErrorMessage: "Invalid token",
				},
				Error: "unauthorized",
			})
			return
		}

		expiry := claims["exp"].(float64)
		if expiry < float64(time.Now().Unix()) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, model.ErrorAPIResponse{
				Message: model.ErrorMessage{
					Key:          "token",
					ErrorMessage: "token is already expired",
				},
				Error: "unauthorized",
			})
			return
		}

		userID, err := redisClient.Get(ctx.Request.Context(), tokenString).Result()
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, model.ErrorAPIResponse{
				Message: model.ErrorMessage{Key: "session", ErrorMessage: "session expires or logged out"},
				Error:   "login failed",
			})
			return
		}

		id, err := strconv.Atoi(userID)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
			return
		}
		userIDInt := int64(id)

		//concurrent session control -
		// only one login per user is allowed,
		// checking if the token in the header is same a active token in redis -
		// if YES: flow continues,
		// if NOT: current token will be set as new active session and older one will become invalid and return with the following response boty
		sessionKey := fmt.Sprintf("userId:%d", userIDInt)
		sessionToken, err := redisClient.Get(ctx, sessionKey).Result()
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
			return
		}

		if sessionToken != tokenString {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, model.ErrorAPIResponse{
				Message: model.ErrorMessage{Key: "token", ErrorMessage: "session expired"},
				Error:   "login detected on another device",
			})
			return
		}

		ctx.Set("userId", userIDInt)
		ctx.Next()
	}
}
