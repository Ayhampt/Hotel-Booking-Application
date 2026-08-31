package services

import (
	db "ReviewService/db/repositories"
	dto "ReviewService/dto"
	"ReviewService/models"
	"fmt"
)

type ReviewService interface {
	CreateReview(payload *dto.CreateReviewRequestDto) (*models.Review, error)
}

type ReviewServiceImpl struct {
	ReviewRepository db.ReviewRepository
}

func NewReviewService(_ReviewRepository db.ReviewRepository) ReviewService {
	return &ReviewServiceImpl{
		ReviewRepository: _ReviewRepository,
	}
}

func (r *ReviewServiceImpl) CreateReview(payload *dto.CreateReviewRequestDto) (*models.Review, error) {
	fmt.Println("Creating Review now at service layer ")

	if payload.Rating < 1 || payload.Rating > 5 {
		return nil, fmt.Errorf("Rating must be between 1 and 5")
	}
	review, err := r.ReviewRepository.Create(payload.UserId, payload.HotelId, payload.BookingId, payload.Comment, int64(payload.Rating))
	if err != nil {
		fmt.Println("Error in creating review at service layer", err)
		return nil, err
	}
	fmt.Println("Review created successfully")
	return review, nil

}
