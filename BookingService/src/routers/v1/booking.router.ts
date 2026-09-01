import express from "express"
import { createBookingHandler,confirmBookingHandler, getBookingById } from "../../controllers/booking.controller";
import { validateRequestBody } from "../../validators/index.validator";
import { bookingSchema } from "../../validators/booking.validator";

const bookingRouter = express.Router()

bookingRouter.post('/',validateRequestBody(bookingSchema),createBookingHandler);
bookingRouter.post('/confirm/:idempotencyKey',confirmBookingHandler)
bookingRouter.get('/:bookingId', getBookingById)
export default bookingRouter