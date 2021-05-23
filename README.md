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

## Custom format error

### Method Not Found
```
curl -i localhost:3000/foo
HTTP/1.1 404 Not Found
Content-Length: 58
Content-Type: application/json
Date: Thu, 13 May 2021 21:52:52 GMT

{
  "error": "the requested resource could not be found"
}
```

### Method Not Allowed
```
curl -i -X PUT localhost:3000/v1/movies/1
HTTP/1.1 405 Method Not Allowed
Allow: GET, OPTIONS
Content-Length: 66
Content-Type: application/json
Date: Thu, 13 May 2021 21:52:52 GMT

{
  "error": "the PUT method is not supported for this resource"
}
```

## Create new Movie

Create a BODY variable conatining the JSON data that we want to send.
```
$ BODY='{"title":"Moana","year":2016,"runtime":107, "genres":["animation","adventure"]}”
```

Use the -d flag to send the contents of the BODY variable as the HTTP request body.

Note that curl will default to sending a POST request when then -d flag is used.
```
curl -i -d "$BODY" localhost:4000/v1/movies
HTTP/1.1 200 OK
Date: Mon, 17 May 2021 15:52:26 GMT
Content-Length: 65
Content-Type: text/plain; charset=utf-8

{Title:Moana Year:2016 Runtime:107 Genres:[animation adventure]}
```

## Read JSON Helper

Send some XML as the request body
```
curl -i -d '<?xml version="1.0" encoding="UTF-8"><note><to>Mavro</to></note>' localhost:4000/v1/movies
HTTP/1.1 400 Bad Request
Content-Type: application/json
Date: Mon, 17 May 2021 22:30:41 GMT
Content-Length: 67

{
        "error": "body contains badly-formated JSON (at character 1)"
}
```

Send some malformed JSON (notice the trailing comma)
```
curl -i -d '<?xml version="1.0" encoding="UTF-8"><note><to>Mavro</to></note>' localhost:4000/v1/movies
HTTP/1.1 400 Bad Request
Content-Type: application/json
Date: Mon, 17 May 2021 22:30:41 GMT
Content-Length: 67

{
        "error": "body contains badly-formated JSON (at character 1)"
}
```

Send a JSON array instead of an object
```
curl -i -d '<?xml version="1.0" encoding="UTF-8"><note><to>Mavro</to></note>' localhost:4000/v1/movies
HTTP/1.1 400 Bad Request
Content-Type: application/json
Date: Mon, 17 May 2021 22:30:41 GMT
Content-Length: 67

{
        "error": "body contains badly-formated JSON (at character 1)"
}
```

Send a numeric 'title' value (instead of string)
```
curl -i -d '{"title": 123 }' localhost:4000/v1/movies   
HTTP/1.1 400 Bad Request
Content-Type: application/json
Date: Mon, 17 May 2021 22:37:05 GMT
Content-Length: 74

{
        "error": "body contains incorrect JSON type for the field \"title\""
}
```

Send an empty request body
```
curl -i -d '<?xml version="1.0" encoding="UTF-8"><note><to>Mavro</to></note>' localhost:4000/v1/movies
HTTP/1.1 400 Bad Request
Content-Type: application/json
Date: Mon, 17 May 2021 22:30:41 GMT
Content-Length: 67

{
        "error": "body contains badly-formated JSON (at character 1)"
}
```

Send an unknown field request body
```
curl -d '{"title": "Moana", "rating":"PG"}' localhost:4000/v1/movies
{
   "error": "body contains unknown field \"rating\""
}
```

Send additional data in the request body
```
curl -d '{"title": "Moana"}{"title":"TopGun"}' localhost:4000/v1/movies
{
   "error": "body must only cotain a single JSON value"
}
```

Send additional data in the request body
```
curl -d '{"title": "Moana"} :~()' localhost:4000/v1/movies
{
   "error": "body must only cotain a single JSON value"
}
```

Send request with a very large JSON body
```
curl -d @./cmd/api/largefile.json localhost:4000/v1/movies
{
   "error": "body must not be large than 1048576 bytes"
}
```

