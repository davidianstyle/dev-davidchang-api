# dev-davidchang-api
RESTful API for interacting with David Chang programmatically

## Build & run locally (http://localhost:80)
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
