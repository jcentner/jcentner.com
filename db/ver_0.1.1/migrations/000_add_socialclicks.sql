CREATE TABLE IF NOT EXISTS socialclicks(
	socialclick_id uuid DEFAULT uuid_generate_v4 (),
	socialclick_timestamp TIMESTAMP DEFAULT NOW(),
	visitor_ip INET,
	button TEXT,
	page TEXT,
	PRIMARY KEY (socialclick_id)
);
