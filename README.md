# GoRESTapi
 Implementing a REST API in Go 

 Update 10/20/21: 

 We should have this working properly now, in accordance with Chapter 16 of the "for Dummies" book. 

 Easy command line interface with the API is accomplished by running the following example commands:

 ```
$ curl http://localhost:5000/api/v1/courses

$ curl http://localhost:5000/api/v1/courses/IOS301

$ curl -X GET http://localhost:5000/api/v1/courses/IOS301

$ curl -H "Content-Type:application/json" -X POST http://localhost:5000/api/v1/courses/IOS102 -d "{\"title\":\"SwiftUI Programming\"}"

$ curl -H "Content-Type:application/json" -X PUT http://localhost:5000/api/v1/courses/IOS102 -d "{\"title\":\"SwiftUI2 Programming\"}"

$ curl -X DELETE http://localhost:5000/api/v1/courses/IOS301
```

I added and modified a few things, and also have the Go debugger working in VSCode. Feeling good about this! 

Git Update 10/20/21:

We did not start with a main branch. Should have done that first. This commit allows us to use Main branch. 