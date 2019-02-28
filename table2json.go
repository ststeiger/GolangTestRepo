
package main


import (
    "encoding/json"
    "fmt"
    "database/sql"
    _ "github.com/lib/pq"
    //"time"
    // "github.com/go-sql-driver/mysql"
    "github.com/elgs/gosqljson"
)



// CREATE ROLE gowebservices LOGIN PASSWORD 'testgo' SUPERUSER INHERIT NOCREATEDB NOCREATEROLE NOREPLICATION;
const (
    DB_USER     = "gowebservices"
    DB_PASSWORD = "testgo"
    DB_NAME     = "test"
)


type foo struct  
{
	headers []string `json:"headers"`
	data [][]string `json:"data"`
}


func main() {
    dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
        DB_USER, DB_PASSWORD, DB_NAME)
    db, err := sql.Open("postgres", dbinfo)
    checkErr(err)
    defer db.Close()
    
    
    
    theCase := "lower" // "lower", "upper", "camel" or the orignal case if this is anything other than these three
    headers, data, _ := gosqljson.QueryDbToArray(db, theCase, "SELECT uid, username, departname, created FROM userinfo;")
	//fmt.Println(headers)
	// ["id","name"]
	//fmt.Println("daten:")
	//fmt.Println(data)
    
    var P foo = foo{headers, data}
    // P.headers = headers
    // P.data = data
    // fmt.Println(P)
   
    
    
    fmt.Println("struct")
     b, err := json.Marshal(P)
    var hallo string = string(b)
    fmt.Println(hallo)
    // fmt.Println(a)
    
    
    fmt.Println("map")
    dataMap, _ := gosqljson.QueryDbToMap(db, theCase, "SELECT uid, username, departname, created FROM userinfo;")
	fmt.Println(dataMap)
    
    
    fmt.Println("datamap")
     bytearray, err := json.Marshal(dataMap)
    var jsondata string = string(bytearray)
    fmt.Println(jsondata)
    // fmt.Println(a)
    
    
    fmt.Println("# Inserting values")
	/*
    var lastInsertId int
    err = db.QueryRow("INSERT INTO userinfo(username,departname,created) VALUES($1,$2,$3) returning uid;", "astaxie", "研发部门", "2012-12-09").Scan(&lastInsertId)
    checkErr(err)
    fmt.Println("last inserted id =", lastInsertId)

    fmt.Println("# Updating")
    stmt, err := db.Prepare("update userinfo set username=$1 where uid=$2")
    checkErr(err)

    res, err := stmt.Exec("astaxieupdate", lastInsertId)
    checkErr(err)

    affect, err := res.RowsAffected()
    checkErr(err)

    fmt.Println(affect, "rows changed")

    fmt.Println("# Querying")
    rows, err := db.Query("SELECT * FROM userinfo")
    checkErr(err)

    for rows.Next() {
        var uid int
        var username string
        var department string
        var created time.Time
        err = rows.Scan(&uid, &username, &department, &created)
        checkErr(err)
        fmt.Println("uid | username | department | created ")
        fmt.Printf("%3v | %8v | %6v | %6v\n", uid, username, department, created)
    }

	*/

    // fmt.Println("# Deleting")
    //stmt, err = db.Prepare("delete from userinfo where uid=$1")
    //checkErr(err)

    //res, err = stmt.Exec(lastInsertId)
    //checkErr(err)

    // affect, err = res.RowsAffected()
    // checkErr(err)

    // fmt.Println(affect, "rows changed")
}


func checkErr(err error) {
    if err != nil {
		fmt.Println("Error:\n", err)
        panic(err)
    }
}
