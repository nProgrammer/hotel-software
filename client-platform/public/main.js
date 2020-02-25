function postClientRegName() {
    const nameR = document.getElementById("nameR").value;
    return nameR
}
function postClientRegSName() {
    const snameR = document.getElementById("snameR").value;
    return snameR
}
function postClientRegSN() {
    const snR =  document.getElementById("snR").value;
    return snR
}
function postClientRegRoomNumber() {
    const room =  document.getElementById("roomNumberR").value;
    return room
}
function postClientRegCarID() {
    const carID =  document.getElementById("carIDR").value;
    return carID
}
function postClientRegPassM() {
    const passMR =  document.getElementById("passMR").value;
    return passMR
}
function postClientDelPassM() {
    const passMD =  document.getElementById("passMD").value;
    return passMD
}
function postClientDelSN() {
    const snD =  document.getElementById("snD").value;
    return snD
}
function CreateQueryRegisterClient() {
    const nameR = postClientRegName();
    const snameR = postClientRegSName();
    const snR = postClientRegSN();
    const room = postClientRegRoomNumber();
    const car = postClientRegCarID();
    const passM = postClientRegPassM();
    if (snR.length === 9) {
        window.location.href = ApiURL + 'register-client?name=' + nameR + "&surname=" + snameR + "&sn=" + snR + "&roomN=" + room + "&carId=" + car + "&passM=" + passM + "&token=" + TOKEN;
    } else if (snR < 9) {
        document.getElementById("errorR").innerHTML = "TO SHORT SERIAL NUMBER";
    } else {
        document.getElementById("errorR").innerHTML = "TO LONG SERIAL NUMBER";
    }
}
function CreateQueryDelClient() {
    const snD = postClientDelSN();
    const passD = postClientDelPassM();
    window.location.href = ApiURL + 'delete-client?sn=' + snD + "&passM=" + passD + "&token=" + TOKEN;
}
function iframeConfig() {
    document.getElementById("clientScreen").innerHTML ='<iframe style="height: 90vh" src="' + ApiURL + 'show-clients' + "?token=" + TOKEN + '"></iframe>';
}
function doE() {
    iframeConfig();
}

window.onload = doE;
