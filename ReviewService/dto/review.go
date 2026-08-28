package dto

type CreateReviewRequestDto struct {
	User_id    int64  `json:"user_id" validate:"required"`
	Hotel_id   int64  `json:"hotel_id" validate:"required"`
	Booking_id int64  `json:"booking_id" validate:"required"`
	Comment    string `json:"comment" validate:"required"`
	Rating     int64  `json:"rating" validate:"required"`
}
