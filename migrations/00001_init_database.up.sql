CREATE TYPE task_status AS ENUM ('backlog', 'open', 'progress', 'review', 'completed');

CREATE TABLE IF NOT EXISTS tasks(
            id              SERIAL          PRIMARY KEY,
            title           VARCHAR(255)    NOT NULL,
            description     TEXT,
            due_date        TIMESTAMPTZ,
            status          task_status     NOT NULL DEFAULT 'backlog',
            created_at      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
            updated_at      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
            deleted_at      TIMESTAMP
);