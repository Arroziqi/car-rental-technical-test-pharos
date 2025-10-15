-- Migration 002: Add memberships table and relation to customers

-- 1️⃣ Create memberships table
CREATE TABLE IF NOT EXISTS memberships (
                                           id SERIAL PRIMARY KEY,
                                           name VARCHAR(50) NOT NULL UNIQUE,
                                           discount_rate NUMERIC(5,2) NOT NULL
);

-- 2️⃣ Add nullable membership_id column to customers
ALTER TABLE customers
    ADD COLUMN IF NOT EXISTS membership_id INT NULL;

-- 3️⃣ Add foreign key constraint (ON DELETE SET NULL)
ALTER TABLE customers
    ADD CONSTRAINT fk_customers_membership
        FOREIGN KEY (membership_id) REFERENCES memberships(id)
            ON DELETE SET NULL;

-- 4️⃣ Seed membership data
INSERT INTO memberships (name, discount_rate) VALUES
                                                  ('Bronze', 4.00),
                                                  ('Silver', 7.00),
                                                  ('Gold', 15.00)
ON CONFLICT (name) DO NOTHING;