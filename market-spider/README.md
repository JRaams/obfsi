# market-spider

[![Go Report](https://goreportcard.com/badge/github.com/JRaams/obfsi)](https://goreportcard.com/report/github.com/JRaams/obfsi)

Steam market spider that runs a crobjob every hour to fetch recent item prices

### Running

`$ cd market-spider`
`../market-spider$ go run .`

### OCI container

Using [Posman](https://podman.io/getting-started/installation)

Building the image:
`$ cd market-spider`
`$ ../market-spider$ podman build -t market-spider:latest .`

Running a container with said image:
`$ podman run -d localhost/market-spider:latest`

Inspecting the container:
`$ podman logs -l`

Copy all files in container for debugging:
`$ podman cp <containername>:/ ./.docker/`