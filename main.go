package main

import (
	"os"

	router "Modifa2/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	gin.DisableConsoleColor()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(gin.ErrorLogger())
	// r.Use(CORSMiddleware())
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	// OPTIONS method for ReactJS
	config.AllowMethods = []string{"POST", "GET", "OPTIONS"}
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Origin", "x-access-token", "content-type", "Content-Length", "Authorization", "Cache-Control"}
	config.ExposeHeaders = []string{"Content-Length"}
	r.Use(cors.New(config))
	// r.Use(gzip.Gzip(gzip.DefaultCompression))

	/************/
	/* Users */
	/************/
	router.Init(r)

	return r
}

func setupConfigs() {
	os.Setenv("CURRENTDOMAIN", "https://1c81-102-32-23-123.in.ngrok.io")

	os.Setenv("EmailFrom", "projectcommunication123@gmail.com")

	os.Setenv("InfoSMSAuthorization", "44847052db9d5bf604282652a9151d60-d73376a9-f31c-47fa-8edf-c99fe52e2adb")
	os.Setenv("InfoSMSAuthorizationUsername", "TebogoMampa")
	os.Setenv("InfoSMSAuthorizationPassword", "modifa193*")

	os.Setenv("InfoSMSBASEURL", "https://19rnq9.api.infobip.com")
	os.Setenv("InfoSMSURL", "api/sms/SMS")
	// os.Setenv("InfoSMSNofifyURL", os.Getenv("CURRENTDOMAIN")+"/api/v1/sms/NotificationStatus/")

	os.Setenv("ProjectMain", "postgres://cogjgedlgavael:cf43a86f559ebdd296331ca10991a0bfc87dfcf1fb7c83d3407698719348a669@ec2-18-204-74-74.compute-1.amazonaws.com:5432/d7jnruc4m8g23q")
	os.Setenv("WEBSERVER_PORT", "8080")

}

func main() {
	//Uncommented When Not Debugging
	// gin.SetMode(gin.ReleaseMode)
	// export GIN_MODE=release

	// gocron.Start()
	// s := gocron.NewScheduler()
	// gocron.Every(2).Seconds().Do(c.CheckNewUser)
	//  <-s.Start()

	r := setupRouter()

	setupConfigs()

	r.Run(":" + os.Getenv("WEBSERVER_PORT"))
}
