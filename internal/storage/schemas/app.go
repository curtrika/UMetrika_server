package schemas

type AppSchema struct {
	ID     int32  `json:"id"`
	Name   string `json:"name"`
	Secret string `json:"secret"`
}
