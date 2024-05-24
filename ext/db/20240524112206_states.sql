-- +goose Up

CREATE TABLE greeting (
	id uuid PRIMARY KEY,
	state jsonb NOT NULL
);

CREATE TABLE greeting_event (
	id uuid PRIMARY KEY,
	greeting_id uuid NOT NULL REFERENCES greeting(id),
  sequence int NOT NULL,
	data jsonb NOT NULL,
  cause jsonb NOT NULL,
	timestamp timestamptz NOT NULL
);

-- +goose Down

DROP TABLE greeting_event;
DROP TABLE greeting;