Send request with Runtime format "<runtime> mins"
```
curl -d '{"title": "Moana", "runtime":"107 mins"}' localhost:4000/v1/movies
{
   Title:Moana Year:0 Runtime:107 Genres:[]
}
```

Send request with Runtime format "<runtime>"
```
curl -d '{"title": "Moana", "runtime":107}' localhost:4000/v1/movies
{
   "error": "invalid runtime format"
}
```


Send request with Runtime format "<runtime> minutes"
```
curl -d '{"title": "Moana", "runtime":"107 minutes"}' localhost:4000/v1/movies
{
    "error": "invalid runtime format"
}
```

Send request data for check validator
```
BODY='{"title":"","year":1000,"runtime":"-123 mins","genres":["sci-fi","sci-fi"]}'
curl -i -d "$BODY" localhost:4000/v1/movies
HTTP/1.1 422 Unprocessable Entity
Content-Type: application/json
Date: Thu, 20 May 2021 21:10:14 GMT
Content-Length: 180
{
   "error": {
        "genres": "must not contain duplicate values",
        "runtime": "must be a positive integer",
        "title": "must be provided",
        "year": "must be greater than 1888"
   }
}
```

# Data store

## Database Postgres

### Config Postgres and PgAmin
 Create docker-compose.yml file
```
version: '3.5'

services:
  postgres:
    container_name: postgres_greenlight_container
    image: postgres
    environment:
      POSTGRES_USER: ${POSTGRES_USER:-postgres}
      POSTGRES_PASSWORD: ${POSTGRES_PASSWORD:-changeme}
      PGDATA: /data/postgres
    volumes:
       - postgres:/data/postgres
    ports:
      - "5432:5432"
    networks:
      - postgres
    restart: unless-stopped
  
  pgadmin:
    # container_name: pgadmin_container
    image: dpage/pgadmin4:5.3
    environment:
      PGADMIN_DEFAULT_EMAIL: ${PGADMIN_DEFAULT_EMAIL:-pgadmin4@pgadmin.org}
      PGADMIN_DEFAULT_PASSWORD: ${PGADMIN_DEFAULT_PASSWORD:-admin}
      PGADMIN_CONFIG_SERVER_MODE: 'False'
    volumes:
       - pgadmin:/root/.pgadmin

    ports:
      - "${PGADMIN_PORT:-9090}:80"
    networks:
      - postgres
    restart: unless-stopped

networks:
  postgres:
    driver: bridge

volumes:
    postgres:
    pgadmin:
```

### Running Database Postgres and PgAmin
```
$ docker-compose up
```

### Remove Docker container 
```
$ docker-compose down
```

### Connect to Postgres server
```
$ psql -h localhost -p 5432 -U postgres
```

### Help commands database Postgres
```
postgres=#\?
```

### Create database greenlight
```
postgres=# CREATE DATABASE greenlight;
CREATE DATABASE
postgres=#\c greenlight
You are now connected to database "greenlight" as user postgres
```

### Create user greenlight
```
postgres=# CREATE ROLE greenlight WITH LOGIN PASSWORD 'pa55w0rd';
CREATE ROLE
```

### Add extension citext
```
greenlight=# CREATE EXTENSION IF NOT EXISTS citext;
CREATE EXTENSION
```

### Connect to databse greenlight with user greenlight
```
psql --host=localhost --dbname=greenlight --username=greenlight
Password for user greenlight:
psql (13.3, server 13.2 (Debian 13.2-1.pgdg100+1))
Type "help" for help
```

### Add environment variable to $HOME/.profile or $HOME/.bashrc or $HOME/.zshrc
```
export GREENLIGHT_DB_DSN='postgres://greenlight:pa55w0rd@localhost/greenlight?sslmode=disable'
```

### Get help from command-line api
```
$ go run ./cmd/api -help
  -db-dsn string
        PostgreSQL DSN (default "postgres://greenlight:pa55w0rd@localhost/greenlight?sslmode=disable")
  -env string
        Environment (development|staging|production) (default "development")
  -port int
        API Server port (default 4000)
```
