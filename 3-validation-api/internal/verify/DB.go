package verify

type RequestVerify struct {
	Email string `json:"email"`
	Hash  string `json:"hash"`
}

type RequestSend struct {
	Email string `json:"email"`
}

var DB = []RequestVerify{}
