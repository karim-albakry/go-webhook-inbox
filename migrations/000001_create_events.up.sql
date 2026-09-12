CREATE TABLE events (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,

    source TEXT NOT NULL CHECK (btrim(source) <> ''),
    event_id TEXT NOT NULL CHECK (btrim(event_id) <> ''),
    event_type TEXT NOT NULL CHECK (btrim(event_type) <> ''),

    payload JSONB NOT NULL
        CHECK (jsonb_typeof(payload) = 'object'),

    received_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT events_source_event_id_unique
        UNIQUE (source, event_id)
);