document.addEventListener('DOMContentLoaded', () => {
    const form = document.getElementById('registerForm');

    form.addEventListener('submit', function(event) {
        event.preventDefault();

        const email = document.getElementById('email').value;
        const password = document.getElementById('password').value;
        const address = document.getElementById('address').value;
        const name = document.getElementById('firstName').value;
        const lastName = document.getElementById('lastName').value;
        const phone = document.getElementById('phone').value;

        // Obtener la imagen como cadena Base64 y enviar los datos
        getBase64Image('Public_Service', function(publicService) {
            const data = {
                Person: {
                    Address: address,
                    Name: name,
                    LastName: lastName,
                    TNumber: phone,
                    Email: email,
                    Password: password,
                },
                PublicService: publicService
            };

            sendToAPI(data);
        });
    });
});

function getBase64Image(inputId, callback) {
    const input = document.getElementById(inputId);
    if (!input || input.files.length === 0) {
        callback(null);
        return;
    }

    const file = input.files[0];
    const reader = new FileReader();

    reader.onloadend = function() {
        const base64String = reader.result.split(',')[1];
        callback(base64String);
    }

    reader.readAsDataURL(file);
}

function sendToAPI(data) {
    fetch('http://localhost:3000/users/create', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(data)
    })
    .then(response => response.json())
    .then(data => {
        if (data.success) {
            showMessage('Se ha creado con éxito', 'success');
            setTimeout(() => {
                location.reload();
            }, 2000);
        } else {
            showMessage('Error al crear', 'error');
        }
    })
    .catch((error) => {
        console.error('Error:', error);
        showMessage('Error al crear', 'error');
    });
}

function showMessage(message, type) {
    const messageDiv = document.getElementById('message');
    messageDiv.textContent = message;
    messageDiv.className = type;
    console.log(showMessage)
}
