import express from "express";
import { validateRequestBody } from "../../validators/index.validator";
import { roomGenerationJobSchema } from "../../dto/roomGeneration.dto";
import { generateRoomHandler } from "../../controllers/roomGeneration.controller";

const roomGenerationRouter = express.Router();

roomGenerationRouter.post(
  "/",
  validateRequestBody(roomGenerationJobSchema),
  generateRoomHandler,
);

export default roomGenerationRouter;
