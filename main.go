package main // package main adalah package wajib dan akan dieksekusi pertama kali

import ( // kalo cuma 1 import a perlu kurung
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"personal-web/connection"
	"strconv"
	"text/template"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type Template struct { // type struct adalah sebuah kumpulan definisi variabel atau property ataupun fungsi yg nantinya dibungkus sebagai tipe data baru dengan nama tertentu / struck digunakan untuk mendefinisakan sebuah field
	templates *template.Template // membuat variabel templates /data akan diambil dari *template/ mengambil method dari Template
}

type Project struct {
	// buat menampung
	ID          int
	Title       string
	StartDate   time.Time
	EndDate     time.Time
	Node        bool
	Next        bool
	React       bool
	Type        bool
	Description string
	Duration    string
	image       string
}

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

type SessionData struct {
	IsLogin bool
	Name    string
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// pada function ini mendeklarasikan alias (t *Template) dengan memanggil Template/ Template diambil dari struct diatas/ dialiaskan sebagai t / Render(bla bla) adalah funtionnya/ function render memiliki 4 parameter yg akan digunakan yaitu w yg menaliaskan io writer, name dengan tipe data string, data menaliaskan interface / c.echo context wajib ada didalam parameternya/ error akan mengembalikan pesan error ketika funtion ini tidak berjalan
	return t.templates.ExecuteTemplate(w, name, data)
	// ini adalah komen yg akan dijalankan function tsb/ dia mengembalikan Template yng diambil dari struct/ lalu mengambil ExecuteTemplate yaitu methode yg digunakan untuk mengeksekusi Template/
}

func main() {
	connection.DatabaseConnect()
	e := echo.New() // membuat request ke server. e adalah variabel. New() adalah method untuk menginisiasi agar bisa menerima dari http

	e.Static("/assets", "assets")

	e.Use(session.Middleware(sessions.NewCookieStore([]byte("session"))))

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
	e.GET("/delete-project/:id", deleteProject)
	e.GET("/edit-project/:id", editProject)
	e.POST("/update-project/:id", updateProject)
	e.GET("/form-login", formLogin)
	e.POST("/login", login)
	e.GET("/form-register", formRegister)
	e.POST("/register", addRegister)
	e.GET("/logout", logout)

	fmt.Println("Server berjalan di port 5000") // sebuah pesan saja di terminal
	e.Logger.Fatal(e.Start("localhost:5000"))   // inisialisasi dari echo yg digunakan untuk menjalankan server local. ketika gagal, maka akan menampilkan pesan fatal
}

func home(c echo.Context) error {
	sess, _ := session.Get("session", c)

	flash := map[string]interface{}{
		"FlashStatus":  sess.Values["isLogin"],
		"FlashMessage": sess.Values["message"],
		"FlashName":    sess.Values["name"],
	}
	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())

	data, _ := connection.Conn.Query(context.Background(), "SELECT * FROM  tb_project;")
	var result []Project
	for data.Next() {
		var each = Project{}

		err := data.Scan(&each.ID, &each.Title, &each.StartDate, &each.EndDate, &each.Node, &each.Next, &each.React, &each.Type, &each.Description, &each.image)
		if err != nil {
			fmt.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"Message ": err.Error()})
		}

		diff := each.EndDate.Sub(each.StartDate)
		var Durations string

		if diff.Hours()/24 < 7 {
			Durations = strconv.FormatFloat(diff.Hours()/24, 'f', 0, 64) + " Days"
		} else if diff.Hours()/24/7 < 4 {
			Durations = strconv.FormatFloat(diff.Hours()/24/7, 'f', 0, 64) + " Weeks"
		} else if diff.Hours()/24/30 < 12 {
			Durations = strconv.FormatFloat(diff.Hours()/24/30, 'f', 0, 64) + " Months"
		} else {
			Durations = strconv.FormatFloat(diff.Hours()/24/30/12, 'f', 0, 64) + " Years"
		}

		each.Duration = Durations

		result = append(result, each)
	}

	projects := map[string]interface{}{
		"Projects": result,
		"Flash":    flash,
	}

	return c.Render(http.StatusOK, "index.html", projects)
}

