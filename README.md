# gonnpass

## Installation
### go get
```
go get github.com/anraku/gonnpass
```

## Usage
```
$ gonnpass help
NAME:
   gonnpass - A new cli application

USAGE:
   gonnpass [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --and value, -a value  use as and condition
   --or value, -o value   use as or condition
   --update-order, --uo   Sort events order by update_at
   --start-order, --so    Sort events order by start_at (default order)
   --new-order, --no      Sort events order by new_at
   -n value               number of events to print (default: 10)
   --help, -h             show help
   --version, -v          print the version
```

## How to use
### search evetns about golang
```
gonnpass -a golang
```
