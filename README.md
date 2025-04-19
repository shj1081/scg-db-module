# DB Module with Golang

## Prerequisites
- Go 1.21 or higher
- MySQL Server

## Setup & Configuration

1. Clone the repository
2. Install dependencies (not necessary):

   ```bash
   make setup
   ```
3. Set up environment variables by copying following `.env.example` to `.env`:

   ```dosini
   # .env.example
   # DB
   DB_DSN=root:1234@tcp(localhost:3306)/ # need tcp declaration
   DB_MAX_CONNS=25
   DB_MAX_IDLE_CONNS=25
   DB_CONN_MAX_LIFETIME=10m

   # Server
   PORT=8080
   ENV=development # adjust logging level by environment

   # Auth
   AUTH_PROXY_URL=http://localhost:8080/validate
    ```
4. Run the development server with hot reload (auto install dependencies):

   ```bash
   make dev
   ```
5. Build for production (auto install dependencies):

   ```bash
   make build
   ```
6. Run server (auto install dependencies):

   ```bash
   make run
   ```

## Swagger API Docs

after running the server, go to

```
http://localhost:8080/swagger/index.html
```


