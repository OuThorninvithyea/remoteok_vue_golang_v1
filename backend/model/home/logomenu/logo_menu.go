package logomenu

import (
	"backend/model/home/logo"
	"backend/model/home/menu"
	"backend/model/home/salary"
)

type LogoMenu struct {
	LogoMenu     []logo.LogoEntry   `json:"logoMenu"`
	SearchMenu   []menu.MenuItem    `json:"searchMenu"`
	LocationMenu []logo.LogoEntry   `json:"locationMenu"`
	SalaryMenu   salary.SalaryConfig `json:"salaryMenu"`
	BenefitsMenu []menu.MenuItem    `json:"benefitsMenu"`
	SortMenu     []menu.MenuItem    `json:"sortMenu"`
	Categories   []menu.MenuItem    `json:"categories"`
}
