# 🎵 Groupie Tracker

## 📌 Description

**Groupie Tracker** is a web project written in Go that consumes an external API containing information about music bands and artists (names, members, concert dates and locations, etc.) and displays this data through a dynamic and user-friendly website.

The goal is to retrieve, manipulate, and organize the received data, then present it using various visualizations (maps, tables, lists…) within a clean and responsive user interface.

---

## 🧠 Objectives

- Consume and manipulate a RESTful API providing:
  - **Artists**: name(s), image, creation date, first album, members.
  - **Locations**: cities where concerts have taken or will take place.
  - **Dates**: dates of past or upcoming concerts.
  - **Relation**: links artists to their concerts (locations + dates).

- Build a **Go-based website** (backend) that:
  - Displays artist information.
  - Supports interaction with the server through client-side events.
  - Handles errors gracefully and never crashes.

---

## 🛠️ Features

- 🔍 Search by artist name or member.
- 📅 View concert dates.
- 🌍 Map or list of concert locations.
- 👥 Detailed artist view (biography, albums, members…).
- 🔁 Real-time interactions with the server via HTTP requests.

---

## 🧪 Testing

- Unit test files are recommended.
- Use standard Go tests (`go test`).

---

## ✅ Requirements

- The **backend must be written in Go** only.
- The site must be **robust** (no crashes).
- All errors must be **handled properly**.
- Follow **best development practices**.
- **No external libraries allowed** — only Go standard packages.

---

## 🚀 Running the Project

### Prerequisites

- [Go](https://golang.org/dl/)
- A web browser

### Commands

```bash
# Clone the project
git clone https://your-repo-url/groupie-tracker.git
cd groupie-tracker

# Run the server
go run main.go
```
