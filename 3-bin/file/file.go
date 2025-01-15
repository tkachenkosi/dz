package file

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func ReadFile(name string) ([]byte, error) {

	if _, err := os.Stat(name); err != nil {
		fmt.Println(err)
		return nil, err
	}

	if filepath.Ext(name) != ".json" {
		return nil, errors.New("Ошибка. Файла имеет расширение не json")
	}

	data, err := os.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data, nil
}

func WriteFile(content []byte, name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Запись успешная")
}

/*


if _, err := os.Stat("file-exists.go"); err == nil {
    fmt.Printf("File exists\n");
  } else {
    fmt.Printf("File does not exist\n");
  }
}




path := "/media/godfather.mp4"

fileExtension := filepath.Ext(path)

if fileExtension != ".mp4" {
    panic("File extension ins't equal to .mp4")
}

*/
