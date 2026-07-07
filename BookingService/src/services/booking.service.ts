import {
  confirmBooking,
  createBooking,
  createIdempotencyKey,
  finalizeIdempotencyKey,
  getIdempotencyKey,
} from "../repositories/booking.repository";
import { createBookingDto } from "../dto/booking.dto";
import { generateIdempotencyKey } from "../utils/generateIdempotencyKey";
import { BadRequestError, NotFoundError } from "../utils/errors/app.error";

export async function createBookingService(createBookingDto: createBookingDto) {
  const booking = await createBooking({
    userId: createBookingDto.userId,
    hotelId: createBookingDto.hotelId,
    totalGuests: createBookingDto.totalGuests,
    bookingAmount: createBookingDto.bookingAmount,
  });

  const idempotencyKey = generateIdempotencyKey();

  await createIdempotencyKey(idempotencyKey, booking.id);

  return {
    bookingId: booking.id,
    idempotencyKey: idempotencyKey,
  };
}

export async function confirmBookingService(idempotencyKey: string) {
  const idempotencyKeyData = await getIdempotencyKey(idempotencyKey);

  if (!idempotencyKeyData) {
    throw new NotFoundError("idempotencyKey not found");
  }
  if (idempotencyKeyData.finalized) {
    throw new BadRequestError("idempotencyKey already finalized");
  }

  const booking = await confirmBooking(idempotencyKeyData.bookingId);
  await finalizeIdempotencyKey(idempotencyKey);

  return booking;
}
