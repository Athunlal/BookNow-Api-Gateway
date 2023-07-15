package domain

type User struct {
	Id          int64  `json:"id"`
	Username    string `json:"username"  `
	Password    string `json:"password" `
	Email       string `json:"email" `
	Phone       string `json:"phone"`
	Otp         string `json:"Otp"`
	Gender      string `json:"gender"`
	Dateofbirth string `json:"dateofbirth"`
}

type Password struct {
	Oldpassword string `json:"oldpassword"`
	Newpassword string `json:"newpassword"`
}

type Address struct {
	Addressid       uint   `JSON:"addressid" gorm:"primarykey;unique"`
	User            User   `gorm:"ForeignKey:uid"`
	Uid             uint   `JSON:"uid"`
	Type            string `JSON:"type" gorm:"not null"`
	Locationaddress string `JSON:"locationaddress" gorm:"not null"`
	Completeaddress string `JSON:"completeAddress" gorm:"not null"`
	Landmark        string `JSON:"landmark" gorm:"not null"`
	Floorno         string `JSON:"floorno" gorm:"not null"`
}
