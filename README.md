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