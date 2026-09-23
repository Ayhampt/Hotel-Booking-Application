import {roomGenerationJob } from "../dto/roomGeneration.dto";
import { roomGenerationQueue } from "../queues/roomGeneration.queue";

export const ROOM_GENERATION_PAYLOAD = "payload-room-generation";

export const addRoomGenerationJobToQueue = async (
  payload: roomGenerationJob,
) => {
  await roomGenerationQueue.add(ROOM_GENERATION_PAYLOAD, payload);
};
