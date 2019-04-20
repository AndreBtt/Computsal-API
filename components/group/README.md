# Group

## Get all groups

* HTTP Request : ```GET http://api.com/groups```
* Return a list object in json format as follow
    ``` 
    [
        {
            "group_number":    int    
        },...
    ]
    ```

## Get groups detail

* HTTP Request : ```GET http://api.com/groupsDetail```
* Return a list object in json format as follow
    ``` 
    [
        {
            "group_number": int,
            "teams": [
                {
                    "id":               int,
                    "name":             string,
                    "photo":            string,
                    "win":              int,
                    "lose":             int,
                    "draw":             int,
                    "goals_pro":        int,
                    "goals_against":    int,
                    "points":           int
                },...
            ]
        }
    ]
    ```
* Teams are ordered first by points than winnings, draws, goals_pro and finaly goals_against

## Update a group

* HTTP Request : ```PUT http://api.com/groups/{groupNumber}```
* Send an array of teams in the request body in the follow format
``` 
    [
        {  
            "team_name":     string,
            "action":   int
        },...
    ]
```
* Action is 1 if you want to add the team in the group or 0 if you want to delete the team from the group
* http StatusCreated (201) will be sent if the group has been updated correctly

## Delete a group
* HTTP Request : ```DELETE http://api.com/group/{groupNumber}```
* http StatusOK (200) will be sent if the group has been deleted correctly

## Create a group
* HTTP Request : ```POST http://api.com/groups```
* Send data in the request body in the follow format 
``` 
    [
        {  
            "team_name": string
        },...
    ]
```
* http StatusCreated (201) will be sent if the group has been created correctly

## Get group

* HTTP Request : ```GET http://api.com/groups/{groupNumber}```
* Return a list object in json format as follow
    ``` 
    {
        "group_number": int,
        "teams": [
            {
                "id":               int,
                "name":             string,
                "photo":            string,
                "win":              int,
                "lose":             int,
                "draw":             int,
                "goals_pro":        int,
                "goals_against":    int,
                "points":           int
            },...
        ]
    }
    ```
* Teams are ordered first by points than winnings, draws, goals_pro and finaly goals_against