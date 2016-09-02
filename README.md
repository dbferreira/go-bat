# go-bat
Basic implementation of a browser based cricket simulation game in Go and AngularJS 2.0

Front-end Angular2 code: https://github.com/dbferreira/go-bat-www

## Infrastructure
* The frontend will be implemented using AngularJS 2.0 and styled with Bootstrap.
* The backend server will be done with Go.  The backend will be responsible for the interface with the DB, and will execute algorithms to generate game results.
* The database will be Firebase -> https://github.com/zabawaba99/firego

## Planned Features
The following features are planned for the initial version:

1. Public home screen
2. User login screen, with basic authentication
3. Random team assigned to new users
4. Random players stats
5. Interface allows viewing of teams, nets and fixtures
6. Allow editing and saving of nets
7. Simulate basic games (no commentary, only results)
8. Basic user maintenance (update user fields)

## Firebase frontend/backend integration
* Implement handler similar to: https://firebase.googleblog.com/2015/05/introducing-firebase-queue_97.html
* Front-end to push to a task handler location on Firebase
* Back-end to listen on that location for new tasks and perform transaction
* Back-end to push to new location on Firebase on complete
* Front-end to listen for changes on the processed state data location

Sanichat is a nice example of how it works (Node.js) -> https://github.com/mcdonamp/sanichat

# Building

`go get github.com/dbferreira/go-bat`

`go run main.go`

