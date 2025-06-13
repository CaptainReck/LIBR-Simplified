# 📚 LIBR Simplified Prototype

A very minimal simulation of the LIBR system implementing decentralized moderation using Go concurrency, PostgreSQL, and RESTful APIs.

---

## 🚀 Objective

Simulate decentralized moderation using goroutines, channels, and context timeouts. Provide basic APIs for:

- Submitting a message (`POST /submit`)
- Retrieving messages by timestamp (`GET /fetch/{timestamp}`)
- Retrieve all stored messages regardless of timestamp (`GET /fetchall`)

---

## 📁 Folder Structure

```
LIBR Prototype/
├── cmd/
│   └── main.go                 # Starts the server
├── controllers/
│   └── handler.go              # HTTP handler logic
├── db/
│   ├── db.go                   # DB connection setup
│   └── message_repo.go         # DB operations
├── moderator/
│   └── moderator.go            # Simulates moderation with goroutines
├── model/
│   └── message.go              # Data models
├── router/
│   └── router.go               # Route registration
├── .env                        # PostgreSQL credentials
├── go.mod                      # Module file
└── README.md
```

---

## 🔧 Setup Instructions

### 1️⃣ Install PostgreSQL

```bash
sudo apt update
sudo apt install postgresql postgresql-contrib
```

### 2️⃣ Create the Database & Table

```sql
CREATE DATABASE libr;

\c libr

CREATE TABLE messages (
  id UUID PRIMARY KEY,
  content TEXT NOT NULL,
  timestamp BIGINT NOT NULL,
  status VARCHAR(10) NOT NULL
);
```

### 3️⃣ Configure Environment Variables

Create a `.env` file in the root directory:

```env
DB_URL=
```

Replace credentials with yours if needed.

---

## 🛠 Installation

```bash
git clone <your-repo-url>
cd "LIBR Prototype"

go mod tidy
go build -o server ./cmd
./server
```

Server will start on: `http://localhost:4000`

---

## 🌐 API Endpoints

### 📨 `POST /submit`

Submit a message for moderation.

**Request Payload:**
```json
{
  "content": "This is a test message."
}
```

**Sample Response (Approved):**
```json
{
  "id": "generated-uuid",
  "timestamp": 1744219507,
  "status": "approved"
}
```

**Sample Response (Rejected):**
```json
{
  "id": "generated-uuid",
  "timestamp": 1744219507,
  "status": "rejected"
}
```

---

### 📥 `GET /fetch/{timestamp}`

Retrieve all messages submitted at a specific timestamp.

**Sample Response:**
```json
[
  {
    "id": "unique-id",
    "content": "This is a test message.",
    "timestamp": 1744219507,
    "status": "approved"
  }
]
```

### 📥 `GET /fetch`

Retrieve all messages regardless of timestamp.

**Sample Response:**
```json
[
  {
    "id": "unique-id",
    "content": "This is a test message.",
    "timestamp": 1744219507,
    "status": "approved"
  }
  {
    "id": "unique-id",
    "content": "This is a test message.",
    "timestamp": 1744219507,
    "status": "approved"
  }
  {
    "id": "unique-id",
    "content": "This is a test message.",
    "timestamp": 1744219507,
    "status": "approved"
  }
]
```

---

## 🧠 Simulated Moderation

Each message is sent to 3 simulated moderators. Moderators:

- Use **goroutines** for concurrency.
- Introduce random delays (`1-3s`) to mimic real processing.
- Randomly **approve/reject** the message.
- Respond via **channels**.
- Have a **3s timeout** using `context.WithTimeout`.

The message is **approved if at least 2 moderators approve**, otherwise it's rejected.

---

## 🛠 Libraries Used

- [`gorilla/mux`](https://github.com/gorilla/mux): Routing
- [`pgx`](https://github.com/jackc/pgx): PostgreSQL driver
- [`google/uuid`](https://pkg.go.dev/github.com/google/uuid): UUID generator
- [`joho/godotenv`](https://github.com/joho/godotenv): Load environment variables

---

## 📸 Screenshots

### ✅ Successful Submit (Postman)

![POST Success](https://github.com/CaptainReck/LIBR-Simplified/blob/main/screenshots/Screenshot%20from%202025-06-13%2018-46-13.png)

### ❌ Rejected Submit

![POST Rejected](https://github.com/CaptainReck/LIBR-Simplified/blob/main/screenshots/Screenshot%20from%202025-06-13%2018-46-52.png)

### 📤 Fetch All Messages

![GET Fetch](https://github.com/CaptainReck/LIBR-Simplified/blob/main/screenshots/image.png)

### 📤 Fetch Messages

![GET Fetch](https://github.com/CaptainReck/LIBR-Simplified/blob/main/screenshots/image%20copy.png)

## 🧪 Testing

Use [Postman](https://www.postman.com/) to test:


## 🔐 Security Note

- DB credentials are stored in `.env`. Do **not** commit `.env` to version control.
- This is a prototype. Avoid using it in production environments.

---
