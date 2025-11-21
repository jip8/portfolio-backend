-- Add thumbnail_id column to projects table
ALTER TABLE portfolio.projects ADD COLUMN IF NOT EXISTS thumbnail_id INTEGER;

-- Add thumbnail_id column to articles table
ALTER TABLE portfolio.articles ADD COLUMN IF NOT EXISTS thumbnail_id INTEGER;

-- Add foreign key constraints (optional, allows null)
ALTER TABLE portfolio.projects
    ADD CONSTRAINT fk_projects_thumbnail
    FOREIGN KEY (thumbnail_id)
    REFERENCES portfolio.attachments(id)
    ON DELETE SET NULL;

ALTER TABLE portfolio.articles
    ADD CONSTRAINT fk_articles_thumbnail
    FOREIGN KEY (thumbnail_id)
    REFERENCES portfolio.attachments(id)
    ON DELETE SET NULL;
