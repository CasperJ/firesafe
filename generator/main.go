package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"text/template"
)

func main() {

	modelPathPtr := flag.String("modelpath", "./models", "Relative path")
	outputPathPtr := flag.String("outfile", "./firestore.rules", "Output file")
	// numbPtr := flag.Int("numb", 42, "an int")
	// forkPtr := flag.Bool("fork", false, "a bool")

	flag.Parse()

	absPath, err := filepath.Abs(*modelPathPtr)
	if err != nil {
		panic(err)
	}
	fmt.Println("modelpath:", absPath)
	outputPath, err := filepath.Abs(*outputPathPtr)
	if err != nil {
		panic(err)
	}

	fmt.Println("outfile:", outputPath)

	baseModel, err := getCombinedFiresafeModel(*modelPathPtr)

	f, err := os.Create(outputPath)
	defer f.Close()
	if err != nil {
		panic(err)
	}

	err = generateFirestoreSecurityRules(baseModel, f)
	if err != nil {
		panic(err)
	}
}

func generateFirestoreSecurityRules(model *DocumentRoot, writer io.Writer) error {
	t := template.Must(template.New("main.tmpl").ParseGlob("templates\\*.tmpl"))
	return t.Execute(writer, *model)
}

func getCombinedFiresafeModel(dirPath string) (*DocumentRoot, error) {
	baseModel := &DocumentRoot{
		Models: []*Model{},
	}

	modelFiles, err := findFiresafeModelFiles(dirPath)
	if err != nil {
		return nil, err
	}

	for _, modelFile := range modelFiles {
		log.Println(modelFile)
		bytes, err := os.ReadFile(modelFile)
		if err != nil {
			return nil, err
		}
		var rootModel DocumentRoot
		err = json.Unmarshal(bytes, &rootModel)
		if err != nil {
			return nil, err
		}
		for _, m := range rootModel.Models {
			baseModel.Models = append(baseModel.Models, m)
		}
	}
	return baseModel, nil
}

func findFiresafeModelFiles(dirPath string) (files []string, path error) {

	filepaths := []string{}

	libRegEx, e := regexp.Compile("^.+firesafemodel\\.json$")
	if e != nil {
		log.Fatal(e)
	}

	e = filepath.WalkDir(dirPath, func(path string, dir fs.DirEntry, err error) error {
		if err == nil && libRegEx.MatchString(dir.Name()) {
			absPath, _ := filepath.Abs(path)
			filepaths = append(filepaths, absPath)
		}
		return nil
	})
	if e != nil {
		return nil, e
	}
	return filepaths, nil
}
