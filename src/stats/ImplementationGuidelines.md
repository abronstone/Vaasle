# Implementation Guidelines

## Introduction

Please implement the following endpoints in the `src/stats/api.go` file (functions can be in a separate file to make the api.go file easier to read if you want). Please implement all of your changes on the `122-create-stats-endpoints` branch and create a PR and add any of us as reviewers when you are done.

## General Information 
See `Vaasle Architecture 2.pdf` for some slightly outdated information on our container architecture. 

### Setup
1. Please take the `secrets.env` file and place it in the `src/mongo` folder to connect to the DB 
2. Please take the `.env` file and place it in the `src/frontend` folder to login to the FE (you run likely won't need to run the FE but it would not hurt  to see what its like cli container is outdated)
3. Run the following command to start up all containers while cded into the top level directory:
```
docker-compose up --build
```

### Communication
- Please reach out to Aidan through text if you have any questions

## Endpoints To Implement

- GET `/getStats/:userId` (MVP scope)

  - Get the following relevant stats for a certain user based on the specified `userId` and present them in the specified format for the FE. 
    - Games Played 
    - Win % 
    - Most guessed word 
  - FE JSON format (as a Go struct that is in `src/structs/shared_structs.go`):
    ```
    type IndividualUserStats struct {
	    GamesPlayed     int   `json:"gamesPlayed"`
      WinPercentage   int   `json:"winPercentage"`
      mostGuessedWord string`json:"mostGuessedWord"`
    }
    ```
  - In addition to code writing code in `src/stats/api.go` and the `src/stats` directory, this will require making a new endpoint in `src/mongo/mongo.go` and writing new queries in `src/mongo/user_queries.go` and `src/mongo/game_queries.go`. 

- GET `getLeaderBoard/:maxNumUsers` (Desireable scope)
  -  Returns the stats above stats for the top `n` amount of users that is defined by the `maxNumUsers` parameter.
  - Return this as an array of `IndividualUserStats` structs 

  


