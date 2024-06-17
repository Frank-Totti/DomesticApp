
function isAuthenticated() {
    const cookies = document.cookie.split(';');
    for (let i = 0; i < cookies.length; i++) {
        const cookie = cookies[i].trim();
        if (cookie.startsWith('token=')) {
            return true; 
        }
    }
    return false; 
}


function checkAuthentication() {
    if (!isAuthenticated()) {
        console.log("manito...")
        window.location.href = '../HTML/log_in_service.html'; 
    }
}


window.onload = function() {

    checkAuthentication();
}


function logout() {
    fetch('http://localhost:3000/users/logout', {
        method: 'GET',
        credentials: 'same-origin'
    })
    .then(response => response.json())
    .then(data => {
        if (data.message === "Logout successful") {
            window.location.href = '../HTML/log_in_service.html'; 
        }
    })
    .catch(error => {
        console.error('Error:', error);
    });
}


document.addEventListener('DOMContentLoaded', function() {
    const logoutButton = document.getElementById('logoutButton');
    if (logoutButton) {
        logoutButton.addEventListener('click', logout);
    }
});
