import {
  GetAvailableRoomsDto,
  UpdateBookingIdToRoomsDto,
} from "../dto/room.dto";
import { RoomRepository } from "../repositories/room.repository";

const roomRepository = new RoomRepository();

export async function getAvailableRoomsService(
  getAvailableRoomsDto: GetAvailableRoomsDto,
) {
  const rooms = await roomRepository.findByCategoryIdAndDateRange(
    getAvailableRoomsDto.roomCategoryId,
    new Date(getAvailableRoomsDto.checkInDate),
    new Date(getAvailableRoomsDto.checkOutDate),
  );

  return rooms;
}

export async function updateBookingToRoomsService(
  updateBookingToRooms: UpdateBookingIdToRoomsDto,
) {

  return await roomRepository.updateBookingIdToRooms(
    updateBookingToRooms.bookingId,
    updateBookingToRooms.roomIds,
  );
}
