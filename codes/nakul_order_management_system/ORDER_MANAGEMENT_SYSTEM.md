# Order Management System - Problem Statement

## Overview
An Order Management System is used to manage customer orders in a simple and organized way. It allows users to place orders, view their order details, make changes, or cancel orders when needed.

---

## Core Features

1. **Login/Logout**
2. **Create orders**
3. **View their orders**
4. **Update or cancel orders**

---

## API Endpoints

### Auth APIs (2 Marks)

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/auth/login` | User login |
| POST | `/auth/logout` | User logout |

### Order APIs - Protected (4 Marks)

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/orders/create` | Create a new order |
| GET | `/orders/list` | Get all orders for the user |
| PUT | `/orders/update/{id}` | Update an existing order |
| DELETE | `/orders/cancel/{id}` | Cancel an order |

---

## Database Schema

### TABLE: Users

| Column | Type | Constraints |
|--------|------|-------------|
| `id` | Integer | Primary Key |
| `name` | String | - |
| `email` | String | Unique |
| `password` | String | Hashed |
| `created_at` | Timestamp | - |

### TABLE: Orders

| Column | Type | Constraints |
|--------|------|-------------|
| `id` | Integer | Primary Key |
| `user_id` | Integer | Foreign Key → Users.id |
| `product_name` | String | - |
| `quantity` | Integer | - |
| `price` | Decimal | - |
| `status` | String | CREATED / PLACED / CANCELLED |
| `created_at` | Timestamp | - |
| `updated_at` | Timestamp | - |

---

## Redis Configuration

| Field | Value |
|-------|-------|
| **KEY** | `JWT_TOKEN` |
| **VALUE** | `user_id` |
| **TTL** | JWT expiry time |

---

## Edge Cases (2 Marks)

### 🔐 Authentication / Authorization

- Password must contain at least one uppercase letter, one lowercase letter, one number, and one special character
- Using an expired JWT token
- Using an invalid / tampered JWT token
- User trying to access another user's order by ID
- APIs called with already invalid/expired token

### 📦 Order Creation

- Creating order with quantity = 0 or negative
- Creating order with price = 0 or negative
- Missing required fields (product_name, quantity, price)
- Sending invalid data types (string instead of number)
- Creating order without authentication

### 📋 Get Orders

- User tries to fetch an order ID that does not exist
- User tries to fetch another user's order (security issue)
- Invalid order ID format (string instead of number)

### ✏️ Update Order

- Updating an order that does not exist
- Updating an order that belongs to another user
- Updating a cancelled order
- Updating with invalid fields (e.g., negative quantity)

### ❌ Cancel Order

- Cancelling an order that does not exist
- Cancelling already cancelled order
- Cancelling another user's order

---

## ⭐ Features (2 Marks)

1. Add pagination for `/orders/list`
2. Add filtering on get orders (by status, date range)
3. Add soft delete instead of hard delete
4. Add Role-Based Access Control (RBAC)
   - 👤 Roles: `USER`, `ADMIN`

---

## Scoring Breakdown

- **Auth APIs**: 2 Marks
- **Order APIs**: 4 Marks
- **Edge Cases**: 2 Marks
- **Features**: 2 Marks

**Total**: 10 Marks 