# golang-rest-postgre
Rest CRUD operations against an PostgreSQL table in Golang.

Example in Go of a Restful we services against a simple PostgreSQL data model. It includes test against all operations, that can be tested with
`` go test -v ``

To statically build and package:

`` go build .``

To run after build:

``export APP_DB_HOST = postgre_host``

``export APP_DB_PORT = postgre_port``

``export APP_DB_USERNAME = postgre_user``

``export APP_DB_PASSWORD = postgre_pwd``

``export APP_DB_NAME = dbname``

``./golang-rest-postgre``

To invoke:

`` curl localhost:8000/products ``

`` curl localhost:8000/product/${id} ``


Modified from https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql
work licensed under a Creative Commons Attribution-NonCommercial-ShareAlike 4.0 International License. Thanks Kulshekhar Kabra!!
