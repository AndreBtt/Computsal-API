# Previous Matches

## Get all previous matches

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


## Create a previous match

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

## Update a previous match

* HTTP Request : ```PUT http://api.com/previousMatches/{id}```
* ID is the match's ID you want to update
* Send player's score data in the request body in the follow format
``` 
        {  
            "player_id":      int,
            "score":         int, 
            "yellowCard":    int, 
            "redCard":       int
        }
```
* Even if you want to update just one field you need to fill all others in order to update correctly
* http StatusCreated (201) will be sent if the team has been updated correctly

## Get a previous match

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