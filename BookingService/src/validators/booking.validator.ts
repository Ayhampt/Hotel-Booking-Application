import { z } from "zod";

export const bookingSchema = z.object({
  userId: z.number({ message: "userId must be present" }),
  hotelId: z.number({ message: "hotelId must be present" }),
  bookingAmount: z.number({ message: "bookingAmount must be present" }),
  totalGuests: z.number({ message: "totalGuests must be present" }),
  roomCategoryId: z.number({ message: "roomCategoryId must be present" }),
  checkInDate: z.string({ message: "checkInDate must be present" }),
  checkOutDate: z.string({ message: "checkOutDate must be present" }),
});
