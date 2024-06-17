document.addEventListener('DOMContentLoaded', () => {
    const loginForm = document.getElementById('loginForm');

    if (loginForm) {
        loginForm.addEventListener('submit', function(event) {
            event.preventDefault();

            const email = document.getElementById('email').value;
            const password = document.getElementById('password').value;

            const credentials = {
                email: email,
                password: password
            };

            loginUser(credentials);
        });
    }
});

function loginUser(credentials) {
    fetch('http://localhost:3000/users/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(credentials)
    })
    .then(response => response.json())
    .then(data => {
        if (data.token) {
            document.cookie = `token=${data.token};path=/`;
            showMessage('Login successful', 'success');
            setTimeout(() => {
                window.location.href = '../HTML/log_in_service.html'; 
            }, 2000);
        } else {

            showMessage('Login failed', 'error');
        }
    })
    .catch((error) => {
        console.error('Error:', error);
        showMessage('Login failed', 'error');
    });
}

function showMessage(message, type) {
    const messageDiv = document.getElementById('message');
    messageDiv.textContent = message;
    messageDiv.className = type;
}
