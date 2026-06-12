import { QueryInterface } from "sequelize";
export default {
  async up(queryInterface: QueryInterface) {
    queryInterface.sequelize.query(`
      CREATE TABLE IF NOT EXISTS hotels (
        id INT NOT NULL AUTO_INCREMENT,
        name VARCHAR(255) NOT NULL,
        address VARCHAR(255) NOT NULL,
        location VARCHAR(255) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        PRIMARY KEY (id)
      );
      `);
  },

  async down(queryInterface: QueryInterface) {
    queryInterface.sequelize.query(`
      DROP TABLE IF EXISTS hotels;
      `);
  },
};
