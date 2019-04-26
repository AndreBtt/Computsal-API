# Team

## Get all teams
* HTTP Request : ```GET http://api.com/teams```
* Return a list of object in json format as follow
    ``` 
        [
            {
                "id":        int,    
                "name":      string, 
                "photo":     string,
                "group":     int    
            },...
        ]
    ```
* http StatusOK (200) will be sent if the team has been sent correctly

## Create team
* HTTP Request : ```POST http://api.com/teams```
* Send team's data in the request body in the follow format 
``` 
    {
        "name":     string,
        "photo":    string,
        "players": 
            [
                {
                    "name" : string		
                },...
            ],
        "captain_email" : string
    }
```
* Send at least one player which is the captain
* The first player will be the captain
* Players need a unique name
* http StatusCreated (201) will be sent if the team has been created correctly

## Delete a team
* HTTP Request : ```DELETE http://api.com/teams/{name}```
* Name is the team's name you want to delete
* All players from this team will also be deleted
* http StatusOK (200) will be sent if the team has been deleted correctly

## Update team
* HTTP Request : ```PUT http://api.com/teams```
* Send team's data in the request body in the follow format
``` 
        {  
            "id":        int,
            "name":      string,
            "photo":     string
            "players": [
                {  
                    "id":       int,
                    "name":     string 
                },...
            ]
        }
```
* Even if you want to update just one field you need to fill all others in order to update team correctly
* http StatusCreated (201) will be sent if the team has been updated correctly

## Get a team
* HTTP Request : ```GET http://api.com/teams/{teamName}```
* teamName is the team's name you want to get information
* Return a object in json format as follow
    ``` 
        {
            "id":               int,
            "name":             string,
            "photo":            string,
            "group":            int,
            "win":              int,
            "lose":             int,
            "draw":             int,
            "goals_pro":        int,
            "goals_against":    int,
            "next_match": {
                "name": string,
                "time": time
            },
            "captain":  string,
            "players": [
                {
                    "id":           int,
                    "name":         string,
                    "score":        int,
                    "yellowCard":   int,
                    "redCard":      int
                },...
            ],
            "previous_matches": [
                {
                    "id":       int,
                    "team1":    string,
                    "team2":    string,
                    "score1":   int,
                    "score2":   int,
                    "type":     int,
                    "phase":    int
                },...
            ]
        }
    ```
* time type is a string in the follow format "HH:MM:SS" where HH is hour, MM is minutes and SS seconds
* http StatusOK (200) will be sent if the team has been sent correctly
