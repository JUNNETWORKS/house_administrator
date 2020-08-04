CREATE TABLE users
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    created_at timestamp,
    updated_at timestamp
);


CREATE TABLE rooms
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    owner_id INTEGER REFERENCES users(id),
    created_at timestamp,
    updated_at timestamp
);

CREATE TABLE permitted_users
(
    room_id INTEGER REFERENCES rooms(id),
    user_id INTEGER REFERENCES users(id),
    PRIMARY KEY(room_id, sensor_id)
);


CREATE TABLE sensor_types
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    unit VARCHAR(100) NOT NULL
);

CREATE TABLE sensors
(
    id SERIAL PRIMARY KEY,
    room_id INTEGER REFERENCES rooms(id) NOT NULL,
    sensor_type_id INTEGER REFERENCES sensor_types(id) NOT NULL,
    created_at timestamp,
    updated_at timestamp
);

CREATE TABLE measurements
(
    id SERIAL PRIMARY KEY,
    sensor_id INTEGER REFERENCES sensors(id) NOT NULL,
    value real NOT NULL,
    created_at timestamp,
);

CREATE TABLE controller_types
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    created_at timestamp,
    updated_at timestamp
);

CREATE TABLE controllers
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    room_id INTEGER REFERENCES rooms(id) NOT NULL,
    controller_type_id INTEGER REFERENCES controller_types(id) NOT NULL,
    created_at timestamp,
    updated_at timestamp
);

CREATE TABLE commands
(
    id SERIAL PRIMARY KEY,
    controller_id INTEGER REFERENCES controllers(id),
    name VARCHAR(100) NOT NULL,
    created_at timestamp,
    updated_at timestamp
);