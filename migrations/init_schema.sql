CREATE SCHEMA IF NOT EXISTS portfolio;

CREATE TABLE IF NOT EXISTS experiences (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    function VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    initial_date DATE NOT NULL,
    end_date DATE NOT NULL
);
