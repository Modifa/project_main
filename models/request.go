package models

import (
	"Modifa2/utils"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
)

//Admin Request
type AdminRequest struct {
	First_Name     string `json:"first_name"`
	Last_Name      string `json:"last_name"`
	Id_Number      string `json:"id_number"`
	Mobile_Number  string `json:"mobile_number"`
	Email_Address  string `json:"email_address"`
	Fk_UserType_Id string `json:"fk_usertype_id"`
	Password_      string `json:"password"`
	Title          string `json:"title"`
	Gender         string `json:"gender"`
}

func (a AdminRequest) AdminValidate() error {

	e := validation.ValidateStruct(&a,
		validation.Field(&a.Title, validation.Required),
		validation.Field(&a.First_Name, validation.Required),
		validation.Field(&a.Last_Name, validation.Required),
		validation.Field(&a.Id_Number, validation.Required),
		validation.Field(&a.Email_Address, validation.Required),
		validation.Field(&a.Mobile_Number, validation.Required, validation.Length(10, 11), validation.By(utils.IsFormatMobileNumber)),
		validation.Field(&a.Fk_UserType_Id, validation.Required),
		validation.Field(&a.Password_, validation.Required))
	if e, ok := e.(validation.InternalError); ok {
		// an internal error happened
		fmt.Println(e.InternalError())
	}

	return e
}

//Activate Account
type ActicateRequest struct {
	Id_Number string `json:"Id_Number"`
}

//Delete Account
type DeleteAccount struct {
	DeletedBy int64  `db:"deletedBy"`
	Id_Number string `db:"id_number"`
}

//Forgot Password
type ForgotPassword struct {
	IDNumber      string `json:"id_Number_"`
	Email_Address string `json:"email_address"`
	Password      string `json:"passoword"`
}

func (a ForgotPassword) ForgotPasswordValidate() error {

	e := validation.ValidateStruct(&a,
		validation.Field(&a.Email_Address, validation.Required),
		validation.Field(&a.Password, validation.Required))
	if e, ok := e.(validation.InternalError); ok {
		// an internal error happened
		fmt.Println(e.InternalError())
	}

	return e
}

// Email Request
type Messages struct {
	// Destinations []Destinations `json:"destinations,omitempty"`
	Text string `json:"text"`
	To   string `json:"to"`
}

type Destinations struct {
	To        string `json:"to"`
	MessageId string `json:"messageId"`
}

/*InfoBipMessages ...*/
type InfoBipMessages struct {
	From                 string `json:"from"`
	RecipientPhoneNumber string `json:"recipientPhoneNumber"`
	To                   string `json:"to"`
	Text                 string `json:"text"`
}

/* Add New Property ...*/
type PropertyRequest struct {
	PostalCode     int64  `json:"postalcode"`
	HouseNumber    int64  `json:"housenumber"`
	StreetNumber   int64  `json:"streetnumber"`
	StreetName     string `json:"streetname"`
	Surburb        string `json:"surburb_"`
	City           string `json:"city_"`
	Province       string `json:"province_"`
	Country        string `json:"country_"`
	System_User_Id int64  `json:"fk_system_user_id_"`
	AddedBy        int64  `json:"added_by"`
}

func (a PropertyRequest) PropertyValidate() error {

	e := validation.ValidateStruct(&a,
		validation.Field(&a.PostalCode, validation.Required),
		validation.Field(&a.HouseNumber, validation.Required),
		validation.Field(&a.StreetName, validation.Required),
		validation.Field(&a.Surburb, validation.Required),
		validation.Field(&a.City, validation.Required),
		validation.Field(&a.Province, validation.Required),
		validation.Field(&a.Country, validation.Required),
		validation.Field(&a.System_User_Id, validation.Required),
		validation.Field(&a.AddedBy, validation.Required))
	if e, ok := e.(validation.InternalError); ok {
		// an internal error happened
		fmt.Println(e.InternalError())
	}

	return e
}

/* Get Property ...*/
type GetPropertyRequest struct {
	Id_Number string `db:"id_number_"`
}

/* Get All Properties ...*/
type GetPropertyAllRequest struct {
	Functionname string `json:"function_name"`
}

/* Delete Property ...*/
type DeleteAllRequest struct {
	HouseNumber int64 `json:"housenumber"`
	Deleted_By_ int64 `json:"deleted_by"`
}

/* Delete Property ...*/
type DeleteUserRequest struct {
	IdNumber    string `json:"IdNumber"`
	Deleted_By_ int64  `json:"deleted_by"`
}

/* Get User ...*/
type GetUserRequest struct {
	Email     string `json:"email"`
	Password_ string `json:"password_"`
}

func (a GetUserRequest) UserValidate() error {

	e := validation.ValidateStruct(&a,
		validation.Field(&a.Email, validation.Required),
		validation.Field(&a.Password_, validation.Required))
	if e, ok := e.(validation.InternalError); ok {
		// an internal error happened
		fmt.Println(e.InternalError())
	}

	return e
}

//use function
type GetAllType struct {
}

type EmptyGetter struct {
	FucntionName string `json:"fucntion_name"`
}
