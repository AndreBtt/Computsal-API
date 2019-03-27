# Score

## Get all players score
* HTTP Request : ```GET http://api.com/scores```
* Return a list object in json format as follow ordered by decreasing score order
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

## Create a player score
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

## Get players score in a match
* HTTP Request : ```GET http://api.com/scores/{matchID}```
* matchID is the match's ID you want to get information
* Return a list object in json format as follow ordered by decreasing score order
    ``` 
        {
            "player" {
                "id":    int,
                "name":  string,
                "team":  string
            }
            "score":         int,
            "yellowCard":    int,
            "redCard":       int
        }
    ```