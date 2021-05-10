# Greenlight RESTful API

## Folders

* The **bin** directory will contain out compiled application binaries, ready for deployment to a production server.

* The **cmd/api** directory will contain the application-specific code for our Greenlight API application.

* The **internal** directory will contain various ancillary packages used by out API. It will contain the code for interacting with our database, doing data validation, sending emails and so on. Basically, any code which isn't application-specific and can potentially be reused will live in here.

* The **migrations** directory will contain the SQL migration files for our database.

* The **remote** directory will contain the configuration files and setup scripts for our production server.

* The **go.mod** file will declare out project depencies, versions and module path.

* The **Makefile** will contain recipes for automating common administrative tasks - like auditing our Go code, building binaries, and executing database migrations.

