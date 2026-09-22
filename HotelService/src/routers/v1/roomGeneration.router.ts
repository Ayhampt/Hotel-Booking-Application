import express from "express";
import { validateRequestBody } from "../../validators/index.validator";
import { roomGenerationJobSchema } from "../../dto/roomGeneration.dto";
import { generateRoomsHandler } from "../../controllers/roomGeneration.controller";

const roomGenerationRouter = express.Router();

roomGenerationRouter.post(
  "/",
  validateRequestBody(roomGenerationJobSchema),
  generateRoomsHandler,
);

export default roomGenerationRouter;
