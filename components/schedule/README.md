# Schedule

## Get team schedule

* HTTP Request : ```GET http://api.com/schedule/{teamID}```
* TeamID is the team's id you want to get information
* Return an array of objects in json format as follow
    ``` 
    [
        {
            "time_id":      int,
            "time":         time,
            "availability": bool
        },...
    ]
    ```
* time type is a string in the follow format "HH:MM:SS" where HH is hour, MM is minutes and SS seconds
* availability is true if the team want to play on that time or false if not

## Update team schedule

* HTTP Request : ```PUT http://api.com/schedule/{teamID}```
* TeamID is the team's id you want to get information
* Send data in the request body in the follow format
    ``` 
    [
        {
            "time_id":      int,
            "availability": bool
        },...
    ]
    ```
* availability is true if the team want to play on that time or false if not
* ensure that the data is correct, just send data that is to `update`
* http StatusOK (200) will be sent if time has been updated correctly