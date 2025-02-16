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
