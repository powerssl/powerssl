$(document).ready(function () {
    $("form").submit(function (event) {
        event.preventDefault();

        var username = $("#username").val();
        $("form").trigger("reset");

        $.get("/jwt?sub=" + username, function (data) {
            window.parent.postMessage(data, webAppURI);
        });
    });
});
