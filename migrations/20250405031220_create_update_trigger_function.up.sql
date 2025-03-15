-- Update trigger function to only change updated_at if data changes
CREATE OR REPLACE FUNCTION set_updated_at_column() 
RETURNS TRIGGER AS $$
BEGIN
  IF ROW(NEW.*) IS DISTINCT FROM ROW(OLD.*) THEN
    NEW.updated_at = NOW();
  END IF;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;