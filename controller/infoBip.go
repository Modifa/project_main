package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"Modifa2/models"
	s "Modifa2/services"
)

/*SMS'S...*/
func SMS(c *gin.Context) {
	SMSservices := s.SMSservices()

	msgblock := models.InfoBipMessages{}

	// var rb models.Returnblock

	if err := c.ShouldBindBodyWith(&msgblock, binding.JSON); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	SMSservices.SendSMS(msgblock)
}
