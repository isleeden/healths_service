package presenters

import (
	"github.com/Shemistan/healths_service/internal/entities"
	"github.com/rodaine/table"
)

const (
	nameHeader   = "NAME"
	statusHeader = "STATUS"
	urlHeader    = "URL"
)

func PrintHealthTable(statuses []entities.ServiceStatus) {
	tbl := table.New(nameHeader, statusHeader)

	for _, status := range statuses {
		tbl.AddRow(status.Name, status.Status, status.Url)
	}

	tbl.Print()
}
