-- +goose Up

CREATE TABLE greeting (
  greeting_id uuid,
  state jsonb NOT NULL,
  CONSTRAINT greeting_pk PRIMARY KEY (greeting_id)
);

CREATE TABLE greeting_event (
  id uuid,
  greeting_id uuid NOT NULL,
  timestamp timestamptz NOT NULL,
  sequence int NOT NULL,
  data jsonb NOT NULL,
  state jsonb NOT NULL,
  CONSTRAINT greeting_event_pk PRIMARY KEY (id),
  CONSTRAINT greeting_event_fk_state FOREIGN KEY (greeting_id) REFERENCES greeting(greeting_id)
);

-- +goose Down

DROP TABLE greeting_event;
DROP TABLE greeting;

