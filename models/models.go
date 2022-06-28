package models

type Token struct {
    Token string `json:"token"`
}

type Blocks struct {
    Data []string `json:"data"`
}

type CheckRequest struct {
    Blocks []string `json:"blocks"`
	Encoded string `json:"encoded"`
}

type CheckResponse struct {
    Message bool `json:"message"`
}