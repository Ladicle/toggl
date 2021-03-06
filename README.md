# toggl CLI client

[toggl](https://toggl.com/) is a time tracking web application.
This program will let you use the toggl in CLI.

## Install

Quick install for OSX

```bash
$ brew tap Ladicle/toggl
$ brew install toggl
```

Build yourself

```bash
$ git clone https://github.com/Ladicle/toggl.git
$ cd toggl && make install
```

## Usage

![demo image](https://cloud.githubusercontent.com/assets/6121271/21588531/0108bd18-d12b-11e6-9fdc-e65aa1f15768.gif)

```
$ toggl --help
NAME:
   toggl - Toggl API CLI Client

USAGE:
   toggl [global options] command [command options] [arguments...]

VERSION:
   0.3.0

COMMANDS:
     start, a       Start time entry
     stop, s        End time entry
     current, c     Show current time entry
     workspaces, w  Show workspaces
     project, p     Options for projects
     local          Set current dir workspace
     global         Set global workspace
     help, h        Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --cache
   --csv
   --help, -h     show help
   --version, -v  print the version
```

### Register API token

When you run `toggl` first time, you will be asked your toggl API token.
Please input toggl API token and register it.
