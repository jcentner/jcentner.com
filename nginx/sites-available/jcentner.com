# server configuration
#

server {

	listen 80 ; # always route http to https
	# listen [::]:80 ; # enable for ipv6
	server_name www.jcentner.com jcentner.com;
	return 302 https://$host$request_uri;	

    if ($host = www.jcentner.com) {
        return 302 https://$host$request_uri;
    }

    if ($host = jcentner.com) {
        return 302 https://$host$request_uri;
    } 
}

server {

	# listen [::]:443 ssl ipv6only=on; # enable for ipv6
	listen 443 ssl; 

	# root /home/ubuntu/jcentner.com/html;
	root /var/www/html;

	index index.html;
	server_name www.jcentner.com jcentner.com;

	location / {
		try_files $uri $uri/ $uri.html /index.html;
	}

	location /index.html {
		expires 30s;
	}

	location /api {

		proxy_set_header		Host $host;
		proxy_set_header		X-Real-IP $remote_addr;
		proxy_set_header		X-Forwarded-For $proxy_add_x_forwarded_for;
		proxy_set_header		X-Forwarded-Proto $scheme;

		proxy_pass				http://localhost:9001;
		proxy_read_timeout		90;
	}


	# Certbot SSL/HTTPS termination 

	ssl_certificate /etc/letsencrypt/live/jcentner.com/fullchain.pem; # managed by Certbot
	ssl_certificate_key /etc/letsencrypt/live/jcentner.com/privkey.pem; # managed by Certbot
	include /etc/letsencrypt/options-ssl-nginx.conf; # managed by Certbot
	ssl_dhparam /etc/letsencrypt/ssl-dhparams.pem; # managed by Certbot

	access_log		/var/log/nginx/jcentner.com.access.log;
	error_log		/var/log/nginx/jcentner.com.error.log;
}
