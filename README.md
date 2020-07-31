# Static Website Lambda Commenting System

[![Go Report Card](https://goreportcard.com/badge/github.com/Fale/swlcs)](https://goreportcard.com/report/github.com/Fale/swlcs)
[![Build Status](https://travis-ci.org/Fale/swlcs.svg?branch=master)](https://travis-ci.org/Fale/swlcs)

## Build for AWS Lambda

    GOOS=linux go build main.go && rm function.zip && zip function.zip main
