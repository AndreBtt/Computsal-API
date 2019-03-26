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
            }
        ]
    ```

## Create a player
* HTTP Request : ```POST http://api.com/players```
* Send player's data in the request body in the follow format 
``` 
        {  
            "name": string, 
            "team" string
        },...
```
* http StatusCreated (201) will be sent if the player has been created correctly
    
## Delete a player
* HTTP Request : ```DELETE http://api.com/players/{ID}```
* ID is the player's ID you want to delete
* http StatusOK (200) will be sent if the player has been deleted correctly

## Update a player
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

## Get a player
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