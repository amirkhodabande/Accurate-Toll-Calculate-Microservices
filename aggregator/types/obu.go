package types

type OBU struct {
	OBUID           int     `json:"obuID"`
	Lat             float64 `json:"lat"`
	Long            float64 `json:"long"`
	DestinationLat  float64 `json:"destination_lat"`
	DestinationLong float64 `json:"destination_long"`
}
