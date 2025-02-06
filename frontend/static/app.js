async function register() {
    const user = {
        username: document.getElementById('username').value,
        email: document.getElementById('email').value,
        password: document.getElementById('password').value
    };

    const response = await fetch('http://localhost:8080/api/register', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(user)
    });
    
    if (response.ok) {
        loadUsers();
    }
}

async function loadUsers() {
    const response = await fetch('http://localhost:8080/api/users');
    const users = await response.json();
    
    const list = document.getElementById('users-list');
    list.innerHTML = users.map(u => 
        `<div class="user">
            <h3>${u.username}</h3>
            <p>Registered: ${new Date(u.created_at).toLocaleDateString()}</p>
        </div>`
    ).join('');
}

// Загружаем пользователей при старте
loadUsers();
