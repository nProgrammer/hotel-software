package Utils

import "API/Model"

func LoginU(status string) Model.Login {
	var loginS Model.Login
	loginS.Status = status
	return loginS
}
