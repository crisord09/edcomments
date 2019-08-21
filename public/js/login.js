//Obtendremos lo que tenemos dentro del formulario HTML y hacer la peticion AJAX
let formLogin = document.getElementById('form-login'),
    email = $('email'),
    password = $('password'),
    mensajeLogin = $('mensaje-login'),
    btnLogin = $('btnLogin');

formLogin.addEventListener('submit', e => {
    e.preventDefault();
    let obj = {
        email: email.value,
        password: password.value
    };

    peticionAjax(formLogin.method, formLogin.action, JSON.stringify(obj))
        .then(respuesta => {
            if (respuesta.status === 200) {
                mensajeLogin.textContent = 'Ingresaste';
                sessionStorage.setItem('tokened', respuesta.response.token);
                console.log(respuesta.response);
            }else{
                mensajeLogin.textContent = respuesta.response.message;
                console.log(respuesta.response);
            }
        })
        .catch(error => {
            console.log(error);
        });

});