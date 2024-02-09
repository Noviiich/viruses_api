CREATE TABLE viruses
(
    virus_id SERIAL NOT NULL UNIQUE,
    virus_name VARCHAR(30) NOT NULL,
    virus_type VARCHAR(30) NOT NULL,
    infection_method VARCHAR(30) NOT NULL,
    severity VARCHAR(30) NOT NULL
);

CREATE TABLE sites
(
    site_id SERIAL NOT NULL UNIQUE,
    site_name VARCHAR(30) NOT NULL,
    hack_date TIMESTAMP NOT NULL,
    virus_id INT REFERENCES viruses(virus_id) ON DELETE CASCADE not null
);

