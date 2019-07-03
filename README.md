### Using simple echo framework
-   structure folders
    ```
    echo-framework
        ├── bin
        ├── pkg
        ├── README.md
        └── src
            ├── config
            │   ├── config.go
            │   └── config.json
            ├── controllers
            │   └── user
            │       └── user.go
            ├── helpers
            │   ├── db.go
            │   └── user.go
            ├── models
            │   └── users.go
            ├── routes.json
            └── server.go
    ```
    -   config - for storing configuration application
        -   config.json
            ```
            {
                "DbHost": "localhost",
                "DbPort": 3306,
                "DbType": "mysql",
                "DbName": "test",
                "DbUsername": "test",
                "DbPassword": "test"
            }

            ```
    -   helper - for helping utils library
        -   DbGorm - global database connection and using for all modules
    -   models - define models for application
        -   users.go - user model
    -   controllers - folder for storing all logical controllers
        -   users - folder to store all user logical function
            -   user.go - executing all user's requests from clients
    -   server.go - entry point for starting echo framework
        -   DbGorm.SetLogger
            -   Can use with https://github.com/uber-go/zap to write to log file
            -   log.New(os.Stdout, "\r\n", 0) for terminal
        -   http.StatusOK - this is standard library from golang

### References
-   https://www.naimibrahim.me/2018/11/20/echo-framework-gorm-blazing-fast-golang-server-side-application-authentication-example/
