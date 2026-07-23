import { QueryInterface } from "sequelize";

export default {
  async up(queryInterface: QueryInterface) {
    await queryInterface.sequelize.query(`
      ALTER TABLE room_categories
      ADD CONSTRAINT fk_hotel_id
      FOREIGN KEY (hotel_id)
      REFERENCES hotels(id)
      ON DELETE CASCADE
      ON UPDATE CASCADE;
      `);
  },
  async down(queryInterface: QueryInterface) {
    await queryInterface.sequelize.query(
      `
      ALTER TABLE room_categories
      DROP FOREIGN KEY fk_hotel;
      `,
    );
  },
};
