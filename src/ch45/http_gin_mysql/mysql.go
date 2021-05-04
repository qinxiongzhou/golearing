package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
	"time"
)

type User struct {
	Id int
	Name string
	Habits string
	CreateTime string
}

var tpl = `<html>
<head>
<title></title>
</head>
<body>
<form action="/info" method="post">
	用户名:<input type="text" name="username">
	兴趣爱好:<input type="text" name="habits">
	<input type="submit" value="提交">
</form>
</body>
</html>`

var db *sql.DB

var err error

func init() {
	db,err = sql.Open("mysql","root:pitaya@tcp(127.0.0.1:3306)/user?charset=utf8")
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/form", submitForm)
	http.HandleFunc("/info", userInfo)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}

func userInfo(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	if r.Method == "POST" {
		user1 := User{Name: r.Form.Get("username"),Habits: r.Form.Get("habits")}
		store(user1)
		fmt.Fprintf(w," %v",queryByName(r.Form.Get("username")))
	}

}

func queryByName(name string) User {
	user := User{}
	stmt,err := db.Prepare("select * from user where name=?")
	checkErr(err)

	rows,_ :=stmt.Query(name)

	for rows.Next(){
		var id int
		var name string
		var habits string
		var createdTime string
		err = rows.Scan(&id, &name, &habits, &createdTime)
		checkErr(err)
		fmt.Printf("[%d, %s, %s, %s]\n", id, name, habits, createdTime)
		user = User{id,name,habits,createdTime}
		break
	}

	return user
}

func store(user User) {
	//插入数据
	stmt,err := db.Prepare("INSERT INTO user SET name=?,habits=?,created_time=?")
	t := time.Now().UTC().Format("2006-01-02")
	res,err := stmt.Exec(user.Name,user.Habits,t)

	checkErr(err)

	id,err := res.LastInsertId()
	checkErr(err)

	fmt.Println("last insert id is: %d \n",id)
}

func submitForm(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:",r.Method)//获取请求的方法
	var t *template.Template
	t = template.New("Products")//创建一个模板
	t,_ = t.Parse(tpl)
	log.Println(t.Execute(w,nil))
}
