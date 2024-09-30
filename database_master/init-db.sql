DO $$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_database WHERE datname = 'medicat_db') THEN
        CREATE DATABASE medicat_db;
    END IF;
END
$$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_user WHERE usename = 'admin1') THEN
        CREATE USER admin1 WITH PASSWORD 'admin123';
    END IF;
END
$$;


GRANT ALL PRIVILEGES ON DATABASE medicat_db TO admin1;
ALTER ROLE admin1 WITH SUPERUSER;
ALTER ROLE admin1 WITH CREATEROLE;


CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    profile_picture VARCHAR(255),
    fullname VARCHAR(100) NOT NULL,
    age INTEGER,
    bio TEXT,
    is_verified BOOLEAN DEFAULT FALSE,
    github_account VARCHAR(100) UNIQUE,
    linkedin_account VARCHAR(100) UNIQUE,
    google_account VARCHAR(100) UNIQUE,
    job VARCHAR(100),
    fav_email VARCHAR(100),
    location VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE comments (
    id SERIAL PRIMARY KEY,
    post_id INTEGER REFERENCES posts(id) ON DELETE CASCADE,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE likes (
    id SERIAL PRIMARY KEY,
    post_id INTEGER REFERENCES posts(id) ON DELETE CASCADE,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    UNIQUE (post_id, user_id)
);


CREATE TABLE friendships (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    friend_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (user_id, friend_id)
);


CREATE TABLE messages (
    id SERIAL PRIMARY KEY,
    friendship_id INTEGER REFERENCES friendships(id),
    user_id INTEGER REFERENCES users(id),
    friend_id INTEGER REFERENCES users(id),
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE blacklist (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    email VARCHAR(100) REFERENCES users(email),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
