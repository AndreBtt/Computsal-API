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

## Update a group

* HTTP Request : ```PUT http://api.com/groups/{id}```
* ID is the group ID you want to update
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
* HTTP Request : ```DELETE http://api.com/group/{ID}```
* ID is the group ID you want to delete
* http StatusOK (200) will be sent if the group has been deleted correctly

## Create a group
* HTTP Request : ```POST http://api.com/groups```
* Send data in the request body in the follow format 
``` 
    [
        {  
            "team_name": string, 
        },...
    ]
```
* http StatusCreated (201) will be sent if the group has been created correctly
  