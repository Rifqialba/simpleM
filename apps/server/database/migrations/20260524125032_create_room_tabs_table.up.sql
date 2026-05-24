CREATE TABLE room_tabs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    room_id UUID NOT NULL REFERENCES rooms(id) ON DELETE CASCADE,

    created_by UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,

    type TEXT NOT NULL,

    title TEXT NOT NULL,

    position INTEGER NOT NULL DEFAULT 0,

    is_active BOOLEAN NOT NULL DEFAULT false,

    metadata JSONB,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_room_tabs_room_id
ON room_tabs(room_id);

CREATE INDEX idx_room_tabs_created_by
ON room_tabs(created_by);