-- Write your migrate up statements here
ALTER TABLE
  fen_pv
ADD
  COLUMN mate int;

ALTER TABLE
  fen_pv
ALTER COLUMN
  eval DROP NOT NULL;

---- tern: disable-tx ----
CREATE INDEX CONCURRENTLY unique_fen_index ON fen_pv (fen);

---- create above / drop below ----
ALTER TABLE
  fen_pv DROP COLUMN mate;

ALTER TABLE
  fen_pv
ALTER COLUMN
  eval
SET
  NOT NULL;

DROP INDEX unique_fen_index;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.