package controllers

import (
	"ReviewService/services"
)

type ReviewController struct {
	ReviewService services.ReviewService
}

func NewReviewController(_ReviewService services.ReviewService) *ReviewController {
	return &ReviewController{
		ReviewService: _ReviewService,
	}
}



