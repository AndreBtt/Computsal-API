# Time

## Get times
* HTTP Request : ```GET http://api.com/times```
* Return a list of object in json format as follow
    ``` 
        [
            {
                "id":   int,    
                "time": time 
            },...
        ]
    ```
* time type is a string in the follow format "HH:MM:SS" where HH is hour, MM is minutes and SS seconds


## Create times
* HTTP Request : ```POST http://api.com/times```
* Send data in the request body in the follow format 
``` 
    [
        {  
            "time": time
        },...
    ]
```
* time type is a string in the follow format "HH:MM:SS" where HH is hour, MM is minutes and SS seconds
* http StatusCreated (201) will be sent if times have been created correctly

## Update and Delete times

* HTTP Request : ```PUT http://api.com/times```
* Send an array of times in the request body in the follow format
``` 
    [
        {  
            "id":       int,
            "time":     time,
            "action":   int
        },...
    ]
```
* time type is a string in the follow format "HH:MM:SS" where HH is hour, MM is minutes and SS seconds
* action is 1 if you want to update the time or 0 if you want to delete it
* http StatusOK (201) will be sent if times have been updated or deleted correctly