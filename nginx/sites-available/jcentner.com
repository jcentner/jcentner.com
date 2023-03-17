# server configuration
#

server {

	root /var/www/html;

	index index.html index.htm index.nginx-debian.html;
	server_name www.jcentner.com jcentner.com; # managed by Certbot


	location / {
		# First attempt to serve request as file, then
		# as directory, then fall back to displaying a 404.
		try_files $uri /index.html =404;
	}

	location /index.html {
		expires 30s;

	}


	listen [::]:443 ssl ipv6only=on; # managed by Certbot
	listen 443 ssl; # managed by Certbot
	ssl_certificate /etc/letsencrypt/live/jcentner.com/fullchain.pem; # managed by Certbot
	ssl_certificate_key /etc/letsencrypt/live/jcentner.com/privkey.pem; # managed by Certbot
	include /etc/letsencrypt/options-ssl-nginx.conf; # managed by Certbot
	ssl_dhparam /etc/letsencrypt/ssl-dhparams.pem; # managed by Certbot



}
server {
    if ($host = www.jcentner.com) {
        return 301 https://$host$request_uri;
    } # managed by Certbot


    if ($host = jcentner.com) {
        return 301 https://$host$request_uri;
    } # managed by Certbot


	listen 80 ;
	listen [::]:80 ;
	server_name www.jcentner.com jcentner.com;
	return 404; # managed by Certbot

}

