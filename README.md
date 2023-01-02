# docker-list-pets
This is my first build docker in golang use Dokerfile.
My mini API get list pets use golang and Gorilla Mux.

### Docker
you can use script.sh to build docker image and container

### Gorilla Mux
```
go get -u github.com/gorilla/mux
```
### Get Pets
```
http://localhost:8181/pets
```
### Get Pets by ID
```
http://localhost:8181/pet/2
```
