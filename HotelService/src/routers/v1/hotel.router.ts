import { hotelSchema } from "../../validators/hotel.validator";
import express from "express";
import { createHotelHandler, getAllHotelsHandler,deleteHotelHandler, updateHotelHandler } from "../../controllers/hotel.controller";
import { validateRequestBody } from "../../validators/index.validator";
import { getHotelByIdHandler } from "../../controllers/hotel.controller";

const hotelRouter = express.Router();

hotelRouter.post("/", validateRequestBody(hotelSchema), createHotelHandler);
hotelRouter.get("/:id", getHotelByIdHandler);
hotelRouter.get("/",getAllHotelsHandler);
hotelRouter.delete("/:id", deleteHotelHandler);
hotelRouter.put("/:id", validateRequestBody(hotelSchema), updateHotelHandler);

export default hotelRouter;
