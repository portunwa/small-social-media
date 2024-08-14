# Small Social Media

## Instructions

1. Clone this repository and add the following environment variables to a `.env` file.

2. Start database server

```bash
docker compose up
```

3. Start the server

```bash
cd server
go mod download
go run .
```

4. Start the client

```bash
cd client
npm install
npm run dev
```

5. Run the tests

```bash
cd server
go test ./controllers -v
```

## Explore `docs` folder for more information
