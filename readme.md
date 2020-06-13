# Silhouette


## Descriptions
This project is a simple project about microservices, the project implementing about `Golang`, `Microservices`, `Clean Arch`, `gRPC`, and `Kafka`, why i use the tech stack because i more familiar.

For clean archtecture is refer to several friend repo and my ex-companny.
- https://github.com/ecojuntak/gorb
- https://github.com/bxcodec/go-clean-arch


### Preparation
Here what you need prepare before run the applications:
1. Docker
2. Go
3. Mysql


### Run Application
#### Kafka
```
cd kafka
docker-compose -d
docker-compose up
```

#### User Service
User Service is application that handle authentication and as producer to publish message to profile service and notification service
```
cd services
cd user_service

// run migrations first
cd database/migrations
make migration_up

// back to user_service
make run

```
#### Profile Service
Profile Service is appication that handle for profile management as consumer, triggered by event, and provice RPC
```
cd services
cd profile_service

// run migrations first
cd database/migrations
make migration_up

// back to profile_service
make run

```

#### Notification Service
Notification Service is application that handle serving notification to client
```
cd services
cd notification_service

// back to notificaton_service
make run

```

### Endpoints
here list endpoint that you can access when all app is running..
```
GET api/v1/users
POST api/v1/users
GET api/v1/users/:id
PUT api/v1/users/:id
DELETE api/v1/users/:id
```
