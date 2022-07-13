# Go-Todo-API

A simple CRUD API built with Go's [Fiber](https://github.com/gofiber/fiber) and [GORM](https://github.com/go-gorm/gorm) libraries. It also uses [Docker](https://www.docker.com/) for containerization.

## Features

1. Simple CRUD functionality (Create, Read, Delete, but not Update yet)
2. Every `todo` has the following format:
```json
{
    "ID": 2,
    "CreatedAt": "2022-07-13T23:29:22.425341Z",
    "UpdatedAt": "2022-07-13T23:29:22.425341Z",
    "DeletedAt": null,
    "title": "Study",
    "description": "MAT-42",
    "priority": 5,
    "checked": false
}
```

The first four fields (`ID`, `CreatedAt`, `UpdatedAt` and `DeletedAt`) are created by GORM's model, which [embeds its own fields into the structs](https://gorm.io/docs/models.html#embedded_struct).

The other four fields are what actually describe a "todo" item: `title`, `description`, `priority` and `checked`.

## Next steps

- [?] Improve development workflow: 
    - Currently, developing the API demands the (re)creation of the Docker images (to recompile the code every time it changes). Can this be done more efficiently?
- Integration tests for the routes
    - [?] Integrate tests with the Docker containers
    - Return 201 status when an entry is created
    - Return the object that was created in the response body
    - Verify if the object was really created in the database
- Add an update/put/patch route
- Many different task lists