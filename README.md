### Go skeleton API environment

Personally I am using this environment to up Go API fast. It's based on docker-compose and includes:
 
 - mysql init
 - mysql db migrations
 - beanstalkd init
 - swagger ui

### Usage

```
docker-compose up
```

### Ports

 - API - [127.0.0.1:8080](http://127.0.0.1:8080)
 - Swagger UI - [127.0.0.1:8081](http://127.0.0.1:8081)