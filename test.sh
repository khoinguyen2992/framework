#!/bin/bash

go test -cover $(go list ./... | grep -vE 'migrations|vendor|test')
go tool vet -structtags=false .