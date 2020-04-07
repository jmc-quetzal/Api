
-- psql -U username -a -f filename
DROP DATABASE IF EXISTS quetzal;
CREATE DATABASE quetzal;
\c quetzal;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id uuid DEFAULT uuid_generate_v4() UNIQUE NOT NULL,
    username VARCHAR(30) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    birthdate DATE NOT NULL,
    PRIMARY KEY (id)
);