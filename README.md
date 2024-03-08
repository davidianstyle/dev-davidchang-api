# David Chang API
RESTful API for interacting with me programmatically!

## Build & run locally (http://localhost:80)

### Configure, build, and run everything using `docker compose`
`$ ./init_env.sh`
- Fill out all the details for your local database (replace examples in brackets):
```
# Shared environment variables for ./api-app, ./webhook-app, ./reverse-proxy, and ./db
ENVIRONMENT=[development]

# ReadMe.io API key secret
README_API_KEY=[""]						# Get from https://dash.readme.com/project/api-davidchang-dev/ > Developer Dashboard > Set Up API Keys

# Database auth/connection variables
DB_USER=[root]
DB_PASSWORD=[password]
DB_NAME=[db]
DB_CONNECTION_NAME=$DB_CONTAINER_NAME:$MYSQL_CONTAINER_PORT
DB_CONNECTION_STRING="$DB_USER:$DB_PASSWORD@tcp($DB_CONNECTION_NAME)/$DB_NAME?charset=utf8&parseTime=True&loc=Local"

# Docker variables

## Network
NETWORK_NAME=[api-bridge]

## Database (built with MySQL locally and Cloud SQL in GCP)
DB_IMAGE_NAME=[mysql-db-image]
DB_CONTAINER_NAME=[db]
DB_HOST_PORT="3306"						# Must match db/Dockerfile
DB_CONTAINER_PORT="3306"				# Must match db/Dockerfile

## API Provider (built with Go)
API_IMAGE_NAME=[go-app-image]
API_CONTAINER_NAME=[api-app]
API_HOST_PORT="8080"					# Must match api-app/Dockerfile
API_CONTAINER_PORT="8080"				# Must match api-app/Dockerfile

## Webhook Provider (built with NodeJS)
WEBHOOK_IMAGE_NAME=[nodejs-app-image]
WEBHOOK_CONTAINER_NAME=[webhook-app]
WEBHOOK_HOST_PORT="8000"				# Must match webhook-app/Dockerfile
WEBHOOK_CONTAINER_PORT="8000"			# Must match webhook-app/Dockerfile

## Reverse Proxy (built with Nginx locally and Cloud Endpoints in GCP)
REVERSE_PROXY_IMAGE_NAME=[nginx-server-image]
REVERSE_PROXY_CONTAINER_NAME=[reverse-proxy]
REVERSE_PROXY_HOST_PORT="80"			# Must match reverse-proxy/Dockerfile
REVERSE_PROXY_CONTAINER_PORT="80"		# Must match reverse-proxy/Dockerfile
```
- Propagate environments to application and database
`$ ./propagate_env.sh`
- Generate the db/init.sql file
`$ cd db`
`$ ./init.sh`
- Build
`$ docker compose build`
- Run
`$ docker compose up`
- Tear down
`$ docker compose down`

### Configure, build, and run each service individually
#### Set up environment
`$ ./init_env.sh`
- Fill out all the details for your local database (replace examples in brackets):
```
# Shared environment variables for ./api-app, ./webhook-app, ./reverse-proxy, and ./db
ENVIRONMENT=[development]

# ReadMe.io API key secret
README_API_KEY=[""]						# Get from https://dash.readme.com/project/api-davidchang-dev/ > Developer Dashboard > Set Up API Keys

# Database auth/connection variables
DB_USER=[root]
DB_PASSWORD=[password]
DB_NAME=[db]
DB_CONNECTION_NAME=$DB_CONTAINER_NAME:$MYSQL_CONTAINER_PORT
DB_CONNECTION_STRING="$DB_USER:$DB_PASSWORD@tcp($DB_CONNECTION_NAME)/$DB_NAME?charset=utf8&parseTime=True&loc=Local"

# Docker variables

## Network
NETWORK_NAME=[api-bridge]

## Database (built with MySQL locally and Cloud SQL in GCP)
DB_IMAGE_NAME=[mysql-db-image]
DB_CONTAINER_NAME=[db]
DB_HOST_PORT="3306"						# Must match db/Dockerfile
DB_CONTAINER_PORT="3306"				# Must match db/Dockerfile

## API Provider (built with Go)
API_IMAGE_NAME=[go-app-image]
API_CONTAINER_NAME=[api-app]
API_HOST_PORT="8080"					# Must match api-app/Dockerfile
API_CONTAINER_PORT="8080"				# Must match api-app/Dockerfile

## Webhook Provider (built with NodeJS)
WEBHOOK_IMAGE_NAME=[nodejs-app-image]
WEBHOOK_CONTAINER_NAME=[webhook-app]
WEBHOOK_HOST_PORT="8000"				# Must match webhook-app/Dockerfile
WEBHOOK_CONTAINER_PORT="8000"			# Must match webhook-app/Dockerfile

## Reverse Proxy (built with Nginx locally and Cloud Endpoints in GCP)
REVERSE_PROXY_IMAGE_NAME=[nginx-server-image]
REVERSE_PROXY_CONTAINER_NAME=[reverse-proxy]
REVERSE_PROXY_HOST_PORT="80"			# Must match reverse-proxy/Dockerfile
REVERSE_PROXY_CONTAINER_PORT="80"		# Must match reverse-proxy/Dockerfile
```
- Propagate environments to application and database
`$ ./propagate_env.sh`
#### Create a shared network for your api-app, webhook-app, reverse-proxy and db
`$ docker network create $NETWORK_NAME`
#### Build & run a local MySQL database
- Build & run
`$ cd ./db`
`$ ./init.sh`
`$ docker build -t $DB_IMAGE_NAME .`
`$ docker run -d -p $DB_HOST_PORT:$DB_CONTAINER_PORT --name $DB_CONTAINER_NAME --network $NETWORK_NAME -e MYSQL_ROOT_PASSWORD=$DB_PASSWORD -e DB_NAME=$DB_NAME $DB_IMAGE_NAME`
- Test
(TODO)
#### Build & run API app code
- Build & run
`$ cd ./api-app`
`$ docker build -t $API_IMAGE_NAME .`
`$ docker run -p $API_HOST_PORT:$API_CONTAINER_PORT -it --rm --name $API_CONTAINER_NAME --network $NETWORK_NAME $API_IMAGE_NAME`
- Test
`$ curl --request GET --url localhost:$API_HOST_PORT/resumes --header 'accept: application/json'`
- Ok if you get a 200 response
#### Build & run webhook app code
- Build & run
`$ cd ./webhook-app`
`$ docker build -t $WEBHOOK_IMAGE_NAME .`
`$ docker run -p $WEBHOOK_HOST_PORT:$WEBHOOK_CONTAINER_PORT -it --rm --name $WEBHOOK_CONTAINER_NAME --network $NETWORK_NAME $WEBHOOK_IMAGE_NAME`
- Test
(TODO)
#### Build & run Nginx reverse-proxy
`$ cd ./reverse-proxy`
`$ docker build -t $REVERSE_PROXY_IMAGE_NAME .`
`$ docker run -p $REVERSE_PROXY_HOST_PORT:$REVERSE_PROXY_CONTAINER_PORT -it --rm --name $REVERSE_PROXY_CONTAINER_NAME --network $NETWORK_NAME $REVERSE_PROXY_IMAGE_NAME`
- Test
(TODO)
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
