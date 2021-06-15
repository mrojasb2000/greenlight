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
$ docker-compose up -d
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
psql --host=localhost --port=5432 --dbname=greenlight --username=greenlight
Password for user greenlight:
psql (13.3, server 13.2 (Debian 13.2-1.pgdg100+1))
Type "help" for help
```

### Add environment variable to $HOME/.profile or $HOME/.bashrc or $HOME/.zshrc
```
export GREENLIGHT_DB_DSN='postgres://greenlight:pa55w0rd@localhost:5432/greenlight?sslmode=disable'
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

### Using the DSN with psql
```
$ psql $GREENLIGHT_DB_DSN
psql (13.3, server 13.2 (Debian 13.2-1.pgdg100+1))
Type "help" for help.
```

### Custom settings PostgreSQL Engine
- Maximum Open Connections:
- Maximum Idle Connections:
- Maximum Idle Time:

```
go run ./cmd/api -help                
  -db-dsn string
        PostgreSQL DSN (default "postgres://greenlight:pa55w0rd@localhost/greenlight?sslmode=disable")
  -db-max-idle-conns int
        PostgreSQL max idle connections (default 25)
  -db-max-idle-time string
        PostgreSQL max connection idle time (default "15m")
  -db-max-open-conns int
        postgreSQL max open connections (default 25)
  -env string
        Environment (development|staging|production) (default "development")
  -port int
        API Server port (default 4000)
```


## Migrations

Using migrations to menage your database schema, rather than amnually executing the SQL statements yourself. Each pair of migration files is numbered sequentially, usually 0001, 0002, 0003... or with a Unix timestamp, to indicate the order in which migrations should be applied to a database.

### Install Migration tools
```
$ brew install golang-migrate
```

### Create movies table
Creating a new movies table in our database, generate a pair of migration file using the migrate create command.

flags:
* -seq flag inicates that we want to use secuential numbering like 0001, 0002, ... for the migration files.
* -ext flag indicates that we want to give the migration files the extenson .sql.
* -dir flag indicates that we want to store the migration files in the ./migration directory.
* create_movies_table is a description label that we give the migration files to signify their contents.

```
$ migrate create -seq -ext=.sql -dir=./migrations create_movies_table
... 000001_create_movies_table.down.sql
... 000001_create_movies_table.up.sql
```

### Add SQL statement migrations files
PostgreSQL DDL apply statements sql filename contains .up. pattern
```
CREATE TABLE IF NOT EXISTS movies (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    title text NOT NULL,
    year integer NOT NULL,
    runtime integer NOT NULL,
    genres text[] NOT NULL,
    version integer NOT NULL DEFAULT 1
);
```

PostgreSQL DDL roll back stattements sql filename contains .down. pattern
```
DROP TABLE IF EXISTS movies;
```

### Executing the migrations
Require DSN parameter from your environment variable.
```
$ migrate -path=./migrations -database=$GREENLIGHT_DB_DSN up
1/u create_movies_table (23.828006ms)
2/u add_movies_check_constraints (42.078909ms)
```

### Using the DSN with psql
```
$ psql $GREENLIGHT_DB_DSN
psql (13.3, server 13.2 (Debian 13.2-1.pgdg100+1))
Type "help" for help.
```

### List tables
```
greenlight=> \dt
                List of relations
| Schema |       Name        | Type  |   Owner     |
|:-------|:------------------|:------|:------------|
| public | movies            | table | greenlight  |
| public | schema_migrations | table | greenlight  |
```

### Describe table
```
greenlight=> \d movies
                                        Table "public.movies"
|   Column   |            Type             | Collation | Nullable |              Default                |
|:-----------|:----------------------------|:----------|:---------|:------------------------------------|
| id         | bigint                      |           | not null | nextval('movies_id_seq'::regclass)  |
| created_at | timestamp(0) with time zone |           | not null | now()                               |
| title      | text                        |           | not null |                                     |
| year       | integer                     |           | not null |                                     |
| runtime    | integer                     |           | not null |                                     |
| genres     | text[]                      |           | not null |                                     |
| version    | integer                     |           | not null | 1                                   |

- Indexes:
   * "movies_pkey" PRIMARY KEY, btree (id)

- Check constraints:
   * "genres_lenght_check" CHECK (array_length(genres, 1) >= 1 AND array_length(genres, 1) <= 5)
   * "movies_runtime_check" CHECK (runtime >= 0)
   * "movies_year_check" CHECK (year >= 1888 AND year::double precision <= date_part('year'::text, now()))
```

### Check current migration version
If you want to see which migration version your database is currently on you can run the migarte tool's version command.
```
$ migrate -path=./migrations -database=$GREENLIGHT_DB_DSN version
2
```

### Change migration version
You can also migrate up or down to a specific version by using to goto command.
```
$ migrate -path=./migrations -database=$GREENLIGHT_DB_DSN goto 1
```

### Down migration version
You can use the down command ro roll-back by a specific number of migrations.
```
$ migrate -path=./migrations -database=$GREENLIGHT_DB_DSN down 1
```

### Roll-back all migrations
```
$ migrate -path=./migrations -database=$GREENLIGHT_DB_DSN down
Are you sure you want to apply all down migrations? [y/N]
y
Applying all down migrations
2/d add_movies_check_constraints (29.759654ms)
1/d create_movies_table (50.684902ms)
```

