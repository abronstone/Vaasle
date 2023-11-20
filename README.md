# Vaasle: Enhanced Wordle 
**Carleton College CS 347 (Advanced Software Design) - Final Project**

🌟 **Elevate Your Wordle Experience with Vaasle** 🌟

Dive into the world of Wordle solo or with friends with Vaasle, a new take on the classic game of Wordle! 

🚀 **What Sets Vaasle Apart?**
- **Multiplayer Functionality:** Play alongside friends in real-time, adding a dynamic social element to the classic Wordle challenge.
- **Game Analytics:** Delve into fascinating game statistics, gaining insights into your gameplay and enhancing your word puzzle strategies.
- **Modern Tech Stack:** Crafted with an advanced and robust tech stack including Go, React, JavaScript, and MongoDB.
- **Microservices Architecture:** Designed for resilience, Vaasle is built on a microservices framework, ensuring smooth and scalable performance.


**Made with ❤️ by team Vaas**

- [**V**arun Saini](https://github.com/VarunSaini02)
- [**A**aron Bronstone](https://github.com/abronstone)
- [**A**idan Roessler](https://github.com/AidanRoessler)
- [**S**erafín Patiño](https://github.com/spatino1234)

*Thank you to [Peter Kelly](https://github.com/pqk314) for implementing our [Stats API](/src/stats/api.go)!*


## Running Locally
1. Before running the application locally, ask for a `secrets.env` and a `.env.local` file from application owners, and add them to the `src/mongo` and `src/frontend` directories respectively. These contain the database and Auth0 credentials.
2. To start the containers, make sure you have [Docker Desktop](https://docs.docker.com/engine/install/) (or equivalent) installed, and run the following command from the root directory of the project:
```
docker compose up --build
```
3. Go to http://localhost:3000 on your browser to play Wordle!

## Navigating/Playing Vaasle
1. You will be prompted to log in, and you can either play a singleplayer or multiplayer game.
2. For multiplayer games, ensure that each player is logged in to a different account.
3. To test multiplayer locally, you can use two different browsers or windows (for example, one in Incognito/Private mode).
