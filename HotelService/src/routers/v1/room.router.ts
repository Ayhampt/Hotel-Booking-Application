import express from "express";
import {
  getAvailableRoomHandler,
  updateBookingIdToRoomsHandler,
} from "../../controllers/room.controller";
import {
  getAvailableRoomsSchema,
  updateBookingIdToRoomsSchema,
} from "../../validators/room.validator";
import { validateQueryParams, validateRequestBody } from "../../validators/index.validator";

const roomRouter = express.Router();

roomRouter.get(
  "/available",
  validateQueryParams(getAvailableRoomsSchema),
  getAvailableRoomHandler,
);
roomRouter.post(
  "/update-booking-id",
  validateRequestBody(updateBookingIdToRoomsSchema),
  updateBookingIdToRoomsHandler,
);

export default roomRouter;
