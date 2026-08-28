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

func (rs *ReviewServiceImpl) CreateReview(payload *dto.CreateReviewRequestDto) (*models.Review, error) {
	fmt.Println("Creating Review now at service layer ")

	review, err := rs.ReviewRepository.Create(payload.User_id, payload.Hotel_id, payload.Booking_id, payload.Comment, payload.Rating)
	if err != nil {
		fmt.Println("error in passing payload to repository layer", err)
		return nil, err
	}
	return review, nil

}
