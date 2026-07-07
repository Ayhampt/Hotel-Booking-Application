import { Request, Response } from "express";
import {
  createBookingService,
  confirmBookingService,
} from "../services/booking.service";
import { confirmBookingParamsDto } from "../dto/booking.dto";

export const createBookingHandler = async (req: Request, res: Response) => {
  const booking = await createBookingService(req.body);

  res.status(201).json({
    bookingId: booking.bookingId,
    idempotencyKey: booking.idempotencyKey,
  });
};

export const confirmBookingHandler = async (req: Request<confirmBookingParamsDto>, res: Response) => {
  const booking = await confirmBookingService(req.params.idempotencyKey);

  res.status(201).json({
    bookingId: booking.id,
    idempotencyKey: booking.status,
  });
};
