//Peticiones AJAX->XML=>JSON
//Activ object la creo Microsoft 
//Con AJAX no necesitamos recargar toda la pagina solo elementos que requerimos
// XMLHttpRequest se encarga de realizar las peticiones y recibir la respuesta del servidor
// let xhr = new XMLHttpRequest();
//Abrimos la conexión y que método vamos a usar para abrir la conexión
// xhr.open('GET', '/api/login', true);
// xhr.setRequestHeader('Content-Type', 'application/json');
//Agregar el listener(escuchador) para cuando el servidor haya enviado la respuesta
// xhr.addEventListener('load', e=>{// 
    // let self = e.target;
    // let respuesta = {
        // status: self.status,
        // response: self.response
    // }
// });
// Escuchador para los errores de la comunicación
// xhr.addEventListener('error', e=>{
  //  let self = e.target;
    // console.log(self);
// });
// Decirle que contenido vamos a enviarle nosotros
// xhr.send(obj);


// Promesas - Promise => Enviamos metodo el url y el contenido
function peticionAjax(metodo, url, obj) {
    
    return new Promise(function(resolver, rechazar){
        let xhr = new XMLHttpRequest();
        //Abrimos la conexión y que método vamos a usar para abrir la conexión
        xhr.open(metodo, url, true);
        xhr.setRequestHeader('Content-Type', 'application/json');

        //Traer el token del session y solo se va a enviar si tenemos el token
        if(sessionStorage.getItem('tokened')){
            xhr.setRequestHeader('Authorization', sessionStorage.getItem('tokened'));
        }

        //Agregar el listener(escuchador) para cuando el servidor haya enviado la respuesta
        xhr.addEventListener('load', e=>{
            let self = e.target;
            let respuesta = {
                status: self.status,
                response: JSON.parse(self.response)
            };
            resolver(respuesta);
        });
        //Escuchador para los errores de la comunicación
        xhr.addEventListener('error', e=>{
            let self = e.target;
            rechazar(self);
        });
        // Decirle que contenido vamos a enviarle nosotros
        xhr.send(obj);
    });

}

function $(elemento) {
    return document.getElementById(elemento);
}

