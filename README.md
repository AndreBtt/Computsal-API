# Computsal

## API - Endpoints

### **Players**

- **Get all players**
    * HTTP Request : ```GET http://api.com/players```
    * Return a list of object in json format as follow
        ``` 
            {
                ID   int    
	            Name string 
	            Team string
            }
        ```

- **Create a player**
    * HTTP Request : ```POST http://api.com/players```
    * Send player's data in the request body in the follow format 
    ``` 
            {  
	            Name string 
	            Team string
            }
    ```
    * http StatusCreated (201) will be sent if the player has been created correctly
    
- **Delete a player**
    * HTTP Request : ```DELETE http://api.com/players/{ID}```
    * ID is the player's ID you want to delete
    * http StatusOK (200) will be sent if the player has been deleted correctly

- **Update a player**
    * HTTP Request : ```PUT http://api.com/players/{ID}```
    * ID is the player's ID you want to update
    * Send player's data in the request body in the follow format
    ``` 
            {  
	            Name string 
	            Team string
            }
    ```
    * Even if you want to update just one field you need to fill all others in order to update player correctly
    * http StatusCreated (201) will be sent if the player has been updated correctly

- **Get a player**
    * HTTP Request : ```GET http://api.com/players/{ID}```
    * ID is the player's ID you want to get information
    * Return a player object in json format as follow
        ``` 
            {
                ID           int
                Name         string
                Team         string
                TeamPhotoURL string
                Score        int
                YellowCard   int
                RedCard      int
                Captain      bool
            }
        ```

### **Teams**

- **Get all teams**
    * HTTP Request : ```GET http://api.com/teams```
    * Return a list of object in json format as follow
        ``` 
            {
                ID       int    
                Name     string 
                PhotoURL string 
                Group    int    
            }
        ```
- **Create a team**
    * HTTP Request : ```POST http://api.com/teams```
    * Send team's data in the request body in the follow format 
    ``` 
            {  
                Name     string 
                PhotoURL string 
                Group    int
            }
    ```
    * http StatusCreated (201) will be sent if the team has been created correctly

- **Delete a team**
    * HTTP Request : ```DELETE http://api.com/teams/{name}```
    * Name is the team's name you want to delete
    * All players from this team will also be deleted
    * http StatusOK (200) will be sent if the team has been deleted correctly

- **Update a team**
    * HTTP Request : ```PUT http://api.com/teams/{name}```
    * Name is the team's name you want to update
    * Send team's data in the request body in the follow format
    ``` 
            {  
                Name     string 
                PhotoURL string 
                Group    int
            }
    ```
    * Even if you want to update just one field you need to fill all others in order to update team correctly
    * http StatusCreated (201) will be sent if the team has been updated correctly

- **Get a team**
    * todo

- **Get team's players**
    * HTTP Request : ```GET http://api.com/teams/{name}/players```
    * Name is the team's name you want to get players
    * Return a list of player object in json format as follow
        ``` 
            {
                ID   int    
	            Name string 
	            Team string
            }
        ```

### **Score**

- **Get all players score**
    * HTTP Request : ```GET http://api.com/score```
    * Return a list object in json format as follow ordered by ascending score order
        ``` 
            {
                Player: {
                    id   int
                    name string
                    team string
                }
                Score int
            }
        ```

- **Get players score in a match**
    * HTTP Request : ```GET http://api.com/score/{matchID}```
    * matchID is the match's ID you want to get information
    * Return a list object in json format as follow ordered by ascending score order
        ``` 
            {
                Player: {
                    id   int
                    name string
                    team string
                }
                Score int
            }
        ```
