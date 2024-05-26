CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    full_name varchar(255) NOT NULL,
    username varchar(255) NOT NULL,
    phone_number varchar(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);