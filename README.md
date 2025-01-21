## Rakia.ai - technical assessment

Golang developer skill assessment by RAKIA.ai

## Assessment #1

To start follow the next steps:
* navigate console to `./assessment1` directory
* run `go run ./main.go` command


## Assessment #2

Project includes a Docker Compose and Make script for local deployment.

To start it locally follow the next steps:
* make sure Docker service is started on your laptop
* navigate console to `./assessment2/project` directory
* run `make up` command to start service
* test API using http://localhost:8080/ (for example with Postman)
* run `make down` command after you will finish to stop all docker containers

To run tests for project follow the next steps:
* navigate console to `./assessment2/project` directory
* run `make test` command to run tests

### API routes supported

<code>GET</code> <code><b>/v1/healthcheck</b></code> - service healthcheck route

<code>GET</code> <code><b>/v1/posts</b></code> - get a filtered list of posts, case insentive filteing by "title", pagination with "page" and "limit" query params

<code>GET</code> <code><b>/v1/posts/{id}</b></code> - get specific post, 404 if post not found

<code>POST</code> <code><b>/v1/posts</b></code> - add a new post, "title", "content", "author" needs to be specified

<code>PUT</code> <code><b>/v1/posts/{id}</b></code> - update specific post, "title", "content", "author" can be specified

<code>DELETE</code> <code><b>/v1/posts/{id}</b></code> - delete specific post
