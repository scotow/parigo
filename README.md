# parigo 🥗

Lille University's Pariselle restaurant menu API

Simple API that fetch and parse the menu of the Pariselle restaurant from their website.

There is two example commands in the [`cmd`](https://github.com/Scotow/parigo/tree/master/cmd) directory:

- [`parigo`](https://github.com/Scotow/parigo/tree/master/cmd/parigo) prints today menu (or the nearest future one if there is no service for today) in an ascii table. The `-a` prints all the available days rather than just today. 
- [`web`](https://github.com/Scotow/parigo/tree/master/cmd/web) starts a web server on port 8080 and return the same ascii table as the parigo command.

NB: This project is just a 2 hours project to fetch the menu of a university restaurant and should not be considered as a serious project.
