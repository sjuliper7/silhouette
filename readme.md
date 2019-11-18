# Silhouette


## Descriptions
Silhouette project is simple project about microservices, the project implementing  about `golang`, `microserviecs`, `clean arch`, `grpc` and `kafka`.First what is microservice? microservices is not about number of service, different of programming languages that implement  on service, but microservices is software as a service. Why silhoutte is a simple microservices? because the silhouette project contains several services like `user-service`, `profile-service` and `notification` service where each service be a software as a service that serve each others, lets say, user registrations `user-service` wil be save data for authentication and give ani next. Also contain `commons` folder as library that store/ contains config  and `protocs` that use for handling gRPC each service. Wow, what is gRPC? gRPC is a remote procedure call develop by goole for communications, means of communications is communication betwen client and server. So each service will .

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

Next the design of implementation of `simple microservices`, `clean acrhitecture`, `kafka`, `API` like `gRPC` and `REST` 
![silhouette (1)](https://user-images.githubusercontent.com/29673571/68541910-09c5d200-03d8-11ea-9c55-eb345347f696.png)



For more explanation about project soon as possible

Thank you