### Fixing errors in SQL migrations
What happens when you make a syntax error in your SQL migration files, then you need to manually roll-back the partially applied migration. 

Once that's done, then you must also 'force' the version number in the schema_migrations table to the correct value.

Once you force the version, the database is considered 'clean' and you should be able to run migrations again without any problem.

```
$ migrate -path=./migrations -database=$GREENLIGHT_DB_DSN force 1
```

### Running migrations on application startup
If you want, it is also posible to use the golang-migrate/migrate Go package to automatically execute your databse migrations on application start up.

```
package main

import (
  "context"
  "database/sql"
  "flag"
  "fmt"
  "log"
  "net/http"
  "os"
  "time"

  "github.com/golang-migrate/migrate/v4"
  "github.com/golang-migrate/migrate/v4/database/postgres"
  "github.com/golang-migrate/migrate/v4/source/file"

  _ "github.com/lib/v4"
)

func main() {
  db, err := openDB(cfg)
  if err != nil {
    logger.Fatal(err)
  }
  defer db.Close()

  logger.Printf("database connection pool established")

  migrationDriver, err := postgres.WithInstance(db, &postgres.Confgi{})
  if err != nil {
    logger.PrintFatal(err, nil)
  }

  migrator, err := migrate.NewWithDatabaseInstance("file:///path/to/you/migrations", "postgres", migrationDriver)
  if err != nil {
    logger.PrintFatal(err, nil)
  }

  err = migrator.Up()
  if err != nil && err != migrate.ErrnoChange {
    logger.PrintFatal(err, nil)
  }
  logger.Printf("database migrations applied")
  ...
}
```

### Implement create new movie (CRUD operations)
```
$ migrate -path=./migrations -database=$GREENLIGHT_DB_DSN up


BODY='{"title":"Moana","year":2016,"runtime":"107 mins","genres":["animation","adventure"]}'

curl -i -d "$BODY" localhost:4000/v1/moviesHTTP/1.1 201 Created
Content-Type: application/json
Location: /v1/movies/1
Date: Thu, 27 May 2021 23:01:38 GMT
Content-Length: 221

{
  "movie": {
    "id": 1,
    "title": "Moana",
    "year": 2016,
    "runtime": "107 mins",
    "genres": [
            "animation",
            "adventure"
    ],
    "version": 1
  }
}
```

### Implement read movie (CRUD operations)
Find movie with ID that exist in the database
```
$ curl -i localhost:4000/v1/movies/2 
HTTP/1.1 200 OK
Content-Type: application/json
Date: Fri, 28 May 2021 02:32:20 GMT
Content-Length: 161

{
  "movie": {
    "id": 2,
    "title": "Black Panther",
    "year": 2018,
    "runtime": "134 mins",
    "genres": [
            "action",
            "adventure"
    ],
    "version": 1
  }
}
```
Find movie with ID that doesn't exist in the database
```
$ curl -i localhost:4000/api/v1/movies/3
HTTP/1.1 404 Not Found
Content-Type: application/json
Date: Fri, 28 May 2021 02:31:23 GMT
Content-Length: 58

{
  "error": "the requested resource could not be found"
}
```

### Implement UPDATE movie (CRUD operations)
Find movie with ID that exist in the database
```
$ curl localhost:4000/v1/movies/2
{
  "movie": {
    "id": 2,
    "title": "Black Panther",
    "year": 2018,
    "runtime": "134 mins",
    "genres": [
            "action",
            "adventure"
    ],
    "version": 1
  }
}
```

Update movie record by ID
```
$ BODY='{"title":"Black Panther","year":2018,"runtime":"134 mins","genres":["sci-fi","action","adventure"]}'
$ curl -X PUT -d "$BODY" localhost:4000/v1/movies/2
{
  "movie": {
    "id": 2,
    "title": "Black Panther",
    "year": 2018,
    "runtime": "134 mins",
    "genres": [
            "sci-fi",
            "action",
            "adventure"
    ],
    "version": 2
  }
}
```

### Implement DELETE movie (CRUD operations)
Remove movie with ID that exist in the database
```
$ curl -X DELETE localhost:4000/v1/movies/3
{
   "message": "movie successfully deleted"
}
```

Remove movie with ID that doesn't exist in the database
```
$ curl -X DELETE localhost:4000/v1/movies/3
{
   "error": "the requested resource could not be found"
}
```

### Performing the partial update
```
$ curl -X PATCH -d '{"year":1985}' localhost:4000/v1/movies/4
{
  "movie": {
    "id": 4,
    "title": "The Breakfast Club",
    "year": 1985,
    "runtime": "96 mins",
    "genres": [
            "drama"
    ],
    "version": 2
  }
}
```

### Performing the partial update with key/value empty
```
$ curl -X PATCH -d '{"title":""}' localhost:4000/v1/movies/4
{
  "error": {
          "title": "must be provided"
  }
}
```

### Adding a query timeout
```
$ curl -w '\nTime: %{time_total}s \n' localhost:4000/v1/movies/1
{
        "error": "the server encontered a problem and could not process your request"
}

Time: 0,027932s
```