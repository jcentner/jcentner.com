--------------- Create extensions ---------------

-- uuid-ossp generates uuids (uuid_generate_v4 ())

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

--------------- Create tables ---------------

-- visits table tracks each visit to each page

CREATE TABLE IF NOT EXISTS visits(
	visit_id uuid DEFAULT uuid_generate_v4 (),
	visit_timestamp TIMESTAMP DEFAULT NOW(),
	visitor_ip INET,
	page TEXT,
	referrer TEXT,
	PRIMARY KEY (visit_id)
);	
