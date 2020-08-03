INSERT INTO users
    (id, name, created_at, updated_at)
VALUES
    (0, 'Administrator', '1900-01-01 00:00:00', '1900-01-01 00:00:00'),
    (1, 'JUN', '1999-01-08 04:05:06', '1999-01-08 04:05:06');

INSERT INTO rooms
    (name, description, owner_id, created_at, updated_at)
VALUES('JUN Room', 'this is test room', 1, '2000-01-01 00:00:00', '2000-01-01 00:00:00');