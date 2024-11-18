package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"os/exec"
	"path/filepath"
)

//for conversion or compression
var flag string

func getDir() (string, error) {
	dir, err := os.Getwd()

	//says if error is not null (nil)
	//yadi error xa vane error throw han
	if err != nil {
		return "", err
	}

	return dir, nil
}

func readArgs() {
	
	if len(os.Args) < 2 {
		fmt.Println("Please provide input and output file paths as arguments.")
		return
	}

	flag = os.Args[1]

	if flag == "help" {
		//how to use
		fmt.Println("=== Welcome to Convertify ===")
		fmt.Println("The ultimate tool for converting and compressing documents and images effortlessly.")
		
		fmt.Println("\nUsage Guide:")
		fmt.Println("\n1) Document Conversion:")
		fmt.Println("   - Convert DOCX to PDF:")
		fmt.Println("     $ convertify docsconvert <input.docx> <output.pdf>")
		fmt.Println("   - Convert PDF to DOCX:")
		fmt.Println("     $ convertify docsconvert <input.pdf> <output.docx>")
		
		fmt.Println("\n2) Image Compression:")
		fmt.Println("   - Compress a JPG image:")
		fmt.Println("     $ convertify imgCompress <input.jpg> <output.jpg>")
		
		fmt.Println("\n3) Image Format Conversion:")
		fmt.Println("   - Convert JPG to PNG:")
		fmt.Println("     $ convertify imgConvert <input.jpg> <output.png>")
		fmt.Println("   - Convert PNG to JPG:")
		fmt.Println("     $ convertify imgConvert <input.png> <output.jpg>")
		
		fmt.Println("\nNotes:")
		fmt.Println("1. Replace `<input>` and `<output>` with the file paths for your source and target files.")
		fmt.Println("2. Ensure you have read/write permissions for the specified directories.")
		// fmt.Println("3. Use absolute paths for better reliability.")
		
		// fmt.Println("\nNeed Help?")
		// fmt.Println("Visit our documentation or contact support for assistance.")
		
		fmt.Println("\n--- Thank you for choosing Convertify ---")
		
		
	}

	if flag == "docsconvert" {

		file1 := os.Args[2]
	   file2 := os.Args[3]
	
	   to := filepath.Ext(file1)
	   from := filepath.Ext(file2)
	
	   fmt.Println("Converting from:", to , "to", from)
	
	   err := converter(to,from,file1,file2)
	   if err != nil {
		   fmt.Println("Error during conversion:", err)
	   }
	}

	if flag == "imgCompress" {
		//image compression invoking code
		toCompressImg := os.Args[2]
		ext := filepath.Ext(toCompressImg)
		compressedImg := os.Args[3]
		compressedImg = compressedImg + ext


		fmt.Println("Compressing image....")
		err := compressJpg(toCompressImg,compressedImg,50)
		fmt.Println("Completed image compression")
		if err != nil {
			fmt.Println("Error during compression:", err)
		}
	}

	if flag == "imgConvert" {

		//image type conversion logicc
		//will implement
		file1 := os.Args[2]
	   file2 := os.Args[3]
	
	   to := filepath.Ext(file1)
	   from := filepath.Ext(file2)

	   fmt.Println("Converting from:", to , "to", from)

	   err := imgConverter(to,from,file1,file2)
	   if err != nil {
		   fmt.Println("Error during conversion:", err)
	   }

	}
}

func convertDocxTopdf(inputPath, outputPath string)  error {
	dir , err := getDir()
	if err != nil {
		fmt.Println("Some issue with directory: ", err)
		return err
	}

	inputDir := dir + string(os.PathSeparator) + inputPath
	outputDir := dir + string(os.PathSeparator) + outputPath
	cmd := exec.Command("powershell", "-ExecutionPolicy", "Bypass", "-File", "scripts/convertDocxToPdf.ps1", 
	"-inputPath", inputDir,
	"-outputPath", outputDir, 
)

	result, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error: %v, output: %s", err, string(result))
	}
	fmt.Println(string(result))
	return nil
}

func convertPdfToDocx (inputPath, outputPath string)  error {
	dir , err := getDir()
	if err != nil {
		fmt.Println("Some issue with directory: ", err)
		return err
	}
	inputDir := dir + string(os.PathSeparator) + inputPath
	outputDir := dir + string(os.PathSeparator) + outputPath

	cmd := exec.Command("powershell", "-ExecutionPolicy", "Bypass", "-File", "scripts/convertPdfToDocx.ps1",
	"-inputPath", inputDir,
	"-outputPath", outputDir,
)
	result, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error: %v, output: %s", err, string(result))
	}
	return nil
}


