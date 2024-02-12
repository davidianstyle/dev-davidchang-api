# dev-davidchang-api
RESTful API for interacting with David Chang programmatically

## API Design (https://api.davidchang.dev)
>[!NOTE]
>_The `resumes` endpoint saves and retrieves resumes (binary data) from object storage_
### `/resumes` 
- GET - Get a list of resumes, returned as JSON
- POST - Add a new resume from request data sent as ???
### `/resumes:id` 
- GET - Get a resume by ID, returning resume data as ???

>[!NOTE]
>_The `family-members` endpoint saves and retrieves family member info saved in a relational (?) or document (?) database_
### `/family-members`
- GET - Get a list of Chang family members returned as JSON
- POST - Add a new Chang family member from request data sent as JSON
### `/family-members:id` 
- GET - Get a family member by ID, returning family member data as JSON
