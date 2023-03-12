package main // package main adalah package wajib dan akan dieksekusi pertama kali

import ( // kalo cuma 1 import a perlu kurung
	"fmt"
	"io"
	"net/http"
	"strconv"
	"text/template"

	"github.com/labstack/echo/v4"
)

type Template struct { // type struct adalah sebuah kumpulan definisi variabel atau property ataupun fungsi yg nantinya dibungkus sebagai tipe data baru dengan nama tertentu
	templates *template.Template // membuat variabel templates /data akan diambil dari *template/ mengambil method dari Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// pada function ini mendeklarasikan alias (t *Template) dengan memanggil Template/ Template diambil dari struct diatas/ dialiaskan sebagai t / Render(bla bla) adalah funtionnya/ function render memiliki 4 parameter yg akan digunakan yaitu w yg menaliaskan io writer, name dengan tipe data string, data menaliaskan interface / c.echo context wajib ada didalam parameternya/ error akan mengembalikan pesan error ketika funtion ini tidak berjalan
	return t.templates.ExecuteTemplate(w, name, data)
	// ini adalah komen yg akan dijalankan function tsb/ dia mengembalikan Template yng diambil dari struct/ lalu mengambil ExecuteTemplate yaitu methode yg digunakan untuk mengeksekusi Template/
}

func main() {
	e := echo.New() // membuat request ke server. e adalah variabel. New() adalah method untuk menginisiasi agar bisa menerima dari http

	e.Static("/assets", "assets")

	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}

	e.Renderer = t

	// routing
	e.GET("/", home)                            //localhost:5000
	e.GET("/contact", contact)                  //localhost:5000/contact
	e.GET("/project-detail/:id", projectDetail) //localhost:5000/project-detail/0
	e.GET("/form-add", formAdd)
	e.POST("/add-project", addProject) //localhost:5000/add-project
	// e.GET("/", func(c echo.Context) error {// e memanggil variabel e, .GET adalah methode untuk meminta data, / digunakan untuk menambahkan alamat, "/" disebut juga end point, func(c echo.Context) untuk membuat function berisi parameter yg diinisialisasi dalam variabel c yg berisi echo lalu menjalankan context berupa request dan response. kemudian mengembalikan nilai/ pesan error
	// 	return c.String(http.StatusOK, "Hello World")
	// 	// return mengembalikan response // c.String digunakan untuk mengirimkan text atau string berupa response atau status kodenya
	// })

	// e.GET("/about", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "About Risky")
	// })

	fmt.Println("Server berjalan di port 5000") // sebuah pesan saja di terminal
	e.Logger.Fatal(e.Start("localhost:5000"))   // inisialisasi dari echo yg digunakan untuk menjalankan server local. ketika gagal, maka akan menampilkan pesan fatal
}

func home(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}

func contact(c echo.Context) error {
	return c.Render(http.StatusOK, "contact.html", nil)
}

func formAdd(c echo.Context) error {
	return c.Render(http.StatusOK, "add-project.html", nil)
}

func projectDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id")) // url params | dikonversikan dari string menjadi int/integer

	data := map[string]interface{}{ // data yang akan digunakan/dikirimkan ke html menggunakan map interface
		"Id":           id,
		"Title":        "Pasar Coding di Indonesia Dinilai Masih Menjanjikan",
		"StarDate":     "11 Maret 2023",
		"EndDate":      "12 Maret 2023",
		"Technologies": "Node JS, Next JS, React JS, TypeScript",
		"Description":  "Tentunya, dengan kemauan kamu untuk melakukan eksplorasi, bisa menambah wawasan dan tentunya juga pengetahuan kamu. Seperti yang bisa kamu lihat, tidak semua perusahaan menggunakan bahasa/frameworks yang sama, dan tentunya kamu pasti diminta untuk mempelajari apa yang nantinya akan digunakan di perusahaan tersebut.",
	}

	return c.Render(http.StatusOK, "project-detail.html", data)
}

func addProject(c echo.Context) error {
	projectname := c.FormValue("project-name")
	startdate := c.FormValue("start-date")
	enddate := c.FormValue("end-date")
	description := c.FormValue("input-description")
	nodeSelected := c.FormValue("node")
	nextSelected := c.FormValue("next")
	reactSelected := c.FormValue("react")
	typeSelected := c.FormValue("type")

	println("Title: " + projectname)
	println("StarDate: " + startdate)
	println("EndDate: " + enddate)
	println("Technologies:")
	if nodeSelected != "" {
		println("Node JS")
	}
	if nextSelected != "" {
		println("Next JS")
	}
	if reactSelected != "" {
		println("React JS")
	}
	if typeSelected != "" {
		println("TypeScript")
	}
	println("Description: " + description)

	return c.Redirect(http.StatusMovedPermanently, "/") // mendireksi ulang ke rutingan yg dituju yaitu home

}
