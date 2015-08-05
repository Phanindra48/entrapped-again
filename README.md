# Entrapped
__Tiny little minefield of our dreams__



#### Cloning
```
Inside your $GOPATH/src create the following directory structure

$ mkdir github.com
$ cd github.com
$ mkdir kgthegreat

And then clone

$ git clone git@github.com:kgthegreat/entrapped-again.git

```

#### client setup 
```
change url in src/api/api.js to 'localhost:7000'

npm install 
npm run start

open in browser: http://localhost:3000/webpack-dev-server/

/* production */
npm run build
```

#### server setup
```
change directory paths in entrapped.go

go run main.go (in directory: server/cmd/entrapped)
```

#### TODOs:
##### Client
- Validation on homepage for nickname. (nickname should not be empty before starting the game)
- Handle errors from server.
- Add extra life to minefield.
- Add time limit UI to each game round.
- Add capability to exit a game.
- Animations on losing a life, opening a block.
- JSON data handling.
- Show a bomb icon on places where bomb exploded.

##### Server
- Make user nickname unique.
- Add time limit to each game round.
- Add capability of adding minefield for opponent.
- Add capability to start a new game on same connection.
- JSON data handling.
- Minefield should also have hidden lifes.
- Tests for the server code.

**Surprise us (bonus)**
