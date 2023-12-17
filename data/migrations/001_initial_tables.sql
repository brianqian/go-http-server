-- Write your migrate up statements here
CREATE TABLE users(
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  email varchar NOT NULL UNIQUE,
  password_hash varchar NOT NULL
);

CREATE TABLE pgn_raw(
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  pgn varchar NOT NULL
);

CREATE TYPE color AS ENUM ('Black', 'White');

CREATE TABLE pgn_overview(
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  user_id uuid REFERENCES users(id),
  raw_pgn uuid REFERENCES pgn_raw(id),
  date date NOT NULL,
  color color NOT NULL,
  result varchar(4) NOT NULL,
  moves int,
  time_control varchar,
  source varchar,
  victory_type varchar,
  average_elo int
);

CREATE TABLE fen_positions(
  fen varchar PRIMARY KEY,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  classic_eval int,
  nnue_eval int,
  final_eval int,
  depth int
);

CREATE TABLE fen_pv(
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  fen varchar NOT NULL UNIQUE,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  line varchar NOT NULL,
  eval int NOT NULL,
  knodes int,
  depth int NOT NULL
);