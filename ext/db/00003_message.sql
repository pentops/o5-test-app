-- +goose Up

CREATE TABLE greeting_message (
	greeting_id uuid NOT NULL,
	event_id uuid NOT NULL,
	message_id uuid NOT NULL,
	timestamp timestamptz NOT NULL DEFAULT now()
);

-- +goose Down

DROP TABLE greeting_message;