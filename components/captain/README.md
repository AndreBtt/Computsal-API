# Captain

## Get team's captain
* HTTP Request : ```GET http://api.com/captain/{team}```
* team is the team's name you want to get information
* Return an object in json format as follow
    ``` 
        {
            "player_id":    int,
            "email":        string
        }
    ```


## Get team from captain email
* HTTP Request : ```GET http://api.com/captainTeam/{email}```
* Return an object in json format as follow
    ``` 
        {
            "team_id":        int
        }
    ```