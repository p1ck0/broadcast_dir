package util

import "os"

func CreateFile(filename string) (os.File, error){
    file, err := os.Create(filename) 

    if err != nil {
        return *file, err
    }
    
    file.Close()
    return *file, nil
}