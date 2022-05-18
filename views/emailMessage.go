package views

type Messages struct {
	RegisteredMessage     string
	ActivateMessage       string
	ForgotPasswordMessage string
}

/*New ... */
func (r *Messages) New() *Messages {
	t := new(Messages)

	t.RegisteredMessage = `Welcome {FullName} \n\n\n
	
	Thank You For Registering For Project Main
	
	Please Click On The link Below To Activate Your Account`

	return t
}
