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

| URL Pattern     |  Handler     | Action                       |
|:----------------|:------------:|-----------------------------:|
| /v1/healthcheck |  healthcheck | Show application information |

## Running application
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