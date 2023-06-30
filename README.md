# GO Fiber ERP

Get all your tools to handle your project/company in one place ! 

## I - Prerequisite

- Linux compatible
- Go language >= 1.20.4
- Postgresql >= 15.3

## II - Configure and prepare dev environnement

### Env file

---

Copy `.env.example` file and rename it into `.env`. Then, fill all env vars or leave default values.
```sh
# Postgres database config
DB_HOST=localhost
DB_NAME=fiber-erp
DB_USER=fiber-erp-user
DB_PASSWORD=fiber-erp-pwd
DB_PORT=5432

# Change it for production env!
JWT_SECRET=my_secret
```

### Postgresql

---

Launch `postgresql` service

```sh
sudo service postgresql status
sudo service postgresql start
```

Enter psql with user postgres to create database and role for your application.

```sh
sudo -u postgres psql
```

Create database and role and grant privilege on the database. **Be aware: if you changed values of your `.env` file, you should change database name, user name and user password in the following commands.**

```sh
# psql

CREATE DATABASE fiber-erp;
CREATE ROLE fiber-erp-user WITH fiber-erp-pwd;

\l  # List all databases
\du # List all users/roles

# Move into the fiber-erp database
\c fiber-erp

# Grant all privileges for the database and its public schema to the created user.
GRANT ALL PRIVILEGES ON DATABASE fiber-erp TO fiber-erp-user;
GRANT ALL PRIVILEGES IN SCHEMA public TO fiber-erp-user;

exit
```

### Run server

---
To run the server, simply use the go command:

```sh
go run main.go
```

The server will be listening on the port `8000` by default, but you can change it inside the `.env` file.