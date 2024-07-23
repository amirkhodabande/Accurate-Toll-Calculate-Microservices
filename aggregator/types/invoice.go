package types

type Invoice struct {
	OBUID    int     `json:"obu_id"`
	Distance float64 `json:"dist"`
	Amount   float64 `json:"amount"`
}
