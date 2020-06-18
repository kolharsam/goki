# goki

A toy non-persistent Key-Value data store

### Main Components

- `Server` - is a webserver that you can start with the command `goki-server start 8080`
- `Client` - is a REPL client that connects to the `goki-server`. Usage: `goki-client connect 8080`

##### Feature Tracker

- [x] Make Client REPL
- [x] Client Connects with server
- [x] Fixed Request and Response types
- [x] Add support for commands like `get` & `set` to start off
- [ ] Print the results and errors in a better format on client REPL
- [ ] Improve Logging format on server. Richer with the actual details
- [ ] Add helper for similar commands? Provide suggestions like `Did you mean? set get...`
- [ ] Show all commands supported helper
- [ ] Other commands - `expiry`, a new data structure `lists` - `l/r-push` & `l\r-pop`
- [ ] Add tests (for both client and server)

##### Supported Methods

- `set`
- `get`
- `del`-ete
- `exists`

More information will be updated here soon.
