# server configuration
#

server {

	listen 80 ; # always route http to https
	listen [::]:80 ;
	server_name www.jcentner.com jcentner.com;
	return 404;


    if ($host = www.jcentner.com) {
        return 301 https://$host$request_uri;
    }

    if ($host = jcentner.com) {
        return 301 https://$host$request_uri;
    } 
}

server {

	listen [::]:443 ssl ipv6only=on; 
	listen 443 ssl; 

	root ~/jcentner.com/html/;

	index index.html index.htm index.nginx-debian.html;
	server_name www.jcentner.com jcentner.com;


	location / {
		# First attempt to serve request as file, then
		# index, so no directory access. Should never get 404
		try_files $uri $uri.html /index.html =404;
	}

	location /index.html {
		expires 30s;
	}

	# Certbot SSL/HTTPS termination 

	ssl_certificate /etc/letsencrypt/live/jcentner.com/fullchain.pem; # managed by Certbot
	ssl_certificate_key /etc/letsencrypt/live/jcentner.com/privkey.pem; # managed by Certbot
	include /etc/letsencrypt/options-ssl-nginx.conf; # managed by Certbot
	ssl_dhparam /etc/letsencrypt/ssl-dhparams.pem; # managed by Certbot

	access_log /var/log/nginx/jcentner.com.access.log;

}
