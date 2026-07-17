import Redis from "ioredis";
import { Redlock } from "@sesamecare-oss/redlock";
import { serverConfig } from ".";

function connectToRedis() {
  try {
    let connection: Redis;
    return () => {
      if (!connection) {
        connection = new Redis(serverConfig.REDIS_URL);
        return connection;
      }
      return connection;
    };
  } catch (error) {
    console.log("Error to connecting to Redis");
    throw error;
  }
}

export const getRedisConnectionObject = connectToRedis();

export const redLock = new Redlock([getRedisConnectionObject()], {
  driftFactor: 0.01,
  retryCount: 10,
  retryDelay: 200,
  retryJitter: 200,
});
