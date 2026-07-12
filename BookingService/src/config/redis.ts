import Redis from "ioredis";
import { Redlock } from "@sesamecare-oss/redlock";
import { serverConfig } from ".";

const redisClient = new Redis(serverConfig.REDIS_URL);

export const redLock = new Redlock([redisClient], {
  driftFactor: 0.01,
  retryCount: 10,
  retryDelay: 200,
  retryJitter: 200,
});
