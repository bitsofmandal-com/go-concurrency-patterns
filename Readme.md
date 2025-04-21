# Golang Concurrency Patterns

[![Golang Concurrency Patterns](https://img.shields.io/badge/Watch%20Video-YouTube-red)](https://www.youtube.com/watch?v=mm4ztXwysLk)

<div align="center">
  <a href="https://www.youtube.com/watch?v=mm4ztXwysLk">
    <img src="https://i.ytimg.com/vi/mm4ztXwysLk/maxresdefault.jpg" alt="Golang Concurrency Patterns Video Tutorial" />
  </a>
  <p><em>5 Most Used Golang Concurrency Patterns That 10x Your Application Performance</em></p>
</div>

## Table of Contents

- [Golang Concurrency Patterns](#golang-concurrency-patterns)
  - [Table of Contents](#table-of-contents)
  - [About This Repository](#about-this-repository)
  - [Introduction](#introduction)
  - [Pattern 1: Worker Pool](#pattern-1-worker-pool)
  - [Pattern 2: Pipelines](#pattern-2-pipelines)
  - [Pattern 3: Fan-Out, Fan-In](#pattern-3-fan-out-fan-in)
  - [Pattern 4: Mutex Pattern](#pattern-4-mutex-pattern)
  - [Pattern 5: Semaphore Pattern](#pattern-5-semaphore-pattern)
  - [Contributing](#contributing)
  - [License](#license)


## About This Repository

This repository contains practical examples and implementations of essential concurrency patterns in Go. These patterns will help you build scalable, efficient, and robust applications by leveraging Go's powerful concurrency features.

👉 **[Subscribe to our YouTube channel](https://www.youtube.com/channel/UCdOb1Se6eXYpYHD8HWFXtWg)** for more Go programming tutorials!


## Introduction

Concurrency patterns are well-established solutions to common problems encountered in concurrent programming. They help us write robust and efficient concurrent code by providing structured approaches to manage goroutines, synchronize data access, and facilitate communication between concurrent tasks.

Most Go developers don't fully utilize concurrency patterns while writing concurrent code. This repository aims to demonstrate how these patterns can be implemented effectively to solve real-world problems.

## Pattern 1: Worker Pool

The Worker Pool pattern is one of the most common and straightforward concurrency patterns in Go. It creates a fixed number of worker goroutines that process tasks from a shared queue.

**Key Components:**
- Task Queue (channel)
- Worker goroutines
- Result channel
- Synchronization (WaitGroup)

**Code Example:** [/workerpool/example1.go](https://github.com/bitsofmandal-com/go-concurrency-patterns/blob/main/workerpool/example1.go)

**Use Cases:**
- Web servers handling multiple client requests
- Batch data processing
- Rate-limited operations

## Pattern 2: Pipelines

Pipelines are perfect for operations that involve a series of stages where each stage depends on the output of the previous one.

**Key Components:**
- Independent stages as goroutines
- Channels connecting the stages
- Proper error handling and propagation

**Code Example:** [/pipelines/example1.go](https://github.com/bitsofmandal-com/go-concurrency-patterns/blob/main/pipelines/example1.go)

**Use Cases:**
- ETL (Extract, Transform, Load) processes
- Image processing with multiple transformations
- Multi-stage data validation and enrichment

## Pattern 3: Fan-Out, Fan-In

When you need maximum parallelism, especially for CPU-intensive or I/O-bound operations, the Fan-Out, Fan-In pattern is ideal.

**Key Components:**
- Input channel for distributing work
- Multiple worker goroutines processing in parallel
- Output channel for collecting results
- Synchronization for coordinating completion

**Code Example:** [/faninout/example1.go](https://github.com/bitsofmandal-com/go-concurrency-patterns/blob/main/faninout/example1.go)


**Use Cases:**
- Parallel data processing
- Concurrent API requests
- File operations across multiple sources

## Pattern 4: Mutex Pattern

The Mutex pattern is essential when multiple goroutines need to access shared resources safely.

**Key Components:**
- Mutex for exclusive access
- Critical section protection
- Shared resource management

**Code Example:** [/mutex/example1.go](https://github.com/bitsofmandal-com/go-concurrency-patterns/blob/main/mutex/example1.go)

**Use Cases:**
- Shared counters and state
- In-memory caches
- Connection pools

## Pattern 5: Semaphore Pattern

The Semaphore Pattern limits the number of goroutines that can simultaneously access a resource, preventing system overload.

**Key Components:**
- Counting semaphore (buffered channel)
- Acquire and release operations
- Controlled concurrency

**Code Example:** [/semaphores/example1.go](https://github.com/bitsofmandal-com/go-concurrency-patterns/blob/main/semaphores/example1.go)

**Use Cases:**
- Database connection limiting
- API rate limiting
- Resource-intensive operations
- External service protection
<!-- 
## Use Cases

These patterns can be applied in various real-world scenarios:

- **Data Processing (ETL)** - [examples/use-cases/etl.go](examples/use-cases/etl.go)
- **Image Processing** - [examples/use-cases/image-processing.go](examples/use-cases/image-processing.go)
- **Web Search Engine** - [examples/use-cases/search-engine.go](examples/use-cases/search-engine.go)
- **Web Scrapers** - [examples/use-cases/web-scraper.go](examples/use-cases/web-scraper.go) -->

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request or open an Issue if you have any suggestions, questions, or improvements.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

<div align="center">
  <p>If you found this helpful, please star the repository and consider <a href="https://www.youtube.com/channel/UCdOb1Se6eXYpYHD8HWFXtWg?sub_confirmation=1">subscribing to our YouTube channel</a>!</p>
  
  <p>
    <a href="https://twitter.com/sourabhmandal_">
      <img src="https://img.shields.io/twitter/follow/sourabhmandal_?style=social" alt="Twitter Follow" />
    </a>
    <a href="https://github.com/sourabhmandal">
      <img src="https://img.shields.io/github/followers/sourabhmandal?style=social" alt="GitHub followers" />
    </a>
  </p>
</div>
