import { createHotelDto } from "../dto/hotel.dto";

import { HotelRepository } from "../repositories/hotel.repository";

const hotelRepository = new HotelRepository();

export async function createHotelService(hotelData: createHotelDto) {
  const hotel = await hotelRepository.create(hotelData);
  return hotel;
}

export async function getHotelByIdService(id: number) {
  const hotel = await hotelRepository.findById(id);
  return hotel;
}

export async function getAllHotelsService() {
  const hotels = await hotelRepository.findAlHotel();
  return hotels;
}

export async function softDeleteHotelService(id: number) {
  await hotelRepository.softDeleteHotel(id);
  return true;
}

export async function updateHotelService(
  id: number,
  hotelData: createHotelDto,
) {
  const hotel = await hotelRepository.update(id, hotelData);
  return hotel;
}
