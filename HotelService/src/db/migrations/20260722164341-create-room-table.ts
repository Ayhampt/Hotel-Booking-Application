import { QueryInterface } from "sequelize";

export default {
  async up(queryInterface: QueryInterface) {
    await queryInterface.sequelize.query(`
      CREATE TABLE IF NOT EXISTS rooms (
        id INT NOT NULL AUTO_INCREMENT,
        hotel_id INT NOT NULL,
        booking_id INT,
        room_category INT NOT NULL,
        date_of_availability DATE NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        deleted_at TIMESTAMP DEFAULT NULL,
        PRIMARY KEY (Id)
      );
      `);
  },
  async down(queryInterface: QueryInterface) {
    await queryInterface.sequelize.query(`
      DROP TABLE IF EXISTS rooms;
      `);
  },
};
