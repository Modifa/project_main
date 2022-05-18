package services

import (
	"Modifa2/models"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/go-resty/resty/v2"
)

/*Notifyurl ...*/
var (
	Notifyurl string
)

/*SMS ... */
type SMS struct {
}

type SMSService interface {
	Send(message models.InfoBipMessages) error
	SendSMS(message models.InfoBipMessages)
}

type smsservice struct {
	auth, url string
}

// auth-jwt
func SMSservices() SMSService {
	return &smsservice{
		auth: os.Getenv("InfoSMSAuthorization"),
		url:  os.Getenv("InfoSMSBASEURL"),
	}
}

/*Send ... */
func (f *smsservice) Send(message models.InfoBipMessages) error {

	// Create a Resty Client
	client := resty.New()

	r := &models.Returnblock{}
	r.Status = true
	r.Message = "Success"

	//fmt.Println("PayLoad Info:", string(b))

	// POST JSON string
	// No need to set content type, if you have client level setting
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(message).
		SetHeader("Authorization", "App "+f.auth).
		Post(f.url)

	if err != nil {
		r.Status = false
		r.Message = ""
		r.Data = err
		fmt.Println(err)
	} else {
		// Explore response object

		// fmt.Println("Response Info:")
		// fmt.Println("  Error      :", err)
		// fmt.Println("  Status Code:", resp.StatusCode())
		// fmt.Println("  Status     :", resp.Status())
		// fmt.Println("  Proto      :", resp.Proto())
		// fmt.Println("  Time       :", resp.Time())
		// fmt.Println("  Received At:", resp.ReceivedAt())
		// fmt.Println("  Body       :\n", resp)
		// fmt.Println()
		r.Data = string(resp.Body())
	}

	return err

}

func (f *smsservice) SendSMS(message models.InfoBipMessages) {

	r := &models.Returnblock{}
	r.Status = true
	r.Message = "Success"

	baseUrl := f.url
	apiKey := f.auth
	recipientPhoneNumber := message.RecipientPhoneNumber
	from := message.From
	messageText := message.Text

	client := &http.Client{}

	jsonBody := fmt.Sprintf(`{"messages":[{"from":"%v","destinations":[{"to":"%v"}],"text":"%v"}]}`, from, recipientPhoneNumber, messageText)

	req, err := http.NewRequest("POST", fmt.Sprintf("%v/sms/2/text/advanced", baseUrl), strings.NewReader(jsonBody))
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Set("Authorization", fmt.Sprintf("App %v", apiKey))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		r.Status = false
		r.Message = ""
		r.Data = err
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print(string(body))

}
