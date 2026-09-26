import { Request, Response } from "express";
import {
  getAvailableRoomsService,
  updateBookingToRoomsService,
} from "../services/room.service";
import { StatusCodes } from "http-status-codes";

export async function getAvailableRoomHandler(req: Request, res: Response) {
  const rooms = await getAvailableRoomsService({
    roomCategoryId: Number(req.query.roomCategoryId),
    checkInDate: req.query.checkInDate as string,
    checkOutDate: req.query.checkOutDate as string,
  });
  return res.status(StatusCodes.OK).json(rooms);
}

export async function updateBookingIdToRoomsHandler(
  req: Request,
  res: Response,
) {
  const response = await updateBookingToRoomsService(req.body);
  return res.status(StatusCodes.OK).json({
    message: "Booking Id updated to room successfully",
    data: response,
    success: true,
  });
}
