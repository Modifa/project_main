package models

/*Newpayload ...*/
type Newpayload struct {
	Bulkid   string            `json:"bulkId"`
	Messages []InfoBipMessages `json:"messages"`
}

/*New .... */
func (n *Newpayload) New() *Newpayload {
	return &Newpayload{}
}
