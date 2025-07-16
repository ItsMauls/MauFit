-- Buat file ini: /db/migrations/000001_create_users_table.up.sql

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(16) NOT NULL DEFAULT 'member',
    fingerprint_id VARCHAR(255) UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS user_profiles (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL UNIQUE REFERENCES users(id) ON DELETE CASCADE,
    photo_profile_url TEXT,
    address TEXT,
    phone VARCHAR(32),
    bio TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Buat file ini: /db/migrations/000001_create_users_table.down.sql
-- File 'down' ini untuk membatalkan migrasi jika diperlukan.

DROP TABLE IF EXISTS user_profiles;
DROP TABLE IF EXISTS users;
