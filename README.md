Project structure

-   /cfg/ is for configuration files, usually json or yaml saved in plain text files, as they should be checked into git too (except for passwords, private keys, etc).
-   /middleware for all middleware.
-   /routes routes get grouped and nested using directories that mirror the API application’s RESTFul-like surface.
-   /webserver contains all shared HTTP structs and interfaces (Broker, configuration, Server, etc).
-   main.go where the application is bootstrapped (New(), Start()).
-   pipeline.go where the BuildPipeline() function lives.
