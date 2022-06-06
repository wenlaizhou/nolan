package nolan

type Stage struct {
	Name string
}

type Pipeline struct {
	Stages []Stage `json:"stages"`
}
