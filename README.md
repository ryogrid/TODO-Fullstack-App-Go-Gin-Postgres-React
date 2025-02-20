# TODO-Fullstack-App-Go-Gin-SamehadaDB-React (forked for demo of SamehadaDB)
![Test Coverage](backend/api/coverage_badge.png)

This fullstack application creates a TODO List Web Page using the Go/Gin/SamehadaDB/React Stack.

![Screen Shot](App.png)

## Go server

Go is used to spin up the server, define routing, and interact with the database.

## Gin router

Gin is used to define the TODO API with functionality such as:

1. Listing all TODO items.
2. Creating a new TODO item and adding to the database.
3. Updating a TODO item with its completed condition.
4. Deleting a TODO item from the database.
5. Later being able to filter TODO items.

And serve static files of frontend.

## Database of backend server

[SamehadaDB](https://github.com/ryogrid/SamehadaDB) is used in embeded DB form to store the TODO items by saving rows in as id, item-text, and done condition.

## React

React is used here to create the frontend fully responsive application on the client side and is built using components.
