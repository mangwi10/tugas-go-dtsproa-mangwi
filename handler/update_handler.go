package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func UpdateHandler(c echo.Context) error {
	// Please note the the second parameter "about.html" is the template name and should
	// be equal to one of the keys in the TemplateRegistry array defined in main.go
	r := c.Request()
	return c.Render(http.StatusOK, "update.html", map[string]interface{}{
		"name":         "Update",
		"Id_task":      r.URL.Query()["Id_task"][0],
		"Deskripsi":    r.URL.Query()["Deskripsi"][0],
		"Nama_pegawai": r.URL.Query()["Nama_pegawai"][0],
		"Deadline":     r.URL.Query()["Deadline"][0],
	})
}
