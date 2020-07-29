/* 
userは予約語なので使えない
https://stackoverflow.com/questions/22256124/cannot-create-a-database-table-named-user-in-postgresql
*/
CREATE TABLE Account
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);


CREATE TABLE Room
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description text,
    ownerId INTEGER REFERENCES Account(id)
);

CREATE TABLE SensorType
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    unit VARCHAR(100) NOT NULL
);

CREATE TABLE Sensor
(
    id SERIAL PRIMARY KEY,
    sensorType TEXT REFERENCES SensorType(id) NOT NULL
);

CREATE TABLE SensorExtension
(
    roomId INTEGER REFERENCES Room(id),
    sensorId INTEGER REFERENCES Sensor(id),
    PRIMARY KEY(roomId, sensorId)
);

CREATE TABLE Measurement
(
    id SERIAL PRIMARY KEY,
    sensorId INTEGER REFERENCES Sensor(id) NOT NULL,
    value real NOT NULL
);