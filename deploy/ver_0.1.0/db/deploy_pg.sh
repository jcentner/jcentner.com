#!/bin/bash

packages=('postgresql' 'postgresql-contrib')
username='pgserver'
database='pgserver'
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

echo "Creating roles and database..."
sudo -u postgres createuser -s $username >> $logfile
sudo -u postgres createdb $username >> $logfile

for file in ${initdb[@]}; do
	echo "Running $file..."
	psql -U $username -d $database -a -o $logfile -f file # add -q for quiet if logfile gets welcome messages, etc. 
done

echo "Deployment done; $database database created with user $username"
