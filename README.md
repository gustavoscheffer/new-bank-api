# NEW-BANK-API

## Endpoints

* GET /v1/user/ - Get all users
* GET /v1/user/:id - Get by id
* POST /v1/user/ - Create user
* PATCH /v1/user/:id - Update user by id
* DELETE /v1/user/:id - Delete user by id
* GET /docs/ - Swagger documentation

## Technologies

* Golang
* Express
* MongoDB
* Swagger

## Tips

- Build and run application:
```
go run .
```

- Create Swagger docs automatically
```
swag init --parseDependency 
```

## Environment Variables:
MONGO_URI
MONGO_DB
APP_PORT

## Improvements

- [x] Improve all Endpoints based on Rest full Api's
- [x] Configure logs
- [x] Implement swagger(docs) 
- [x] Improve README.md
- [ ] Unit Test
- [ ] Docker instructions
- [ ] Insomnial test colections
- [ ] Implement authentication
- [ ] GitLab Pipeline (CI/CD)