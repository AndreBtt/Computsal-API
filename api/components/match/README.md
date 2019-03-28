# Previous Matches

## Get all previous matches

* HTTP Request : ```GET http://api.com/previousMatches```
* Return a list object in json format as follow
    ``` 
    [
        {
            "id":        int,    
            "team1":     string, 
            "team2":     string,
            "score1":    int,    
            "score2":    int,    
            "type":      int,    
            "phase":     int    
        },...
    ]
    ```

## Create a previous match

* HTTP Request : ```POST http://api.com/previousMatches```
* Send data in the request body in the follow format 
``` 
    {
        "next_match_id":    int,
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
* The match which has the "next_match_id" value will be deleted
* http StatusCreated (201) will be sent if the player's score has been created correctly

## Update a previous match

* HTTP Request : ```PUT http://api.com/previousMatches/{id}```
* ID is the match ID you want to update
* Send an array of player's score data in the request body in the follow format
``` 
    [
        {  
            "player_id":      int,
            "score":         int, 
            "yellowCard":    int, 
            "redCard":       int
        },...
    ]
```
* Even if you want to update just one field you need to fill all others in order to update correctly
* Players that don't have neither score or cards will be deleted from previous match data
* http StatusCreated (201) will be sent if the match has been updated correctly

## Get a previous match

* HTTP Request : ```GET http://api.com/previousMatches/{id}```
* ID is the match ID you want to get information
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

# Next Matches

## Get all next matches

* HTTP Request : ```GET http://api.com/nextMatches```
* Return a list object in json format as follow
    ``` 
    [
        {
            "id":        int,    
            "team1":     string, 
            "team2":     string,
            "type":      int,    
            "time":      int    
        },...
    ]
    ```

## Update next matches

* HTTP Request : ```PUT http://api.com/nextMatches```
* Send **all** next matches in the request body in the follow format
``` 
    [
        {  
            "team1":    string,
            "team2":    string,
            "type":     int,
            "time":     int
        },...
    ]
```
* It's important to send all matches because in this update specifically 
we delete all the previous data and than insert the new matches