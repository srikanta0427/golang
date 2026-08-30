package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("example.txt")

	if err != nil {
		panic(err)
	}
	fileInfo, errs := f.Stat()
	if errs != nil {
		panic(errs)
	}
	fmt.Println(fileInfo.Name(), fileInfo.Size(), fileInfo.ModTime())

	//read file
	defer f.Close()

	buf := make([]byte, 128)
	n, err := f.Read(buf)
	if err != nil {
		panic(err)
	}
	fmt.Printf("size of file", n)
	fmt.Println(n, string(buf[:n]))

	/// another way
	data, err := os.ReadFile("example.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))

	//reading directories
	dir, err := os.Open(".")
	if err != nil {
		panic(err)
	}

	defer dir.Close()
	fileInfos, err := dir.Readdir(-1)
	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}

	//create a file
	cf, err := os.Create("hello.txt")
	if err != nil {
		panic(err)
	}
	defer cf.Close()

	_, err = cf.WriteString("Hello, World! Wow\n")
	if err != nil {
		return
	}

	// using byte[]

	bytes := []byte("hello world")
	cf.Write(bytes)

	sourceFile, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	defer sourceFile.Close()

	destFile, err := os.Create("example2.txt")
	if err != nil {
		panic(err)
	}
	defer destFile.Close()

	reader := bufio.NewReader(sourceFile)
	writer := bufio.NewWriter(destFile)

	for {
		b, err := reader.ReadByte()

		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		err = writer.WriteByte(b)
		if err != nil {
			panic(err)
		}
	}
	err = writer.Flush()
	if err != nil {
		panic(err)
	}

	// to delete a file
	err = os.Remove("example2.txt")
	if err != nil {
		panic(err)
	}

}
