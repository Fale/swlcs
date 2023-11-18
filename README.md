# Static Website Lambda Commenting System

[![Go Report Card](https://goreportcard.com/badge/github.com/Fale/swlcs)](https://goreportcard.com/report/github.com/Fale/swlcs)
[![Build Status](https://travis-ci.org/Fale/swlcs.svg?branch=master)](https://travis-ci.org/Fale/swlcs)

## Build for AWS Lambda

    GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -tags lambda.norpc -o bootstrap main.go && rm -f function.zip && zip function.zip bootstrap
