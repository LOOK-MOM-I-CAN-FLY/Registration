# Registration
The key idea is to create a registration simulation, after which information about registered users will be posted on the site.

# Stucture 
```
.
├── backend
│   ├── cmd          # main.go
│   ├── internal
│   │   ├── db       # PostgreSQL interaction
│   │   ├── handlers # HTTP handler
│   │   └── models   # Data structures
├── frontend
│   ├── static       # CSS/JS
│   └── templates    # HTML
├── migrations       # SQL scripts
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
