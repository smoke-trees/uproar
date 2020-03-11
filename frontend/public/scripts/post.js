function onNewPostSubmit(event) {
    event.preventDefault()
    formData = getFormData($("#post-form"))

    postContent = formData.postContent
    token = getCookie("jwt")
    console.log(token)
    data = {
        postContent: postContent, jwt: token
    }
    let ajax = $.ajax({
        url: "http://localhost:3000/forum/post/new",
        method: "POST",
        data: data,
    })
    ajax.done((msg) => {
        console.log(msg)
        window.location = "http://localhost:5000/feed"
    })
}

function getFormData($form) {
    var unindexed_array = $form.serializeArray();
    var indexed_array = {};

    $.map(unindexed_array, function (n, i) {
        indexed_array[n['name']] = n['value'];
    });

    return indexed_array;
}

function getCookie(cname) {
    var name = cname + "=";
    var decodedCookie = decodeURIComponent(document.cookie);
    var ca = decodedCookie.split(';');
    for (var i = 0; i < ca.length; i++) {
        var c = ca[i];
        while (c.charAt(0) == ' ') {
            c = c.substring(1);
        }
        if (c.indexOf(name) == 0) {
            return c.substring(name.length, c.length);
        }
    }
    return "";
}