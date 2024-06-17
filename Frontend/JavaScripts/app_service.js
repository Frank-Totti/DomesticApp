document.addEventListener('DOMContentLoaded', function() {
  // URL de la API desde donde obtendremos los datos
  const apiUrl = 'http://localhost:3000/professional/request';

  var data1 = {

    id : 11,
    state : true


  };

  var data2 = {

    id : 20,
    state : false


  };

  

  // Realizar la solicitud fetch para obtener los datos
  fetch(apiUrl,{
    method: 'POST',
    headers: {
        'Content-Type': 'application/json',
    },
    body: JSON.stringify(data1)
})
      .then(response => response.json())
      .then(data => {
          // Actualizar la información del perfil en la barra lateral
          const profilePicElement = document.getElementById('profile-pic');
          const userNameElement = document.getElementById('user-name');
          const userLastNameElement = document.getElementById('user-lastname');
          const userEmailElement = document.getElementById('user-email');
          const userAddressElement = document.getElementById('user-address');
          const userPhoneElement = document.getElementById('user-phone');

          // Suponiendo que la API devuelve estos datos directamente
          profilePicElement.src = data.profilePicUrl; // URL de la foto de perfil
          userNameElement.textContent = data.name; // Nombre del cliente
          userLastNameElement.textContent = data.lastname; // Apellido del cliente
          userEmailElement.textContent = data.email; // Correo electrónico
          userAddressElement.textContent = data.address; // Dirección
          userPhoneElement.textContent = data.phone; // Número de teléfono

          // Actualizar las listas de solicitudes recibidas y completadas
          const requestListElement = document.getElementById('request-list');
          const requestCompletedListElement = document.getElementById('request-completed-list');

          // Limpiar las listas antes de agregar nuevos elementos
          requestListElement.innerHTML = '';
          requestCompletedListElement.innerHTML = '';

          // Iterar sobre los datos recibidos y agregar elementos a las listas
          data.history.forEach(item => {
              const listItem = document.createElement('li');
              listItem.textContent = `ID: ${item.ID}, Estado: ${item.Request.State}`;

              if (item.Request.State) {
                  requestCompletedListElement.appendChild(listItem); // Agregar a completados
              } else {
                  requestListElement.appendChild(listItem); // Agregar a recibidos
              }
          });
      })
      .catch(error => console.error('Error al cargar datos desde la API:', error));
});
