# goki

A toy non-persistent(in-memory) Key-Value data store. Redis is the inspiration.
I'll try to model this repo as close to redis as possible.

### Allowed Data Types

- Strings
- Integer
- Float (although redis doesn't really support, I've added it anyway)

### Main Components

- `Server` - is a webserver that you can start with the command `goki-server start 8080`
- `Client` - is a REPL client that connects to the `goki-server`. Usage: `goki-client connect 8080`

##### Feature Tracker

- [x] Make Client REPL
- [x] Client Connects with server
- [x] Fixed Request and Response types
- [x] Print the results and errors in a better format on client REPL
- [ ] Check Mutex locking and unlocking behaviour for the methods that interact with the store
- [ ] Improve Logging format on server. Make it richer with the essential details alone
- [ ] Add helper for similar commands? Provide suggestions like `Did you mean? set get...`
- [ ] Show all commands supported helper
- [ ] Good to have : Instead of using JSON over http, try and use some other lighter format for requests and response, like `bencode`
- [ ] Commands tracker: `persist` for certain keys, a new data structure `lists` - `l/r-push` & `l\r-pop`
- [ ] Add tests (for both client and server)

##### Supported Client Commands

- `set`
- `get`
- `del`
- `exists`
- `expire`
- `ttl`
- `incr`
- `decr`

### Follow these steps to contribute or set up the dev env
 
 - fork/clone this repo
 - make you sure you have `go` installed and your `GOPATH` setup since we use go modules.
 - to run the **server** you can use the command `go run server/main.go start 9000`
 - to run the **client** you can use the command `go run client/main.go connect 9000`
 - Make your changes! :)

Also, it'll be great if you could help me set up `realize` or any other golang task runners for this!

Any feedback, suggestions and criticism's are very much welcome! :)

More information will be updated here soon.
