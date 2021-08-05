# TODO APP

## Installation

```sh
git clone https://github.com/titan2903/todoapp-refactory.git
cd todoapp-refactory
go run main.go
```

## Application Flow Explenation
   - User must register first, after the registration is complete, the user gets a token so that the user can directly enter the application.
   - The user can also enter the application via the login page if the user wants.
   - After the user enters the application, the user can create as many tasks as possible.
   - Users can only do their own update, Create, get, get one or delete tasks. Users cannot create, get, get one or delete tasks belonging to other users.


## DB Diagram

this is the database schema between the users and todos tables:

- [dbdiagram] - <= CLICK dbdiagram


## Postman TODO APP Collection

- [postman] - Collection Todo App


## License

MIT

**Free Software, Hell Yeah!**

[//]: # (These are reference links used in the body of this note and get stripped out when the markdown processor does its job. There is no need to format nicely because it shouldn't be seen. Thanks SO - http://stackoverflow.com/questions/4823468/store-comments-in-markdown-syntax)

[dbdiagram]: <https://dbdiagram.io/d/60f97a8cb7279e4123366c2e>
[postman]: <https://www.getpostman.com/collections/3c74741b5a92859ec9f3>