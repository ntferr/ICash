package entities

type Card struct {
	ID       int    `json:"id"`
	Type     Debt   `json:"debt"`
	Bank     Bank   `json:"bank"`
	Number   string `json:"number"`
	ExpireAt string `json:"expire_at"`
}
