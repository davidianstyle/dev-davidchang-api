# dev-davidchang-api
RESTful API for interacting with David Chang programmatically

## Build & run locally (http://localhost:80)
### Build
`$ docker build -t dev-davidchang-api .`
### Run
`$ docker run -p 8080:80 -it --rm --name dev-dc-api dev-davidchang-api`

## API Design (https://api.davidchang.dev)
>[!NOTE]
>_The `resumes` endpoint retrieves and saves resumes (binary data) from/to object storage_
### `/resumes` 
- GET - Get a list of resumes, returned as JSON
- POST - Add a new resume from request data sent as JSON
### `/resumes:id` 
- GET - Get a resume by ID, returning resume data as JSON

>[!NOTE]
>_The `family-members` endpoint saves and retrieves family member info saved in a relational (?) or document (?) database_
### `/family-members`
- GET - Get a list of Chang family members returned as JSON
- POST - Add a new Chang family member from request data sent as JSON
### `/family-members:id` 
- GET - Get a family member by ID, returning family member data as JSON
