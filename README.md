# jcentner.com
Personal website

## Description

Simple personal website to demonstrate some full-stack web dev, and to serve as a repository for projects and past. 

Deployed on an EC2 instance in AWS using nginx, with go middleware to track access (with more features penciled in), and simple HTML/CSS/Bootstrap.js content.

Planned features and exercises:

- Shinier interface/frontend design
- Full deployment documentation
- Project pages w/ descriptions, documentations, challenges, images/videos, etc.
- Rate limiting on API requests with Redis as an exercise
- AWS EBS for personal fileserver
- Add admin dashboard for reports on website visits, such as top referrers

## Deployment

Some work left for a clean deployment script
- Start from EC2 Ubuntu instance
- Set up deployment key from repo and DNS settings
- Pull repo
- Script: Postgres setup and conf (done)
- Script: Server start (todo)
- Script: Nginx install, conf, certbot (todo) 
