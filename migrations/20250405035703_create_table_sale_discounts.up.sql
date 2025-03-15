CREATE TABLE IF NOT EXISTS sale_discounts (
    id SERIAL PRIMARY KEY,
    sale_id INTEGER REFERENCES sales(id) ON DELETE CASCADE,
    discount_id INTEGER REFERENCES discounts(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Trigger for updated_at if needed
CREATE TRIGGER trigger_set_updated_at_sale_discounts
BEFORE UPDATE ON sale_discounts
FOR EACH ROW
EXECUTE FUNCTION set_updated_at_column();