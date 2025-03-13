# Proof of concept: Fault tolerance in a microservices architecture

This project is a Proof of Concept, in Go, for fault tolerance in a microservices architecture.  
It is composed of two services, and includes a 20mn presentation, with a demo. It explores the concept of fault tolerance in a distributed service architecture through the Circuit Breaker, Retry, Timeout, and Fallback patterns.

WIP README:  
(not in order, or not always relevant parts of this readme:) What, why, how, repo/project structure, project dependencies, instructions to build and run the demo.

## What is fault tolerance?

## Why fault tolerance?

Let's say you have two services. For example:

- a **backend** service, which stores and serves messages via an exposed REST API endpoint.

- a **UI** service, which consumes the backend service's endpoint, and serves a web page to the user, displaying the messages with style sheets.

The point is, the UI service is dependent on the backend service's ability to serve the message, and generally be available through the network.

Now a number of fundamentally different issues could arise in this system:

- The backend service could crash and become unavaible, if no measures for high-availability, such replication, are in place.
- The network could fail, rendering communication impossible between the UI and backend services.
- The backend service could be waiting indefinitely for a resource such as storage or network, or for any other reason, is not able to close its network requests, until it exhausts its allocated thread pool and cannot handle any more requests.
- more

We can classify/categorize these issues into two types/classes/categories:

- good (unavaiable, handle directly. Fast. Cache them and test on them.)
- bad (timeout, have to wait)
- more?

## How is fault tolerance implemented?

---

Presentation:

Presentation of the concepts/technologies + demo. Max 15/20mn length.

- Cover the topic.
- Simple working demo illustrating the presented concepts.
- Brief presentation to introduce the project.

Demo implemented in Golang.

- Comprehensive readme that clearly and concisely explain:
  - project setup,
  - dependencies,
  - instructions on how to build and run the demo.

---

This presentation should demonstrate:

- how to implement and test fault tolerance patterns in a microservices architecture.
- Two services communicating with each other, one dependent on the other. Patterns to implement (and explain(=understand)):
  - Circuit Breaker,
  - Retry,
  - Timeout,
  - and Fallback.

Include the following features :

- Main service calls a dependent service.
- Failure simulation in the dependent service.
- API endpoints to trigger various failure scenarios (protected/unprotected to certain faults), or other way? Proxy? Maybe only just talk abt the proxy (not in demo).

- Explore cascading failure prevention techniques (in meshed services, real use case, titanic effect).
