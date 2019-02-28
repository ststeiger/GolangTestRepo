
package main


import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
    //"time"
     "encoding/json"
)



// CREATE ROLE gowebservices LOGIN PASSWORD 'testgo' SUPERUSER INHERIT NOCREATEDB NOCREATEROLE NOREPLICATION;
const (
    DB_USER     = "gowebservices"
    DB_PASSWORD = "testgo"
    DB_NAME     = "test"
)

/*
func dumpTable(w io.Writer, table) {
    // ...

    rows, err := Query(db, fmt.Sprintf("SELECT * FROM %s", table))
    checkError(err)
    columns, err := rows.Columns()
    checkError(err)

    scanArgs := make([]interface{}, len(columns))
    values   := make([]interface{}, len(columns))

    for i := range values {
        scanArgs[i] = &values[i]
    }

    for rows.Next() {
        err = rows.Scan(scanArgs...)
        checkError(err)

        record := make(map[string]interface{})

        for i, col := range values {
            if col != nil {
                fmt.Printf("\n%s: type= %s\n", columns[i], reflect.TypeOf(col))

                switch t := col.(type) {
                default:
                    fmt.Printf("Unexpected type %T\n", t)
                case bool:
                    fmt.Printf("bool\n")
                    record[columns[i]] = col.(bool)
                case int:
                    fmt.Printf("int\n")
                    record[columns[i]] = col.(int)
                case int64:
                    fmt.Printf("int64\n")
                    record[columns[i]] = col.(int64)
                case float64:
                    fmt.Printf("float64\n")
                    record[columns[i]] = col.(float64)
                case string:
                    fmt.Printf("string\n")
                    record[columns[i]] = col.(string)
                case []byte:   // -- all cases go HERE!
                    fmt.Printf("[]byte\n")
                    record[columns[i]] = string(col.([]byte))
                case time.Time:
                    // record[columns[i]] = col.(string)
                }
            }
        }

        s, _ := json.Marshal(record)
        w.Write(s)
        io.WriteString(w, "\n")
    }
}
*/


func getJSON(sqlString string) (string, error) {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()
	
	
	rows, err := db.Query(sqlString)
	if err != nil {
	  return "", err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
	  return "", err
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	
	jsonData, err := json.Marshal(tableData)
	if err != nil {
		return "", err
	}
	
	// fmt.Println(string(jsonData))
	return string(jsonData), nil 
}


func main() {
    dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
        DB_USER, DB_PASSWORD, DB_NAME)
    db, err := sql.Open("postgres", dbinfo)
    checkErr(err)
    defer db.Close()
    
    
    res, err := getJSON("select * from userinfo");
    fmt.Println(res);
    
    
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
        panic(err)
    }
}
