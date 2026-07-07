import { PrismaMariaDb } from "@prisma/adapter-mariadb";
import { PrismaClient } from "../generated/prisma/client";
import { DBConfig } from "../config";

const adapter = new PrismaMariaDb({
  host: DBConfig.DB_HOST,
  user: DBConfig.DB_USER,
  password: DBConfig.DB_PASSWORD,
  database: DBConfig.DB_NAME,
  connectionLimit: 5,
});
export default new PrismaClient({ adapter });

