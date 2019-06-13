package main

import (
	"github.com/pkg/errors"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {

	srcDir := "/Users/duocnguyen/Pictures/emoji/"
	dstDir := "/Users/duocnguyen/Pictures/emoji-loop/"
	fileList, err := getFileList(srcDir)
	if err != nil {
		log.Panicln(errors.Wrap(err, "[ERROR] can not get file list"))
	}
	for imageName, imagePath := range fileList {
		outputPath := filepath.Join(dstDir, imageName)
		if err := resizeImage(imagePath, outputPath); err != nil {
			log.Printf("Error(%s): %s", imageName, err)
		} else {
			log.Printf("Success: %s", imageName)
		}
	}
	log.Println(fileList)
}

func getFileList(rootDir string) (map[string]string, error) {
	files := map[string]string{}
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(info.Name()) != ".gif" {
			return nil
		}
		files[info.Name()] = path
		return nil
	})
	return files, err
}

func resizeImage(inputPath string, outputPath string) error {
	sourcePath := filepath.Dir(inputPath)
	args := []string{
		//"--resize-fit-width",
		//"100",
		"--lossy=0",
		"-l",
		"-i",
		"--careful",
		"--crop-transparency",
		"--output",
		outputPath,
		inputPath,
	}
	runCommand(sourcePath, "gifsicle", args, nil, nil)
	return nil
}

func runCommand(srcPath string, command string, args []string, env []string, out io.Writer) {
	cmd := exec.Command(command, args...)
	if srcPath != "" {
		cmd.Dir = srcPath
	}
	cmd.Env = os.Environ()
	if env != nil {
		cmd.Env = append(cmd.Env, env...)
	}
	cmd.Stderr = os.Stdout
	if out != nil {
		cmd.Stdout = out
	} else {
		cmd.Stdout = os.Stdout
	}
	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}
}
