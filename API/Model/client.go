package Model

type Client struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Sn         string `json:"sn"`
	RoomNumber string `json:"room"`
	CarID      string `json:"carID"`
}

type ClientDel struct {
	Sn    string `json:"sn"`
	Passm string `json:"passm"`
}

type ClientLogin struct {
	Login string `json:"login"`
	Pass  string `json:"pass"`
}
