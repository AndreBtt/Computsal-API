# Score

## Get all players score
* HTTP Request : ```GET http://api.com/scores```
* Return a list object in json format as follow ordered by decreasing score order
    ``` 
        [
            {
                "player" {
                    "id":    int,
                    "name":  string,
                    "team":  string
                }
                "score":         int,
                "yellowCard":    int,
                "redCard":       int
            },...
        ]
    ```
* HTTP Request : ```GET http://api.com/scores/{matchID}```
* matchID is the match ID you want to get information
* Return a list object in json format as follow ordered by decreasing score order
    ``` 
        {
            "player" {
                "id":    int,
                "name":  string,
                "team":  string
            }
            "score":         int,
            "yellowCard":    int,
            "redCard":       int
        }
    ```