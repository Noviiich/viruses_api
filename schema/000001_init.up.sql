CREATE TABLE IF NOT EXISTS users
(
    id SERIAL NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL UNIQUE,
    username VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS virus
(
    virus_id SERIAL NOT NULL UNIQUE,
    virus_name VARCHAR(255) NOT NULL,
    virus_type VARCHAR(255) NOT NULL,
    infection_method VARCHAR(255) NOT NULL,
    severity VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS site
(
    site_id SERIAL NOT NULL UNIQUE,
    site_name VARCHAR(255) NOT NULL,
    security_level VARCHAR(255),
    owner_contact VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS attack_list
(
    id SERIAL NOT NULL UNIQUE,
    site_id INT REFERENCES site(site_id) ON DELETE CASCADE NOT NULL,
    virus_id INT REFERENCES virus(virus_id) ON DELETE CASCADE NOT NULL,
    hack_date TIMESTAMP NOT NULL
);