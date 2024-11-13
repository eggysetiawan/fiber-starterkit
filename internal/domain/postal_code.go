package domain

import (
	"strconv"

	"github.com/eggysetiawan/fiber-starterkit/internal/dto"
)

type PostalCode struct {
	PostalCode  string `json:"kodePos"`
	AreaCode    string `json:"kodeWilayah,omitempty"`
	Province    string `json:"namaProvinsi,omitempty"`
	City        string `json:"namaKabupaten,omitempty"`
	SubDistrict string `json:"namaKecamatan,omitempty"`
	Village     string `json:"namaKelurahan,omitempty"`
}

func (pc *PostalCode) ToDto() dto.PostalCodeResponse {
	pcInt, err := strconv.Atoi(pc.PostalCode)
	if err != nil {
		panic("failed to convert postal code str to int ToDto: " + err.Error())
	}

	resp := dto.PostalCodeResponse{
		PostalCode:  pcInt,
		AreaCode:    pc.AreaCode,
		Province:    pc.Province,
		City:        pc.City,
		SubDistrict: pc.SubDistrict,
		Village:     pc.Village,
	}

	return resp

}
