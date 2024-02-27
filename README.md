# dev-davidchang-api
RESTful API for interacting with David Chang programmatically

## Build & run locally (http://localhost:80)
### Configure
`$ cp .env.example .env`
- Fill out all the details for your local database (replace examples in brackets):
	- MYSQL_PORT=[3306]
	- DB_USER=[root]
	- DB_PASSWORD=[password]
	- DB_NAME=[local_db]
	- DB_CONNECTION_NAME=localhost:$MYSQL_PORT
	- DB_CONNECTION_STRING=$DB_USER:$DB_PASSWORD@tcp($DB_CONNECTION_NAME)/$DB_NAME?charset=utf8&parseTime=True&loc=Local
- Populate your local environment
`$ source .env`
- Run a local MySQL database
`$ docker run -d -p $MYSQL_PORT:$MYSQL_PORT --name local_mysql -e MYSQL_ROOT_PASSWORD=$DB_PASSWORD mysql:latest`
- Create local database
`$ docker exec -it local_mysql mysql -u $DB_USER -p`
(enter password when prompted $DB_PASSWORD)
`CREATE DATABASE [$DB_NAME]` (replace [$DB_NAME] with password you set up)
### Build
`$ docker build -t dev-davidchang-api .`
### Run
`$ docker run -p 80:8080 -it --rm --name dev-dc-api dev-davidchang-api`

## API Design (https://api.davidchang.dev)
>[!NOTE]
>_The `resumes` endpoint retrieves and saves resumes (meta data) from/to a database. Future extension to retrieve/save resumes as binary data from/to object storage_
### `/resumes` 
- GET - Get a list of resumes, returned as JSON
- POST - Add a new resume from request data sent as JSON
### `/resumes:id` 
- GET - Get a resume by ID, returning resume data as JSON
