# Telega Forward Messages

### Use

This  package used go-tdlib and is dependencies on Tdlib telegram client.
How is install  [Readme](https://github.com/zelenin/go-tdlib/blob/master/README.md)

Next use sh file how is it:
```
#!/bin/bash
set -e
set -a;

source ./.env;
go run ./cmd/main.go
```

Or use [docker-compose](docker-compose.yml)

`>$ docker-compose up`
