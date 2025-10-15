CREATE TABLE IF NOT EXISTS customers (
                                         id SERIAL PRIMARY KEY,
                                         name VARCHAR(100) NOT NULL,
    nik VARCHAR(20) NOT NULL UNIQUE,
    phone_number VARCHAR(20) NOT NULL
    );

CREATE TABLE IF NOT EXISTS cars (
                                    id SERIAL PRIMARY KEY,
                                    name VARCHAR(100) NOT NULL,
    stock INT NOT NULL DEFAULT 0,
    daily_rent NUMERIC(12,2) NOT NULL DEFAULT 0
    );

CREATE TABLE IF NOT EXISTS bookings (
                                        id SERIAL PRIMARY KEY,
                                        customer_id INT NOT NULL REFERENCES customers(id) ON DELETE CASCADE,
    car_id INT NOT NULL REFERENCES cars(id) ON DELETE CASCADE,
    start_rent DATE NOT NULL,
    end_rent DATE NOT NULL,
    total_cost NUMERIC(12,2),
    finished BOOLEAN DEFAULT FALSE
    );

-- Seed V1 customers
INSERT INTO customers (name, nik, phone_number) VALUES
                                                    ('Wawan Hermawan', '3372093912739', '081237123682'),
                                                    ('Philip Walker', '3372093912785', '081237123683'),
                                                    ('Hugo Fleming', '3372093912800', '081237123684')
ON CONFLICT DO NOTHING;

-- Seed V1 cars
INSERT INTO cars (name, stock, daily_rent) VALUES
                                               ('Toyota Camry', 2, 500000),
                                               ('Toyota Avalon', 2, 500000),
                                               ('Toyota Yaris', 2, 400000)
ON CONFLICT DO NOTHING;

