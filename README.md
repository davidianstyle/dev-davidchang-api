# David Chang API
RESTful API for interacting with me programmatically!

## Build & run locally (http://localhost:80)

### Configure, build, and run using `docker compose`
`$ cp .env.example .env`
- Fill out all the details for your local database (replace examples in brackets):
	- ENVIRONMENT=[development]
	- MYSQL_PORT=[3306]
	- DB_USER=[root]
	- DB_PASSWORD=[password]
	- DB_NAME=[local_db]
	- DB_CONNECTION_NAME=db:$MYSQL_PORT
	- DB_CONNECTION_STRING="$DB_USER:$DB_PASSWORD@tcp($DB_CONNECTION_NAME)/$DB_NAME?charset=utf8&parseTime=True&loc=Local"
- Propagate environments to application and database
`$ cp .env app/.env`
`$ cp .env db/.env`
- Build
`$ docker compose build`
- Run
`$ docker compose up`
- Tear down
`$ docker compose down`

### Configure, build, and run manually
#### Set up environment
`$ cp .env.example .env`
- Fill out all the details for your local database (replace examples in brackets):
	- ENVIRONMENT=[development]
	- MYSQL_PORT=[3306]
	- DB_USER=[root]
	- DB_PASSWORD=[password]
	- DB_NAME=[db]
	- DB_CONNECTION_NAME=db:$MYSQL_PORT
	- DB_CONNECTION_STRING="$DB_USER:$DB_PASSWORD@tcp($DB_CONNECTION_NAME)/$DB_NAME?charset=utf8&parseTime=True&loc=Local"
- Propagate environments to application and database
`$ cp .env app/.env`
`$ cp .env db/.env`
- Populate your local environment
`$ source .env`
#### Create a shared network for your app and db
`$ docker network create $NETWORK_NAME`
#### Build & run a local MySQL database
- Build & run
`$ cd ./db`
`$ docker build -t mysql .`
`$ docker run -d -p $MYSQL_PORT:$MYSQL_PORT --name db --network $NETWORK_NAME -e MYSQL_ROOT_PASSWORD=$DB_PASSWORD mysql`
- Create local database
`$ docker exec -it db mysql -u $DB_USER -p$DB_PASSWORD`
(enter password when prompted $DB_PASSWORD)
`CREATE DATABASE [$DB_NAME];` (replace [$DB_NAME] with password you set up)
#### Build & run app code
- Build & run
`$ cd ./app`
`$ docker build -t dev-davidchang-api .`
`$ docker run -p 80:8080 -it --rm --name app --network $NETWORK_NAME dev-davidchang-api`
- Test
`$ curl --request GET --url localhost/resumes --header 'accept: application/json'`
- Ok if you get a 200 response

## API Design (https://api.davidchang.dev)
>[!NOTE]
>_The `resumes` endpoint retrieves and saves resumés (meta data) from/to a database. Future extension to retrieve/save resumés as binary data from/to object storage_
### `/resumes/latest`
- GET - Get the most recent resumé
### `/resumes` 
- GET - Get all resumés, returned as JSON
- POST - Add a new resumé from request data sent as JSON
### `/resumes:id` 
- GET - Get a resumé by ID, returning resumé data as JSON
- PATCH - Update a resumé by ID
- DELETE - Delete a resumé by ID

## Documentation

Full documentation can be found at https://docs.davidchang.dev
