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

## Create a team
* HTTP Request : ```POST http://api.com/teams```
* Send team's data in the request body in the follow format 
``` 
        {  
            "name":      string,
            "photo":     string, 
            "group":     int
        }
```
* http StatusCreated (201) will be sent if the team has been created correctly

## Delete a team
* HTTP Request : ```DELETE http://api.com/teams/{name}```
* Name is the team's name you want to delete
* All players from this team will also be deleted
* http StatusOK (200) will be sent if the team has been deleted correctly

## Update a team
* HTTP Request : ```PUT http://api.com/teams```
* Send team's data in the request body in the follow format
``` 
        {  
            "id":        int,
            "name":      string,
            "photo":     string, 
            "group":     int
        }
```
* It will be replace using the ID value, ensure it is correct
* Even if you want to update just one field you need to fill all others in order to update team correctly
* http StatusCreated (201) will be sent if the team has been updated correctly

## Get a team
* todo
