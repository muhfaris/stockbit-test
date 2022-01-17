## Please write a microservice to search movies from http://www.omdbapi.com/



The microservice should be able to handle two transports : REST JSON HTTP and GRPC

Access credentials :
```
OMDBKey : "faf7e5bb&s"
URL : http://www.omdbapi.com/
```

* Example url call to search is 
   
   --> GET http://www.omdbapi.com/?apikey=faf7e5bb&s=Batman&page=2

Functions to be implemented are :
- Search with pagination 
   
   --> 2 parameters : "pagination" and "searchword"

- Get single detail of the movie
- Log each search calls to a dummy DB eg. let's just say we have a MySQL DB table for this.

Important aspects :
- Readability of code
- Good display on the knowledge of "Separation of Concerns for Codes"
- Write unit tests on some of the important files. (For Bigger plus points see below)
- Good use of asynchronousy with Go-Routine

Plus points:
- Implementation of Clean Architecture is a BIG plus
- Complete Unit tests on all codes

### Answer
#### Description
Get movie information from imdbd. You can search movie and you can get detail information of movie. The app support for Rest API and GRCP, 

#### Configs 
You can change configuration application at `configs/config.toml`.
```
[app]
  name         = "movie"
  port         = 8989
  environment  = "dev"

  [app.http]
    write_timeout = 10
    read_timeout  = 10
    idle_timeout  = 10

[storage]
  [storage.database]
    host       = "127.0.0.1"
    port       = "3306"
    username   = "userapp"
    password   = "userapp"
    name       = "movies_db"

[api]
  [api.ombdb]
    api        = "http://www.omdbapi.com"
    secret_key = "faf7e5bb"
```


#### How to run 
`go run main.go`

#### API
##### Search movie 
You can search with `searchword` and combination with `pagination`.

``` 
curl --request GET \
  --url 'http://localhost:8989/api/v1/movies/search?searchword=batman&pagination=5'

```

##### Get detail movie
 ```
 curl --request GET \
  --url http://localhost:8989/api/v1/movies/tt1213218
  ```
  
##### Grpc client 
the client default search batman movie. You can run with command below:
```
cd grpc/client

go run client.go
```
