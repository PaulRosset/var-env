# var-env <img src="computer.png" alt="Logo" height="35" align="top" />

#### Create persistent env variables easily and quickly for testing locally.

[![Travis CI Build Status](https://travis-ci.org/PaulRosset/var-env.svg?branch=master)](https://travis-ci.org/PaulRosset/var-env)
[![Go Report Card](https://goreportcard.com/badge/github.com/PaulRosset/var-env)](https://goreportcard.com/report/github.com/PaulRosset/var-env)
[![](https://cover.run/go/github.com/PaulRosset/var-env.svg)](https://cover.run/go/github.com/PaulRosset/var-env)
[![MIT Licence](https://badges.frapsoft.com/os/mit/mit.svg?v=103)](https://opensource.org/licenses/mit-license.php)

## Install

`go get github.com/PaulRosset/var-env/cmd/varenv`

## Usage

```sh
$ varenv
   varenv - Create persistent env variables easily and quickly for testing locally.

USAGE:
   cli [global options] command [command options] [arguments...]

VERSION:
   1.0.0

COMMANDS:
     load        The yaml file where the variables live in, from `FILE`
     list, l     List all the persistent environment variables on your computer.
     add, a      Quick commands to add an env variables through the cli interface
     remove, rm  Quick commands to remove an env variables through the cli interface
     help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

You can either load a `yaml` file of the following format:

```
[NAME_VAR]: [value]
```

Example:

```
ENV: prod
TOKEN_API: azeqcvhqsd677T4aze
```

Then:
`varenv load ./path/of/the/yaml/file`

or either load manually the variables with the `add` command, like the following:

`varenv add ENV=PROD TOKEN_API=azeqcvhqsd677T4aze`

Finally, if you want the change effective on your current shell, you need to do the following in order to refresh: `source ~/.bashrc`

## License

MIT Paul Rosset
