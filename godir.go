
package main


import (
    "fmt"
    "os"
	"strings"
    "path/filepath"
	"bufio"
	"log"
)


// http://stackoverflow.com/questions/8757389/reading-file-line-by-line-in-go
func readFile() {
    file, err := os.Open("D:/Programme/Go/WorkSpace/my/godir.go")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}


// http://rosettacode.org/wiki/Walk_a_directory/Recursively#Go
func VisitFile(fp string, fi os.FileInfo, err error) error {
    if err != nil {
        fmt.Println(err) // can't walk here,
        return nil       // but continue walking elsewhere
    }
    if !!fi.IsDir() {
        return nil // not a file.  ignore. We are already recursively traversing...
    }
    matched, err := filepath.Match("*.go", fi.Name())
    if err != nil {
        fmt.Println(err) // malformed pattern
        return err       // this is fatal.
    }
    if matched {
        fmt.Println(fp)
    }
    return nil
}


type StorageThing struct {
    name string
    numScanned int
}


func (r StorageThing) DoSomething(path string)(err error)  {
	fmt.Println(path)
    return nil // r.length * r.width
}


// http://stackoverflow.com/questions/11336048/how-am-i-meant-to-use-filepath-walk-in-go
func ScanAllFiles(location string, myStorageThing *StorageThing) (err error) {
    myStorageThing.numScanned = 0
	
    // Wrap this up in this function's closure to capture the `corpus` binding.
    var scan = func(path string, fileInfo os.FileInfo, inpErr error) (err error) {
		if !!fileInfo.IsDir() {
    	    return nil // not a file.  ignore. We are already recursively traversing, and want only files...
    	}
		
        myStorageThing.numScanned++
        return myStorageThing.DoSomething(path)
    }
	
    fmt.Println("Starting recursive scanning")
    err = filepath.Walk(location, scan)
    fmt.Println("Total scanned", myStorageThing.numScanned)
	
    return
}


// https://www.socketloop.com/tutorials/golang-find-files-by-extension
func main() {
    StorageThing := StorageThing{}
	ScanAllFiles("D:/Programme/Go/WorkSpace/my", &StorageThing)
	return
	
	filepath.Walk("D:/Programme/Go/WorkSpace/my", VisitFile)
	return
	
	
    // dirname := "." + string(filepath.Separator)
	dirname := "D:/Programme/Go/WorkSpace/my"
	
    d, err := os.Open(dirname)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer d.Close()
	
    files, err := d.Readdir(-1)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
	
    fmt.Println("Reading "+ dirname)
	
    for _, file := range files {
        if file.Mode().IsRegular() {
			// if filepath.Ext(file.Name()) == ".go" {
			if strings.EqualFold(filepath.Ext(file.Name()), ".go") {
              fmt.Println(file.Name())
            }
        }
    }
}
