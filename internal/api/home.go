package api

type HomeResponse struct {
	BuildTime string `json:"build_time"`
	Revision  string `json:"revision"`
	Release   string `json:"release"`
}
