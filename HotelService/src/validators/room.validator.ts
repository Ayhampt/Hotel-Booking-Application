import { z } from "zod";

export const getAvailableRoomsSchema = z.object({
  roomCategoryId: z.string({ message: "Room category id must be a number" }),
  checkInDate: z.string({ message: "Check-in date must be a valid date" }),
  checkOutDate: z.string({ message: "Check-out date must be a valid date" }),
});

export const updateBookingIdToRoomsSchema = z.object({
  bookingId: z.number({ message: "Booking Id must be a number" }),
  roomIds: z.array(z.number({ message: "Room Id Id must be a number" })),
});
