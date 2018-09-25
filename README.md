# Goes

a set of tools to test server

## Server

a simple server to test http port, there are some environments as follow:

name |  useage | default | example
---- | ---- | ---- | ----
REDIRECT_TO| if visit /redirect it will sed a get request to defined location | | http://server2:9090
VERSION | the response will print the value | | v2

simple run with docker:

```bash
docker run -p 9000:9000 vinkdong/server
```


## How to build

- server (on Mac)

```bash
APP=server make image
```