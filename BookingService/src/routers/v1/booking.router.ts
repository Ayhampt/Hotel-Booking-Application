import express from "express"
import { createBookingHandler,confirmBookingHandler } from "../../controllers/booking.controller";
import { validateRequestBody } from "../../validators/index.validator";
import { bookingSchema } from "../../validators/booking.validator";

const bookingRouter = express.Router()

bookingRouter.post('/',validateRequestBody(bookingSchema),createBookingHandler);
bookingRouter.post('/confirm/:idempotencyKey',confirmBookingHandler)

export default bookingRouter