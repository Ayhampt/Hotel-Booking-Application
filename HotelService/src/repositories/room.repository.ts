import { CreationAttributes, Op } from "sequelize";
import Room from "../db/models/room";
import BaseRepository from "./base.repository";

export class RoomRepository extends BaseRepository<Room> {
  constructor() {
    super(Room);
  }
  async findByRoomCategoryIdAndDate(roomCategoryId: number, currentDate: Date) {
    return await this.model.findOne({
      where: {
        roomCategoryId: roomCategoryId,
        dateOfAvailability: currentDate,
        deletedAt: null,
      },
    });
  }
  async bulkCreate(roomsToCreate: CreationAttributes<Room>[]) {
    return await this.model.bulkCreate(roomsToCreate);
  }

  async findLatestAvailabilityPerCategory() {
    return await this.model.findAll({
      attributes: [
        "roomCategoryId",
        [
          this.model.sequelize!.fn(
            "MAX",
            this.model.sequelize!.col("dateOfAvailability"),
          ),
          "lastDate",
        ],
      ],
      group: ["roomCategoryId"],
      raw: true,
    });
  }
  async findByCategoryIdAndDateRange(
    roomCategoryId: number,
    checkInDate: Date,
    checkOutDate: Date,
  ) {

    return await this.model.findAll({
      where: {
        roomCategoryId: roomCategoryId,
        bookingId: null,
        dateOfAvailability: {
          [Op.between]: [checkInDate, checkOutDate],
        },
      },
    });
  }

  async updateBookingIdToRooms(bookingId: number, roomIds: number[]) {
    return await this.model.update(
      { bookingId },
      {
        where: {
          id: {
            [Op.in]: roomIds,
          },
        },
      },
    );
  }
}
