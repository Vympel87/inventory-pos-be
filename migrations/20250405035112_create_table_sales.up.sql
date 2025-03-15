-- UP
CREATE TABLE IF NOT EXISTS sales (
    id SERIAL PRIMARY KEY,
    store_id UUID REFERENCES stores(id) ON DELETE CASCADE,
    user_id UUID REFERENCES users(id) ON DELETE SET NULL,
    sale_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    sale_number VARCHAR(50) UNIQUE NOT NULL,
    payment_method VARCHAR(50),
    total_amount DECIMAL(10,2) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_sales_sale_date ON sales (sale_date);
CREATE INDEX idx_sales_store_id ON sales (store_id);
CREATE INDEX idx_sales_user_id ON sales (user_id);

CREATE TRIGGER trigger_set_updated_at_sales
BEFORE UPDATE ON sales
FOR EACH ROW
EXECUTE FUNCTION set_updated_at_column();