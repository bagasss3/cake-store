-- +goose Up
CREATE TABLE IF NOT EXISTS cakes (
  id INT AUTO_INCREMENT PRIMARY KEY,
  title VARCHAR(60) NOT NULL,
  description TEXT,
  rating FLOAT,
  image TEXT,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  deleted_at timestamp NULL
);

-- +goose Down
DROP TABLE IF EXISTS cakes;
