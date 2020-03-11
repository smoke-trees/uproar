function getFormData($form) {
    var unindexed_array = $form.serializeArray();
    var indexed_array = {};

    $.map(unindexed_array, function (n, i) {
        indexed_array[n['name']] = n['value'];
    });

    return indexed_array;
}

function onRegisterFormSubmit(event) {
    event.preventDefault();
    data = getFormData($("#register-form"))
    let ajax = $.ajax({
        url: "http://localhost:3000/forum/register",
        method: "POST",
        data: data,
    })

    let ajax2 = $.ajax({
        url: "http://localhost:3001/register",
        method: "POST",
        data: data,
    })

    ajax2.done(msg => {
        console.log(msg)
    })
    ajax2.fail(() => {
        console.log("fail auth")
    })
    ajax.done(msg => {
            window.location = "http://localhost:5000/login"
        }
    )
}