import { createHotelDto } from "../dto/hotel.dto";
import {
  createHotel,
  getAllHotels,
  getHotelById,
  softDeleteHotel,
  updateHotel,
} from "../repositories/hotel.repository";

export async function createHotelService(hotelData: createHotelDto) {
  const hotel = await createHotel(hotelData);
  return hotel;
}

export async function getHotelByIdService(id: number) {
  const hotel = await getHotelById(id);
  return hotel;
}

export async function getAllHotelsService() {
  const hotels = await getAllHotels();
  return hotels;
}

export async function softDeleteHotelService(id: number) {
  await softDeleteHotel(id);
  return true;
}

export async function updateHotelService(id:number,hotelData:createHotelDto) {
  const hotel = await updateHotel(id,hotelData);
  return hotel
}