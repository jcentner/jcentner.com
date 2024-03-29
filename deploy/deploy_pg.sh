#!/bin/bash

packages=('postgresql' 'postgresql-contrib')
if [ -z "$USERNAME" ] ; then
	USERNAME=$USER
fi
dburl='../db/ver_0.2.0' # convert this to take an argument --version -v 

# username='pgserver'
# database='pgserver'
# pg_hba='/etc/postgresql/14/main/pg_hba.conf'


logfile='/tmp/pg-install.log'
echo 'new log' > $logfile
chmod o+rw $logfile

echo "Updating server..."
export DEBIAN_FRONTEND=noninteractive
sudo apt-get update -y >> $logfile

echo "Installing dependencies..."
sudo apt-get install ${packages[@]} -y >> $logfile

echo "Starting service..."
sudo systemctl start postgresql.service >> $logfile
sudo systemctl enable postgresql.service >> $logfile
sleep 10

echo "Gathering deployment scripts..."
shopt -s nullglob
initdb=("$dburl"/*.sql)

echo "Creating user, role, and database..."
# for added security, make username/database a newuser (and sudo useradd newuser)
# then define and use a password for newuser with an environment variable
sudo -i -u postgres psql >> $logfile << xx
CREATE ROLE $USERNAME SUPERUSER LOGIN;
CREATE DATABASE $USERNAME;
GRANT ALL PRIVILEGES ON DATABASE $USERNAME TO $USERNAME;
xx

for file in ${initdb[@]}; do
	echo "Running $file..."
	psql -a -f $file >> $logfile
done

cp $logfile .
echo "Deployment done"

