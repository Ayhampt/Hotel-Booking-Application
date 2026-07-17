import {
  confirmBooking,
  createBooking,
  createIdempotencyKey,
  finalizeIdempotencyKey,
  getIdempotencyKeyWithLock,
} from "../repositories/booking.repository";
import { createBookingDto } from "../dto/booking.dto";
import { generateIdempotencyKey } from "../utils/generateIdempotencyKey";
import {
  BadRequestError,
  InternalServerError,
  NotFoundError,
} from "../utils/errors/app.error";
import PrismaClient from "../prisma/client";
import { redLock } from "../config/redis.config";
import { serverConfig } from "../config";

export async function createBookingService(createBookingDto: createBookingDto) {
  const ttl = serverConfig.LOCK_TTL;
  const bookingResource = `hotelId:${createBookingDto.hotelId}`;

  try {
    await redLock.acquire([bookingResource], ttl);
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
  } catch (error) {
    throw new InternalServerError("Failed to acquire lock for booking resource");
  }
}

export async function confirmBookingService(idempotencyKey: string) {
  return await PrismaClient.$transaction(async (tx) => {
    const idempotencyKeyData = await getIdempotencyKeyWithLock(
      tx,
      idempotencyKey,
    );
    if (!idempotencyKeyData) {
      throw new NotFoundError("idempotencyKey not found");
    }
    if (idempotencyKeyData.finalized) {
      throw new BadRequestError("idempotencyKey already finalized");
    }

    const booking = await confirmBooking(tx, idempotencyKeyData.bookingId);
    await finalizeIdempotencyKey(tx, idempotencyKey);

    return booking;
  });
}