func converter(toExt string, fromExt string, inputPath string, outputPath string) ( error) {
	if toExt == ".docx" && fromExt == ".pdf" {
	err :=	convertDocxTopdf(inputPath, outputPath)
	if err != nil {
        fmt.Println("Conversion failed:", err)
    } else {
        fmt.Println("Conversion completed successfully!")
    }
	}

	if toExt == ".pdf" && fromExt == ".docx" {
		//convert from pdf to docx
		err := convertPdfToDocx(inputPath, outputPath)
		if err != nil {
			fmt.Println("Conversion failed:", err)
		} else {
			fmt.Println("Conversion completed successfully!")
		}

	}
	return nil
}

func imgConverter(toExt string, fromExt string, inputPath string, outputPath string) ( error) {
	if toExt == ".jpg" && fromExt == ".png" {
		err :=	convertJpgToPng(inputPath, outputPath)
		if err != nil {
			return err
		} else {
			fmt.Println("Conversion completed successfully!")
		}
		}

		if toExt == ".png" && fromExt == ".jpg" {
			//convert from pdf to docx
			err := convertPngToJpg(inputPath, outputPath)
			if err != nil {
				return err
			} else {
				fmt.Println("Conversion completed successfully!")
			}
	
		}

		return nil
}

func convertJpgToPng(inputPath string , outputPath string) error {
	dir , err := getDir()
	if err != nil {
		fmt.Println("Some issue with directory: ", err)
		return err
	}
	//getting the absolute path...
	inputDir := dir + string(os.PathSeparator) + inputPath
	outputDir := dir + string(os.PathSeparator) + outputPath

	file , err := os.Open(inputDir)
	if err != nil {
		return fmt.Errorf("failed to open JPG file: %v", err)
	}
	defer file.Close()

	//decoding jpg
	img, err := jpeg.Decode(file)
	if err != nil {
		return fmt.Errorf("failed to decode JPG: %v", err)
	}

	//create png file
	output , err := os.Create(outputDir)
	if err != nil {
		return fmt.Errorf("failed to create PNG file: %v", err)
	}
	defer output.Close()

	//encode to png
	if err := png.Encode(output,img); err != nil {
		return fmt.Errorf("failed to encode PNG: %v", err)
	}

	fmt.Println("Converted JPG to PNG successfully!")
	return nil

}

func convertPngToJpg(inputPath string , outputPath string) error {
	dir , err := getDir()
	if err != nil {
		fmt.Println("Some issue with directory: ", err)
		return err
	}
	//getting the absolute path...
	inputDir := dir + string(os.PathSeparator) + inputPath
	outputDir := dir + string(os.PathSeparator) + outputPath

	file, err := os.Open(inputDir)
	if err != nil {
		return fmt.Errorf("failed to open PNG file: %v", err)
	}
	defer file.Close()

	// Decode PNG
	img, err := png.Decode(file)
	if err != nil {
		return fmt.Errorf("failed to decode PNG: %v", err)
	}

	output , err := os.Create(outputDir)
	if err != nil {
		return fmt.Errorf("failed to create JPG file: %v", err)
	}
	defer output.Close()

	//encode to jpg with default quality
	options := &jpeg.Options{Quality: 85}
	if err := jpeg.Encode(output, img, options); err != nil {
		return fmt.Errorf("failed to encode JPG: %v", err)
	}

	fmt.Println("Converted PNG to JPG successfully!")
	return nil


}

func compressJpg(inputPath string, outputPath string, quality int) error {
	dir , err := getDir()
	if err != nil {
		fmt.Println("Some issue with directory: ", err)
		return err
	}

	inputDir := dir + string(os.PathSeparator) + inputPath
	outputDir := dir + string(os.PathSeparator) + outputPath
	inputFile, err := os.Open(inputDir)
	if err != nil {
        return fmt.Errorf("failed to open input file: %v", err)
    }
    defer inputFile.Close()

	img,_,err := image.Decode(inputFile)
	if err != nil {
        return fmt.Errorf("failed to decode image: %v", err)
    }

	//creating output image file
	outputFile, err := os.Create(outputDir)
	if err != nil {
        return fmt.Errorf("failed to create output file: %v", err)
    }
    defer outputFile.Close()

   // Set the options for JPEG encoding with the desired quality
   options := &jpeg.Options{Quality: quality}

   err = jpeg.Encode(outputFile, img, options)
   if err != nil {
	   return fmt.Errorf("failed to encode image: %v", err)
   }
   return nil
}

func main() {
	readArgs()
}