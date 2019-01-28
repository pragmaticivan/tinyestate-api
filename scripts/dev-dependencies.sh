#!/bin/bash

set -euo pipefail
IFS=$'\n\t'

function main() {
    passGoDevDependencies
}

function passGoDevDependencies() {
  echo "Installing developer tools."

  echo "gin https://github.com/codegangsta/gin"
  go get -u github.com/codegangsta/gin

  echo "goveralls https://github.com/mattn/goveralls"
  go get -u github.com/mattn/goveralls

  echo "errcheck https://github.com/kisielk/errcheck"
  go get -u github.com/kisielk/errcheck

  echo "golint https://github.com/golang/lint/golint"
  go get -u github.com/golang/lint/golint

  echo "megacheck https://github.com/dominikh/go-tools/tree/master/cmd/megacheck"
  go get -u honnef.co/go/tools/cmd/megacheck

  echo "swag github.com/swaggo/swag/cmd/swag"
  go get -u github.com/swaggo/swag/cmd/swag
}

main
