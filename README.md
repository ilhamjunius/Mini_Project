# ToDoList App
- Todo List App is an application used for task management
- A mini project RESTful API by Group 4 Alterra Immersive Program Batch 5


Run project with: 
```
go run main.go
```

## Stack-tech :dart:
- [x] RESTful API using Go, Echo, MySQL, Gorm

## Open Endpoints

Open endpoints require no Authentication.

* Register : `POST /users/register`
* Login : `POST /users/login`

## Endpoints that require Authentication

Closed endpoints require a valid Token to be included in the header of the request. A Token can be acquired from the Login view above.

### Current User related

Each endpoint manipulates or displays information related to the User whose Token is provided with the request:

- Get user by ID : `GET /users/:id`
- Update user by ID : `PUT /users/:id`
- Delete user by ID : `DELETE /users/:id`

### Task related

Each endpoint manipulates or displays information related to the Task whose Token is provided with the request:

- Get tasks : `GET /tasks`
- Get task by ID : `GET /tasks/:id`
- Create task : `POST /tasks`
- Update task : `PUT /tasks/:id`
- Delete task : `DELETE /tasks/:id`

### Project related

Each endpoint manipulates or displays information related to the Project whose Token is provided with the request:

- Get project by UserID : `GET /projects`
- Get project by ID : `GET /projects/:id`
- Create project : `POST /projects`
- Update project : `PUT /projects/:id`
- Delete project :  `DELETE /projects/:id`


