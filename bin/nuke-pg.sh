apt autoremove --purge postgresql* -y
rm /var/lib/postgresql -r
userdel -r postgres
whereis postgres
which psql
