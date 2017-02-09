package main

import(
    "path/filepath"
    "fmt"
    "flag"
    "os"
)


func main() {
    flag.Parse()
    path := flag.Arg(0)
    dir, _ := os.Open(path)
    defer dir.Close()

    files, err := dir.Readdir(-1)
    if err != nil {
        panic(err)
    }

    for _, fileinfo := range files {
        if fileinfo.IsDir() {
           fullpath := filepath.Clean(path + string(os.PathSeparator) + fileinfo.Name())
           subdir, _ := os.Open(fullpath)
           defer subdir.Close()
           files, _ := subdir.Readdir(1)
           if len(files) == 0 {
               fmt.Println(fullpath)
               os.Remove(fullpath)
           }
        }
    } 
}

