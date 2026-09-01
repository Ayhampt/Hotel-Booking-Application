package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"ReviewService/dto"
	"io"
)

const bookingPath = "/api/v1/booking/%d"

type BookingClient struct {
	BaseURL    string
	HTTPClient *http.Client
}

func NewBookingClient(baseURL string) *BookingClient {
	return &BookingClient{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (c *BookingClient) GetBooking(ctx context.Context, bookingID int64, token string) (*dto.BookingResponseDto, error) {
	url := fmt.Sprintf(c.BaseURL+bookingPath, bookingID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("booking service returned status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Print the actual response from BookingService
	fmt.Println("Booking service response:", string(body))

	var booking dto.BookingResponseDto
	if err := json.Unmarshal(body, &booking); err != nil {
		return nil, err
	}

	return &booking, nil
}
