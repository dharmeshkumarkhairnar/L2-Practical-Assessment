package middleware

import (
	// "log"
	// "net/http"
	// "practical-assessment/constant"
	// "practical-assessment/utils"

	// // "practical-assessment/utils/redis"
	// "strconv"
	// "strings"
	// "time"

	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// 	authHeader := c.GetHeader("Authorization")
		// 	log.Println("Authorization header:", authHeader)
		// 	if authHeader == "" {
		// 		c.JSON(http.StatusUnauthorized, gin.H{
		// 			"message": "Header not Found",
		// 		})
		// 		c.Abort()
		// 		return
		// 	}
		// 	log.Println("Authorization header:", authHeader)
		// 	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		// 	log.Println("Authorization header:", authHeader)
		// 	c.Set(constant.Token, tokenString)

		// 	tokenExpiry, err := utils.ValidateToken(tokenString)
		// 	if err != nil {
		// 		c.JSON(http.StatusUnauthorized, gin.H{
		// 			"message": err.Error(),
		// 		})
		// 		c.Abort()
		// 		return
		// 	}
		// 	expiry, _ := strconv.Atoi(tokenExpiry)
		// 	exp := int64(expiry)
		// 	c.Set(constant.Expiry, exp)

		// 	// cacheKey := fmt.Sprintf("JWT_TOKEN:%s", tokenString)
		// 	// redisClient := redis.GetRedisClient()
		// 	// if redisClient != nil {
		// 	// 	exists, err := redisClient.Get(c, cacheKey).Result()
		// 	// 	if err != nil {
		// 	// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		// 	// 			"message": constant.RedisOperationFailedError,
		// 	// 		})
		// 	// 		return
		// 	// 	}
		// 	// 	if exists {

		// 	// 	}
		// 	// }

		// 	username, err := utils.ValidateToken(tokenString)
		// 	if err != nil {
		// 		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		// 			"message": err.Error(),
		// 		})
		// 		return
		// 	}
		// 	c.Set(constant.FieldName, username)

		// 	c.Next()
		duration := time.Since(start)
		log.Printf("Completed in %v", duration)

	}
}
