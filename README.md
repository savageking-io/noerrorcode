# noerrorcode

Game backend solution for everyone. This project is in active development.
Supports Steam. More platforms could be added in the future.
So, with NoErrorCode you can enrich your games with:

* Achievements
* Lobbies
* In-game economy
* Character progression
* Friends
* Matchmaking
* ...and probably more (It's WIP, you know)

NoErrorCode uses WebSockets for communication, meaning your game must support this protocol. Most of the popular engines already support it.

### Setup and configuration

Setting up NoErrorCode is pretty easy. Get the sources, run `go get` and then run `make`. Now it's up to you how to run it - as a service, daemon, 
or just start it in your terminal. The only thing is needed is configuration that can be passed to the app using `--config=/path/to/config.yaml` argument.

An example config file:

```

```

