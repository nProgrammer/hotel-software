package Utils

import "API/Model"

func LoginU(status string) interface{} {
	var loginS Model.Login
	loginS.Status = status
	return nil
}
