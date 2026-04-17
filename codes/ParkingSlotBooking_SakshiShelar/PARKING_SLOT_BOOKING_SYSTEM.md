# Parking Slot Booking System - Problem Statement

## Overview

A Parking Slot Booking System allows users to reserve parking slots in advance, view their bookings, and manage their reservations. It helps optimize parking space utilization and provides a convenient way to secure parking spots before arrival.

---

## Core Features

1. **Login/Logout**
2. **View available parking slots**
3. **Book parking slots**
4. **View and cancel bookings**

---

## API Endpoints

### Auth APIs (2 Marks)


| Method | Endpoint       | Description |
| ------ | -------------- | ----------- |
| POST   | `/auth/login`  | User login  |
| POST   | `/auth/logout` | User logout |


### Parking APIs - Protected (4 Marks)


| Method | Endpoint                | Description                 |
| ------ | ----------------------- | --------------------------- |
| GET    | `/slots/available`      | Get available parking slots |
| POST   | `/bookings/create`      | Book a parking slot         |
| GET    | `/bookings/my-bookings` | Get user's bookings         |
| DELETE | `/bookings/cancel/{id}` | Cancel a booking            |


---

## Business Rules

### Parking Constraints

- **Total parking slots**: 100 (numbered 1 to 100)
- **Slot availability**: A slot is available if booked reservation doesn't exist
- **Slot availability management**: Use Redis cache or database to track free slots

### Booking Status Flow

- **BOOKED**: Initial status when booking is created
- **CANCELLED**: When user cancels the booking

---

## Database Schema

### TABLE: Users


| Column       | Type      | Constraints |
| ------------ | --------- | ----------- |
| `id`         | Integer   | Primary Key |
| `name`       | String    | -           |
| `email`      | String    | Unique      |
| `password`   | String    | Hashed      |
| `created_at` | Timestamp | -           |


### TABLE: Bookings


| Column           | Type      | Constraints                    |
| ---------------- | --------- | ------------------------------ |
| `id`             | Integer   | Primary Key                    |
| `user_id`        | Integer   | Foreign Key → Users.id         |
| `slot_number`    | Integer   | Range: 1-100                   |
| `vehicle_number` | String    | -                              |
| `status`         | String    | BOOKED / CANCELLED             |
| `created_at`     | Timestamp | -                              |
| `updated_at`     | Timestamp | -                              |

### TABLE: Slots

| Column | Type | Constraints |
|--------|------|-------------|
| `id` | Integer | Primary Key |
| `slot_number` | Integer | Unique, Range: 1-100 |
| `status` | String | BOOKED / FREE |


---

## Redis Configuration

### Session Management


| Field     | Value           |
| --------- | --------------- |
| **KEY**   | `JWT_TOKEN`     |
| **VALUE** | `user_id`       |
| **TTL**   | JWT expiry time |


### Available Slots Cache


| Field     | Value                                   |
| --------- | --------------------------------------- |
| **KEY**   | `available_slots`                       |
| **VALUE** | Array of available slot numbers [1-100] |


---

## Edge Cases ( 2 Marks)

### 🔐 Authentication / Authorization

- Password must contain at least one uppercase letter, one lowercase letter, one number, and one special character
- Using an expired JWT token
- Using an invalid / tampered JWT token
- APIs called with already invalid/expired token

### 🚗 Slot Availability

- Requesting slot number outside valid range (1-100)

### 📅 Booking Creation

- Booking a slot number outside valid range (not between 1-100)
- Booking a slot that is already booked 
- Missing required fields (slot_number, vehicle_number)
- Invalid vehicle_number format
- Booking without authentication

### 📋 Get Bookings

- User cannot see another user's booking (security issue)
- Invalid booking ID format (string instead of number)
- Fetching bookings with invalid filters

### ❌ Cancel Booking

- Cancelling a booking that does not exist
- Cancelling another user's booking
- Cancelling an already cancelled booking

---

## ⭐ Features ( 2 Marks)

1. Add pagination and filtering for `/bookings/my-bookings` (by status, date range)
2. Add search functionality for bookings (by slot number, vehicle number)
3. Add soft delete instead of hard delete

---

## Scoring Breakdown

- **Auth APIs**: 2 Marks
- **Parking APIs**: 4 Marks
- **Edge Cases**: 2 Marks
- **Features**: 2 Marks 

**Total**: 10 Marks 

---



