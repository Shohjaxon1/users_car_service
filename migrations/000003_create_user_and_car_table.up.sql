CREATE TABLE IF NOT EXISTS car_and_user (
    car_id UUID,
    user_id UUID,
    PRIMARY KEY (car_id, user_id),
    FOREIGN KEY (car_id) REFERENCES car(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);