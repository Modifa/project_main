package controller

import (
	"Modifa2/models"
	s "Modifa2/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

/* Register New Admin...*/
func DeleteUser(c *gin.Context) {
	db := s.DB{}
	//
	var rb models.Returnblock
	var u models.DeleteUserRequest
	if err := c.ShouldBindBodyWith(&u, binding.JSON); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	//use function
	rows, err := db.SaveOnDB("project_main.System_User_Delete", u)
	if err != nil {
		fmt.Println("QueryRow failed: ", err.Error())
		errormessage := fmt.Sprintf("%v\n", err)
		c.JSON(http.StatusBadRequest, errormessage)
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, rb.New(false, err.Error(), err))
	} else {
		c.JSON(http.StatusOK, rb.New(true, "", rows))
	}
}

//get_system_user

/* Register New Admin...*/
func GetUser(c *gin.Context) {
	db := s.DB{}
	//
	var rb models.Returnblock
	var u models.GetUserRequest
	if err := c.ShouldBindBodyWith(&u, binding.JSON); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	//	//Validate Json Input
	if err := u.UserValidate(); err != nil {
		c.JSON(http.StatusBadRequest, rb.New(false, "", err))
		fmt.Println("Inavlid Json Input", err.Error())
		return
	}
	//use function
	rows, err := db.GetUser("project_main.get_system_user", u)
	if err != nil {
		fmt.Println("QueryRow failed: ", err.Error())
		errormessage := fmt.Sprintf("%v\n", err)
		c.JSON(http.StatusBadRequest, errormessage)
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, rb.New(false, err.Error(), err))
	} else {
		c.JSON(http.StatusOK, rb.New(true, "", rows))
	}
}
