package controller

import (
	"Modifa2/models"
	s "Modifa2/services"
	v "Modifa2/views"
	"strings"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

/*Email ...*/
func Email(c *gin.Context) {
	var u models.Messages

	vr := *new(v.Messages).New()

	if err := c.ShouldBindBodyWith(&u, binding.JSON); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	fullname := "Tebogo Modifa Mampa"
	message := strings.ReplaceAll(vr.RegisteredMessage, "{FullName}", fullname)

	u.Text = message
	s.Send(u.To, u.Text)
}

//
func RegisterEmail(ToEmail string, FullName string) {
	vr := *new(v.Messages).New()

	fullname := FullName
	message := strings.ReplaceAll(vr.RegisteredMessage, "{FullName}", fullname)

	Text := message
	s.Send(ToEmail, Text)
}
