package controller

import (
	"Modifa2/models"
	s "Modifa2/services"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

/* Register New Admin...*/
func RegisterAdmin(c *gin.Context) {
	db := s.DB{}
	//
	var rb models.Returnblock
	var u models.AdminRequest
	if err := c.ShouldBindBodyWith(&u, binding.JSON); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	//Validate Json Input
	if err := u.AdminValidate(); err != nil {
		c.JSON(http.StatusBadRequest, rb.New(false, "", err))
		fmt.Println("Inavlid Json Input", err.Error())
		return
	}

	//use function
	rows, err := db.SaveOnDB("project_main.add_system_user", u)
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
		email := u.Email_Address
		FullName := u.First_Name + ` ` + u.Last_Name
		RegisterEmail(email, FullName)
	}

}

/* Activate New Account...*/
func ActivateAccount(c *gin.Context) {
	db := s.DB{}
	//
	var rb models.Returnblock
	var u models.ActicateRequest
	if err := c.ShouldBindBodyWith(&u, binding.JSON); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	//use function
	rows, err := db.ActivateAccount("project_main.System_User_Activate", u)
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

/* Delete New Account...*/
func DeleteAccount(c *gin.Context) {
	db := s.DB{}
	//
	var rb models.Returnblock
	var u models.ActivateResponse
	if err := c.ShouldBindBodyWith(&u, binding.JSON); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	//use function
	rows, err := db.DeleteAccount("project_main.system_users_delete", u)
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

/* Forgot Password...*/
func FortgotPassword(c *gin.Context) {
	db := s.DB{}
	//
	var rb models.Returnblock
	var u models.ForgotPassword
	if err := c.ShouldBindBodyWith(&u, binding.JSON); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	//Validate Json Input
	if err := u.ForgotPasswordValidate(); err != nil {
		c.JSON(http.StatusBadRequest, rb.New(false, "", err))
		fmt.Println("Inavlid Json Input", err.Error())
		return
	}

	//use function
	rows, err := db.ForgotPassword("project_main.system_user_forgotpassword", u)
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

/* Get All Admins...*/
func GetAllAdmins(c *gin.Context) {
	db := s.DB{}
	//
	var rb models.Returnblock
	var u models.GetPropertyAllRequest
	if err := c.ShouldBindBodyWith(&u, binding.JSON); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	//use function
	rows, err := db.GetAllAdmins("project_main.get_system_users_internal_system_admin", u)
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
		c.JSON(http.StatusOK, rb.New(true, "Success", rows))
		fmt.Fprintf(os.Stdout, "QueryRow Success: %v\n", rows)
	}
}

func GetAssociations(c *gin.Context) {
	db := s.DB{}

	var rb models.Returnblock

	var l models.EmptyGetter
	if err := c.ShouldBindJSON(&l); err != nil {
		c.JSON(http.StatusUnprocessableEntity, rb.New(false, "Invalid json provided", err))
	}
	rows, err := db.ReturnAssocuations("AGM.get_all_associations", l)
	// rows, err := db.ReturnAssociations("AGM.get_all_associations", l)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		errormessage := fmt.Sprintf("%v\n", err)
		c.JSON(http.StatusBadRequest, rb.New(false, errormessage, err))
	} else {
		c.JSON(http.StatusOK, rb.New(true, "Success", rows))
		fmt.Println(rows)
		fmt.Fprintf(os.Stdout, "QueryRow Success: %v\n", rows)
	}
}
