import logger from "../config/logger";
import { NotFoundError } from "../utils/errors/app.error";
import Hotel from "./../db/models/hotel";
import BaseRepository from "./base.repository";

export class HotelRepository extends BaseRepository<Hotel> {
  constructor() {
    super(Hotel);
  }

  async findAlHotel() {
    const hotels = await this.model.findAll({
      where: {
        deletedAt: null,
      },
    });
    if (!hotels) {
      logger.error("no hotels found");
      throw new NotFoundError("No Hotels Found");
    }
    return hotels;
  }

  async softDeleteHotel(id: number) {
    const hotel = await this.model.findByPk(id);
    if (!hotel) {
      logger.error(`No hotel found:${id}`);
      throw new NotFoundError(`Hotel with ${id} not found`);
    }
    hotel.deletedAt = new Date();
    await hotel.save();
    return true;
  }
}
