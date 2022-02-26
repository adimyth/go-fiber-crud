# Go Fiber CRUD App

---

## CRUD App

---

A simple todo crud app.

| Method | Endpoint     | Description    |
| ------ | ------------ | -------------- |
| GET    | `/todos`     | Get all todos  |
| GET    | `/todos/:id` | Get todo by id |
| POST   | `/todos`     | Create todo    |
| DELETE | `/todos/:id` | Delete todo    |
| PUT    | `/todos/:id` | Update todo    |

1. `models/Todo.go` contains model definition
2. `controllers/Todo.go` contains the app logic

## Dependencies

---

1. [Go Fiber](https://gofiber.io) - An Express-inspired web framework written in Go.
2. [Mongo Models](https://github.com/Kamva/mgm) - Mongo Go Models (mgm) is a fast and simple MongoDB ODM for Go (based on official Mongo Go Driver)
3. [Air](https://github.com/cosmtrek/air) - Live reload for Go apps. Equivalent to `nodmeon` in Node.js

## Running the App

---

1. Clone the repository
2. Install required dependencies for the app & Air
    ```bash
    go mod tidy
    go get github.com/cosmtrek/air
    ```
3. Run the app
    ```bash
    air serve .
    ```

## DB Setup

---

1. Install mongodb community version
2. Start mongodb
3. Create `todos` database & `todo` collection
4. Create a user with read & write access to `todos` database

    ```bash
    use todos;

    db.createCollection("todo");

    db.createUser({
      user: "test",
      pwd: "test123",
      roles: [{
        role: "readWrite",
        db: "todos"
      }]
    });
    ```

For some reason the authentication fails even when able to login via mongo shell. Have raised an [issue here](https://github.com/Kamva/mgm/issues/66)
