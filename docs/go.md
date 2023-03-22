Go Setup

- Install go
- Install goimports
- go mod init
- go mod tidy for new imports
- vim-go and vim-goimports for automation
- IF USING go server to serve static content, rather than just API, configure user perms
	- create new user for go server with its own group, then add user to www-data group
	- then user has same perms as www-data
	- again: API server has no need for server group perms if not accessing index.html, images, etc	

API Server

- gin-gonic/gin for fast, simple server
- Don't serve static content, only an API server
- Call API from page JS to log access
