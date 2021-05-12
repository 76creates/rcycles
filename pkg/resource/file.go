package resource

import (
	"fmt"
	"io/ioutil"
)

type File struct {
 	Base
}

type fileResource struct {
	filename string
	filepath string
}

type fileResources []*fileResource


func (r File)Lookup() error {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		return err
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}
	return nil
}