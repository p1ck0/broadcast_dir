package util

import "os"

func CreateDir(dirname string) {
    if !CheckFile(dirname) {
        err := os.Mkdir(dirname, os.ModePerm)
        if err != nil {
            panic(err)
        }
    }
}