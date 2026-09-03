-- +goose Up
CREATE TABLE IF NOT EXISTS roles(
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  description VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

--seeder data
INSERT INTO roles (name, description) VALUES
('admin', 'Administrator role with full access'),
('user', 'Regular user role with limited access'),
('moderator', 'Moderator role with elevated privileges');


-- +goose Down
DROP TABLE IF EXISTS roles;
