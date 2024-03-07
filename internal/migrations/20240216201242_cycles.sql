-- +goose Up
-- +goose StatementBegin
CREATE TABLE Cycles(
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    price INT CHECK ( price > 0 ),
    description VARCHAR NULL,
    properties JSONB
);

CREATE TABLE Cycles_view(
    id SERIAL PRIMARY KEY,
    cycle_id INT NOT NULL,
    color VARCHAR(30) CHECK ( color IN ('Bright sun', 'Concrete', 'Turquoise', 'Swamp green',
                                        'Bronze', 'Burgundy', 'Picton Blue', 'Icterine', 'India Green',
                                        'Coral', 'Camelopardalis', 'Red', 'Neon Green', 'Onyx', 'Sea Buckthorn',
                                        'Platinum', 'Falcon', 'Pink', 'X11 Gray', 'Gray', 'RYB Blue', 'Tundora',
                                        'Oxford Blue', 'Blue Violet', 'Black') ),
    picture TEXT[],
    FOREIGN KEY (cycle_id) REFERENCES Cycles(id) ON DELETE CASCADE
);

CREATE TABLE Cycles_view_difference(
    id SERIAL PRIMARY KEY,
    cycle_view_id INT NOT NULL,
    size VARCHAR(3) CHECK (size IN ('S', 'M-L', 'M', 'L', 'XL')) NOT NULL,
    amount INT NOT NULL,
    FOREIGN KEY (cycle_view_id) REFERENCES Cycles_view(id) ON DELETE CASCADE
);

CREATE TABLE accessories(
    id SERIAL PRIMARY KEY,
    type VARCHAR CHECK (type IN ('accessories', 'parking', 'equipment', 'parts')) NOT NULL,
    price INT NOT NULL,
    description VARCHAR,
    amount INT NOT NULL
);

-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE Cycles_view_difference;

DROP TABLE Cycles_view;

DROP TABLE Cycles;

DROP TABLE accessories;
-- +goose StatementEnd