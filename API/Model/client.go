package Model

type Client struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Sn         string `json:"sn"`
	RoomNumber string `json:"room"`
	CarID      string `json:"carID"`
}
