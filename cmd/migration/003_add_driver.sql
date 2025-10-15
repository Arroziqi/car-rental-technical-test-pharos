-- Migration 003: Create drivers table

-- 1️⃣ Create table drivers
CREATE TABLE IF NOT EXISTS drivers (
                                       id SERIAL PRIMARY KEY,
                                       name VARCHAR(100) NOT NULL,
    nik VARCHAR(20) NOT NULL UNIQUE,
    phone_number VARCHAR(20) NOT NULL,
    daily_cost NUMERIC(12,2)
    );

-- 2️⃣ Seed data
INSERT INTO drivers (name, nik, phone_number, daily_cost) VALUES
    ('Stanley Baxter', '3220132938273', '81992048712', 150000),
    ('Halsey Quinn', '3220132938293', '081992048713', 135000),
    ('Kingsley Alvarez', '3220132938313', '081992048714', 150000),
    ('Cecilia Flowers', '3220132938330', '081992048715', 155000),
    ('Clarissa Brown', '3220132938351', '081992048716', 145000),
    ('Zeph Larson', '3220132938372', '081992048717', 130000),
    ('Zach Reynolds', '3220132938375', '081992048718', 140000),
    ('Zach Reynolds', '3220132938375', '081992048718', NULL)
    ON CONFLICT (nik) DO NOTHING;
