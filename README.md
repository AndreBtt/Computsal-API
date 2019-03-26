# Computsal

## API - Endpoints

### **Players**

- **Get all players**
    * HTTP Request : ```GET http://api.com/players```
    * Return a list of object in json format as follow
        ``` 
            [
                {
                    "id":   int,    
                    "name": string, 
                    "team": string
                }
            ]
        ```

- **Create a player**
    * HTTP Request : ```POST http://api.com/players```
    * Send player's data in the request body in the follow format 
    ``` 
            {  
	            "name": string, 
	            "team" string
            }
    ```
    * http StatusCreated (201) will be sent if the player has been created correctly
    
- **Delete a player**
    * HTTP Request : ```DELETE http://api.com/players/{ID}```
    * ID is the player's ID you want to delete
    * http StatusOK (200) will be sent if the player has been deleted correctly

- **Update a player**
    * HTTP Request : ```PUT http://api.com/players```
    * Send player's data in the request body in the follow format
    ``` 
            {  
                "id":       int,
	            "name":     string, 
	            "team":     string
            }
    ```
    * It will be replace using the ID value, ensure it is correct
    * Even if you want to update just one field you need to fill all others in order to update player correctly
    * http StatusCreated (201) will be sent if the player has been updated correctly

- **Get a player**
    * HTTP Request : ```GET http://api.com/players/{ID}```
    * ID is the player's ID you want to get information
    * Return a player object in json format as follow
        ``` 
            {
                "id":            int,
                "name":          string,
                "team":          string,
                "teamPhoto":     string,
                "score":         int,
                "yellowCard":    int,
                "redCard":       int,
                "captain":       bool
            }
        ```

### **Teams**

- **Get all teams**
    * HTTP Request : ```GET http://api.com/teams```
    * Return a list of object in json format as follow
        ``` 
            [
                {
                    "id":        int,    
                    "name":      string, 
                    "photo":     string,
                    "group":     int    
                }
            ]
        ```

- **Create a team**
    * HTTP Request : ```POST http://api.com/teams```
    * Send team's data in the request body in the follow format 
    ``` 
            {  
                "name":      string,
                "photo":     string, 
                "group":     int
            }
    ```
    * http StatusCreated (201) will be sent if the team has been created correctly

- **Delete a team**
    * HTTP Request : ```DELETE http://api.com/teams/{name}```
    * Name is the team's name you want to delete
    * All players from this team will also be deleted
    * http StatusOK (200) will be sent if the team has been deleted correctly

- **Update a team**
    * HTTP Request : ```PUT http://api.com/teams```
    * Send team's data in the request body in the follow format
    ``` 
            {  
                "id":        int,
                "name":      string,
                "photo":     string, 
                "group":     int
            }
    ```
    * It will be replace using the ID value, ensure it is correct
    * Even if you want to update just one field you need to fill all others in order to update team correctly
    * http StatusCreated (201) will be sent if the team has been updated correctly

- **Get a team**
    * todo

- **Get team's players**
    * HTTP Request : ```GET http://api.com/teams/{name}/players```
    * Name is the team's name you want to get players
    * Return a list of player object in json format as follow
        ```
            [ 
                {
                    "id":    int,    
                    "name":  string, 
                    "team":  string
                },...
            ]
        ```

### **Score**

- **Get all players score**
    * HTTP Request : ```GET http://api.com/scores```
    * Return a list object in json format as follow ordered by ascending score order
        ``` 
            [
                {
                    "player" {
                        "id":    int,
                        "name":  string,
                        "team":  string
                    }
                    "score":         int,
                    "yellowCard":    int,
                    "redCard":       int
                },...
            ]
        ```

- **Create a player score**
    * HTTP Request : ```POST http://api.com/scores```
    * Send player's score data in the request body in the follow format 
    ``` 
            {  
                "id":            int,
                "playerID":      int, 
                "matchID":       int, 
                "score":         int, 
                "yellowCard":    int, 
                "redCard":       int 
            }
    ```
    * http StatusCreated (201) will be sent if the player's score has been created correctly

- **Delete a player score**
    * HTTP Request : ```DELETE http://api.com/scores/{id}```
    * ID is the player's score ID you want to delete
    * http StatusOK (200) will be sent if the team has been deleted correctly

- **Update a player score**
    * HTTP Request : ```PUT http://api.com/scores```
    * Send player's score data in the request body in the follow format
    ``` 
            {  
                "id":            int,
                "playerID":      int, 
                "matchID":       int, 
                "score":         int, 
                "yellowCard":    int, 
                "redCard":       int,
            }
    ```
    * It will be replace using the ID value, ensure it is correct
    * Even if you want to update just one field you need to fill all others in order to update correctly
    * http StatusCreated (201) will be sent if the team has been updated correctly

- **Get players score in a match**
    * HTTP Request : ```GET http://api.com/scores/{matchID}```
    * matchID is the match's ID you want to get information
    * Return a list object in json format as follow ordered by ascending score order
        ``` 
            {
                "player" {
                    "id":    int,
                    "name":  string,
                    "team":  string,
                }
                "score":         int,
                "yellowCard":    int,
                "redCard":       int
            }
        ```


### **Previous Matches**

- **Get all previous matches**
    * HTTP Request : ```GET http://api.com/previousMatches```
    * Return a list object in json format as follow
        ``` 
            {
                "id":        int,    
                "team1":     string, 
                "team2":     string,
                "score1":    int,    
                "score2":    int,    
                "type":      int,    
                "phase":     int    
            }


- **Create a previous match**
    * HTTP Request : ```POST http://api.com/previousMatches```
    * Send data in the request body in the follow format 
    ``` 
        {
            "team1" : string,
            "team2" : string,
            "type"  : int,
            "players" : [
                {
                    "player_id": int,
                    "score": int,
                    "yellowCard": int,
                    "redCard": int
                },...
            ]
        }
    ```
    * http StatusCreated (201) will be sent if the player's score has been created correctly

- **Get a previous match**
    * HTTP Request : ```GET http://api.com/previousMatches/{id}```
    * ID is the match's ID you want to get information
    * Return a list object in json format as follow
        ``` 
            {
                "id": int,
                "team1": string,
                "team2": string,
                "yellowCard1": int,
                "yellowCard2": int,
                "redCard1": int,
                "redCard2": int,
                "score1": int,
                "score2": int,
                "type": int,
                "phase": int,
                "players": [
                    {
                        "player": {
                            "id": int,
                            "name": string,
                            "team": string
                        },
                        "score": int,
                        "yellowCard": int,
                        "redCard": int
                    },...
                ]
            }
        ```