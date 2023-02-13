package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/linothomas14/exercise-course-api/helper"
	"github.com/linothomas14/exercise-course-api/service"
)

//AuthorizeJWT validates the token user given, return 401 if not valid
func AuthorizeJWT(jwtService service.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.BuildResponse("Use Bearer token please", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		if authHeader == "" {
			response := helper.BuildResponse("No token found", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		authHeader = strings.Replace(authHeader, "Bearer ", "", -1)

		token, err := jwtService.ValidateToken(authHeader)

		if err != nil {
			response := helper.BuildResponse(err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		if !token.Valid {
			response := helper.BuildResponse("Token is not valid", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		c.Set("claims", claims)
		c.Set("user_id", claims["user_id"])
		c.Set("role", claims["role"])

		c.Next()

	}

}

func AuthorizeJWTAdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		role := GetRoleFromClaims(c)
		if role != "admin" {
			response := helper.BuildResponse("You must be admin to access this endpoint", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

	}
}

func GetUserIdFromClaims(ctx *gin.Context) int {
	userClaims, ok := ctx.Get("user_id")
	if !ok {
		response := helper.BuildResponse("Cant get user_id from claims", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return 0
	}

	id, ok := userClaims.(float64)

	if !ok {
		response := helper.BuildResponse("Cant Parsing user_id", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return 0
	}
	userID := int(id)

	return userID
}

func GetRoleFromClaims(ctx *gin.Context) string {
	userClaims, ok := ctx.Get("role")

	if !ok {
		response := helper.BuildResponse("Cant get role from claims", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return ""
	}

	str := fmt.Sprintf("%v", userClaims)

	return str
}
