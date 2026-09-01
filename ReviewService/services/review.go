package services

import (
	client "ReviewService/client"
	db "ReviewService/db/repositories"
	dto "ReviewService/dto"
	"ReviewService/models"
	"context"
	"fmt"
)

type ReviewService interface {
	CreateReview(payload *dto.CreateReviewRequestDto) (*models.Review, error)
}

type ReviewServiceImpl struct {
	ReviewRepository db.ReviewRepository
	BookingClient    client.BookingClient
}

func NewReviewService(_ReviewRepository db.ReviewRepository, _BookingClient client.BookingClient) ReviewService {
	return &ReviewServiceImpl{
		ReviewRepository: _ReviewRepository,
		BookingClient:    _BookingClient,
	}
}

func (r *ReviewServiceImpl) CreateReview(payload *dto.CreateReviewRequestDto) (*models.Review, error) {
	fmt.Println("Creating Review now at service layer ")

	if payload.Rating < 1 || payload.Rating > 5 {
		return nil, fmt.Errorf("Rating must be between 1 and 5")
	}
	booking, err := r.BookingClient.GetBooking(context.Background(), payload.BookingId, "")
	if err != nil {
		fmt.Println("Error in fetching the booking")
		return nil, err
	}
	if booking == nil {
		fmt.Println("No booking found")
		return nil, fmt.Errorf("No booking found for the given booking ID")
	}

	review, err := r.ReviewRepository.Create(booking.BookingId, booking.HotelId, payload.BookingId, payload.Comment, int64(payload.Rating))
	if err != nil {
		fmt.Println("Error in creating review at service layer", err)
		return nil, err
	}
	fmt.Println("Review created successfully")
	return review, nil

}
