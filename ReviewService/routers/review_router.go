package routers

import (
	"ReviewService/controllers"
	"ReviewService/middlewares"
	"github.com/go-chi/chi/v5"
)

type ReviewRouter struct {
	ReviewController *controllers.ReviewController
}

func NewReviewRouter(_reviewController *controllers.ReviewController) *ReviewRouter {
	return &ReviewRouter{
		ReviewController: _reviewController,
	}
}

func (rr *ReviewRouter) Register(r chi.Router) {
	r.With(middlewares.ReviewCreateRequestValidator).Post("/reviews", rr.ReviewController.CreateReview)
}
