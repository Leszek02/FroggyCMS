# FROGGY CMS

### Local run
run `cp .env.example .env`
in /cms folder run `go mod download`
in /cms/cmd folder run `go run .`

Recommend building and running database via docker-compose by running `docker compose build` & `docker compose up db`

If you want a local postgres, then `create_database.sql` file creates and populates tables.

### Docker run
run `docker compose build` & `docker compose up`


### Credentials
Admin: email: `a@b.c` password: `admin123`
Client: email: `a@b.c` password: `qwerty`

### Links
Main page: `localhost:8080/shop/main`
Admin panel: `localhost:8080/admin/login`

# DISCLAIMER
In admin panel under "pages" tab editing works only for "Products" template