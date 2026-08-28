package models

type Review struct {
	Id         int64
	User_id    int64
	Hotel_id   int64
	Booking_id int64
	Comment    string
	Rating     int64
	CreatedAt string
	UpdatedAt string
}
