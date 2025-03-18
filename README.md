# Proof of concept: Fault tolerance in a microservices architecture

This project is a Proof of Concept, in Go, for fault tolerance in a microservices architecture.  
It is composed of two services, and includes a 20mn presentation, with a demo. It explores the concept of fault tolerance in a distributed service architecture through the Circuit Breaker, Retry, Timeout, and Fallback patterns.

## What this project demonstrates

This project aims to demonstrate the main mechanisms and patterns of fault tolerance in a distributed architecture, including the Retry, Timeout and Fallback patterns with a manual UI panel to try them out in different use cases; as well as an circuit breaker mechanism, and a fault-intolerant service.

This project is made of a monorepo containing two Go programs, a "backend" service, and a "ui" service, showing text messages in the browser.

## How to use this demo

tag 1

tag 5

tag 6-
```
.
├── backend
│   ├── go.mod
│   ├── go.sum
│   └── main.go
├── LICENSE
├── README.md
└── ui
    ├── go.mod
    ├── main.go
```

flaky response, backend on at first, then off to demonstrate full circuit breaking staying open.

