-- DOWN
DROP TRIGGER IF EXISTS trigger_set_updated_at_sales ON sales;
DROP INDEX IF EXISTS idx_sales_sale_date;
DROP INDEX IF EXISTS idx_sales_store_id;
DROP INDEX IF EXISTS idx_sales_user_id;
DROP TABLE IF EXISTS sales;