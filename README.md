# About the Project
This is basic learning project for Authetication using a Go framework named as [GIN](https://gin-gonic.com/) and an ORM (Object Relational Mapping) library [GORM](https://gorm.io/)

## Prerequisit
install [Go](https://go.dev/)
install [postgres](https://www.postgresql.org/) database, download pgAdmin (client for postgres) or download a client extension in [vs code](https://code.visualstudio.com/) editor

# How to run locally
1. clone this repository
```sh
https://github.com/xerofenix/gin-jwt
```
or
```sh
https://gitlab.com/xerofenix/git-jwt
```
2. Make ```.env``` in "gin-jwt" directory and add the following with your credentials
```sh
PORT=3000(or any port that you want to access on)
DB_URL="host=your-db-host user=-db-user password=-db-password dbname=db-name port=-db-port sslmode=disable TimeZone=Asia/Shanghai"
SECRET=any-secret-code
```
3. Open termianl in the "gin-jwt" directory and run
```sh
go run main.go
```

4. Open the browser and goto [https://localhost:300](https://localhost:300) or ```http://localhost:you-port (if chaged in .env)```
