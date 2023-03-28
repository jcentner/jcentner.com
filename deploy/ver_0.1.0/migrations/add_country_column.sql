ALTER TABLE visits
ADD COLUMN visitor_country TEXT;

UPDATE visits
SET visitor_country = 'United States'
WHERE visitor_ip = '127.0.0.1';
