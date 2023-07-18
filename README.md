# API

[![Static Badge](https://img.shields.io/badge/project%20use%20codesystem-green?link=https%3A%2F%2Fgithub.com%2Fgofast-pkg%2Fcodesystem)](https://github.com/gofast-pkg/codesystem)
![Build](https://github.com/gofast-pkg/api/actions/workflows/ci.yml/badge.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/gofast-pkg/api.svg)](https://pkg.go.dev/github.com/gofast-pkg/api)
[![codecov](https://codecov.io/gh/gofast-pkg/api/branch/main/graph/badge.svg?token=7TCE3QB21E)](https://codecov.io/gh/gofast-pkg/api)
[![Release](https://img.shields.io/github/release/gofast-pkg/api?style=flat-square)](https://github.com/gofast-pkg/api/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/gofast-pkg/api)](https://goreportcard.com/report/github.com/gofast-pkg/api)
[![codebeat badge](https://codebeat.co/badges/2d3d16b7-86b5-45c2-93e3-1fe94d6edbf8)](https://codebeat.co/projects/github-com-gofast-pkg-api-main)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/gofast-pkg/api/blob/main/LICENSE)

API is a package to provide http server with commons features. The core use the [echo framework](https://echo.labstack.com/).

## Install

``` bash
$> go get github.com/gofast-pkg/api@latest
```

## Usage

``` Golang
// Read documentation for more details usages
// With New(), api assume that you initialize the configuration with viper.
// Else, call the NewWithConfig() method to get you api instance.
import (
    "github.com/spf13/viper"
    "github.com/go-fast/api"
)

func main() {
    viper.AutomaticEnv()

    api, err := server.New()
    if err != nil {
        panic(err)
    }
    if err = s.Start(); err != nil {
        panic(err)
    }
}
```

## Contributing

&nbsp;:grey_exclamation:&nbsp; Use issues for everything

Read more informations with the [CONTRIBUTING_GUIDE](./.github/CONTRIBUTING.md)

For all changes, please update the CHANGELOG.txt file by replacing the existant content.

Thank you &nbsp;:pray:&nbsp;&nbsp;:+1:&nbsp;

<a href="https://github.com/gofast-pkg/api/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=gofast-pkg/api" />
</a>

Made with [contrib.rocks](https://contrib.rocks).

## Licence

[MIT](https://github.com/gofast-pkg/api/blob/main/LICENSE)
