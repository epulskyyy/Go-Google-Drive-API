package main

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/api/option"
	"io"
	"log"
	"os"

	"google.golang.org/api/drive/v3"
)

// func createFolder(service *drive.Service, name string, parentId string) (*drive.File, error) {
// 	d := &drive.File{
// 		Name:     name,
// 		MimeType: "application/vnd.google-apps.folder",
// 		Parents:  []string{parentId},
// 	}

// 	file, err := service.Files.Create(d).Do()

// 	if err != nil {
// 		log.Println("Could not create dir: " + err.Error())
// 		return nil, err
// 	}

// 	return file, nil
// }

func createFile(service *drive.Service, name string, mimeType string, content io.Reader, parentId string) (*drive.File, error) {
	f := &drive.File{
		MimeType: mimeType,
		Name:     name,
		Parents:  []string{parentId},
	}
	file, err := service.Files.Create(f).Media(content).Do()

	if err != nil {
		log.Println("Could not create file: " + err.Error())
		return nil, err
	}

	return file, nil
}

func getFiles(service *drive.Service)(*drive.FileList, error){
	files, err := service.Files.List().Do()
	if err != nil {
		log.Println("Could not create file: " + err.Error())
		return nil, err
	}
	return files , nil
}

func main() {
	//Open  file
	f, err := os.Open("contoh.txt")

	if err != nil {
		panic(fmt.Sprintf("cannot open file: %v", err))
	}

	defer f.Close()

	ctx := context.Background()
	srv, err := drive.NewService(ctx,option.WithCredentialsFile("client_secret.json"))
	if err != nil {
		log.Fatalf("Unable to retrieve drive Client %v", err)
	}

	// Create directory
	// dir, err := createFolder(srv, "New Folder", "root")
	// if err != nil {
	// 	panic(fmt.Sprintf("Could not create dir: %v\n", err))
	// }

	//give your folder id here in which you want to upload or create new directory
	folderId := "1rQq4sE3JYLDXdtRuM7iW6fhmGzVI1LwQ"

	//create the file and upload
	file, err := createFile(srv, f.Name(), "application/octet-stream", f, folderId)

	if err != nil {
		panic(fmt.Sprintf("Could not create file: %v\n", err))
	}
	files , errFiles := getFiles(srv)
	if errFiles != nil {
		panic(fmt.Sprintf("Could not create file: %v\n", err))
	}

	out, err := json.Marshal(files.Files)
	if err != nil {
		panic (err)
	}
	fmt.Printf("File '%s' udah ke upload", file.Name)
	fmt.Printf("Files '%s'get files cuy", out)
	fmt.Printf("\nFile Id: '%s' ", file.Id)

}
