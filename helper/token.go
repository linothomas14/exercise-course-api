package helper

import (
	"crypto/rand"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TokenGenerator() string {
	b := make([]byte, 3)
	fmt.Println(b)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func GetUserIdFromClaims(ctx *gin.Context) int {
	userClaims, ok := ctx.Get("user_id")

	if !ok {
		response := BuildResponse("Cant get user_id from claims", EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return 0
	}

	id, ok := userClaims.(float64)

	if !ok {
		response := BuildResponse("Cant Parsing user_id", EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return 0
	}
	userID := int(id)

	return userID
}
