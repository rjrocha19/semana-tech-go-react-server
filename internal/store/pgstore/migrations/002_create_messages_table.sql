CREATE TABLE IF NOT EXISTS messages (
  "id"              uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
  "room_id"         uuid NOT NULL REFERENCES rooms(id) ON DELETE CASCADE,  -- Added ON DELETE CASCADE
  "message"         varchar(255) NOT NULL,
  "reaction_count"  BIGINT NOT NULL DEFAULT 0,
  "answered"        boolean NOT NULL DEFAULT false
)
---- create above / drop below ----
DROP TABLE IF EXISTS messages;