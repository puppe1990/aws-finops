ALTER TABLE cost_lines ADD COLUMN usage_type TEXT NOT NULL DEFAULT '';

-- down
-- SQLite cannot drop a column in older versions; leave usage_type in place on rollback.
