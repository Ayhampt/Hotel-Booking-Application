export type GetAvailableRoomsDto = {
  roomCategoryId: number;
  checkInDate: string;
  checkOutDate: string;
};

export type UpdateBookingIdToRoomsDto = {
  bookingId: number;
  roomIds: number[];
}