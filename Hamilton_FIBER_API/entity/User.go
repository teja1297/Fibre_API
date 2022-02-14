package entity

// type User struct {
// 	_id         string    `json:"id,omitempty"`
// 	FirstName   string    `json:"firstname,omitempty"`
// 	LastName    string    `json:"lastname,omitempty"`
// 	Email       string    `json:"email,omitempty"`
// 	Age         int       `json:"age,omitempty"`
// 	DOB         time.Time `json:"DOB,omitempty"`
// 	AccountType string    `json:"Savings,omitempty"`
// 	Gender      string    `json:"Gender,omitempty"`
// 	PAN         string    `json:"PAN,omitempty"`
// 	mobile      int64     `json:"mobile,omitempty"`
// 	Balance     int64     `json:"Balance,omitempty"`
// }

type User struct {
	Firstname   string `json:"Firstname"`
	Lastname    string `json:"Lastname"`
	Age         int    `json:"Age"`
	Email       string `json:"Email"`
	Dob         string `json:"DOB"`
	AccountType string `json:"AccountType"`
	Gender      string `json:"Gender"`
	Pan         string `json:"PAN"`
	Mobile      int64  `json:"mobile"`
	Balance     int    `json:"balance"`
}
