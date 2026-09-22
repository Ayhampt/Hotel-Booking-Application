import { Request, Response, NextFunction } from "express";
import { generateRooms } from "../services/roomGeneration.service";
import { StatusCodes } from "http-status-codes";

export async function generateRoomsHandler(
  req: Request,
  res: Response,
  next: NextFunction,
) {
  const result = await generateRooms(req.body);
  res.status(StatusCodes.OK).json({
    message: " Rooms generated successfully",
    data: result,
    success: true,
  });
}
