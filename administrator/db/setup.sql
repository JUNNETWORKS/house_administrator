/* 
userは予約語なので使えない
https://stackoverflow.com/questions/22256124/cannot-create-a-database-table-named-user-in-postgresql
*/
CREATE TABLE accounts
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
    description text,
    owner_id INTEGER REFERENCES accounts(id),
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
    sensorType INTEGER REFERENCES SensorType(id) NOT NULL,
    created_at timestamp,
    updated_at timestamp
);

CREATE TABLE rooms_sensors
(
    roomId INTEGER REFERENCES Room(id),
    sensorId INTEGER REFERENCES Sensor(id),
    PRIMARY KEY(roomId, sensorId)
);

CREATE TABLE Measurement
(
    id SERIAL PRIMARY KEY,
    sensorId INTEGER REFERENCES Sensor(id) NOT NULL,
    value real NOT NULL,
    created_at timestamp,
    updated_at timestamp
);