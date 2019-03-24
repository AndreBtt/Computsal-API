# Computsal

### API - Endpoints

#### **Players**

- **Get all players**
    * HTTP Request : ```GET http://api.com/players```
    * Return a list of player object in json format as follow
        ```json 
            {
                ID   int    
	            Name string 
	            Team string
            }
        ```

- **Create a player**
    * HTTP Request : ```POST http://api.com/players```
    * Send player data in the request body, in the follow format 
    ```json 
            {  
	            Name string 
	            Team string
            }
    ```
    * http StatusCreated (201) will be sent if the player has been created correctly
    
- **Delete a player**
    * HTTP Request : ```DELETE http://api.com/players/{ID}```
    * ID is the player ID you want to delete
    * http StatusOK (200) will be sent if the player has been deleted correctly

- **Update a player**
    * HTTP Request : ```PUT http://api.com/players/{ID}```
    * ID is the player ID you want to update
    * Send player data in the request body, in the follow format
    ```json 
            {  
	            Name string 
	            Team string
            }
    ```
    * Even if you want to update just one field, you need to fill all others in order to update player correctly
    * http StatusCreated (201) will be sent if the player has been updated correctly

- **Get a player**
    * HTTP Request : ```GET http://api.com/players/{ID}```
    * ID is the player ID you want to get information
    * Return a player object in json format as follow
        ```json 
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


