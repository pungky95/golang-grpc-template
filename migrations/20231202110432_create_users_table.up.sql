CREATE TYPE user_status AS ENUM ('ACTIVE', 'INACTIVE');

CREATE TABLE IF NOT EXISTS users (
    id UUID DEFAULT uuid_generate_v4 () NOT NULL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ,
    name VARCHAR (300) NOT NULL,
    email VARCHAR (250),
    phone VARCHAR (30),
    password VARCHAR (250) NOT NULL,
    status user_status DEFAULT 'ACTIVE' :: user_status NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_users_deleted_at ON users (deleted_at);