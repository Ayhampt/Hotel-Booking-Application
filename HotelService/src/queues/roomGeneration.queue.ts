import { Queue } from "bullmq";
import { getRedisConnectionObject } from "../config/redis.config";

export const ROOM_GENERATION_QUEUE = "queue-room-generation";

export const roomGenerationQueue = new Queue(ROOM_GENERATION_QUEUE, {
  connection: getRedisConnectionObject(),
});
