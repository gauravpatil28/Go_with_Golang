package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("File Handling in GO")

	content := "This needs to be gone inside the File we create."

	file, err := os.Create("./CreatedFile")

	/*he varchi command ne file create hote ./ mhanje current directory madhe havi aahe file
	ani hya example madhe mi file mhanun file create karat aahe ani content mhanje techyat kay content thevaych te aahe*/

	/*OS package ne apan create karto file ani mag IO use karun techyat given values takto*/

	if err != nil { //Commonly used to check if there is an error in the code, If there is any then the
		panic(err) //Panic commond stops the program and displays the err.
	}

	display, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	checkNilerr(err)

	fmt.Println(display)

	defer file.Close()
	/*Apan basically defer vaparto close file chya adhi karan incase apan ajun operation kele tya file var tar
	te execute hoyla have mhanun defer so at the ending of the file the file gets close
	Its a good practice to follow.*/

	readfile("./CreatedFile")
	/*ikde tula bracket madhe ek proper file path deyla lagto magach file read hote*/
}

/*
Apan yeun jaun kadhi tasa vaprat baslo tar khup space jate ani mhanun apn ek function banvun gheto ani apan
mag after every comma ok synatx can just call the function and check if there is an error or not
*/
func checkNilerr(err error) {
	if err != nil {
		panic(err)
	}
}

func readfile(filename string) {
	show, err := os.ReadFile(filename)

	checkNilerr(err)

	fmt.Println("The Given content in Your file is \n", string(show))
}

/*Ani show mhanje majhya file la kadhi mi asa dakhavla tar te mala bytes dakhvel techyatla data nahi
mhanun mala tyala string madhe wrapped dakhvayla lagel if i want to see the data.*/
