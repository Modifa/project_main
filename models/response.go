package models

type AdminResponse struct {
	ID string `db:"id"`
}

//
type ActivateResponse struct {
	ID int64 `db:"id"`
}

//
type DeleteResponse struct {
	ID int64 `db:"id"`
}

type DeleteUserResponse struct {
	Message string `db:"message"`
}

//
type ForgotPasswordResponse struct {
	ID int64 `db:"id"`
}

//
type PropertyResponse struct {
	ID int64 `db:"id"`
}
type GetPropertyResponse struct {
	Firstname    string `db:"firstname"`
	Lastname     string `db:"lastname"`
	Dateadded    string `db:"dateadded"`
	City         string `db:"city"`
	Country      string `db:"country"`
	Housenumber  int64  `db:"housenumber"`
	Postalcode   int64  `db:"postalcode"`
	Province     string `db:"province"`
	Steetname    string `db:"steetname"`
	StreetNunber int64  `db:"streetnunber"`
	Surburb      string `db:"surburb"`
}

//
type GetUserResponse struct {
	System_User_ID int64  `db:"id"`
	Title          string `db:"title"`
	FirstName      string `db:"firstname"`
	LastName       string `db:"lastname"`
	IDNumber       string `db:"id_number"`
	EmainAddress   string `db:"email_address"`
	MobileNumber   string `db:"mobile_number"`
	UserTyoe       string `db:"user_type"`
	DateAdded      string `db:"date_added"`
	Active         string `db:"active"`
	Gender         string `db:"gender"`
}

type Associations struct {
	AssociationName string `db:"a_name_"`
	Arrived         int    `db:"arrived_"`
	Expected        int    `db:"expected_"`
}
type GetAllAdminResponse struct {
	Firstname    string `db:"firstname"`
	LastName     string `db:"lastname"`
	EmailAddress string `db:"emailaddress"`
	IdNumber     string `db:"idnumber"`
	MobileNumber string `db:"mobilenumber"`
	DateAdded    string `db:"date"`
}

// firstname text,
// lastname text,
// active boolean,
// emailaddress text,
// idnumber text,
// mobilenumber text
