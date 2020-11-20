$(document).ready(function() {
    $.ajax({
        url: "/api/test"
    }).then(function(data) {
       $('.randNum').append(data.Randnum);
    });
});
