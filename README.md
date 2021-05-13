# Greenlight RESTful API

## Skeleton structure for our project

* The **bin** directory will contain out compiled application binaries, ready for deployment to a production server.

* The **cmd/api** directory will contain the application-specific code for our Greenlight API application.

* The **internal** directory will contain various ancillary packages used by out API. It will contain the code for interacting with our database, doing data validation, sending emails and so on. Basically, any code which isn't application-specific and can potentially be reused will live in here.

* The **migrations** directory will contain the SQL migration files for our database.

* The **remote** directory will contain the configuration files and setup scripts for our production server.

* The **go.mod** file will declare out project depencies, versions and module path.

* The **Makefile** will contain recipes for automating common administrative tasks - like auditing our Go code, building binaries, and executing database migrations.

## Running application
For running application open terminal use the go run command to compile and execute the code in the cmd/api package.
```
$ go run ./cmd/api
Hello World!
```

## Endpoints

| Method | URL Pattern     |  Handler            | Action                                 |
|:-------|:----------------|:--------------------|:---------------------------------------|
| GET    | /v1/healthcheck | healthcheckHandler  | Show application information           |
| GET    | /v1/movies      | listMoviesHandler   | Show the details of all movies         |
| POST   | /v1/movies      | createMoviesHandler | Create new movie                       |
| GET    | /v1/movies/:id  | showMovieHandler    | Show the details of a specific movie   |
| PUT    | /v1/movies/:id  | editMovieHandler    | Update the details of a specific movie |
| DELETE | /v1/movies/:id  | deleteMovieHandler  | Delete a specific movie                |


| Method | Usage                                                                                                                                         | 
|:-------|:----------------------------------------------------------------------------------------------------------------------------------------------|
| GET    | Use for actions that retrieve information only and don't change the state of your application or any data.                                    |
| POST   | Use for non-idempotent actions that modify state. In the context of a REST API, POST is generally used for actions that create a new resource.|
| PUT    | Use for idempotent actions that modify the state of a resource at a specific URL. In the context of a REST API, PUT is generally used for actions that replace or update an existing resource. |
| PATCH  | Use for actions that partially update a resource at a specific URL. It's OK for the action to be either idempotent or non-idempotent. |
| DELETE | Use for actions that delete a resource at a specific URL. |








## Running application without parameters
For running application open terminal use the go run command to compile and execute the code in the cmd/api package.
```
$ go run ./cmd/api
2021/05/10 20:08:47 starting development server on :4000
```

## Get healthcheck endpoint
```
$ curl -i localhost:4000/v1/healthcheck
HTTP/1.1 200 OK
Date: Tue, 11 May 2021 00:12:18 GMT
Content-Length: 58
Content-Type: text/plain; charset=utf-8

status: available
environment: development
version: 0.0.1
```

## Running application with parameters
For running application open terminal use the go run command to compile and execute the code in the cmd/api package.
```
$ go run ./cmd/api -port=3030 -env=production
2021/05/10 20:15:34 starting production server on :3030
```
## Get healthcheck endpoint
```
curl -i localhost:3030/v1/healthcheck
HTTP/1.1 200 OK
Date: Tue, 11 May 2021 00:15:48 GMT
Content-Length: 57
Content-Type: text/plain; charset=utf-8

status: available
environment: production
version: 0.0.1
```

## Implementation movies actions
```
$ go run ./cmd/api
2021/05/10 20:15:34 starting production server on :4000
```

### Get healthcheck endpoint
```
curl -i localhost:4000/v1/healthcheck
HTTP/1.1 200 OK
Date: Tue, 11 May 2021 00:15:48 GMT
Content-Length: 57
Content-Type: text/plain; charset=utf-8

status: available
environment: production
version: 0.0.1
```

### Post movies endpoint
```
curl -i -X POST localhost:4000/v1/movies
HTTP/1.1 200 OK
Date: Tue, 11 May 2021 00:15:48 GMT
Content-Length: 57
Content-Type: text/plain; charset=utf-8

create a new movie
```

### Get a movie endpoint
```
curl -i  localhost:4000/v1/movies/1
HTTP/1.1 200 OK
Date: Tue, 11 May 2021 00:15:48 GMT
Content-Length: 57
Content-Type: text/plain; charset=utf-8

show the details of a movie 1
```

## Unsupported HTTP method
You might also want to try making some requests for a particular URL using an unsupported HTTP method. The httprouter packge has automatically sent a 405 Method Not Allowed response for us, including an Allow header which lists the HTTP methods that are supported for the endpoint.

### POST a healthcheck
```
curl -i -X POST localhost:4000/v1/healthcheck/
HTTP/1.1 405 Method Not Allowed
Allow: GET, OPTIONS
Content-Type: text/plain; charset=utf-8
X-Content-Type-Options: text/plain nosniff
Date: Tue, 11 May 2021 00:15:48 GMT
Content-Length: 19

Method Not Allowed
```

## Get allowed HTTP method supported
You can make an OPTIONS request to a specific URL and httprouter will send back a response with an Allowed header detailing the supported HTTP methods.

```
curl -i -X OPTIONS localhost:4000/v1/healthcheck/
HTTP/1.1 200 OK
Allow: GET, OPTIONS
Date: Tue, 11 May 2021 00:15:48 GMT
Content-Length: 0
```

## Live reloading Go application

Install gin is a simple command line utility for live-reloading Go web applications. Just run gin in your app directory and your web app will be served with gin as a proxy. gin will automatically recompile your code when it detects a change. Your app will be restarted the next time it receives an HTTP request.

### How to install
```
$ go get https://github.com/codegangsta/gin
```

### Basic usage
```
Options:
    --laddr value, -l value       listening address for the proxy server
   --port value, -p value        port for the proxy server (default: 3000)
   --appPort value, -a value     port for the Go web server (default: 3001)
   --bin value, -b value         name of generated binary file (default: "gin-bin")
   --path value, -t value        Path to watch files from (default: ".")
   --build value, -d value       Path to build files from (defaults to same value as --path)
   --excludeDir value, -x value  Relative directories to exclude
   --immediate, -i               run the server immediately after it's built
   --all                         reloads whenever any file changes, as opposed to reloading only on .go file change
   --godep, -g                   use godep when building
   --buildArgs value             Additional go build arguments
   --certFile value              TLS Certificate
   --keyFile value               TLS Certificate Key
   --logPrefix value             Setup custom log prefix
   --notifications               enable desktop notifications
   --help, -h                    show help
   --version, -v                 print the version
```

### Running application
```
$ gin --appPort 4000 --port 3000 --path ./cmd/api
```

### Get healthcheck endpoint with JSON response
```
curl -i localhost:3000/v1/healthcheck
HTTP/1.1 200 OK
Content-Length: 68
Content-Type: application/json
Date: Thu, 13 May 2021 21:52:52 GMT

{
  "status": "available",
  "environment": "development",
  "version": "0.0.1"
}
```