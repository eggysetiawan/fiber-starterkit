package dto

type MachineRequest struct {
	Identifier string `json:"identifier"`
}

type MachineDetails struct {
	Label string `json:"label"`
	Value string `json:"value"`
}
type MachineResponse struct {
	Machines []MachineDetails `json:"machines"`
}
