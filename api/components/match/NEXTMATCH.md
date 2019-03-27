# Next Matches

## Get all next matches

* HTTP Request : ```GET http://api.com/nextMatches```
* Return a list object in json format as follow
    ``` 
    [
        {
            "id":        int,    
            "team1":     string, 
            "team2":     string,
            "type":      int,    
            "time":      int    
        },...
    ]
    ```

## Update next matches

* HTTP Request : ```PUT http://api.com/nextMatches```
* Send **all** next matches in the request body in the follow format
``` 
    [
        {  
            "team1":    string,
            "team2":    string,
            "type":     int,
            "time":     int
        },...
    ]
```
* It's important to send all matches because in this update specifically 
we delete all the previous data and than insert the new matches