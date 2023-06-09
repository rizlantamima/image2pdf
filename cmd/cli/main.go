package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/rizlantamima/image2pdf"
)

func main()  {

	helpFlag := flag.Bool("h", false, "Menampilkan penjelasan tentang aplikasi ini")
	
	flag.Parse()

	if *helpFlag {
		printHelp()
		return
	}

	if(len(os.Args) < 2){
		fmt.Println("at least you need provide 1 image file as arguments, example : image2pdf <path/to/image>")
		os.Exit(0);
	}

	inputImage :=  getArgsWithoutFlags(os.Args[1:])

	config := image2pdf.New(inputImage, "")

	err := config.Convert()
	if err != nil {
		panic(err)
	}
	
}

func printHelp() {
	fmt.Println("Image2Pdf")
	fmt.Println("Image2Pdf is application that help user to convert any image to pdf files")
	fmt.Println("instead of upload the image to the internet (using free web services) we dont know where that file will go")	
	fmt.Println("Usage: ./image2pdf <multiple-destination-file-path-at-least-one> [-h]")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -h    Show the application's explanation")
}

func getArgsWithoutFlags(args []string) []string {
	var filteredArgs []string

	for _, arg := range args {
		if !strings.HasPrefix(arg, "-") {
			filteredArgs = append(filteredArgs, arg)
		}
	}

	return filteredArgs
}