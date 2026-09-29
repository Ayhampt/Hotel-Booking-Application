# ⭐ Review Service

The Review Service is a Go-based microservice responsible for managing user feedback and hotel ratings. It is designed with strict data integrity rules, ensuring that reviews can only be created against actual, verifiable bookings, while deferring heavy cross-service updates using an eventual consistency model.

## 🏗 Key Architectural Decisions

- **Raw SQL & Goose:** Built with Go and `database/sql` (no ORM) for maximum execution speed and memory efficiency. Database migrations are managed via [Goose](https://github.com/pressly/goose).
- **Verified Stays & Anti-Spoofing:** When a review is submitted, the service ignores client-provided user or hotel IDs. Instead, it makes a synchronous HTTP call to the `BookingService` using the provided `booking_id`. If the booking exists, it extracts the true `userId` and `hotelId` from the booking response, completely preventing users from spoofing reviews for properties they never booked.
- **Eventual Consistency (`is_synced` Flag):** To prevent slowing down the review submission process, updates to a hotel's overall average rating are not performed synchronously. Newly created reviews are flagged with `is_synced = FALSE`, scaffolding a lightweight **Outbox Pattern** for a future background worker to aggregate and push rating updates to the `HotelService` asynchronously.

---

## 🗄 Database Schema (ERD)

The database schema utilizes raw SQL constraints and soft foreign keys to maintain loose coupling with other microservices.

```mermaid
erDiagram
    REVIEW {
        BIGINT id PK "AUTO_INCREMENT"
        BIGINT user_id "Soft FK -> AuthInGo"
        BIGINT hotel_id "Soft FK -> HotelService"
        BIGINT booking_id "Soft FK -> BookingService"
        INT rating "CHECK (rating >= 1 AND rating <= 5)"
        TEXT comment "Max 1000 chars"
        BOOLEAN is_synced "DEFAULT FALSE"
        TIMESTAMP created_at
        TIMESTAMP updated_at
        TIMESTAMP deleted_at "Nullable (Soft Delete)"
    }
```

---

## 🔄 The "Verified Stay" Workflow

This sequence diagram illustrates the synchronous validation pattern used to enforce authentic reviews and extract untamperable relational IDs.

```mermaid
sequenceDiagram
    participant C as Client
    participant RS as Review Service (Go)
    participant BS as Booking Service (Node)
    participant DB as MySQL (reviews)

    C->>RS: POST /reviews (booking_id, rating, comment)
    Note over RS: Payload Validated via<br/>ReviewCreateRequestValidator

    %% Verified Stay Check
    RS->>BS: GET /api/v1/booking/{booking_id}
    Note over RS: 5-second timeout threshold

    alt Booking Not Found
        BS-->>RS: 404 Not Found
        RS-->>C: 400 Bad Request (Invalid/Missing Booking)
    else Booking Exists
        BS-->>RS: 200 OK (Returns actual userId & hotelId)

        %% Anti-Spoofing Injection
        Note over RS: Overrides any client-provided IDs<br/>with authoritative Booking data
        RS->>DB: INSERT INTO reviews (userId, hotelId, is_synced: false)
        DB-->>RS: Success (Returns new review ID)
        RS-->>C: 201 Created
    end
```

---

## 📖 API Reference

### 1. Submit a Review
`POST /reviews`

Creates a new review. The client only needs to provide the `booking_id`, the `rating`, and the `comment`. The system automatically derives the `user_id` and `hotel_id` by communicating with the `BookingService`.

**Request Body:**
```json
{
  "booking_id": 42,
  "comment": "Great stay, very clean! Highly recommend the breakfast.",
  "rating": 5
}
```

**Success Response (201 Created):**
```json
{
  "message": "Review submitted successfully",
  "review_id": 108
}
```

