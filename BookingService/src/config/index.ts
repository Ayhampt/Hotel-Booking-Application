// This file contains all the basic configuration logic for the app server to work
import dotenv from "dotenv";

type serverConfig = {
  PORT: number;
  REDIS_URL: string;
  LOCK_TTL: number;
};
type DBConfig = {
  DB_HOST: string;
  DB_USER: string;
  DB_PASSWORD: string;
  DB_NAME: string;
};
function loadEnv() {
  dotenv.config();
  console.log("environment variables loaded");
}
loadEnv();

export const serverConfig: serverConfig = {
  PORT: Number(process.env.PORT) || 8080,
  REDIS_URL: process.env.REDIS_URL || "redis://localhost:6379",
  LOCK_TTL: Number(process.env.LOCK_TTL) || 60000,
};
export const DBConfig: DBConfig = {
  DB_HOST: process.env.DB_HOST || "localhost",
  DB_USER: process.env.DB_USER || "root",
  DB_PASSWORD: process.env.DB_PASSWORD || "root",
  DB_NAME: process.env.DB_NAME || "test",
};
