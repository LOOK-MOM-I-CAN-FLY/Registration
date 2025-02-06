# Registration
The key idea is to create a registration simulation, after which information about registered users will be posted on the site.

# Stucture 
```
.
├── backend
│   ├── cmd          # main.go
│   ├── internal
│   │   ├── db       # PostgreSQL взаимодействие
│   │   ├── handlers # HTTP handlers
│   │   └── models   # Структуры данных
├── frontend
│   ├── static       # CSS/JS
│   └── templates    # HTML
├── migrations       # SQL скрипты
├── docker-compose.yml
└── Dockerfile
```
