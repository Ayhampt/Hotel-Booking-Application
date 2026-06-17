import logger from "../config/logger";
import { createHotelDto } from "../dto/hotel.dto";
import { NotFoundError } from "../utils/errors/app.error";
import Hotel from "./../db/models/hotel";

export async function createHotel(hotelData: createHotelDto) {
  const hotel = await Hotel.create({
    name: hotelData.name,
    address: hotelData.address,
    location: hotelData.location,
    rating: hotelData.rating,
    ratingCount: hotelData.rating_count,
  });
  logger.info(`hotel created:${hotel.id}`);
  return hotel;
}

export async function getHotelById(id: number) {
  const hotel = await Hotel.findByPk(id);
  if (!hotel) {
    logger.error(`Hotel not found ${id}`);
    throw new NotFoundError("Hotel not found");
  }
  logger.info(`Hotel found:${hotel.id}`);
  return hotel;
}

export async function getAllHotels() {
  const hotels = await Hotel.findAll({
    where: {
      deletedAt: null,
    },
  });
  if (!hotels) {
    logger.error(`No hotels available`);
    throw new NotFoundError("Hotels not found");
  }
  logger.info(`All hotels fetched`);
  return hotels;
}

export async function softDeleteHotel(id: number) {
  const hotel = await Hotel.findByPk(id);

  if (!hotel) {
    logger.error(`No hotel found:${id}`);
    throw new NotFoundError(`Hotel with ${id} not found`);
  }
  hotel.deletedAt = new Date();
  await hotel.save();
  return true;
}

export async function updateHotel(id: number, hotelData: createHotelDto) {
  const hotel = await Hotel.findByPk(id);
  if (!hotel) {
    logger.error(`No hotel found:${id}`);
    throw new NotFoundError(`Hotel with ${id} not found`);
  }
  await hotel.update({
    name: hotelData.name,
    address: hotelData.address,
    location: hotelData.location,
    rating: hotelData.rating,
    ratingCount: hotelData.rating_count,
  });
  logger.info(`Hotel updated:${hotel.id}`);
  return hotel;
}
