# Next Matches

## Get next matches

* HTTP Request : ```GET http://api.com/nextMatches```
* If next matches are in group phase return a list object in json format as follow ordered by time
    ``` 
    [
        {
            "id":        int,    
            "team1":     string, 
            "team2":     string,
            "type":      int,    
            "time":      time    
        },...
    ]
    ```
* time type is a string in the follow format "HH:MM:SS" where HH is hour, MM is minutes and SS seconds
* matches are sorted by time if they don't have time seted yet, it will be sorted by type


## Update next matches

* HTTP Request : ```PUT http://api.com/nextMatches```
* Send an array of next matches data in the request body in the follow format
``` 
    [
        {  
            "id":       int,
            "team1":    string,
            "team2":    string,
            "type":     int,
            "time":     int
        },...
    ]
```
* When update a group phase (type = 0) team1, team2 and time will be updated
* When update a elimination phase (type != 0) only the match time will be updated
* http StatusCreated (201) will be sent if the match has been updated correctly

## Create next matches

* HTTP Request : ```POST http://api.com/nextMatches```
* This method is only available once to create the elimination phase
* Send an array of next matches data in the request body in the follow format
```
[
	{
		"team1" : string,
		"team2" : string
	},...
]
```
* The order of the array matters, it creates the next matches based on this order
* http StatusCreated (201) will be sent if it has been completed correctly


## Generate next matches

* HTTP Request : ```POST http://api.com/generateNextMatches```
* It will create the best match possible based on the schedule from each team
* http StatusOK (200) will be sent if it has been completed correctly