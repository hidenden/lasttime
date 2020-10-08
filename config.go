package lasttime

type target struct {
	Name      string `json:"name"`
	Threshold int    `json:"threshold"`
}

type Config struct {
	Mode   string   `json:"mode"`
	Target []target `json:"target"`
}

func (c Config) AddTarget(name string, threshold int) {
	t := target{
		Name:      name,
		Threshold: threshold,
	}
	c.Target = append(c.Target, t)
}
