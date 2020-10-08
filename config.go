package lasttime

type target struct {
	Name      string `json:"name"`
	Threshold int    `json:"threshold"`
}

type Config struct {
	Mode   string   `json:"mode"`
	Target []target `json:"target"`
}
