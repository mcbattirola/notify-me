CREATE TABLE device (
    token TEXT PRIMARY KEY,
    is_enabled BOOLEAN NOT NULL DEFAULT true
);

CREATE TABLE sender (
    ip TEXT PRIMARY KEY,
    name VARCHAR(64) NOT NULL,
    is_enabled BOOLEAN NOT NULL DEFAULT true
);

CREATE TABLE device_sender (
    device_token TEXT,
    sender_ip TEXT,
    PRIMARY KEY (device_token, sender_ip),
    CONSTRAINT fk_device_sender_device_token FOREIGN KEY (device_token) REFERENCES device(token),
    CONSTRAINT fk_device_sender_sender_ip FOREIGN KEY (sender_ip) REFERENCES sender(ip)
);

CREATE TABLE notification (
    id SERIAL PRIMARY KEY,
    title VARCHAR(150) NOT NULL,
    timestamp TIMESTAMP NOT NULL DEFAULT now(),
    content TEXT NOT NULL,
    mode VARCHAR(10)
);

CREATE TABLE device_notification (
    device_token TEXT,
    notification_id INTEGER,
    read_at TIMESTAMP,
    PRIMARY KEY (device_token, notification_id),
    CONSTRAINT fk_device_notification_device_token FOREIGN KEY (device_token) REFERENCES device(token),
    CONSTRAINT fk_device_notification_notification_id FOREIGN KEY (notification_id) REFERENCES notification(id)
);