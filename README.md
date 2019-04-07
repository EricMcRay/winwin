# WinWin

WinWin simulates football league matches and shows final scores table ordered by scores.

This project is just for fun. Inspired from Atölye15's 2019 Interns Bootcamp tasks.

## Setup

After clone project to you working enviroment, just run `go get`

***This project uses go modules so you can run it from any path.***

### Running

Easiest way, just `go run ./cli``

Also, you can build cli with `go build` and run executable.

### Config

You can change mock teams which located at `./cli/mock.go`

*Right now, there is no options to configurate teams via CLI.*

#### Tuning Team Strength

You can think of team strength as maximum number of goals, the team can do in 90 minutes againest weakest opponent.
