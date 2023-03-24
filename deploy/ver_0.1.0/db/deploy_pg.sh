#!/bin/bash

packages=('postgresql' 'postgresql-contrib')
self='jcentner' # my username - automate this from who ran file, maybe?
username='pgserver'
database='pgserver'
pg_hba='/etc/postgresql/14/main/pg_hba.conf'
logfile='pg-install.log'

echo "Updating server..."
sudo apt-get update -y >> $logfile

echo "Installing dependencies..."
sudo apt-get install ${packages[@]} -y >> $logfile

echo "Starting service..."
sudo systemctl start postgresql.service >> $logfile
sudo systemctl enable postgresql.service >> $logfile
sleep 5

echo "Gathering deployment scripts..."
shopt -s nullglob
initdb=(./*.sql)

echo "Creating user, role, and database..."
# for added security, make username/database a newuser (and sudo useradd newuser)
# then define and use a password for newuser with an environment variable
sudo -u postgres createuser -s $username >> $logfile
sudo -u postgres createdb $username >> $logfile
sudo echo "local $database $self trust" >> $pg_hba # for added security, use peer only for newuser
sudo pg_ctl reload
sleep 3

for file in ${initdb[@]}; do
	echo "Running $file..."
	su $self -c "cd psql -U $username -d $database -a -f $file" # add -q for quiet if logfile gets welcome messages, etc.
done

echo "Deployment done; $database database created with user $username"
