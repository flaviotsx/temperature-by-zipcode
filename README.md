# Temperature by zipcode

A tool for get teperature by zipcode

## Config

```bash
go mod tidy
```

### Running

You must configure your [Weather API](https://weatherapi.com/) key on `.env` file like the example `.env.sample`

```bash
go run cmd/main.go
```
#### Or
```bash
docker compose up -d
```

The server will be running on port `:8080`

## Usage

### Local

```bash
curl http://localhost:8080/28928728/temperature
```
