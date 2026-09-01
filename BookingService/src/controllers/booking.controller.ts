import { Request, Response } from "express";
import {
  createBookingService,
  confirmBookingService,
  getBookingByIdService,
} from "../services/booking.service";
import { confirmBookingParamsDto } from "../dto/booking.dto";

export const createBookingHandler = async (req: Request, res: Response) => {
  const booking = await createBookingService(req.body);

  res.status(201).json({
    bookingId: booking.bookingId,
    idempotencyKey: booking.idempotencyKey,
  });
};

export const confirmBookingHandler = async (
  req: Request<confirmBookingParamsDto>,
  res: Response,
) => {
  const booking = await confirmBookingService(req.params.idempotencyKey);

  res.status(201).json({
    bookingId: booking.id,
    idempotencyKey: booking.status,
  });
};

export const getBookingById = async (req: Request, res: Response) => {
  try {
    const bookingId = Number(req.params.bookingId);
    if (!Number.isInteger(bookingId) || bookingId <= 0) {
      return res.status(400).json({ message: "invalid bookingId" });
    }

    const data = await getBookingByIdService(bookingId);
    if (!data) {
      return res.status(404).json({ message: "booking not found" });
    }

    return res.status(200).json(data);
  } catch (error) {
    return res.status(500).json({ message: "internal server error" });
  }
};
