# Computsal

## Start

* First build docker image `docker build -t computsal:api .`
* Since the image is created get its ID using the command `docker images`
* To start the server run the command `docker run -p 8080:8080/tcp "ImageID"` 

## API - Endpoints

### [Players](components/player/README.md)

### [Teams](components/team/README.md)

### [Groups](components/group/README.md)     

### [Previous Matches](components/previousmatch/README.md)

### [Next Matches](components/nextmatch/README.md)

### [Score](components/score/README.md)

### [Captain](components/captain/README.md)

### [Time](components/time/README.md)

### [Schedule](components/schedule/README.md)
