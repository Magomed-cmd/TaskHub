CREATE TYPE task_status AS ENUM ('pending', 'in_progress', 'completed', 'expired');

CREATE TYPE task_priority AS ENUM ('low', 'medium', 'high', 'critical');

CREATE TABLE IF NOT EXISTS task(
    id BIGSERIAL PRIMARY KEY,
    title varchar(255) NOT NULL,
    detail TEXT,
    status task_status NOT NULL DEFAULT 'pending',
    priority task_priority NOT NULL DEFAULT 'medium',
    due_date TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP,
    user_id BIGINT NOT NULL,
    assignee_id BIGINT,
    parent_task_id BIGINT REFERENCES task(id)
);


CREATE TABLE IF NOT EXISTS users(
    id BIGSERIAL PRIMARY KEY,
    email VARCHAR(100) NOT NULL,
    password TEXT NOT NULL,
    name varchar(20) NOT NULL
);