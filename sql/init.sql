CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY, username VARCHAR(50) UNIQUE NOT NULL, password VARCHAR(255) NOT NULL,
    name VARCHAR(100) NOT NULL, role VARCHAR(20) DEFAULT 'visitor' CHECK (role IN ('admin','guide','visitor','staff')),
    status VARCHAR(20) DEFAULT 'active', created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE scenic_spots (
    id BIGSERIAL PRIMARY KEY, name VARCHAR(200) NOT NULL, code VARCHAR(20) UNIQUE NOT NULL,
    description TEXT, category VARCHAR(20) DEFAULT 'natural' CHECK (category IN ('natural','cultural','theme_park')),
    address TEXT, latitude DECIMAL(10,7), longitude DECIMAL(10,7),
    images JSONB DEFAULT '[]', rating DECIMAL(3,2) DEFAULT 5.00,
    open_time TIME DEFAULT '08:00', close_time TIME DEFAULT '18:00',
    max_capacity INT DEFAULT 5000,
    status VARCHAR(20) DEFAULT 'active', created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE ticket_types (
    id BIGSERIAL PRIMARY KEY, spot_id BIGINT NOT NULL REFERENCES scenic_spots(id),
    name VARCHAR(50) NOT NULL, price DECIMAL(10,2) NOT NULL, description TEXT,
    start_time TIME DEFAULT '08:00', end_time TIME DEFAULT '18:00',
    quota INT DEFAULT 1000, sold INT DEFAULT 0,
    enabled BOOLEAN DEFAULT TRUE, created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE ticket_orders (
    id BIGSERIAL PRIMARY KEY, order_no VARCHAR(50) UNIQUE NOT NULL,
    user_id BIGINT NOT NULL REFERENCES users(id), spot_id BIGINT NOT NULL REFERENCES scenic_spots(id),
    ticket_type_id BIGINT NOT NULL REFERENCES ticket_types(id),
    quantity INT NOT NULL DEFAULT 1, total_amount DECIMAL(10,2) NOT NULL,
    visit_date DATE NOT NULL, qr_code VARCHAR(100),
    status VARCHAR(20) DEFAULT 'paid' CHECK (status IN ('paid','used','expired','refunded')),
    paid_at TIMESTAMP, used_at TIMESTAMP, created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_ticket_orders_user ON ticket_orders(user_id);
CREATE INDEX idx_ticket_orders_date ON ticket_orders(visit_date);

CREATE TABLE visitor_flow (
    id BIGSERIAL PRIMARY KEY, spot_id BIGINT NOT NULL REFERENCES scenic_spots(id),
    date DATE NOT NULL, hour INT NOT NULL CHECK (hour BETWEEN 0 AND 23),
    enter_count INT DEFAULT 0, exit_count INT DEFAULT 0, current_count INT DEFAULT 0
);

CREATE UNIQUE INDEX idx_flow_spot_hour ON visitor_flow(spot_id, date, hour);

CREATE TABLE tour_guides (
    id BIGSERIAL PRIMARY KEY, user_id BIGINT UNIQUE NOT NULL REFERENCES users(id),
    name VARCHAR(100) NOT NULL, phone VARCHAR(20),
    language VARCHAR(50) DEFAULT 'Chinese', rating DECIMAL(3,2) DEFAULT 5.00,
    tour_count INT DEFAULT 0, license_no VARCHAR(50),
    status VARCHAR(20) DEFAULT 'active', created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE guide_schedules (
    id BIGSERIAL PRIMARY KEY, guide_id BIGINT NOT NULL REFERENCES tour_guides(id),
    spot_id BIGINT NOT NULL REFERENCES scenic_spots(id),
    date DATE NOT NULL, start_time TIME NOT NULL, end_time TIME NOT NULL,
    group_size INT DEFAULT 20, status VARCHAR(20) DEFAULT 'scheduled'
);

CREATE TABLE hotels (
    id BIGSERIAL PRIMARY KEY, name VARCHAR(200) NOT NULL, address TEXT,
    star INT DEFAULT 3 CHECK (star BETWEEN 1 AND 5), phone VARCHAR(20),
    description TEXT, rating DECIMAL(3,2) DEFAULT 5.00, images JSONB DEFAULT '[]',
    status VARCHAR(20) DEFAULT 'active', created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE rooms (
    id BIGSERIAL PRIMARY KEY, hotel_id BIGINT NOT NULL REFERENCES hotels(id),
    type VARCHAR(20) NOT NULL CHECK (type IN ('single','double','suite','family')),
    price DECIMAL(10,2) NOT NULL, capacity INT DEFAULT 2,
    total_rooms INT NOT NULL DEFAULT 1, available INT NOT NULL DEFAULT 1,
    status VARCHAR(20) DEFAULT 'active'
);

CREATE TABLE room_bookings (
    id BIGSERIAL PRIMARY KEY, order_no VARCHAR(50) UNIQUE NOT NULL,
    user_id BIGINT NOT NULL REFERENCES users(id), room_id BIGINT NOT NULL REFERENCES rooms(id),
    check_in DATE NOT NULL, check_out DATE NOT NULL, guests INT DEFAULT 1,
    total_price DECIMAL(10,2) NOT NULL,
    status VARCHAR(20) DEFAULT 'confirmed' CHECK (status IN ('confirmed','checked_in','checked_out','cancelled')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE routes (
    id BIGSERIAL PRIMARY KEY, name VARCHAR(200) NOT NULL, description TEXT,
    duration VARCHAR(50), distance_km DECIMAL(6,2),
    spot_ids JSONB DEFAULT '[]', difficulty VARCHAR(10) DEFAULT 'easy' CHECK (difficulty IN ('easy','moderate','hard')),
    rating DECIMAL(3,2) DEFAULT 5.00, created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE events (
    id BIGSERIAL PRIMARY KEY, spot_id BIGINT NOT NULL REFERENCES scenic_spots(id),
    title VARCHAR(200) NOT NULL, description TEXT,
    start_date TIMESTAMP NOT NULL, end_date TIMESTAMP NOT NULL,
    location VARCHAR(200), max_participants INT DEFAULT 100,
    status VARCHAR(20) DEFAULT 'upcoming', created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE feedbacks (
    id BIGSERIAL PRIMARY KEY, user_id BIGINT NOT NULL REFERENCES users(id),
    spot_id BIGINT NOT NULL REFERENCES scenic_spots(id),
    rating INT NOT NULL CHECK (rating BETWEEN 1 AND 5), content TEXT,
    images JSONB DEFAULT '[]', created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE complaints (
    id BIGSERIAL PRIMARY KEY, user_id BIGINT NOT NULL REFERENCES users(id),
    spot_id BIGINT REFERENCES scenic_spots(id),
    type VARCHAR(30) NOT NULL, content TEXT NOT NULL,
    status VARCHAR(20) DEFAULT 'pending' CHECK (status IN ('pending','processing','resolved')),
    handler_id BIGINT REFERENCES users(id), result TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE emergencies (
    id BIGSERIAL PRIMARY KEY, spot_id BIGINT NOT NULL REFERENCES scenic_spots(id),
    type VARCHAR(20) NOT NULL CHECK (type IN ('medical','fire','lost','weather','other')),
    description TEXT, level VARCHAR(10) DEFAULT 'medium' CHECK (level IN ('low','medium','high','critical')),
    location VARCHAR(200),
    status VARCHAR(20) DEFAULT 'reported' CHECK (status IN ('reported','responding','resolved')),
    reporter_id BIGINT NOT NULL REFERENCES users(id), handler_id BIGINT REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, resolved_at TIMESTAMP
);

INSERT INTO users (username, password, name, role) VALUES
('admin', '$2a$10$dummyhash', 'Admin', 'admin');
