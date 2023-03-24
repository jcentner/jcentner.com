Created a deployment key for the instance to use to pull the repo. Symbolic links in /etc/nginx allow the repo to be source.

github deploy key
- created with ssh-keygen
- added as deploy key in repo settings
- symlink nginx sites-available, sites-enabled to repo
- symlink /var/www/html/ to repo html
	- make sure repo /html is readable by www-data group (which is nginx)
	- directory path must have executable privs to be readable (chmod +rx /home/ubuntu)

on merge to main, pull repo from instance
- ensure ssh-agent running: eval `ssh-agent -s`

To automate (I may do this later):
- create custom shell that can only 'git pull'
- create new user that can only use that shell
- use a github action to ssh into the server as that user and pull
- also probably a good idea to plug in Circle CI on github for merges to main
