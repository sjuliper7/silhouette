# Silhouette


## Descriptions
Silhouette project is a simple project about microservices, the project implementing about `Golang`, `Microservices`, `Clean Arch`, `gRPC`, and `Kafka`. First, what is microservice? microservices is not about the number of services, different programming languages that implement on service, but microservices is software as a service. Silhouette project is a simple microservices because the silhouette project contains several services like `user-service`, `profile-service` and `notification-service` where each service be a software as a service that serve each other, lets say, user registrations `user-service` will save data for authentication and another data `like user_id, address, gender, phone_number, etc` will pass to `profile-service` for handling storing data to database and `notification-service` will sent email after receive data from `user-service`. Also, Silhourtte contain commons folder as a library that stores/ contains config and `Protocs` that use for handling `gRPC` each service. Wow, what is `gRPC`? `gRPC` is a remote procedure call develop by google for communications, The means of communication is communication between client and server, so... if user-service need data from `profile-service` then `user-service` will be a client and `profile-service` will be a server. `Kafka` is a message broker, so a lot of message broker, like `rabbitMQ, google pub sub, etc`. When using the Message Broker? if you ask when we to use message broker is depends on case what kind of application we want to develop, but for general when to use message broker is when you design the task just do task does not need to know the task successfully or not [EX [registration]: Register new user where while user has added new user will receive an email confirmation so the task will handle by `user-service` for registration and sending email will be handled by `notification-service`] so process just need waiting until user successfully add to database, not need to wait until email has been sent to new user. Because nothing gets blocked in the client application or service until the service finishes its task. The actual event itself will live on the message broker, possibly, and it will sit there waiting for a consumer to pick it up and process it.[Fire and Forget]

### Case
In this project will handle how user register until the new user will received email  after registration, `user-service` will handle several function like add-user, update-user, delete-user, we can say CRUD[Create, Read, Update, and Delete] for user `profile-service` will have CRUD to, and `notification-service` will handle sending email, the notification will sent email when user do registration and user update the new password.

#### design
```
User.go
- ID          int64
- Email       string
- Password    string
- Role        string
- Name        string
- Profile     Profile

Profile.go
- ID          int64 
-	UserId      int64  
-	Address     string
-	WorkAt      string
-	PhoneNumber string
- Gender      string
```

each service in this project will implement the clean arch, we can see the picture bellow how the structure the clean arch
i find the in on `repository` my friend
![alt text](https://raw.githubusercontent.com/bxcodec/go-clean-arch/master/clean-arch.png)

Next the design of implementation of `simple microservices`, `clean acrhitecture`, `kafka`, `API` like `gRPC` and `REST`\
![silhouette (1)](https://user-images.githubusercontent.com/29673571/68541910-09c5d200-03d8-11ea-9c55-eb345347f696.png)



For more explanation about project soon as possible

Thank you
