-- Create ENUM type first
CREATE TYPE discount_type_enum AS ENUM ('PERCENTAGE', 'FIXED');

CREATE TABLE IF NOT EXISTS discounts (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    discount_type discount_type_enum NOT NULL,
    discount_value DECIMAL(10,2) NOT NULL,
    start_date TIMESTAMP,
    end_date TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Optional trigger if you're tracking updated_at
CREATE TRIGGER trigger_set_updated_at_discounts
BEFORE UPDATE ON discounts
FOR EACH ROW
EXECUTE FUNCTION set_updated_at_column();