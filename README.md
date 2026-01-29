# 🎟️ Ticket Booking CLI

A simple CLI application to **list events** and **book tickets**, with persistent storage using **SQLite**.

This project was built as an upgrade from an **in-memory system** to a **database-backed system**, following real backend engineering practices.

---

## ✨ Features

- ✅ Persistent storage using **SQLite**
- ✅ List / fetch events
- ✅ Book tickets for an event
- ✅ Stores bookings in database
- ✅ Simple CLI workflow (fast + minimal)

---

## 🛠️ Tech Stack

- **Go (Golang)**
- **SQLite**
- **GORM**

---

## 📦 Installation

### 1) Clone the repository

```bash
git clone https://github.com/HarshithRajesh/BookIt_Go.git
cd BookIt_Go
```

### 2) Install dependencies

```bash
go mod tidy
```

---

## ▶️ Run the CLI

```bash
go build
./BookIt_Go --help
```

## ⚙️ Database

This project uses **SQLite** for storage.

A local database file will be created automatically (example):

- `local.db`

> You can delete the `.db` file anytime to reset the application.

---

## 📌 Commands (Examples)

> Update these commands based on your actual CLI interface.

### List all events

```bash
./BookIt_Go getev
```

### Book tickets

```bash
./BookIt_Go booking 1 915xxxxxxx H  122
```

### View bookings (optional)

```bash
./BookIt_Go bi
```

---

## 🚀 Future Improvements (Optional)

- Authentication (user-based bookings)
- Prevent double booking / ticket limit handling
- Upgrade SQLite → PostgreSQL (Neon)
- Booking cancellation
- Admin event creation

---

## 👨‍💻 Author

**Harshith Rajesh**
Backend Developer (Go + Databases + System Design)

---
