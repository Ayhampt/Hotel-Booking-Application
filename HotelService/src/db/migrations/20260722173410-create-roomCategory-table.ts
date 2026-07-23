import { QueryInterface } from "sequelize";

export default {
  async up(queryInterface: QueryInterface) {
    await queryInterface.sequelize.query(
      `
      CREATE TABLE IF NOT EXISTS room_categories (
        id INT NOT NULL AUTO_INCREMENT,
        hotel_id INT NOT NULL,
        price DECIMAL(10, 2) NOT NULL,
        room_type ENUM('SINGLE', 'DOUBLE','FAMILY', 'DELUXE' ,'SUITE') NOT NULL,
        room_count INT NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        deleted_at TIMESTAMP DEFAULT NULL,
        PRIMARY KEY (id)
      );`
    );
  },
  async down(queryInterface: QueryInterface) {
    await queryInterface.sequelize.query(`
      DROP TABLE IF EXISTS room_categories;
      `);
  },
};

