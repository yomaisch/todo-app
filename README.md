# Todo App

This will be a simple todo app created to learn how to create web apps with golang, including CRUD.

We are using postgresql for the database, which is not created on a virtual server such as Docker, so if you want to try it out, you will need to set up a your local environment (but it should be easy :) ).

# DEMO

Top page<br>
<img width="457" alt="top" src="https://user-images.githubusercontent.com/57692216/123016458-98e4f980-d405-11eb-9f86-7489ce0eedb2.png">

Create Task Page<br>
<img width="403" alt="create" src="https://user-images.githubusercontent.com/57692216/123016536-caf65b80-d405-11eb-8264-459a09ca5c24.png">

After creating a task<br>
<img width="520" alt="after create" src="https://user-images.githubusercontent.com/57692216/123016760-46580d00-d406-11eb-8db9-0ae9ee770dfd.png">

After finishing the task (After pressing done button, Here we are, finishing ID 6.)<br>
<img width="435" alt="after delete" src="https://user-images.githubusercontent.com/57692216/123017674-0560f800-d408-11eb-8703-8be0a9bb27a8.png">

# Requirement

- Go v1.15.1
- [pq](https://github.com/lib/pq) v1.10.2

# Installation

```bash
go get github.com/lib/pq
```

# Usage

Go into your database and create a table.

```bash
create table todo(id serial PRIMARY KEY, task varchar(50));
```

After cloning this repository.

Set the environment variables in the main.go and dbConn functions to the values of your environment as appropriate.

And you can run the following command.

```bash
go run main.go
```

Load the following URL

```bash
http://localhost:8080/
```

# License

MIT License. See LICENSE for more information.
