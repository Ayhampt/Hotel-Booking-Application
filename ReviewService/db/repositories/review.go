package db

import (
	"ReviewService/models"
	"database/sql"
	"fmt"
)

type ReviewRepository interface {
	Create(user_id int64, hotel_id int64, booking_id int64, comment string, rating int64) (*models.Review, error)
}

type ReviewRepositoryImpl struct {
	db *sql.DB
}

func NewReviewRepository(_db *sql.DB) ReviewRepository {
	return &ReviewRepositoryImpl{
		db: _db,
	}
}

func (r *ReviewRepositoryImpl) Create(user_id int64, hotel_id int64, booking_id int64, comment string, rating int64) (*models.Review, error) {
	query := "INSERT INTO reviews (user_id, hotel_id, booking_id, comment, rating) VALUES(?,?,?,?,?)"
	result, err := r.db.Exec(query, user_id, hotel_id, booking_id, comment, rating)
	if err != nil {
		fmt.Println("Error in creating new review", err)
		return nil, err
	}
	lastInsertedId, rowErr := result.LastInsertId()
	if rowErr != nil {
		fmt.Println("Error getting last inserted id", rowErr)
		return nil, rowErr
	}
	review := &models.Review{
		Id:        lastInsertedId,
		UserId:    user_id,
		HotelId:   hotel_id,
		BookingId: booking_id,
		Comment:   comment,
		Rating:    rating,
		IsSynced:  false,
	}
	fmt.Println("review created successfully", review)
	return review, nil
}