func contact(c echo.Context) error {
	return c.Render(http.StatusOK, "contact.html", nil)
}

func formAdd(c echo.Context) error {
	return c.Render(http.StatusOK, "add-project.html", nil)
}

func projectDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id")) // url params | dikonversikan dari string menjadi int/integer

	var ProjectDetail = Project{}

	err := connection.Conn.QueryRow(context.Background(), "SELECT id, nama_project, stardate, enddate, node, next, react, type, description from tb_project WHERE id = $1", id).Scan(&ProjectDetail.ID, &ProjectDetail.Title, &ProjectDetail.StartDate, &ProjectDetail.EndDate, &ProjectDetail.Node, &ProjectDetail.Next, &ProjectDetail.React, &ProjectDetail.Type, &ProjectDetail.Description)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"Message ": err.Error()})
	}

	startFormat := ProjectDetail.StartDate.Format("02 February 2006")
	endFormat := ProjectDetail.EndDate.Format("02 February 2006")

	diff := ProjectDetail.EndDate.Sub(ProjectDetail.StartDate)
	var Durations string

	if diff.Hours()/24 < 7 {
		Durations = strconv.FormatFloat(diff.Hours()/24, 'f', 0, 64) + " Days"
	} else if diff.Hours()/24/7 < 4 {
		Durations = strconv.FormatFloat(diff.Hours()/24/7, 'f', 0, 64) + " Weeks"
	} else if diff.Hours()/24/30 < 12 {
		Durations = strconv.FormatFloat(diff.Hours()/24/30, 'f', 0, 64) + " Months"
	} else {
		Durations = strconv.FormatFloat(diff.Hours()/24/30/12, 'f', 0, 64) + " Years"
	}

	detailProject := map[string]interface{}{ // data yang akan digunakan/dikirimkan ke html menggunakan map interface
		"Project":   ProjectDetail,
		"Duration":  Durations,
		"startDate": startFormat,
		"endDate":   endFormat,
	}

	return c.Render(http.StatusOK, "project-detail.html", detailProject)
}

func addProject(c echo.Context) error {
	sess, _ := session.Get("session", c)

	delete(sess.Values, "message")
	sess.Save(c.Request(), c.Response())

	projectname := c.FormValue("project-name")
	startdate := c.FormValue("start-date")
	enddate := c.FormValue("end-date")
	nodeSelected := false
	nextSelected := false
	reactSelected := false
	typeSelected := false
	if c.FormValue("node") != "" {
		nodeSelected = true
	}
	if c.FormValue("next") != "" {
		nextSelected = true
	}
	if c.FormValue("react") != "" {
		reactSelected = true
	}
	if c.FormValue("type") != "" {
		typeSelected = true
	}

	description := c.FormValue("input-description")
	image := "image.png"

	_, err := connection.Conn.Exec(context.Background(), "INSERT INTO tb_project (nama_project, stardate, enddate, node, next, react, type, description, image) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)", projectname, startdate, enddate, nodeSelected, nextSelected, reactSelected, typeSelected, description, image)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"Message ": err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, "/") // mendireksi ulang ke rutingan yg dituju yaitu home

}

func deleteProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id")) //id = 2

	_, err := connection.Conn.Exec(context.Background(), "DELETE FROM tb_project WHERE id=$1", id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"Message ": err.Error()})
	}
	// dataProject = append(dataProject[:id], dataProject[id+1:]...)
	// //

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func editProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var ProjectDetail = Project{}
	err := connection.Conn.QueryRow(context.Background(), "SELECT id, nama_project, stardate, enddate, node, next, react, type, description from tb_project WHERE id = $1", id).Scan(&ProjectDetail.ID, &ProjectDetail.Title, &ProjectDetail.StartDate, &ProjectDetail.EndDate, &ProjectDetail.Node, &ProjectDetail.Next, &ProjectDetail.React, &ProjectDetail.Type, &ProjectDetail.Description)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"Message ": err.Error()})
	}

	detailProject := map[string]interface{}{
		"Project": ProjectDetail,
	}
	return c.Render(http.StatusOK, "edit-project.html", detailProject)
}

func updateProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	projectname := c.FormValue("project-name")
	startdate := c.FormValue("start-date")
	enddate := c.FormValue("end-date")
	nodeSelected := false
	nextSelected := false
	reactSelected := false
	typeSelected := false
	if c.FormValue("node") != "" {
		nodeSelected = true
	}
	if c.FormValue("next") != "" {
		nextSelected = true
	}
	if c.FormValue("react") != "" {
		reactSelected = true
	}
	if c.FormValue("type") != "" {
		typeSelected = true
	}

	description := c.FormValue("input-description")
	image := "image.png"

	_, err := connection.Conn.Exec(context.Background(), "UPDATE tb_project SET nama_project=$1, stardate=$2, enddate=$3, node=$4, next=$5, react=$6, type=$7, description=$8, image=$9 WHERE id = $10", projectname, startdate, enddate, nodeSelected, nextSelected, reactSelected, typeSelected, description, image, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"Message ": err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func formRegister(c echo.Context) error {
	// parameter contect yg diinisialkan c, fungsinya menerima request dan respon / kondisi error apabila gagal
	tmpl, err := template.ParseFiles("views/register.html")
	// methode parsefile gunanya mencari alamat yg dicari yaitu register.html
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message ": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func addRegister(c echo.Context) error {
	err := c.Request().ParseForm()
	// method request dari echo berfungsi menerima request/ parse form digunakan untuk menginisiasi form / kemudian ditampung di variabel error
	if err != nil {
		log.Fatal(err)
		// ketika request gagal, maka akan berhenti
	}

	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	// []byte digunakan untuk mengganti dan mencacah tipe data password

	_, err = connection.Conn.Exec(context.Background(), "INSERT INTO tb_user (name, email, password) VALUES ($1, $2, $3)", name, email, passwordHash)
	if err != nil {
		fmt.Println(err)
		redirectWithMessage(c, "Register failed, please try again", false, "/form-register")
	}

	return redirectWithMessage(c, "Register success", true, "/form-login")
}

func formLogin(c echo.Context) error {
	sess, _ := session.Get("session", c)

	flash := map[string]interface{}{
		"FlashStatus":  sess.Values["status"],
		"FlashMessage": sess.Values["message"],
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	tmpl, err := template.ParseFiles("views/login.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message ": err.Error()})
	}

	return tmpl.Execute(c.Response(), flash)
}

func login(c echo.Context) error {
	err := c.Request().ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	email := c.FormValue("email")
	password := c.FormValue("password")

	user := User{}
	err = connection.Conn.QueryRow(context.Background(), "SELECT * FROM tb_user WHERE email=$1", email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return redirectWithMessage(c, "Email Salah !", false, "/form-login")
	}

	fmt.Println(user)
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return redirectWithMessage(c, "Password Salah !", false, "/form-login")
	}

	sess, _ := session.Get("session", c)
	sess.Options.MaxAge = 10800 //3 jam
	sess.Values["message"] = "Login Success !"
	sess.Values["status"] = true //show alert
	sess.Values["name"] = user.Name
	sess.Values["id"] = user.ID
	sess.Values["isLogin"] = true //akses login
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func logout(c echo.Context) error {
	sess, _ := session.Get("session", c)
	sess.Options.MaxAge = -1
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusTemporaryRedirect, "/")
}

func redirectWithMessage(c echo.Context, message string, status bool, path string) error {
	sess, _ := session.Get("session", c)
	sess.Values["message"] = message
	sess.Values["status"] = status
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusMovedPermanently, path)
}
