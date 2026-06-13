# Tickitz App - Backend

[![License: MIT](https://img.shields.io/badge/License-MIT-blue)](https://opensource.org/license/mit)

## About Tickitz

Tickitz is a fast, secure, and scalable movie ticket booking backend system.
Built with Go, PostgreSQL, and Redis, it handles:

- User management & authentication (JWT + argon2)

- Movie catalog, schedules, genres, and seat availability

- Booking & payment processing

- Order history for users

- Movie inventory management

From browsing movies to confirming your seat. Tickitz works behind the scenes to keep your cinema experience smooth.

## Technologies Used

- [![Go](https://img.shields.io/badge/Go-1.24-00ADD8?logo=go&logoColor=white)](https://go.dev/)
- [![Gin](https://img.shields.io/badge/Gin-Framework-00ADD8?logo=go&logoColor=white)](https://gin-gonic.com/)
- [![PostgreSQL](https://img.shields.io/badge/PostgreSQL-16.13-4169E1?logo=postgresql&logoColor=white)](https://www.postgresql.org/)
- [![Redis](https://img.shields.io/badge/Redis-8.6.3-FF4438?logo=redis&logoColor=white)](https://redis.io/)
- [![JWT](https://img.shields.io/badge/JWT-Auth-000000?logo=jsonwebtokens&logoColor=white)](https://jwt.io/)
- [![Swagger](https://img.shields.io/badge/Swagger-Docs-85EA2D?logo=swagger&logoColor=white)](https://swagger.io/)
- [![Docker](https://img.shields.io/badge/Docker-29.5.2-2496ED?logo=docker&logoColor=white)](https://www.docker.com/)

## Environment

```bash
APP_HOST=<your_app_host>
APP_PORT=<your_app_port>

DB_HOST=<your_database_host>
DB_PORT=<your_database_port>
DB_USER=<your_database_user>
DB_PASS=<your_database_password>
DB_NAME=<your_database_name>

JWT_ISSUER=<your_jwt_issuer>
JWT_SECRET=<your_jwt_secret>

RDB_USER=<your_redis_user>
RDB_PASS=<your_redis_password>

SMTP_PORT=<your_smpt_port>
SMTP_USER=<your_smpt_uset>
SMTP_PASSWORD=<your_smpt_password>
SMTP_FROM_EMAIL=<your_smpt_email>
```

## ⚙️ Installation

1. Clone the project

```sh
$ https://github.com/L1mus/Tickitz-backend.git
```

2. Navigate to project directory

```sh
$ cd Tickitz-backend
```

3. Install dependencies

```sh
$ go mod tidy
```

4. Setup your [environment](##-environment)

5. Install [migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate#installation) for DB migration

6. Do the DB Migration

```sh
$ migrate -database YOUR_DATABASE_URL -path ./db/migrations up
```

or if you install Makefile run command

```sh
$ make migrate-up
```

and seeding data

```sh
$ make db-seed
```

7. Run the project

```sh
$ go run ./cmd/main.go
```

## API Endpoint

| Method | Endpoint                             | Description                                                               |
| ------ | ------------------------------------ | ------------------------------------------------------------------------- |
| POST   | `/auth/register`                     | Register new user                                                         |
| POST   | `/auth`                              | Login user and admin                                                      |
| DELETE | `/auth/logout`                       | Logout                                                                    |
| POST   | `/auth/check-email`                  | Request an OTP to reset user password                                     |
| POST   | `/auth/check-email-verify-otp`       | Verify the OTP sent for password reset                                    |
| POST   | `/auth/check-email/verify-otp/reset` | Set a new password after successful OTP verification                      |
| POST   | `/auth/register/activate`            | Verify OTP to activate the user account                                   |
| POST   | `/auth/register/resend-otp`          | Resend activation OTP to the user's email                                 |
| GET    | `/users/profile`                     | Get the profile data of the logged in user based on the jwt token         |
| PATCH  | `/users/profile`                     | Update user data profile & change password                                |
| GET    | `/users/history`                     | Get list of Order History for logged-in User                              |
| GET    | `/users/history/{id}/detail`         | Get detailed information for specific order history by user logged-in     |
| GET    | `/movies`                            | fetch all movie data to be displayed on the main page                     |
|  |
| GET    | `/movies/:id`                        | Get the Movie detail data                                                 |
| GET    | `/movies/:id/showtime`               | Get the Showtime data filtered result                                     |
| GET    | `/order/seats`                       | Get seats data                                                            |
| POST   | `/order/booking`                     | Create Booking seat on cinema                                             |
| POST   | `/transactions/checkout`             | DOKU Checkout                                                             |
| GET    | `/transactions/doku-callback`        | Web hook notification Listener                                            |
| POST   | `/transactions/confirm`              | Confirm ticket payment                                                    |
| GET    | `/transactions/payment`              | Retrieve payment summary information                                      |
| GET    | `/transactions/qr`                   | Generates a PNG format QR Code                                            |
| GET    | `/transactions/result`               | Retrieve final invoice data for digital tickets after successful payment. |
| GET    | `/admin/movies`                      | Get admin List Movies                                                     |
| POST   | `/admin/movies`                      | Add New Movie                                                             |
| PATCH  | `/admin/movies/{id}`                 | Edit Movie                                                                |
| GET    | `/admin/sales-chart`                 | Get Revenue Sales Chart Data                                              |
| GET    | `/admin/ticket-sales`                | Get Ticket Sales Data                                                     |

Full interactive docs available at `/swagger/index.html` after running the server.

## Changelog
| Version | Description |
| ------- | ----------- |
| 1.0  | Setup Docker multi-stage build and docker-compose orchestration with PostgreSQL & Redis by [Alpha Team](https://github.com/iamhanif11, https://github.com/L1mus, https://github.com/Ilhammursidi, https://github.com/aqilknz, https://github.com/Akmalrian) |

## How to Contribute
- Fork this repository
- Create your changes
- Commit your changes (Please strictly follow the [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) standard: `feat:`, `fix:`, `chore:`, `docs:`)
- Push to the branch
- Open a Pull Request

## Related Project
[Frontend Tickitz Repository](https://github.com/L1mus/Tickitz-frontendd)

Copyright (c) 2026 Alpha Team