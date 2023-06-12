package image2pdf

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"

	"github.com/jung-kurt/gofpdf"
)


type Image2PdfConfig struct {
	fileOrigin     []string
	outputPath     string
}

func New(fileOrigin []string, outputPath string) *Image2PdfConfig  {
	return &Image2PdfConfig{
		fileOrigin: fileOrigin,
		outputPath: outputPath,
	}
}


func (config Image2PdfConfig) Convert() error {
	
	pdf := gofpdf.New("P", "mm", "A4", "")
	
	// var x_pos float64 = 0
	var y_pos float64 = 0
	for _, fileDir := range config.fileOrigin {
		file, err := os.Open(fileDir)
		if err != nil{
			return err;
		}
		defer file.Close()
		
		
		image, format, err := image.DecodeConfig(file)
		if err != nil {
			return err
		}		
		pdf.AddPage()
		pdf.Image(fileDir, 0, y_pos ,float64(image.Width), float64(image.Height),false,format,0,"")
		
		y_pos = y_pos + float64(image.Width)
		
		
	}
	
	err := pdf.OutputFileAndClose("./file.pdf")
	if err != nil {
		log.Println("ERROR", err.Error())
	}
	
	return err
}