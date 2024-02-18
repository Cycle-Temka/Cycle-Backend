-- +goose Up
-- +goose StatementBegin
CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    login VARCHAR UNIQUE NOT NULL,
    name VARCHAR NOT NULL,
    hashed_password VARCHAR NOT NULL
);

CREATE TABLE admin_users(
    id SERIAL PRIMARY KEY,
    login VARCHAR UNIQUE NOT NULL,
    hashed_password VARCHAR NOT NULL
);

CREATE TABLE user_bin(
    id SERIAL PRIMARY KEY,
    cycles_view_difference_id INT NULL,
    user_id INT NOT NULL,
    accessory_id INT NULL,
    amount INT NOT NULL,
    FOREIGN KEY (cycles_view_difference_id) REFERENCES cycles_view_difference(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE wishlist(
    id SERIAL PRIMARY KEY,
    cycles_view_difference_id INT NULL,
    user_id INT NOT NULL,
    accessory_id INT NULL,
    amount INT NOT NULL,
    FOREIGN KEY (cycles_view_difference_id) REFERENCES cycles_view_difference(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (accessory_id) REFERENCES accessories(id) ON DELETE CASCADE
);
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE user_bin;

DROP TABLE wishlist;

DROP TABLE users;

DROP TABLE admin_users;
-- +goose StatementEnd
