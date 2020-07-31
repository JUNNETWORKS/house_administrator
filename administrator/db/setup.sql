/* 
userは予約語なので使えない
https://stackoverflow.com/questions/22256124/cannot-create-a-database-table-named-user-in-postgresql
*/
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

CREATE TABLE sensor_types
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    unit VARCHAR(100) NOT NULL
);

CREATE TABLE sensors
(
    id SERIAL PRIMARY KEY,
    sensor_type_id INTEGER REFERENCES sensor_types(id) NOT NULL,
    created_at timestamp,
    updated_at timestamp
);

CREATE TABLE rooms_sensors
(
    room_id INTEGER REFERENCES rooms(id),
    sensor_id INTEGER REFERENCES sensors(id),
    PRIMARY KEY(room_id, sensor_id)
);

CREATE TABLE measurements
(
    id SERIAL PRIMARY KEY,
    sensor_id INTEGER REFERENCES sensors(id) NOT NULL,
    value real NOT NULL,
    created_at timestamp,
    updated_at timestamp
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
    description TEXT,
    controller_type_id INTEGER REFERENCES controller_types(id) NOT NULL,
    created_at timestamp,
    updated_at timestamp
);

CREATE TABLE rooms_controllers
(
    room_id INTEGER REFERENCES rooms(id),
    controller_id INTEGER REFERENCES controllers(id),
    PRIMARY KEY(room_id, controller_id)
);

CREATE TABLE commands
(
    id SERIAL PRIMARY KEY,
    controller_id INTEGER REFERENCES controllers(id),
    name VARCHAR(100) NOT NULL,
    description TEXT,
    created_at timestamp,
    updated_at timestamp
);