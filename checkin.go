package gomgtclient

type CheckIn struct {
	CheckedIn bool   `json:"CheckedIn"`
	Message   string `json:"Message"`
}
