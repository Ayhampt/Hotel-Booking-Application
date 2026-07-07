export type createBookingDto = {
  userId: number;
  hotelId: number;
  bookingAmount: number;
  totalGuests: number;
};

export type confirmBookingParamsDto = {
  idempotencyKey: string;
};