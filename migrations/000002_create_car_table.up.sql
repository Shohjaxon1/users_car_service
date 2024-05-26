CREATE TABLE IF NOT EXISTS car (
    id UUID PRIMARY KEY,
    model varchar(255) NOT NULL,
    color varchar(255) NOT NULL,
    year DATE NOT NULL,
    price FLOAT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);