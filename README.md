# David Chang API
RESTful API for interacting with me programmatically!

## Build & run locally (http://localhost:80)

### Configure, build, and run using `docker compose`
`$ cp .env.example .env`
- Fill out all the details for your local database (replace examples in brackets):
	- ENVIRONMENT=[development]
	- NETWORK_NAME=[api-bridge]
	- DB_IMAGE_NAME=[mysql-image]
	- DB_CONTAINER_NAME=[db]
	- APP_IMAGE_NAME=[api-image]
	- APP_CONTAINER_NAME=[app]
	- MYSQL_HOST_PORT=[3306]	# Must match db/Dockerfile
	- MYSQL_CONTAINER_PORT=[3306]	# Must match db/Dockerfile
	- API_HOST_PORT=[80]		# Must match app/Dockerfile
	- API_CONTAINER_PORT=[8080]	# Must match app/Dockerfile
	- DB_USER=[root]
	- DB_PASSWORD=[password]
	- DB_NAME=[db]
	- DB_CONNECTION_NAME=$DB_CONTAINER_NAME:$MYSQL_CONTAINER_PORT
	- DB_CONNECTION_STRING="$DB_USER:$DB_PASSWORD@tcp($DB_CONNECTION_NAME)/$DB_NAME?charset=utf8&parseTime=True&loc=Local"
- Propagate environments to application and database
`$ cp .env app/.env`
`$ cp .env db/.env`
- Generate the db/init.sql file
`$ cd db`
`$ ./init.sh`
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
	- NETWORK_NAME=[api-bridge]
	- DB_IMAGE_NAME=[mysql-image]
	- DB_CONTAINER_NAME=[db]
	- APP_IMAGE_NAME=[api-image]
	- APP_CONTAINER_NAME=[app]
	- MYSQL_HOST_PORT=[3306]	# Must match db/Dockerfile
	- MYSQL_CONTAINER_PORT=[3306]	# Must match db/Dockerfile
	- API_HOST_PORT=[80]		# Must match app/Dockerfile
	- API_CONTAINER_PORT=[8080]	# Must match app/Dockerfile
	- DB_USER=[root]
	- DB_PASSWORD=[password]
	- DB_NAME=[db]
	- DB_CONNECTION_NAME=$DB_CONTAINER_NAME:$MYSQL_CONTAINER_PORT
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
`$ ./init.sh`
`$ docker build -t $DB_IMAGE_NAME .`
`$ docker run -d -p $MYSQL_HOST_PORT:$MYSQL_CONTAINER_PORT --name $DB_CONTAINER_NAME --network $NETWORK_NAME -e MYSQL_ROOT_PASSWORD=$DB_PASSWORD -e DB_NAME=$DB_NAME $DB_IMAGE_NAME`
#### Build & run app code
- Build & run
`$ cd ./app`
`$ docker build -t $APP_IMAGE_NAME .`
`$ docker run -p $API_HOST_PORT:$API_CONTAINER_PORT -it --rm --name $APP_CONTAINER_NAME --network $NETWORK_NAME $APP_IMAGE_NAME`
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
