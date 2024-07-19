package types

type Distance struct {
	OBUID int     `json:"obu_id"`
	Value float64 `json:"value"`
	Unix  int64   `json:"unix"`
}
