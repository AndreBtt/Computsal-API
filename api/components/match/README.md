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
* Type is 0 if it is a group match or different then 0 if it is an elimination match

## Create previous match

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

## Update previous match

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

## Get previous match

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

## Get next matches

* HTTP Request : ```GET http://api.com/nextMatches```
* If next matches are in group phase return a list object in json format as follow ordered by time
    ``` 
    [
        {
            "id":        int,    
            "team1":     string, 
            "team2":     string,
            "type":      int,    
            "time":      time    
        },...
    ]
    ```
* time type is a string in the follow format "HH:MM:SS" where HH is hour, MM is minutes and SS seconds 


## Update next matches

* HTTP Request : ```PUT http://api.com/nextMatches```
* Send an array of next matches data in the request body in the follow format
``` 
    [
        {  
            "id":       int,
            "team1":    string,
            "team2":    string,
            "type":     int,
            "time":     int
        },...
    ]
```
* When update a group phase (type = 0) team1, team2 and time will be updated
* When update a elimination phase (type != 0) only the match time will be updated
* http StatusCreated (201) will be sent if the match has been updated correctly

## Create next matches

* HTTP Request : ```POST http://api.com/nextMatches```
* This method is only available once to create the elimination phase
* Send an array of next matches data in the request body in the follow format
```
[
	{
		"team1" : string,
		"team2" : string
	},...
]
```
* The order of the array matters, it creates the next matches based on this order
* http StatusCreated (201) will be sent if it has been completed correctly
