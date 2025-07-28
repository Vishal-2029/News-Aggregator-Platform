## News Aggregator Platform

`A News Aggregator Platform is a software application or service that collects news articles from various sources (like websites, RSS feeds, APIs) and displays them in one place, usually categorized or filtered by topic, source, date, or relevance.`
This project uses the `NewsData.io` platform to fetch news articles from various sources.

## Technologies Used

  - **Golang** – Core programming language for backend development.
  - **Cron Job** – For scheduling automatic tasks like fetching news periodically.
  - **gRPC** – For efficient communication between microservices.
  - **NATS** – For asynchronous pub/sub messaging between services.
  - **REST API** – To expose data externally for clients, built using Fiber.
  - **MySQL** – Relational database to store normalized news data.
  - **NewsData.io API** – External news source used to fetch real-time news.

## Architecture Diagram

```text
                                ┌────────────────────────┐
                                │     Cron Scheduler     │
                                │ (robfig/cron in Go)     │
                                └──────────┬─────────────┘
                                           │
                                           ▼
                                ┌────────────────────────┐
                                │  News Fetcher Service  │
                                │ (Calls NewsData.io API)│
                                └──────────┬─────────────┘
                                           │
                                  Normalize & Format
                                           │
                                           ▼
                                ┌────────────────────────┐
                                │     NATS Publisher      │
                                │ (nats.go in Go)         │
                                └──────────┬─────────────┘
                                           │
                                           ▼
                                ┌────────────────────────┐
                                │     NATS Subscriber     │
                                │   (Database Service)    │
                                └──────────┬─────────────┘
                                           │
                                           ▼
                                ┌────────────────────────┐
                                │     MySQL Database      │
                                │ (news table structure)  │
                                └──────────┬─────────────┘
                            ┌──────────────┴───────────────┐
                            ▼                              ▼
                ┌────────────────────┐         ┌────────────────────┐
                │    gRPC Server     │         │     REST API       │
                │ (google.golang.org │         │   (Fiber in Go)    │
                │       /grpc)       │         └────────────────────┘
                └────────────────────┘
                            ▼
                 ┌────────────────────┐
                 │  External Clients  │
                 │  (Mobile/Web Apps) │
                 └────────────────────┘
```

## Folder Structure

```text 

news-aggregator-platform/
│
├── cmd/                         # Application entry points
│   ├── REST-Api/                # REST API using Fiber
│   │   ├── rest.go              # Fiber setup and routes
│   └── grpc.go                  # gRPC client (if needed in REST)
│   └── main.go                  # Starts cron, NATS, gRPC
│
├── config/                      # Configuration files
│   └── db/                      # DB connection logic
│       └── db.go
│
├── handlers/                    # Business logic handlers
│   └── news_handler.go
│
├── internal/                    # Internal-only packages
│   └── grpc/                    # gRPC server logic
│       └── server.go
│
├── models/                      # Data models (structs)
│   └── model.go
│
├── news-service/                # Generated gRPC files
│   └── pb/
│       ├── news_grpc.pb.go
│       └── news.pb.go
│
├── pkg/                         # Core application logic (reusable)
│   ├── fetcher/                 # News fetching and scheduling
│   │   ├── cron.go              # Cron job setup
│   │   └── news.go              # Fetches news from NewsData.io
│   ├── nats/                    # NATS Pub/Sub logic
│   │   ├── nats_publisher.go
│   │   └── nats_subscriber.go
│
├── proto/                       # Protobuf definitions
│   └── news/
│       └── news.proto
│
├── utility/                     # Utility functions
│   └── logger.go                # Logger setup (Zap or log)
│
├── app.log                      # Log file (optional)
├── go.mod                       # Go module definition
├── go.sum                       # Dependency checksums
├── main.go                      # (optional) alternative entry point
├── Makefile                     # Commands for build/run/test
└── README.md                    # Project documentation
```

## Run the Project 

  1. Fetch News
     `go run main.go`

 2. For REST Api
     `go run cmd/REST-Api/rest.go`
    
 3. Make run
    `make run` - for fetch news
    `make restapi` - for rest api

## REST API Endpoints
  - For Get NEWS
      `- http://localhost:3000/news`
  - For refresh   
      `- http://localhost:3000/refresh`

`In this project, I want to use an API that fetches only 10 news articles at a time. I want to fetch each article one by one every 5 seconds. After fetching all 10 articles, the code should not stop — it should wait for 5 minutes and check again for new articles. If new articles are available, it should start fetching them again.`


  
