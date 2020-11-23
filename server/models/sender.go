package models

type Sender struct {
	ID         int  `json:"id" gorm:"primary_key"`
	Is_Enabled bool `json:"is_enabled"`
}

func ListSenders() []Sender {
	var senders []Sender

	DB.Find(&senders)

	return senders
}
