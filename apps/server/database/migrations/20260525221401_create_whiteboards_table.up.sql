CREATE TABLE whiteboards (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    tab_id UUID NOT NULL UNIQUE
        REFERENCES room_tabs(id)
        ON DELETE CASCADE,

    scene JSONB NOT NULL DEFAULT '{}'::jsonb,

    version BIGINT NOT NULL DEFAULT 1,

    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),

    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_whiteboards_tab_id
ON whiteboards(tab_id);