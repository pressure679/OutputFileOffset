package main
import (
	"fmt"
	"os"
	"flag"
	"bufio"
)
func main() {
	pwd, err := os.Getwd()
	if err != nil { panic(err) }
	file := flag.String("file", "", "What file to read from.")
	offset := flag.Int64("offset", -1, "What offset to start from.")
	n := flag.Uint("number", 0, "How many bytes to read.")
	flag.Parse()
	if *file == "" { fmt.Println("Which file to read from has to be provided."); os.Exit(1) }
	if *offset == -1 { fmt.Println("Which offset to start from has to be provided"); os.Exit(1) }
	if *n == 0 { fmt.Println("How many bytes to read has to be provided"); os.Exit(1) }
	fmt.Println(*file, *offset, *n, pwd)
	// osFile, err := os.Open(pwd + "/" + *file)
	osFile, err := os.Open(*file)
	if err != nil { panic(err) }
	var buffer []byte = make([]byte, int(*n))
	_, err = osFile.Seek(*offset, 0)
	if err != nil { panic(err) }
	reader := bufio.NewReader(osFile)
	_, err = reader.Read(buffer)
	if err != nil { panic(err) }
	fmt.Println(string(buffer))
}
