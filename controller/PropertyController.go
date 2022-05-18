package controller

import (
	"Modifa2/models"
	s "Modifa2/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

/* Register New Property...*/
func RegisterProperty(c *gin.Context) {
	db := s.DB{}
	//
	var rb models.Returnblock
	var u models.PropertyRequest
	if err := c.ShouldBindBodyWith(&u, binding.JSON); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	//Validate Json Input
	if err := u.PropertyValidate(); err != nil {
		c.JSON(http.StatusBadRequest, rb.New(false, "", err))
		fmt.Println("Inavlid Json Input", err.Error())
		return
	}
	//use function
	rows, err := db.SaveOnDB("project_main.add_property", u)
	if err != nil {
		fmt.Println("QueryRow failed: ", err.Error())
		errormessage := fmt.Sprintf("%v\n", err)
		c.JSON(http.StatusBadRequest, rb.New(false, err.Error(), errormessage))
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, rb.New(false, err.Error(), err))
	} else {
		c.JSON(http.StatusOK, rb.New(true, "", rows))
	}
}

/* Get User Property...*/
func GetProperty(c *gin.Context) {
	db := s.DB{}
	//
	var rb models.Returnblock
	var u models.GetPropertyRequest
	if err := c.ShouldBindBodyWith(&u, binding.JSON); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	//use function
	rows, err := db.GetProperty("project_main.get_system_user_properties", u)
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

/* Get All Properties...*/
func GetAllProperties(c *gin.Context) {
	db := s.DB{}
	//
	var rb models.Returnblock
	var u models.GetPropertyAllRequest
	if err := c.ShouldBindBodyWith(&u, binding.JSON); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	//use function
	rows, err := db.GetProperty("project_main.get_all_properties", u)
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

/* Delete Property...*/
func DeleteProperty(c *gin.Context) {
	db := s.DB{}
	//
	var rb models.Returnblock
	var u models.DeleteAllRequest
	if err := c.ShouldBindBodyWith(&u, binding.JSON); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	//use function
	rows, err := db.SaveOnDB("project_main.proper_delete", u)
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
