var isLogged;
module.exports = {
    IsLoggedR: function () {
        document.getElementById("proceedButton").style = "visibility: visible;";
        return Login()
    }
};

function Login () {
    let xhr = new XMLHttpRequest();
    let url = "http://localhost:8000/login?token=4908239482njkb2u3b234b2343094324n23io4bn2i3b452io3b42io3b4i2o3b52uj3i5b23b523ij5b3253ub53523ui523b523b532ub52b5323jb52b523b523j5b235b25jbj33";
    xhr.open("POST", url, true);
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4 && xhr.status === 200) {
            let jsonL = JSON.parse(xhr.responseText);
            if (jsonL.status === "SUCCESS") {
                var IsLogged = jsonL;
            }
            console.log(jsonL.status);
        }
        isLogged = JSON.stringify(IsLogged);
    };
    let data = JSON.stringify({
        "login": document.getElementById("login").value,
        "pass": document.getElementById("pass").value
    });
    xhr.send(data);
    return isLogged
}