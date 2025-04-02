# Task Management System

- Build a task management system using Clean Architecture in Go.
- Added create, read, update, delete and list task apis.
- Used sql for database
- Used go-chi framework for route and middleware handling
- Added pagination and filtering using sql queries

## How to Run ?

1. Clone the repo
2. Set config variables
   1. SQL_ADDR=$USERNAME:$PASSWORD@tcp(127.0.0.1:3306)/$DATABASE?charset=utf8&parseTime=True&loc=Local - replace USERNAME, PASSWORD and DATABASE. I used "root" username, "" password and "task_management", default value already set.
   2. SERVER_PORT - port where you want to run the server.
3. Log in to your mysql db and run these command -

```
CREATE DATABASE task_management;
USE task_management;

CREATE TABLE tasks (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    status VARCHAR(50) DEFAULT 'NEW',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```

4. Go to cloned repo and run `go mod tidy`, it will install all dependency.
5. Run `go run cmd/main.go`

## APIs

1. Create Task Api

#### Request

```
curl --location 'http://localhost:8080/tasks' \
--header 'Content-Type: application/json' \
--data '{
    "title" : "Task 12",
    "description" : "urgent requirement",
    "status" : "NEW"
}'
```

#### Response

```
{"id":11,"title":"Task 12","description":"urgent requirement","status":"NEW","created_at":"","updated_at":""}
```

2.  Get Task by ID

#### Request

```
curl --location 'http://localhost:8080/tasks/1'
```

#### Response

```
{
    "id": 1,
    "title": "Task Updated",
    "description": "done",
    "status": "COMPLETED",
    "created_at": "2025-04-02T18:08:48+05:30",
    "updated_at": "2025-04-02T19:07:07+05:30"
}
```

3.  Update Task

#### Request

```
curl --location --request PUT 'http://localhost:8080/tasks/1' \
--header 'Content-Type: application/json' \
--data '{
   "title" : "Task Updated",
   "description" : "done",
   "status" : "COMPLETED"
}'
```

#### Response

```
{"id":1,"title":"Task Updated","description":"done","status":"COMPLETED","created_at":"","updated_at":""}

```

4.  Delete Task by ID

#### Request

```
curl --location --request DELETE 'http://localhost:8080/tasks/4'
```

5.  List Tasks

#### Request

```
curl --location 'http://localhost:8080/tasks?page=2&size=4&status=NEW'
```

#### Response

```
[
   {
       "id": 11,
       "title": "Task 12",
       "description": "urgent requirement",
       "status": "NEW",
       "created_at": "2025-04-02T20:12:04+05:30",
       "updated_at": "2025-04-02T20:12:04+05:30"
   },
   {
       "id": 12,
       "title": "Task 33",
       "description": "urgent requirement",
       "status": "NEW",
       "created_at": "2025-04-02T20:19:35+05:30",
       "updated_at": "2025-04-02T20:19:35+05:30"
   },
   {
       "id": 13,
       "title": "Task 44",
       "description": "urgent requirement",
       "status": "NEW",
       "created_at": "2025-04-02T20:19:39+05:30",
       "updated_at": "2025-04-02T20:19:39+05:30"
   },
   {
       "id": 14,
       "title": "Task 34",
       "description": "urgent requirement",
       "status": "NEW",
       "created_at": "2025-04-02T20:19:49+05:30",
       "updated_at": "2025-04-02T20:19:49+05:30"
   }
]
```

## Clean Layered Architecture :-

1. **Separation of Concerns** : Each level in your program should be separate by a clear barrier, like application layer, service layer , storage layer ,.etc.
2. **Dependency Inversion Principle** : Code dependencies can only point inward. High-level modules should not depend on low-level modules.
3. **Adaptability to Change** : Organize your code in a modular and flexible way, so that it can be easily introduce new features, refactor existing code, and respond to evolving business requirements.
4. **Encapuslation** : Each layer hides its internal workings from others, exposing only what is necessary.

## Additional Points

1. To handle concurreny, we can use Goroutines and channels.
2. To add a new service like user service, we can directly add one more package in internal folder and implement it.
