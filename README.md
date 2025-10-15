
# 🚗 Car Rental API — v1 & v2

A backend service built with **Golang (Gin + GORM)** and **PostgreSQL** to manage a car rental system.  
This project is divided into two main phases:  
- **Car Rental v1:** CRUD for Customers, Cars, and Bookings  
- **Car Rental v2:** Adds Membership, Drivers, and Discount/Incentive features  

---

## 📖 API Documentation
You can explore and test all endpoints here:  
👉 [Car Rental API Postman Documentation](https://documenter.getpostman.com/view/39133117/2sB3QNonxF)

---

## 🧩 Tech Stack

- **Language:** Go 1.22+
- **Framework:** [Gin Gonic](https://github.com/gin-gonic/gin)
- **ORM:** [GORM](https://gorm.io)
- **Database:** PostgreSQL
- **Environment Management:** [joho/godotenv](https://github.com/joho/godotenv)

---

## 🗂️ Project Structure

```

car-rental/
├── cmd
│   ├── api
│   │   └── main.go                  # Application entrypoint
│   └── migration
│       └── 001_init.sql             # Initial SQL schema migration
├── common
│   └── database
│       └── database.go              # Database connection setup
├── features
│   ├── booking
│   │   ├── adapter/persistence/sql/booking_repository.go
│   │   ├── application/usecase/booking_usecase.go
│   │   ├── domain/entity/booking.entity.go
│   │   ├── domain/repository
│   │   ├── domain/service
│   │   └── presentation
│   │       ├── controller/booking_controller.go
│   │       └── dto
│   ├── car
│   │   ├── adapter/persistence/sql/car_repository.go
│   │   ├── application/usecase/car_usecase.go
│   │   ├── domain/entity/car.entity.go
│   │   ├── domain/repository
│   │   ├── domain/service
│   │   └── presentation
│   │       ├── controller/car_controller.go
│   │       └── dto
│   └── customer
│       ├── adapter/persistence/sql/customer_repository.go
│       ├── application/usecase/customer_usecase.go
│       ├── domain/entity/customer.entity.go
│       ├── domain/repository
│       ├── domain/service
│       └── presentation
│           ├── controller/customer_controller.go
│           └── dto
├── go.mod
└── go.sum

````

> 💡 The structure follows **Clean Architecture (DDD-inspired)** — separating domain, usecase, adapter, and presentation layers for better scalability and maintainability.

---

## 🚀 Getting Started

### 1️⃣ Clone Repository
```bash
git clone https://github.com/Arroziqi/car-rental-technical-test-pharos.git
cd car-rental-technical-test-pharos
````

### 2️⃣ Setup Environment

Create a `.env` file in the project root:

```
export DATABASE_URL="host=localhost user=postgres password=root dbname=car_rental port=5432 sslmode=disable TimeZone=Asia/Jakarta"
export PORT=8080
GIN_MODE=debug
```

> Make sure PostgreSQL is running and the database name matches your `.env`.

---

### 3️⃣ Run Database Migration

```bash
psql -U <username> -d car_rental -f cmd/migration/001_init.sql
```

---

### 4️⃣ Install Dependencies

```bash
go mod tidy
```

---

### 5️⃣ Run the API

```bash
go run cmd/api/main.go
```

You should see logs like:

```
[GIN-debug] Listening and serving HTTP on :8080
```

Now open `http://localhost:8080/api/v1/customers` (or use Postman collection above).

---

## 🧠 Project Overview

### **Car Rental — v1**

This version provides the fundamental API for managing car rentals.

#### ✅ Must Support:

* CRUD operations for:

    * Customers
    * Cars
    * Bookings

#### 🗄️ Expected Database Tables

* `customers`
* `cars`
* `bookings`

#### 🔗 Relationships

* One `customer` can have many `bookings`
* One `car` can have many `bookings`

---

### **Car Rental — v2**

Adds business logic enhancements and new entities.

#### 🚘 New Features

1. **Membership Program**

    * Customers can optionally join:

        * Bronze → 4% discount
        * Silver → 7% discount
        * Gold → 15% discount
    * Discount Formula:

      ```
      Discount = (Days_of_Rent * Daily_car_Rent) * Membership_discount
      ```

2. **Rent with Driver**

    * Customers can choose to rent:

        * Car Only
        * Car + Driver

3. **Driver Management**

    * Store driver data with their daily cost.
    * Calculate incentive per booking:

      ```
      Incentive = (Days_of_Rent * Daily_car_Rent) * 5%
      ```

#### 🗄️ Expected Database Changes

* Add `memberships` table (with discount rate)
* Add `drivers` table (with daily cost & incentive tracking)
* Extend `bookings` table:

    * `membership_id`
    * `driver_id`
    * `discount`
    * `book_type`

---

## 📚 API Collections

| Resource      | Method   | Endpoint                | Description         |
| ------------- | -------- | ----------------------- | ------------------- |
| **Customers** | `GET`    | `/api/v1/customers`     | List all customers  |
|               | `POST`   | `/api/v1/customers`     | Create new customer |
|               | `GET`    | `/api/v1/customers/:id` | Get customer by ID  |
|               | `PUT`    | `/api/v1/customers/:id` | Update customer     |
|               | `DELETE` | `/api/v1/customers/:id` | Delete customer     |
| **Cars**      | `GET`    | `/api/v1/cars`          | List all cars       |
|               | `POST`   | `/api/v1/cars`          | Create new car      |
|               | `GET`    | `/api/v1/cars/:id`      | Get car by ID       |
|               | `PUT`    | `/api/v1/cars/:id`      | Update car          |
|               | `DELETE` | `/api/v1/cars/:id`      | Delete car          |
| **Bookings**  | `GET`    | `/api/v1/bookings`      | List all bookings   |
|               | `POST`   | `/api/v1/bookings`      | Create booking      |
|               | `GET`    | `/api/v1/bookings/:id`  | Get booking by ID   |
|               | `PUT`    | `/api/v1/bookings/:id`  | Update booking      |
|               | `DELETE` | `/api/v1/bookings/:id`  | Delete booking      |

---

## 🧪 Testing

You can import the ready-to-use Postman Collection:

* [Open Postman Docs](https://documenter.getpostman.com/view/39133117/2sB3QNonxF)
* Click **“Run in Postman”** to import requests automatically.

---

## 🧑‍💻 Author

**Ahmad Arroziqi**
🌐 [Website](https://ahmad-arroziqi.vercel.app/)
📧 [LinkedIn](https://www.linkedin.com/in/ahmad-arroziqi-5a0566274/)
💻 [GitHub](https://github.com/Arroziqi)
---

## 📄 License

This project is created for educational and technical assessment purposes.

