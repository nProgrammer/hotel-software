package View

func ShowClients(name string, surname string, sn string, is string, aUrl string, token string) string {
	return " " + "<a style='text-decoration: none; color: black;' href='" + aUrl + "show-client?sn=" + sn + "&token=" + token + "'><b>" + is + ") <br> </b>" + `<b style="margin-left: 30px;"> Client name: </b>` + name + `<br> <b style="margin-left: 30px;"> Client surname: </b>` + surname + `<br> <b style="margin-left: 30px;"> Client S/N: </b>` + sn + "<br> </a>"
}

func ShowClient(name string, surname string, sn string, number string, id string) string {
	return "<b> Name: </b>" + name + "<br><b> Surname: </b>" + surname + "<br><b> Serial Number: </b>" + sn + "<br><b> Room number: </b>" + number + "<br><b> Car ID: </b>" + id
}
