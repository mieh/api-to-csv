package output

import "fmt"

func ToConsole(csv string) error {
	fmt.Println(csv)
	return nil
}
