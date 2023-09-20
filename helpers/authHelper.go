package helpers

import (
	"errors"
	"log"

	"github.com/gin-gonic/gin"
)

func CheckUserType(c *gin.Context, role string) (err error) {
	userType := c.GetString("uid")
	log.Println("userTYPE", userType)

	err = nil
	if userType != role {
		err = errors.New("Unauthorized to access this resource ddfd")
		return err
	}
	return err
}

func MatchUserTypeToUid(c *gin.Context, userId string) (err error) {
	userType := c.GetString("user_type")
	uid := c.GetString("uid")
	err = nil

	log.Println("usertype ", userType)
	log.Println("uid ", uid)

	if userType == "USER" && uid != userId {
		err = errors.New("Unauthorized to access this resource aaaaaaa")
		return err
	}
	err = CheckUserType(c, userType)
	return err
}
