# Registration
The key idea is to create a registration simulation, after which information about registered users will be posted on the site.

# Stucture 
```
.
├── backend
│   ├── cmd
│   │   └── main.go
│   └── internal
│       ├── db
│       │   └── db.go
│       ├── handlers
│       │   ├── api.go
│       │   ├── home.go
│       │   └── register.go
│       └── models
│           └── user.go
├── frontend
│   ├── static
│   │   ├── app.js
│   │   └── style.css
│   └── templates
│       ├── board.html
│       └── index.html
├── migrations
├── docker-compose.yml
└── Dockerfile

```

PostgreSQL query: 
```
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL
);
```
how to start: docker-compose up --build

<img width="1420" alt="Снимок экрана 2025-02-27 в 22 50 41" src="https://github.com/user-attachments/assets/cc0e9a8c-9d76-4fb4-b7b9-3f632910c4c6" />

<img width="1423" alt="Снимок экрана 2025-02-27 в 22 51 23" src="https://github.com/user-attachments/assets/81866e70-1a4a-4cbd-866d-d6d676c82445" />
