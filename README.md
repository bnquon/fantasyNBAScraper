NBA Fantasy Scraper and Email Notifier
This project is a web application built using Next.js and Golang that scrapes NBA Fantasy player data, stores it in a MongoDB database, and sends daily emails containing the top NBA fantasy performers. The project leverages Google Cloud Functions and Cloud Scheduler for automating daily email notifications.

Link to the frontend to signup: https://fantasy-nba-scraper.vercel.app/

Features
Scrapes daily NBA Fantasy player stats from ESPN using Rod (a Golang headless browser library).
Stores player stats and email addresses in a MongoDB Atlas database.
Uses Google Cloud Functions to automatically send daily emails with the top 10 NBA Fantasy performers to subscribed users.
Next.js frontend for user email subscriptions.
Material UI for styling the frontend components.
Automatically triggered by Google Cloud Scheduler to run the scraper and send email notifications every day.
Tech Stack
Frontend: Next.js, Material UI
Backend: Golang (Rod for scraping), MongoDB Atlas (for data storage)
Automation: Google Cloud Functions, Google Cloud Scheduler
Deployment: (Assuming Vercel or Netlify for frontend, Google Cloud for serverless functions)
