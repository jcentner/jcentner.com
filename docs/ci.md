To enable CI/CD for jcentner.com, created a deployment key for the instance to use to pull the repo. Symbolic links allow the repo to be source.
To automate, create custom shell for a new user. The shell only supports 'git pull' and the user can only access that shell. Then, use a github action to ssh into the server as that user and pull.

github deploy key
- created with ssh-keygen
- added as deploy key in repo settings
- 

on merge to main, pull repo from instance
- ensure ssh-agent running: eval `ssh-agent -s`
