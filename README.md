# goki

A toy non-persistent(in-memory) Key-Value data store. Redis is the inspiration.
I'll try to model this as close to it as possible.

### Main Components

- `Server` - is a webserver that you can start with the command `goki-server start 8080`
- `Client` - is a REPL client that connects to the `goki-server`. Usage: `goki-client connect 8080`

##### Feature Tracker

- [x] Make Client REPL
- [x] Client Connects with server
- [x] Fixed Request and Response types
- [x] Add support for commands like `get` & `set` to start off
- [x] Print the results and errors in a better format on client REPL
- [ ] Improve Logging format on server. Richer with the actual details
- [ ] Add helper for similar commands? Provide suggestions like `Did you mean? set get...`
- [ ] Show all commands supported helper
- [ ] Good to have : Instead of using JSON over http, try and use some other lighter format for requests and response, like `bencode`
- [ ] Other commands - `ttl`, `incr`, `decr` (for integers), `persist` for certain keys, a new data structure `lists` - `l/r-push` & `l\r-pop`
- [ ] Add tests (for both client and server)

##### Supported Client Commands

- `set`
- `get`
- `del`-ete
- `exists`
- `expire`

- for `ttl` I'm thinking of changing the data struture slightlt TODO    

More information will be updated here soon.
