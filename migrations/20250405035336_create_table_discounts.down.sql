-- DOWN for discounts
DROP TRIGGER IF EXISTS trigger_set_updated_at_discounts ON discounts;
DROP TABLE IF EXISTS discounts;
DROP TYPE IF EXISTS discount_type_enum;