# dashboard
dadshboard for kaspi api

- for db docker
````
   docker run --name=invoices-db -e POSTGRES_PASSWORD=<dbpassword> -p 5432:5432 -d --rm postgres

````

- for migrations
````
migrate -path ./schema -database 'postgres://postgres:<dbpassword>@localhost:5432/postgres?sslmode=disable' up
migrate -path ./schema -database 'postgres://postgres:<dbpassword>@localhost:5432/postgres?sslmode=disable' down
````

