package domain

import (
	"github.com/eggysetiawan/fiber-starterkit/errs"
	"github.com/eggysetiawan/fiber-starterkit/internal/dto"
)

type Machine struct {
	Tid             string `db:"tid"`
	Ip              string `db:"ip_address"`
	Type            string `db:"type"`
	Brand           string `db:"brand"`
	Management      string `db:"pengelola"`
	Location        string `db:"name"`
	RegionSpv       string `db:"kanwil"`
	BranchSpv       string `db:"branch"`
	RegionReplenish string `db:"kanwil2"`
	BranchReplenish string `db:"branch2"`
}

type MachineRepository interface {
	FindBy(identifier string, value string) (*Machine, *errs.AppError)
}

func (m *Machine) ToDto() *dto.MachineResponse {
	return &dto.MachineResponse{
		Machines: []dto.MachineDetails{
			{
				Label: "TID",
				Value: m.Tid,
			},
			{
				Label: "IP Address",
				Value: m.Ip,
			},
			{
				Label: "Tipe Mesin",
				Value: m.Type,
			},
			{
				Label: "Merk Mesin",
				Value: m.Brand,
			},
			{
				Label: "Pengelola",
				Value: m.Management,
			},
			{
				Label: "Lokasi",
				Value: m.Location,
			},
			{
				Label: "Kantor Wilayah Supervisi",
				Value: m.RegionSpv,
			},
			{
				Label: "Kantor Cabang Supervisi",
				Value: m.BranchSpv,
			},
			{
				Label: "Kantor Wilayah Replenish",
				Value: m.RegionReplenish,
			},
			{
				Label: "Kantor Cabang Replenish",
				Value: m.BranchReplenish,
			},
		},
	}

}
