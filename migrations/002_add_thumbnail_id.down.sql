-- Remove foreign key constraints
ALTER TABLE portfolio.projects DROP CONSTRAINT IF EXISTS fk_projects_thumbnail;
ALTER TABLE portfolio.articles DROP CONSTRAINT IF EXISTS fk_articles_thumbnail;

-- Remove thumbnail_id columns
ALTER TABLE portfolio.projects DROP COLUMN IF EXISTS thumbnail_id;
ALTER TABLE portfolio.articles DROP COLUMN IF EXISTS thumbnail_id;
