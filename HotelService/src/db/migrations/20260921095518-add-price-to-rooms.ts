import { QueryInterface } from "sequelize";

export default {
  async up(queryInterface: QueryInterface) {
    await queryInterface.sequelize.query(`
      ALTER TABLE rooms
      ADD COLUMN price DECIMAL(10, 2) NOT NULL DEFAULT 0.00
      AFTER date_of_availability;
    `);
  },

  async down(queryInterface: QueryInterface) {
    await queryInterface.sequelize.query(`
      ALTER TABLE rooms
      DROP COLUMN price;
    `);
  },
};
