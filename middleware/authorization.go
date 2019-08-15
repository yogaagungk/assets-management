package middleware

import (
	"log"
	"net/http"

	"github.com/gomodule/redigo/redis"

	"github.com/gin-gonic/gin"
	"github.com/yogaagungk/assets-management/common"
	"github.com/yogaagungk/assets-management/util/auth"
)

type AuthService struct {
	redis redis.Conn
}

func ProvideAuthService(redis redis.Conn) AuthService {
	return AuthService{redis}
}

func (authS *AuthService) Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("entering Authorization middleware")

		tokenHeader := c.GetHeader("Authorization")

		if tokenHeader == "" {
			log.Println("Can't found Authorization header")

			c.JSON(http.StatusBadRequest, gin.H{
				"status": common.INVALID_HEADER,
			})
			c.Abort()
			return
		}

		tokenString := tokenHeader[7:len(tokenHeader)]

		username := auth.ParseToken(tokenString)

		if username != "" {
			_, err := redis.String(authS.redis.Do("GET", username))

			if err == redis.ErrNil {
				log.Println("Unauthorization failed because token not found in cache")

				c.JSON(http.StatusUnauthorized, gin.H{
					"status": common.ACCESS_DENIED,
				})
				c.Abort()
				return
			} else if err != nil {
				log.Println(err.Error())
				log.Println("Can't read data from cache")

				c.JSON(http.StatusUnauthorized, gin.H{
					"status": common.ACCESS_DENIED,
				})
				c.Abort()
				return
			}

		} else {
			log.Println("Unauthorization failed because token not valid")

			c.JSON(http.StatusUnauthorized, gin.H{
				"status": common.ACCESS_DENIED,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
