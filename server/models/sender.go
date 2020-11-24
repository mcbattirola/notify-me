package models

type Sender struct {
	ID        uint   `json: "id" gorm:"primary_key"`
	IsEnabled bool   `json: "isEnabled"`
	IP        string `json: "IP" gorm:"not null"`
}

func ListSenders() []Sender {
	var senders []Sender

	DB.Find(&senders)

	return senders
}

func CreateSender(sender Sender) (Sender, error) {
	result := DB.Create(&sender)

	return sender, result.Error
}

func GetSenderById(id uint) Sender {
	sender := Sender{}
	DB.First(&sender, id)
	return sender
}
