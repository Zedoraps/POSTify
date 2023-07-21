# [POST]ify 

The simple and effective solution to POSTify your requests.

Reverse-proxy example/inspiration by https://gist.github.com/JalfResi/6287706

Every GET request will automatically be rewritten to a POST request. 

## RUN

`POSTIFY_REMOTE_URL=https://echo-api.3scale.net/ go run .`

### Configuration

- `POSTIFY_REMOTE_URL` defines the desired remote url

### Test

`curl localhost:8080/this-is-a-get-request`


