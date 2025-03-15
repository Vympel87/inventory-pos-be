-- DOWN
DROP TRIGGER IF EXISTS trigger_set_updated_at_sale_items ON sale_items;
DROP INDEX IF EXISTS idx_sale_items_sale_id;
DROP INDEX IF EXISTS idx_sale_items_product_id;
DROP TABLE IF EXISTS sale_items;