package controllers

import (
	dto "ReviewService/dto"
	"ReviewService/services"
	"ReviewService/utils"
	"fmt"
	"net/http"
)

type ReviewController struct {
	ReviewService services.ReviewService
}

func NewReviewController(_ReviewService services.ReviewService) *ReviewController {
	return &ReviewController{
		ReviewService: _ReviewService,
	}
}

func (rc *ReviewController) CreateReview(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating Review now at controller layer ")
	payload := r.Context().Value("payload").(dto.CreateReviewRequestDto)
	fmt.Println("Payload received at controller layer:", payload)
	review, err := rc.ReviewService.CreateReview(&payload)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Failed to create review", err)
		return
	}
	utils.WriteJsonSuccessResponse(w, http.StatusCreated, "Review created successfully", review)
	fmt.Println("Review created successfully ", review)

}
