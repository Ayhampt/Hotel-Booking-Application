import {
  Model,
  InferAttributes,
  InferCreationAttributes,
  CreationOptional,
} from "sequelize";

import sequelize from "./sequelize";

class Room extends Model<InferAttributes<Room>, InferCreationAttributes<Room>> {
  declare id: CreationOptional<number>;
  declare hotelId: number;
  declare bookingId: number;
  declare roomCategoryId: number;
  declare dateOfAvailability: Date;
  declare createdAt: CreationOptional<Date>;
  declare updatedAt: CreationOptional<Date>;
  declare deletedAt: CreationOptional<Date | null>;
}

Room.init(
  {
    id: {
      type: "INTEGER",
      autoIncrement: true,
      primaryKey: true,
    },
    hotelId: {
      type: "INTEGER",
      allowNull: false,
      references: {
        model: "Hotel",
        key: "id",
      },
    },
    bookingId: {
      type: "INTEGER",
      allowNull: false,
    },
    roomCategoryId: {
      type: "INTEGER",
      allowNull: false,
      references: {
        model: "RoomCategory",
        key: "id",
      },
    },
    dateOfAvailability: {
      type: "DATE",
      allowNull: false,
    },
    updatedAt: {
      type: "DATE",
      defaultValue: new Date(),
    },
    createdAt: {
      type: "DATE",
      defaultValue: new Date(),
    },
    deletedAt: {
      type: "DATE",
      defaultValue: null,
    },
  },
  {
    tableName: "rooms",
    sequelize: sequelize,
    underscored: true,
    timestamps: true,
  },
);

export default Room;
