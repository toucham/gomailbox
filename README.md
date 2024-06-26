# Go Mailbox

A simple forward proxy in golang

## Motivation

A simple project to better understand network programming and Golang.

## Sidetrack 

Other stuff created along the way to understand Go's features and standard libs.

### Go PostService

A small CLI tool for analyzing packets from network to application layers. It is inspired by (httpstat)[https://github.com/davecheney/httpstat].
It have the following features:

- Sends a HTTP request
- Shows the time it takes for each event to take place (e.g. DNS lookup, TCP connection, and etc.)
- Shows header info of each layer (e.g. IP, TCP, and App headers)

