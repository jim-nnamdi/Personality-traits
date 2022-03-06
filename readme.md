![Screenshot 2022-03-06 at 11 10 19 PM](https://user-images.githubusercontent.com/46245794/156944382-b3d1c488-b27f-4a5b-9bd9-59f8ac9e1bb4.png)

# Personality Test

> Without any Framework (i could have used one, but i wanted to go raw. i'll replicate this later by using an ORM (gorm) and a flexible framework (gin) ), am trying to wire up a minimalistic personality test in golang, this test would have just two questions and based on the results and calculations returned by the system would display to an individual if the person is an introvert or an extrovert.

## Architecture

> Am trying to implsement a Service Oriented Architecture :+1: using the MVC paradigm to create this open source project. Trying to make each function standalone and simple, hence the root of the project only carries the routes call that would initialize everything from DB setup to services and controllers.

```
routes.Routes()
``` 


![Screenshot 2022-03-06 at 11 10 54 PM](https://user-images.githubusercontent.com/46245794/156944391-0276b94d-c1e0-4552-be9f-59391d50f17f.png)
![Screenshot 2022-03-06 at 11 11 04 PM](https://user-images.githubusercontent.com/46245794/156944394-9b0b3e9f-cd2d-4924-b0f5-71d588c3ab4c.png)

## Start Project

> To start the project run the command below to install dependencies, and then ensure that you create a database with preferred Options [database name: personalitytest table name: questions] then add columns.
```
go mod tidy
```
