// Cargar servicios al cargar la página
document.addEventListener('DOMContentLoaded', function() {
    loadServices();
});

function loadServices() {
    fetch('http://localhost:3000/services/search/true')
      .then(response => response.json())
      .then(data => {  
        const serviceList = document.querySelector('.service-list');
  
        // Limpiar lista de servicios antes de añadir nuevos
        serviceList.innerHTML = '';
  
        // Iterar sobre los servicios recibidos del JSON
        data.forEach(service => {
          const serviceElement = document.createElement('div');
          serviceElement.classList.add('service-item');
  
          const typeElement = document.createElement('h3');
          typeElement.textContent = service.Type;
          serviceElement.appendChild(typeElement);
  
          const nameElement = document.createElement('p');
          nameElement.textContent = service.Description;
          serviceElement.appendChild(nameElement);
  
          // Agregar evento de clic para cargar proveedores al hacer clic en el servicio
          serviceElement.onclick = function() {
            loadServiceProviders(service.Type); // Envía el SID del servicio al cargar proveedores SERVICE.SID
            highlightService(serviceElement); // Resalta el servicio seleccionado visualmente
          };
  
          serviceList.appendChild(serviceElement);
        });
      })
      .catch(error => console.error('Error al cargar servicios:', error));
  }
  
  // Función para crear un elemento HTML con sus propiedades
  function crearElemento(tipo, propiedades) {
    const elemento = document.createElement(tipo);
    for (const propiedad in propiedades) {
      elemento[propiedad] = propiedades[propiedad];
    }
    return elemento;
  }


function loadServiceProviders(serviceType) {
    const apiUrl = 'http://localhost:3000/professional_offers/search/service/type';
    const body = JSON.stringify({ Type: serviceType });
  
    fetch(apiUrl, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: body
    })
      .then(response => response.json())
      .then(data => {
        const providersList = document.getElementById('service-providers-list');
        providersList.innerHTML = ''; // Limpiar la lista antes de agregar nuevos proveedores
  
        data.forEach(offer => {
          const providerItem = document.createElement('li');
          const provider = offer.Professional; // Acceder a los datos del proveedor
          providerItem.textContent = `${provider.Person.Name} ${provider.Person.LastName}`;
          providerItem.onclick = () => showProviderDetails(offer);
          providersList.appendChild(providerItem);
        });
      })
      .catch(error => console.error('Error al cargar proveedores:', error));
  }
  
  function showProviderDetails(offer) {
    const providerDetails = document.getElementById('provider-details');
    const provider = offer.Professional; // Acceder a los datos del proveedor
    const service = offer.Service; // Acceder a los datos del servicio
  
    providerDetails.innerHTML = `
      <p><strong>Especialidad:</strong> ${offer.Major}</p>
      <p><strong>Precio por Unidad:</strong> ${offer.UnitPrice}</p>
      <p><strong>Precio por Hora:</strong> ${offer.PricePerHour}</p>
      <p><strong>Nombre:</strong> ${provider.Person.Name} ${provider.Person.LastName}</p>
      <p><strong>Tipo de Servicio:</strong> ${service.Type}</p>
    `;
  }

// Función para resaltar visualmente el servicio seleccionado
function highlightService(serviceElement) {
    const serviceItems = document.querySelectorAll('.service-item');
    serviceItems.forEach(item => item.classList.remove('selected'));
    serviceElement.classList.add('selected');
}

// Función para resaltar visualmente el proveedor seleccionado
function highlightProvider(providerElement) {
    const providerItems = document.querySelectorAll('.service-provider-item');
    providerItems.forEach(item => item.classList.remove('selected-provider'));
    providerElement.classList.add('selected-provider');
}
