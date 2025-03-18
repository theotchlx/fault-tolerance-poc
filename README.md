# Proof of concept: Fault tolerance in a microservices architecture

This project is a Proof of Concept, in Go, for fault tolerance in a microservices architecture.  
It is composed of two services, and includes a 20mn presentation, with a demo. It explores the concept of fault tolerance in a distributed service architecture through the Circuit Breaker, Retry, Timeout, and Fallback patterns.

## What this project demonstrates

This project aims to demonstrate the main mechanisms and patterns of fault tolerance in a distributed architecture, including the Retry, Timeout and Fallback patterns with a manual UI panel to try them out in different use cases; as well as an circuit breaker mechanism, and a fault-intolerant service.

This project is made of a monorepo containing two Go programs, a "backend" service, and a "ui" service, showing text messages in the browser.

## How to use this demo

Checkout tag `1-dynamic-ui`. Open two terminals, one in each folder `./backend/` and `./ui/`. Run `go run main.go` in each directory to execute the UI service on `localhost:8080` and the backend service on `localhost:8081`.

```txt
.
├── backend/
│   ├── go.mod
│   ├── go.sum
│   └── main.go
├── LICENSE
├── README.md
└── ui/
    ├── go.mod
    ├── index.html
    ├── main.go
    └── pictures/
        ├── perpetua.jpg
        ├── teal2.jpg
        ├── teal.jpg
        └── x.webp
```

This tag shows a UI with minimal fault tolerance mechanisms. The backend simulates load by waiting 2 seconds before answering, allowing you (the user) to see the loading skeletons, although there's no timeout. You can try turning off the backend - the HTTP request will disconnect and the UI will show a fallback message. That's it.

Checkout tag `5-full-admin-panel`. The UI now features an admin panel that allows you to manually retry, setup auto-retry with a timer, etc.

On this tag, the backend exposes multiple routes that expose different faults. You can give the frontend an argument while running it. No arguments will fetch the backend route that always delivers messages.

- `go run main.go /flaky` to fetch the route that does not deliver the first few messages. To show off retries.
- `go run main.go /unreliable` to deliver only 50% of messages. To show off retries, auto-retries.
- `go run main.go /slow` that takes 20 seconds (instead of 2) to deliver messages. To trigger timeouts.
- `go run main.go /down` that's always down, always returns an error. To show off UI fallback, auto-retries flooding, etc.
- `go run main.go /fallback` always returns a fallback response. To demonstrate its catching by the UI for a seamless user experience.

During the demo, I ran the frontend with the `/flaky` argument to show off the retry mechanism.

Checkout tag `6-circuit-breaker`

Here, there's no interactive/real-time UI (as to code the circuit breaker mechanism in Go and not JS). You can still use the browser to visualize the messages, errors and fallbacks from the UI on port 8080.  
This tag is much simpler and there's no argument to the frontend, but you can still change the route in the ui code to show off circuit breaking behaviour on different backend fault scenarios (flaky, unreliable, down...) or just turn off the backend.

```txt
.
├── backend/
│   ├── go.mod
│   ├── go.sum
│   └── main.go
├── LICENSE
├── README.md
└── ui/
    ├── go.mod
    └── main.go
```

In my demonstration, I ran the UI on the `/flaky` route, to demonstrate the circuit breaker auto retry, then cut off the backend (~=unresponsive service) to show the circuit breaker mechanism in action. Running on route `/down` or `/unreliable` (with a bit of luck for the latter) would have given the same results. Or altering the `/flaky` route to be flakkier than the circuit's breaker maximum fault threshold tolerance constant. Or altering that constant.

As show in the presentation, the circuit breaker goes through each stage of being closed, open and half-open. The duration of the half-open state is set to 10 seconds by default. The number of retries is set to 3 by default.

Other tags show off some additional features to the UI admin panel, such as timeout (automatic or adjustable), and others demonstrated by the multiple backend routes; but it's mostly in JS.
