# CleverBank
Bank Software Simulation with some basic operation..

## Operation

- Create accounts (ahorro o corriente)
- Deposit
- Withdrawal
- Transfer from an account to another
- Login

You must be Logged before make a transaction

Run it with
go run cmd/main.go


## Requirenment
GOLANG 1.18
Postgresql 11


## Envoriment Vars

| Name              | Description                              |
|-------------------|------------------------------------------|
| JWT_SECRET        | Secret key for JWT Token Generation      |
| DB_HOST           | DB Server IP or URL (example: 127.0.0.1) |
| DB_PORT           | DB Port (example: 5432)                  |
| DB_USER           | User db (example: pguser)                |
| DB_PASS           | User password (example: 123456pswd)      |
| DB_NAME           | Name of Databse (example: bank)          |
