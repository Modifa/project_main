package router

import (
	c "Modifa2/controller"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	//Admin
	admin := r.Group("api/admin/")
	{
		admin.POST("RegisterAdmin", c.RegisterAdmin)
		admin.POST("ActivateAccount", c.ActivateAccount)
		admin.POST("DeleteAccount", c.DeleteAccount)
		admin.POST("FortgotPassword", c.FortgotPassword)
		admin.POST("GetAllAdmins", c.GetAllAdmins)
	}
	//Emails
	email := r.Group("api/emails/")
	{
		email.POST("Email", c.Email)
	}
	//SMS
	sms := r.Group("api/sms/")
	{
		sms.POST("SMS", c.SMS)
	}
	//User
	user := r.Group("api/user/")
	{
		user.POST("GetUser", c.GetUser)
		user.POST("DeleteUser", c.DeleteUser)
		// user.POST("GetAssociations", c.GetAssociations)

	}
	//properties
	properties := r.Group("api/properties/")
	{
		properties.POST("RegisterProperty", c.RegisterProperty)
		properties.POST("GetProperty", c.GetProperty)
		properties.POST("GetAllProperties", c.GetAllProperties)
		properties.POST("DeleteProperty", c.DeleteProperty)
	}
	//Done Push Again
	//One More Time

}
