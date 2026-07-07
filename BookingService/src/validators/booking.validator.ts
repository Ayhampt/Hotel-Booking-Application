import { z } from "zod";

export const bookingSchema = z.object({
  userId: z.number({ message: "userId must be present" }),
  hotelId: z.number({ message: "hotelId must be present" }),
  bookingAmount: z.number({ message: "bookingAmount must be present" }),
  totalGuests: z.number({ message: "totalGuests must be present" }),
});
