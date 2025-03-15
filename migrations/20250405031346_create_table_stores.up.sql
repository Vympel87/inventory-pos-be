-- ENUM
CREATE TYPE store_type_enum AS ENUM ('CENTRAL', 'BRANCH');

-- TABLE
CREATE TABLE stores (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    address TEXT,
    phone VARCHAR(50),
    email VARCHAR(255),
    store_type store_type_enum NOT NULL,
    parent_store_id UUID REFERENCES stores(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

-- TRIGGER
CREATE TRIGGER trigger_set_updated_at_stores
BEFORE UPDATE ON stores
FOR EACH ROW
EXECUTE FUNCTION set_updated_at_column();