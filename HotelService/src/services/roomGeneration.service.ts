import { date } from "zod";
import RoomCategory from "../db/models/roomCategory";
import { roomGenerationJob } from "../dto/roomGeneration.dto";
import { RoomRepository } from "../repositories/room.repository";
import { RoomCategoryRepository } from "../repositories/roomCategory.repository";
import { BadRequestError, NotFoundError } from "../utils/errors/app.error";
import Room from "../db/models/room";
import { CreationAttributes } from "sequelize";
import logger from "../config/logger";

const roomCategoryRepository = new RoomCategoryRepository();
const roomRepository = new RoomRepository();

export async function generateRooms(jobData: roomGenerationJob) {
  let totalRoomsCreated = 0;
  let totalDatesProcessed = 0;

  //check if category exists or not
  const roomCategory = await roomCategoryRepository.findById(
    jobData.roomCategoryId,
  );
  if (!roomCategory) {
    throw new NotFoundError(
      `Room category with ID ${jobData.roomCategoryId} not found.`,
    );
  }

  //check if date range is valid
  const startDate = new Date(jobData.startDate);
  const endDate = new Date(jobData.endDate);

  if (startDate >= endDate) {
    throw new BadRequestError(
      `Start date ${jobData.startDate} cannot be after end date ${jobData.endDate}.`,
    );
  }
  if (startDate < new Date()) {
    throw new BadRequestError(
      `Start date ${jobData.startDate} cannot be in the past.`,
    );
  }

  const totalDays = Math.ceil(
    (endDate.getTime() - startDate.getTime()) / (1000 * 60 * 60 * 24),
  );
  logger.info(`Generating rooms for category ID ${jobData.roomCategoryId} from ${jobData.startDate} to ${jobData.endDate}. Total days: ${totalDays}`,
  );
  const batchSize = jobData.bachSize || 100;

  const currentDate = new Date(startDate);

  while (currentDate <= endDate) {
    const batchEndDate = new Date(currentDate);
    batchEndDate.setDate(batchEndDate.getDate() + batchSize);

    if (batchEndDate > endDate) {
      batchEndDate.setTime(endDate.getTime());
    }
    const batchResult = await processDateBatch(
      roomCategory,
      currentDate,
      batchEndDate,
      jobData.priceOverride,
    );
    totalRoomsCreated += batchResult.roomsCreated;
    totalDatesProcessed += batchResult.dateProcessed;

    currentDate.setTime(batchEndDate.getTime());
  }
  return {
    totalRoomsCreated,
    totalDatesProcessed,
  };
}

export async function processDateBatch(
  roomCategory: RoomCategory,
  startDate: Date,
  endDate: Date,
  priceOverride?: number,
) {
  let roomsCreated = 0;
  let dateProcessed = 0;
  const roomsToCreate: CreationAttributes<Room>[] = [];

  const currentDate = new Date(startDate);

  while (currentDate <= endDate) {
    const existingRooms = await roomRepository.findBuyRoomCategoryIdAndDate(
      roomCategory.id,
      currentDate,
    );
    if (!existingRooms) {
      roomsToCreate.push({
        hotelId: roomCategory.hotelId,
        roomCategoryId: roomCategory.id,
        dateOfAvailability: currentDate,
        price: priceOverride || roomCategory.price,
        createdAt: new Date(),
        updatedAt: new Date(),
      });
    }
    currentDate.setDate(currentDate.getDate() + 1);
    dateProcessed++;

    if (roomsToCreate.length >= 0) {
      await roomRepository.bulkCreate(roomsToCreate);
      roomsCreated += roomsToCreate.length;
    }
  }
  return {
    roomsCreated,
    dateProcessed,
  };
}
