# Packform-Test

## Backend Project Setup

1. Install dependencies
```sh
go mod download
```
2. Copy paste `.env.template` and rename it to `.env`
3. Configure environment variables inside `.env` file.
4. Run migration & seed the data from the rood folder
```sh
go run .\migrate\migrate.go
```
5. Run the project
```sh
go run main.go
```

## Frontend Project Setup

1. Go to webapp folder
```sh
cd .\webapp\
```

2. Install dependecies
```sh
npm install
```

3. Copy paste `.env.template` and rename it to `.env`
4. Configure environment variables inside `.env` file. Make sure `VITE_BASE_API_URL` is the same as where the backend is running.
5. Run the project
```sh
npm run dev
```


## Run Unit Tests
### Backend unit test
```sh
go test ./...
```

### Frontend unit test

```sh
cd .\webapp\ 
npm run test:unit
```

## Run End-to-End Tests

1. Before running e2e test, both backend and frontend must be running.
2. Set `E2E_BASE_URL` inside webapp `.env` file. It should be the same as where the Web application running on
3. Install browers for the first run
```sh
npx playwright install
```
4. Run the test
```sh
# Runs the end-to-end tests
npm run test:e2e
```
