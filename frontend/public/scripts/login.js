function getFormData($form) {
    var unindexed_array = $form.serializeArray();
    var indexed_array = {};

    $.map(unindexed_array, function (n, i) {
        indexed_array[n['name']] = n['value'];
    });

    return indexed_array;
}

function onLoginFormSubmit(event) {
    event.preventDefault();
    data = getFormData($("#login-form"))
    let ajax = $.ajax({
        url: "http://localhost:3001/login",
        method: "POST",
        data: data,
    })

    ajax.done(msg => {
            console.log(msg)
            let data = JSON.parse(msg);
            let jwt = data.jwt;
            document.cookie = "jwt=" + jwt
            window.location = "http://localhost:5000/feed"
        }
    )
}
