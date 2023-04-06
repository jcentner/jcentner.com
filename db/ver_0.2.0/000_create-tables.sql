-- version 0.2.0

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

-- ip_data table tracks geolocation info + some others by ip

CREATE TABLE IF NOT EXISTS ip_data(
	ip_data_uuid uuid DEFAULT uuid_generate_v4 (),
	-- populate this 

	PRIMARY KEY (ip_data_uuid)
);
