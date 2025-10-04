CREATE SCHEMA IF NOT EXISTS portfolio;
SET search_path TO portfolio;

CREATE TABLE IF NOT EXISTS experiences (
    id              SERIAL PRIMARY KEY,
    title           VARCHAR(255) NOT NULL,
    function        VARCHAR(255) NOT NULL,
    description     TEXT,
    initial_date    DATE,
    end_date        DATE,
    actual          BOOLEAN DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS about_text (
    id          SERIAL PRIMARY KEY,
    title       VARCHAR(255),
    content     TEXT NOT NULL,

    updated_at  TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS contacts (
    id          SERIAL PRIMARY KEY,
    link        VARCHAR(255) NOT NULL,
    plataform   VARCHAR(255) NOT NULL,
    description TEXT,

    updated_at  TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS projects (
    id              SERIAL PRIMARY KEY,
    title           VARCHAR(255) NOT NULL,
    description     TEXT,
    published_at    DATE,
    revelance       INTEGER,

    updated_at      TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS links (
    id              SERIAL PRIMARY KEY,
    parent_id       INTEGER NOT NULL,
    module          VARCHAR(64) NOT NULL,
    title           VARCHAR(255) NOT NULL,
    link            VARCHAR(255) NOT NULL,
    revelance       INTEGER,
    description     TEXT,

    updated_at      TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS attachments (
    id              SERIAL PRIMARY KEY,
    parent_id       INTEGER NOT NULL,
    module          VARCHAR(64) NOT NULL,
    title           VARCHAR(255) NOT NULL,
    link            VARCHAR(255) NOT NULL,
    description     TEXT,

    updated_at      TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS skills (
    id              SERIAL PRIMARY KEY,
    title           VARCHAR(255) NOT NULL,
    description     TEXT
);

CREATE TABLE IF NOT EXISTS skills_relations (
    id              SERIAL PRIMARY KEY,
    parent_id       INTEGER NOT NULL,
    module          VARCHAR(64) NOT NULL,
    skill_id        INTEGER NOT NULL,
    revelance       INTEGER,
  
    FOREIGN KEY (skill_id) REFERENCES skills(id)
);

CREATE TABLE IF NOT EXISTS courses (
    id              SERIAL PRIMARY KEY,
    title           VARCHAR(255) NOT NULL,
    description     TEXT,
    concluded_at    DATE,
    revelance       INTEGER,

    updated_at      TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS articles (
    id              SERIAL PRIMARY KEY,
    type            VARCHAR(20) NOT NULL,
    title           VARCHAR(255) NOT NULL,
    description     TEXT,
    local           VARCHAR(255),
    published_at    DATE,
    revelance       INTEGER,

    updated_at      TIMESTAMP DEFAULT NOW()
);

