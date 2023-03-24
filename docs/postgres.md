Postgres Setup

- install postgresql: sudo apt install postgresql postgresql-contrib
- login as postgres user: sudo -i -u postgres
- create role in postgres for Linux username: createuser --interactive (not in psql, but as postgres@server)
- create database: createdb [username]

Then, can connect to Postgres with 'psql' as username


