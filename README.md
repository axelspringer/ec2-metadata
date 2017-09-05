[![Taylor Swift](https://img.shields.io/badge/secured%20by-taylor%20swift-brightgreen.svg)](https://twitter.com/SwiftOnSecurity)
[![Volkswagen](https://auchenberg.github.io/volkswagen/volkswargen_ci.svg?v=1)](https://github.com/auchenberg/volkswagen)
[![MIT license](http://img.shields.io/badge/license-MIT-brightgreen.svg)](http://opensource.org/licenses/MIT)

# EC2 Metadata

What you get is a simple server that reads in a json (e.g. `example.json`) from [EC2 Metadata](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html) and allows to use it local.

## Getting Started

> You can install go with `brew install go`

This project requires Go to be installed. You can also use the [Golang Docker](https://hub.docker.com/_/golang/) to build the binary.

Running is it should be as simple as:

```
$ make
$ make restore
$ make build
```

## Available Metadata

The currently supported metadata endpoints:

* [http://169.254.169.254/](http://169.254.169.254/)
* [http://169.254.169.254/latest/meta-data/](http://169.254.169.254/latest/meta-data/)
* [http://169.254.169.254/latest/user-data/](http://169.254.169.254/latest/user-data/)
* [http://169.254.169.254/latest/dynamic/](http://169.254.169.254/latest/dynamic/)

## Subcommands

### `help`

Displays the available options of the server.

### `version`

Displays the current version of the server.

## Configuration on OSX

> https://github.com/threadwaste/finto/wiki/Using-network-redirection

```
sudo ifconfig lo0 alias 169.254.169.254
```

```
docker run --rm -p 80:8111 -v `pwd`/ec2-metadata.json:/ec2-metadata.json pixelmilk/ec2-metadata --data /ec2-metadata.json
```

# License
[MIT](/LICENSE)

