package account

import "time"

type accountResponse struct {
	Previous bool      `json:"previous"`
	Next     bool      `json:"next"`
	Page     int       `json:"page"`
	Count    int64     `json:"count"`
	Data     []User    `json:"data"`
}

type User struct {
	ID          uint      `json:"id"`
	FullName    string    `json:"full_name"`
	Username    string    `json:"username"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	CreatedAt   time.Time `json:"created_at"`
}
