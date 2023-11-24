# kcrps (KASPI CUSTOM REMOTE PAYMENT SYSTEM)

Golang
JWT
Swagger
PostgreSQL
Python
Flask
NOX Player
HTML
Bootstrap CSS

- for db docker
````
   docker run --name=invoices-db -e POSTGRES_PASSWORD=<dbpassword> -p 5432:5432 -d postgres

````

- for migrations
````
migrate -path ./schema -database 'postgres://postgres:<dbpassword>@localhost:5432/postgres?sslmode=disable' up
migrate -path ./schema -database 'postgres://postgres:<dbpassword>@localhost:5432/postgres?sslmode=disable' down
````

