# Narawangsa

Narawangsa is an app for managing your reading schedule. It is such a fantastic app. It can increase your reading intensity.

# Features
- Login and store JWT to the user
- Create, Read, Update, and Delete user
- Create, Read, Update, and Delete book
- Create, Read, Update, and Delete category
- Create, Read, Update, and Delete booklist
- Create and Delete category book
- Read and Update user level
- Create and Read read confirmation

## Installation

First, install [go](https://go.dev/doc/install) to use this backend app.

Check the version of your go. This app works for v1.17.2
```bash
go version
```

This app will be run on docker as well. So, make sure that docker has been installed on local.

Also, check the version of your docker. This app works for v20.10.10
```bash
docker -v
```
## Usage
1. Clone the repository
```bash
git clone https://github.com/rhtyx/narawangsa.git
```
2. Later on, define the app.env

```
DB_DRIVER=postgres
DB_SOURCE=postgresql://narawangsa:narawangsa@localhost:5434/narawangsa_db?sslmode=disable
SERVER_ADDRESS=0.0.0.0:8080
SECRET_KEY=narawangsa_only_narawangsa_only_
ACCESS_TOKEN_DURATION=15m
REFRESH_TOKEN_DURATION=24h
```
3. Create postgres database on docker
```bash
make postgres
```
4. Create the narawangsa database in the container
```bash
make createdb
```

5. Do the migration to build up the tables
```bash
make migrateup
```

6. Run the server
```bash
make run
```

## Test
This is the [link](https://drive.google.com/drive/folders/1Pmmm7uNadgNJjmD2oqOakcMgVTWjdOHu?usp=sharing) where postman environment and collection provided.
