package main
import (
	"fmt"
	"github.com/astaxie/beego/client/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

// Model Struct
type Person struct {
	PersonId int    `orm:"pk"`
	Name   string `orm:"size(100)"`
}

func init() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:pitaya@tcp(127.0.0.1:3306)/user?charset=utf8")

	// register mode
	orm.RegisterModel(new(Person))

	// create table
	orm.RunSyncdb("default", false, true)
}

func main() {
	o := orm.NewOrm()

	person := Person{Name: "aoho"}

	// insert
	id, err := o.Insert(&person)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	// update
	person.Name = "boho"
	num, err := o.Update(&person)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	// read one
	u := Person{PersonId: person.PersonId}
	err = o.Read(&u)
	fmt.Printf("ERR: %v\n", err)

	var maps []orm.Params

	res, err := o.Raw("SELECT * FROM person").Values(&maps)
	fmt.Printf("NUM: %d, ERR: %v\n", res, err)
	for _, term := range maps {
		fmt.Println(term["person_id"], ":", term["name"])
	}
	// delete
	//num, err = o.Delete(&u)
	//fmt.Printf("NUM: %d, ERR: %v\n", num, err)
}