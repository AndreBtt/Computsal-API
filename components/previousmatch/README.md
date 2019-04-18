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
            "photo1":    string, 
            "photo2":    string,
            "score1":    int,    
            "score2":    int,    
            "type":      int,    
            "phase":     int    
        },...
    ]
    ```
* The list is sorted from the most recent games to the oldest ones
* Type is 0 if it is a group match or different then 0 if it is an elimination match
* Type diffente then zero can be values 1,2,4,8,16,...,2<sup>k</sup> where k >= 0
* Phase indicates the round's number which that game happened
* This means that if type is 1 it's the final game, 2 is the semi-final game, and so on


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

