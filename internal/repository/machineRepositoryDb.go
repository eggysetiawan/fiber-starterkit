package repository

import (
	"database/sql"
	"fmt"

	"github.com/eggysetiawan/fiber-starterkit/errs"
	"github.com/eggysetiawan/fiber-starterkit/internal/domain"
	"github.com/gofiber/fiber/v2"
)

type MachineRepositoryDb struct {
	client *sql.DB
}

func NewMachineRepositoryDb(db *sql.DB) *MachineRepositoryDb {
	return &MachineRepositoryDb{db}
}

func (d *MachineRepositoryDb) FindBy(identifier string, value string) (*domain.Machine, *errs.AppError) {
	query := fmt.Sprintf("select tid, ip_address, type, brand, pengelola, name, kanwil, branch, kanwil2, branch2  from atm_mappings WHERE %s = ?", identifier)

	row := d.client.QueryRow(query, value)

	var m domain.Machine

	err := row.Scan(&m.Tid, &m.Ip, &m.Type, &m.Brand, &m.Management, &m.Location, &m.RegionSpv, &m.BranchSpv, &m.RegionReplenish, &m.BranchReplenish)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("TID/IP Address tidak ditemukan")
		} else {
			fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf("error when scanning machine %s", err.Error()))
			return nil, errs.NewUnexpectedError(fmt.Sprintf("Unexpected database error: %s", err.Error()))
		}
	}

	return &m, nil
}
