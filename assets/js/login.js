$(document).ready(function() {
    $.ajax({
        url: "/api/test"
    }).then(function(data) {
       $('.Country').append(data.Content);
       $('.Province').append(data.Province);
       $('.City').append(data.City);
    });
});