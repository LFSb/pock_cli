# Introduction
This is my first attempt at creating a cli using golang. The idea is simple: call this cli with a single url and by doing so, add said url to my pocket reading list. For this, we'll need to:

- Implement Oauth2 to get a token to interact with pocket
- Call the pocket api
- ???
- Profit

## Configuration
By default, the application is going to look for a yaml-formatted file `config` in the following directories:
- `/etc/pock_cli`
- `$HOME/.pock_cli`
- `.` (the directory in which this README file resides)

Look at the [config.example](./config.example) file for an example.

## Development
Each of the subcommands is its own separate file in the `./cmd`. Calling them is as easy as running `go run main.go COMMAND`. So to call `get.go` you run `go run main.go get`.