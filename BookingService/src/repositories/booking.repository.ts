import { IdempotencyKey, Prisma } from "../generated/prisma/client";
import { TransactionClient } from "../generated/prisma/internal/prismaNamespace";
import PrismaClient from "../prisma/client";
import { validate as isValidUUID } from "uuid";
import { BadRequestError, NotFoundError } from "../utils/errors/app.error";

export async function createBooking(bookingInput: Prisma.BookingCreateInput) {
  const booking = await PrismaClient.booking.create({
    data: bookingInput,
  });
  return booking;
}

export async function createIdempotencyKey(key: string, bookingId: number) {
  const idempotencyKey = await PrismaClient.idempotencyKey.create({
    data: {
      idemKey: key,
      booking: {
        connect: {
          id: bookingId,
        },
      },
    },
  });
  return idempotencyKey;
}

export async function getIdempotencyKeyWithLock(
  tx: TransactionClient,
  key: string,
) {
  if (!isValidUUID(key)) {
    throw new BadRequestError("invalid idempotency key");
  }
  const idempotencyKey: Array<IdempotencyKey> =
    await tx.$queryRaw`SELECT * FROM IdempotencyKey WHERE idemKey = ${key} FOR UPDATE ;`;

  if (!idempotencyKey || idempotencyKey.length == 0) {
    throw new NotFoundError("IdempotencyKey Not found");
  }
  return idempotencyKey[0];
}

export async function getBookingById(bookingId: number) {
  const booking = await PrismaClient.booking.findUnique({
    where: {
      id: bookingId,
    },
  });
  return booking;
}

export async function confirmBooking(
  tx: Prisma.TransactionClient,
  bookingId: number,
) {
  const booking = await tx.booking.update({
    where: {
      id: bookingId,
    },
    data: {
      status: "CONFIRMED",
    },
  });
  return booking;
}

export async function cancelBooking(bookingId: number) {
  const booking = await PrismaClient.booking.update({
    where: {
      id: bookingId,
    },
    data: {
      status: "CANCELLED",
    },
  });
  return booking;
}

export async function finalizeIdempotencyKey(
  tx: Prisma.TransactionClient,
  key: string,
) {
  const idempotencyKey = await tx.idempotencyKey.update({
    where: {
      idemKey: key,
    },
    data: {
      finalized: true,
    },
  });
  return idempotencyKey;
}
