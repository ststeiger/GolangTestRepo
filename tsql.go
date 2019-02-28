package main

import (
	"bufio"
	"database/sql"
	"flag"
	"fmt"
	"io"
	"os"
	"time"
 // "reflect"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	var (
		// To enable SSO do not provide a user id and a password
		//userid   = flag.String("U", "AnyServiceWebservicesDE", "login_id")
		//password = flag.String("P", "TopSecret", "password")
		userid   = flag.String("U", "", "login_id")
		password = flag.String("P", "", "password")
		server   = flag.String("S", "CORPDB2005", "server_name[\\instance_name]")
		//server   = flag.String("S", "asshat", "server_name[\\instance_name]")
		database = flag.String("d", "Basic_Demo", "db_name")
	)
	flag.Parse()

	dsn := "server=" + *server + ";user id=" + *userid + ";password=" + *password + ";database=" + *database
	db, err := sql.Open("mssql", dsn)
	if err != nil {
		fmt.Println("Cannot connect: ", err.Error())
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("Cannot connect: ", err.Error())
		return
	}
	defer db.Close()
	r := bufio.NewReader(os.Stdin)
	for {
		_, err = os.Stdout.Write([]byte("> "))
		if err != nil {
			fmt.Println(err)
			return
		}
		cmd, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println()
				return
			}
			fmt.Println(err)
			return
		}
		err = exec(db, cmd)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func exec(db *sql.DB, cmd string) error {
	rows, err := db.Query(cmd)
	if err != nil {
		return err
	}
	defer rows.Close()
	cols, err := rows.Columns()
	if err != nil {
		return err
	}
	if cols == nil {
		return nil
	}
	vals := make([]interface{}, len(cols))
	for i := 0; i < len(cols); i++ {
		vals[i] = new(interface{})
		if i != 0 {
			fmt.Print("\t")
		}
		fmt.Print(cols[i])
	}
	fmt.Println()
	for rows.Next() {
		
		err = rows.Scan(vals...)
		if err != nil {
			fmt.Println(err)
			continue
		}
		for i := 0; i < len(vals); i++ {
			if i != 0 {
				fmt.Print("\t")
			}
			// fmt.Print(cols[i] + ": ")
			printValue(vals[i].(*interface{}))
		}
		fmt.Println()

	}
	if rows.Err() != nil {
		return rows.Err()
	}
	return nil
}

func printValue(pval *interface{}) {
	
	// fmt.Println(reflect.TypeOf(*pval))
	// return
	
	switch v := (*pval).(type) {
	case nil:
		fmt.Print("NULL")
	case bool:
		if v {
			fmt.Print("1")
		} else {
			fmt.Print("0")
		}
	case []byte:
		fmt.Print(string(v))
	case time.Time:
		fmt.Print(v.Format("2006-01-02 15:04:05.000"))
	default:
		fmt.Print(v)
	}
}