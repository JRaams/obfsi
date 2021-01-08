# market-spider

[![Go Report](https://goreportcard.com/badge/github.com/JRaams/obfsi)](https://goreportcard.com/report/github.com/JRaams/obfsi)

Steam market spider that runs a crobjob every hour to fetch recent item prices

### Running

`$ cd market-spider`
`../market-spider$ go run .`

### Migrations

Database migrations are executed/created using [golang-migrate](https://github.com/golang-migrate/migrate)

**Creating a new migration** 
`$ migrate create -ext sql -dir migrations -seq <name>`
A `x_<name>_.up.sql` and `x_<name>_.down.sql` will be created, edit these with the desired SQL statements.

**Executing existing migrations**
Export postgres URL
`$ export POSTGRESQL_URL='postgres://obfsi@0.0.0.0:5432/obfsi?sslmode=disable'`

Run migrations
`$ migrate -database ${POSTGRESQL_URL} -path migrations up`

### OCI container

Using [Podman](https://podman.io/getting-started/installation)

Building the image:
`$ cd market-spider`
`$ ../market-spider$ podman build -t market-spider:latest .`

Running a container with said image:
`$ podman run -d localhost/market-spider:latest`

Inspecting the container:
`$ podman logs -lf`

Copy all files in container for debugging:
`$ podman cp <containername>:/ ./.docker/`