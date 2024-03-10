ALTER TABLE property_type ADD UNIQUE(name);
ALTER TABLE confectionary_type ADD UNIQUE(name);
ALTER TABLE units ADD UNIQUE(name);
ALTER TABLE packaging ADD UNIQUE(name);
ALTER TABLE city ADD UNIQUE(name);