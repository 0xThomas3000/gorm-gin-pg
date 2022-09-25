# Lesson summary of branch 1:
I. Setup Go-GORM-RESTful-API-Project:
   1. What is ORM (Object Relational Mapping)?
      - A technique that allows developers to use their native programming paradigm(mô hình) to query, manipulate data stored in Database.
      - It provides a data mapper pattern to transform data defined in classes / objects to SQL.
      - It came with CRUD functions already implemented making it easier for developers to perform CRUD operations against the database.
      - Golang has built a number of developer-friendly ORM libraries for developers to use JSON key:value pair syntax and encoding to map directly to any SQL database.

   2. Setup the Go project:
      a. Init a new Golang project:
         - Initialize a Go module
      b. Create a Docker PostgreSQL Server (use Docker to run the Postgres server which could contain many databases):
         - Create a PostgreSQL Docker Container:
            * Create "docker-compose.yml" with configs:
               + To manage an instance of PostgreSQL.
               + Map our local port 6500(port at local machine) to the Postgres default port 5432(inside the container):
                  => Allow us to connect and access our running PostgreSQL server outside the container.
               + Create a named "volume" to prevent data loss when deleting the Postgres container.
               + Run "docker-compose up -d": to start the Postgres Docker container in detached mode.
      c. Load and Validate the Environment Variables(app.env) with Golang Viper:
         - "go get github.com/spf13/viper"
      d. Create a Utility/Helper Function to Connect the Golang app to PostgreSQL:
         - "go get -u gorm.io/gorm"
         - "go get gorm.io/driver/postgres"
      e. Data Modeling and Migration with GORM:
         - Build a database model with GORM.
         - Install the UUID OSSP plugin.
         - Run a migration to push the schema to the PostgreSQL database.
         - DB migration: simply refers to the techniques used by developers to track incremental, and reversible changes in DB schemas. We can relate it to Git version control, which allows developers to track changes in a file or source code.
         - Install the UUID OSSP Module for PostgreSQL:
            * Access the bash shell of the running Postgres Docker container (exit to exit):
               + docker exec -it <container_name> bash
            * Enter the Postgres shell (\q to exit):
               + psql -U <postgres_user> <database_name>
            * List all the available extensions (to install "uuid-ossp"):
               + select * from pg_available_extensions;
               + CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


II. Template for environment variables ("app.env" file):
      POSTGRES_HOST = ...
      POSTGRES_USER = ...
      POSTGRES_PASSWORD = ...
      POSTGRES_DB = ...
      POSTGRES_PORT = ...

      PORT = ...
      CLIENT_ORIGIN = ...

      ACCESS_TOKEN_PRIVATE_KEY = ...
      ACCESS_TOKEN_PUBLIC_KEY = ...
      ACCESS_TOKEN_EXPIRED_IN = ...
      ACCESS_TOKEN_MAXAGE = ...

      REFRESH_TOKEN_PRIVATE_KEY = ...
      REFRESH_TOKEN_EXPIRED_IN = ...
      REFRESH_TOKEN_MAXAGE = ...
