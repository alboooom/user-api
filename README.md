# user-api - application
- the routes is in cmd pacage
- file.json as database in db pacage
- models and dd drivers in models pacge
# Run the app
- go run main.go   

# REST API

# get all users:
-Request

 GET /users
 curl -H  http://localhost:1323/users

-Response

Content-Type: text/plain; charset=UTF-8
Date: Wed, 10 Feb 2021 08:43:10 GMT
Content-Length: 115
 
# get user by id:

-Request

 GET /users/id 
 curl -H  http://localhost:1323/users/1

-Response

Content-Type: application/json; charset=UTF-8
Date: Wed, 10 Feb 2021 08:43:04 GMT
Content-Length: 42
{
  "id": 1,
  "name": "alberttt",
  "email": "ast_albert@mail.ru"
}
 
# update user by id:

-Request

 PUT /users/5 HTTP/1.1
 Host: 127.0.0.1:1323
 User-Agent: insomnia/2020.5.2
 Content-Type: application/json
 Accept: */*
 Content-Length: 41

| {
| 	"email": "bbba@mail",
| 	"name": "bba"
| }

-Response

< HTTP/1.1 200 OK
< Date: Wed, 10 Feb 2021 08:41:09 GMT
< Content-Length: 0
  
# create new user:

-Request

 POST /users/new HTTP/1.1
 Host: 127.0.0.1:1323
 User-Agent: insomnia/2020.5.2
 Content-Type: application/json
 Accept: */*
 Content-Length: 53

| {
| 	"id": 5,
| 	"name": "bbb",
| 	"email": "bbb@mail.ru"
| }

-Response
 HTTP/1.1 200 OK
 Content-Type: application/json; charset=UTF-8
 Date: Wed, 10 Feb 2021 08:40:49 GMT
 Content-Length: 44
  
# delete user by id:

-Request

DELETE /users/id
curl -H  http://localhost:1323/users/1

-Response

 HTTP/1.1 200 OK
 Content-Type: application/json; charset=UTF-8
 Date: Wed, 10 Feb 2021 09:22:08 GMT
 Content-Length: 56

