SET search_path TO portfolio;

ALTER TABLE projects DROP CONSTRAINT IF EXISTS fk_projects_thumbnail;
ALTER TABLE articles DROP CONSTRAINT IF EXISTS fk_articles_thumbnail;

-- Drop tables in reverse order
DROP TABLE IF EXISTS articles;
DROP TABLE IF EXISTS courses;
DROP TABLE IF EXISTS skills_relations;
DROP TABLE IF EXISTS skills;
DROP TABLE IF EXISTS attachments;
DROP TABLE IF EXISTS links;
DROP TABLE IF EXISTS projects;
DROP TABLE IF EXISTS contacts;
DROP TABLE IF EXISTS about_text;
DROP TABLE IF EXISTS experiences;
