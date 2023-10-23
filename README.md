# Users API <img height="24" width="24" src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg"/>

## Table of Contents

* [Description](#description)
* [Features](#features)
* [Project Structure](#project-structure)

## Description

The application allows the user to create, read, update, and delete data in a database.

It uses the following technologies: 

* [Go](https://golang.org/) as programming language          
* [Gin](https://github.com/gin-gonic/gin) as web framework
* [MongoDB](https://www.mongodb.com/) as database
* [GoMail](gopkg.in/mail.v2) as smtp service
* [Swagger](https://github.com/go-swagger/go-swagger) for documentation
* [Docker](https://www.docker.com/) for containerization


## Features

&gt;&gt;&gt; [The project is available here](https://go-users-api.onrender.com) &lt;&lt;&lt;

Users can sign up. Once signed in, an email is sent to the user with an activation link. The user can then activate the account. Following activation, the user must log in, so a token is generated. The user can perform various operations such as retrieving data by id, by email, updating, and deleting data if logged in.

## Project Structure

The project is based on the [MVC](https://en.wikipedia.org/wiki/Model%E2%80%93view%E2%80%93controller) pattern.

<image src="./public/mvc.png" />