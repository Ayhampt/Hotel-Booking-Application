export type createBookingDto = {
  userId: number;
  hotelId: number;
  bookingAmount: number;
  totalGuests: number;
  roomCategoryId: number;
  checkInDate: string;
  checkOutDate: string;
};

export type confirmBookingParamsDto = {
  idempotencyKey: string;
};