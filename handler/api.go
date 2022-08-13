package handler

import (
	"fmt"
	_ "mysql-master"
	"net/http"
	"tugas-go-dtsproa-mangwi/server"

	"github.com/labstack/echo"
)

type task struct {
	Id_task      int
	Deskripsi    string
	Nama_pegawai string
	Deadline     string
}

var data []task

func BacaData(c echo.Context) error {
	task_kerja()

	return c.JSON(http.StatusOK, data)
}

func TambahData(c echo.Context) error {
	db, err := server.Koneksi()

	defer db.Close()

	var deskripsi = c.FormValue("Deskripsi")
	var pegawai = c.FormValue("Nama_pegawai")
	var deadline = c.FormValue("Deadline")

	_, err = db.Exec("insert into tb_task values (?,?,?,?)", nil, deskripsi, pegawai, deadline)

	if err != nil {
		fmt.Println("Task gagal ditambahkan")
		return c.HTML(http.StatusOK, "<strong>Gagal Menambahkan Task<strong>")
	} else {
		fmt.Println("Menu Berhasil Ditambahkan")
		return c.HTML(http.StatusOK, "<script> alert('Berhasil membuat Task'); window.location ='http://localhost:1323';</script>")
	}
	return c.Redirect(http.StatusSeeOther, "/")
}

func UbahData(c echo.Context) error {
	db, err := server.Koneksi()

	defer db.Close()

	var id_task = c.FormValue("Id_task")
	var deskripsi = c.FormValue("Deskripsi")
	var pegawai = c.FormValue("Nama_pegawai")
	var deadline = c.FormValue("Deadline")

	_, err = db.Exec("update tb_task set deskripsi = ? , pegawai = ? , deadline = ? where id_task = ? ", deskripsi, pegawai, deadline, id_task)

	if err != nil {
		fmt.Println("Task gagal diubah")
		return c.HTML(http.StatusOK, "<strong>Gagal diubah Task<strong>")
	} else {
		fmt.Println("Menu Berhasil Diubah")
		return c.HTML(http.StatusOK, "<script> alert('Berhasil ubah Task'); window.location ='http://localhost:1323';</script>")
	}
	return c.Redirect(http.StatusSeeOther, "/")
}

func DeleteData(c echo.Context) error {
	db, err := server.Koneksi()

	defer db.Close()

	var id_task = c.FormValue("Id_task")

	_, err = db.Exec("delete from tb_task where id_task = ?", id_task)

	if err != nil {
		fmt.Println("Task hapus diubah")
		return c.HTML(http.StatusOK, "<strong>Gagal diubah Task<strong>")
	} else {
		fmt.Println("Menu Berhasil dihapus")
		return c.HTML(http.StatusOK, "<script> alert('Berhasil dihapus Task'); window.location ='http://localhost:1323';</script>")
	}
	return c.Redirect(http.StatusSeeOther, "/")
}

func task_kerja() {
	data = nil
	db, err := server.Koneksi()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query("select * from tb_task")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		var each = task{}
		var err = rows.Scan(&each.Id_task, &each.Deskripsi, &each.Nama_pegawai, &each.Deadline)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		data = append(data, each)
		fmt.Println(data)
	}
	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}
}
