# Player

## Get all players
* HTTP Request : ```GET http://api.com/players```
* Return a list of object in json format as follow
    ``` 
        [
            {
                "id":   int,    
                "name": string, 
                "team": string
            },...
        ]
    ```

## Create players
* HTTP Request : ```POST http://api.com/players```
* Send player's data in the request body in the follow format 
``` 
    [
        {  
            "name": string, 
            "team" string
        },...
    ]
```
* http StatusCreated (201) will be sent if the player has been created correctly
    
## Delete players
* HTTP Request : ```DELETE http://api.com/players```
* Send data in the request body in the follow format
``` 
    [
        {  
            "id":       int,
        },...
    ]
```
* http StatusOK (200) will be sent if the players have been deleted correctly

## Update players
* HTTP Request : ```PUT http://api.com/players```
* Send data in the request body in the follow format
``` 
    [
        {  
            "id":       int,
            "name":     string 
        },...
    ]
```
* http StatusCreated (201) will be sent if the player has been updated correctly

## Get a player
* HTTP Request : ```GET http://api.com/players/{ID}```
* ID is the player ID you want to get information
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