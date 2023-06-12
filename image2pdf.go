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
	
	for _, fileDir := range config.fileOrigin {
		file, err := os.Open(fileDir)
		if err != nil{
			return err;
		}
		defer file.Close()
		
		
		image, format, err := image.Decode(file)
		if err != nil {
			return err
		}
		
		
		scaledWidth, scaledHeight := rescaleImage(image)

		// Calculate the x and y position to center the image on the PDF page
		x_pos := (210.0 - scaledWidth) / 2.0
		y_pos := (297.0 - scaledHeight) / 2.0
		
		pdf.AddPage()
		pdf.Image(fileDir, x_pos, y_pos ,scaledWidth,scaledHeight,false,format,0,"")
				
		
	}
	
	err := pdf.OutputFileAndClose("./file.pdf")
	if err != nil {
		log.Println("ERROR", err.Error())
	}
	
	return err
}


func rescaleImage(image image.Image) (scaledWidth float64, scaledHeight float64) {
	// Calculate the aspect ratio of the image
	imgWidth := float64(image.Bounds().Dx())
	imgHeight := float64(image.Bounds().Dy())
	aspectRatio := imgWidth / imgHeight
	
	// Set the maximum width and height for the image in mm
	maxWidth := 180.0
	maxHeight := 250.0

	scaledWidth = maxHeight * aspectRatio
	scaledHeight = maxHeight
	
	if (aspectRatio > 1) {
		scaledWidth = maxWidth
		scaledHeight = maxWidth / aspectRatio	
	}

	return scaledWidth, scaledHeight
	
}