package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
)

func NewFileFromPath(path string) (*os.File, error) {
    file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0644)
    if err != nil {
        return nil, fmt.Errorf("failed to create file: %v", err)
    }
    return file, nil
}

func main() {

	app := pocketbase.New()

	app.OnBeforeBootstrap().Add(func(e *core.BootstrapEvent) error {
		log.Println("Hello world, the server is starting, ready ? start !")
		return nil
	})

	app.OnModelAfterCreate().Add(func(e *core.ModelEvent) error {
		go func() {
			if record, ok := e.Model.(*models.Record); ok && record.Collection().Name == "images" {
				fmt.Println("execution is going ...")
				imageId, ok := record.Get("id").(string)
				if !ok {
					fmt.Println("we couldn't get ID")
				}

				fmt.Println("imageId: ")
				fmt.Println(imageId)

				//cmd := exec.Command("python3", "/home/quantumass/Documents/Otterimg/backend/python/script.py", imageId) // replace "ls -la" with your desired command
				//out, err := cmd.Output()
				//if err != nil {
				//	fmt.Println("Error:", err)
				//}
				//fmt.Println(string(out))


				folderToWatch := filepath.Join(os.Getenv("PWD"), "../go/pb_data/storage/c83ukhcc0l9jq90", imageId)
				fmt.Println("Started to watch", folderToWatch)

				record.Set("isProcessing", true)
				app.Dao().SaveRecord(record)


				err := filepath.Walk(folderToWatch, func(path string, info os.FileInfo, err error) error {
					if err != nil {
						return err
					}

					if info.IsDir() && strings.HasPrefix(info.Name(), "thumbs_") {
						return filepath.SkipDir
					}

					if !info.Mode().IsRegular() {
						return nil
					}

					if !strings.HasSuffix(strings.ToLower(info.Name()), ".jpg") && !strings.HasSuffix(strings.ToLower(info.Name()), ".png") && !strings.HasSuffix(strings.ToLower(info.Name()), ".webp") {
						return nil
					}

					if strings.HasSuffix(strings.ToLower(info.Name()), "_scaled.png") {
						return nil
					}

					processedFilePath := filepath.Join(filepath.Dir(path), "folder.processed")
					processingFilePath := filepath.Join(filepath.Dir(path), "folder.processing")

					if _, err := os.Stat(processedFilePath); !os.IsNotExist(err) {
						return nil
					}

					if _, err := os.Stat(processingFilePath); !os.IsNotExist(err) {
						return nil
					}

					fmt.Println("Found file to be processed:", path)

					if err := os.WriteFile(processingFilePath, []byte{}, 0666); err != nil {
						return err
					}

					outputFileName := strings.TrimSuffix(info.Name(), filepath.Ext(info.Name()))+"_scaled.png"

					outputFilePath := filepath.Join(filepath.Dir(path), outputFileName)

					cmd := exec.Command("rembg", "i", path, outputFilePath)
					if err := cmd.Run(); err != nil {
						return err
					}

					if err := os.Remove(processingFilePath); err != nil {
						return err
					}

					if err := os.WriteFile(processedFilePath, []byte{}, 0666); err != nil {
						return err
					}

					fmt.Println("Processed file:", outputFilePath)

					/*file, err := NewFileFromPath(outputFilePath)
					if err != nil {
						fmt.Println("Error on opening file ", err)
					}*/

					record.Set("convertedImage", outputFileName)
					record.Set("isReady", true)
					app.Dao().SaveRecord(record)

					if img, err := imaging.Open(outputFilePath); err == nil {
						if err := imaging.Save(img, outputFilePath, imaging.PNGCompressionLevel(5)); err != nil {
							return err
						}
					} else {
						return err
					}

					return filepath.SkipDir
				})

				if err != nil {
					fmt.Fprintf(os.Stderr, "%v\n", err)
				}



			}
		}()

		return nil

	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
