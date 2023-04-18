# jcentner.com
Personal website

## Description

Simple personal website to demonstrate some full-stack web dev, and to serve as a repository for projects and thoughts.

Deployed on an EC2 instance in AWS using nginx as a webserver, using go middleware to track access (with more features to come), and simple HTML/CSS/Bootstrap.js content.

Planned features and exercises:

- Full deployment documentation
- Project pages w/ descriptions, documentations, challenges, images/videos, etc.
- Rate limiting on API requests with Redis as an exercise
- AWS EBS for personal fileserver
- Thought repo: a repository of thoughts (blog, basically) for personal clarity and public critique. Markets, finance, econ, crypto, blockchain, nfts, AI/ML, cool projects, hype vs reality, WY, problems, etc.
- Add admin dashboard for reports on website visits, such as top referrers

## Deployment

Lots todo here for a clean deployment script
- Start from EC2 Ubuntu instance
- Set up deployment key from repo and DNS settings
- Pull repo
- Script: Postgres setup and conf (done)
- Script: Server start (todo)
- Script: Nginx install, conf, certbot (todo) 
