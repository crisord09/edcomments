let formCommentAdd = $('form-comment-add'),
    commentContent = $('comment-content'),
    mensajeComentario = $('mensaje-comentario');

formCommentAdd.addEventListener('submit', e=>{
    e.preventDefault();
    let obj = {
        content: commentContent.value
    };

    peticionAjax(formCommentAdd.method, formCommentAdd.action, JSON.stringify(obj))
        .then(respuesta => {
            if (respuesta.status === 201){
                mensajeComentario.textContent = respuesta.response.message;

            }else{
                mensajeComentario.textContent = respuesta.response.message;
            }
        })
        .catch(error =>{
            console.log(error);
        });

    });