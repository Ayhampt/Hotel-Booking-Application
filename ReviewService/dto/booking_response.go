package dto

// BookingResponseDto represents the structure of a booking response.
type BookingResponseDto struct {
	BookingId int64  `json:"id"`
	UserId    int64  `json:"userId"`
	HotelId   int64  `json:"hotelId"`
}