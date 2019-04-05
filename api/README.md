# Computsal

## Start

* First build docker image `docker build -t computsal:api .`
* Since the image is created get its ID using the command `docker images`
* To start the server run the command `docker run -p 8080:8080/tcp "ImageID"` 

## API - Endpoints

### [Players](api/components/player/README.md)

### [Teams](api/components/team/README.md)

### [Groups](api/components/group/README.md)     

### [Previous Matches](api/components/previousmatch/README.md)

### [Next Matches](api/components/nextmatch/README.md)

### [Score](api/components/score/README.md)

### [Captain](api/components/captain/README.md)

### [Time](api/components/time/README.md)

### [Schedule](api/components/schedule/README.md)
