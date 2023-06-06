CREATE TABLE users
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    surname VARCHAR(255),
    phone_number VARCHAR(255)
);

CREATE TABLE accounts
(
    id SERIAL PRIMARY KEY,
    login VARCHAR(255) UNIQUE,
    password_hash VARCHAR(255)
);

CREATE TABLE roles
(
    id SERIAL PRIMARY KEY,
    role VARCHAR(255)
);

ALTER TABLE users ADD COLUMN account_id INTEGER REFERENCES accounts(id);
ALTER TABLE accounts ADD COLUMN role_id INTEGER REFERENCES roles(id);