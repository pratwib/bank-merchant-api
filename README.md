# Bank Merchant API

## How to Run

1. Clone the repository
2. Run `go mod tidy`
3. Run `go run main.go`
4. The application will start on `http://localhost:8080`

## API Endpoints

- `POST /auth/register`
- `POST /auth/login`
- `POST /customers/balance/add`
- `POST /customers/payment/pay`
- `POST /merchants`
- `GET /merchants/{id}`
- `DELETE /merchants/{id}`