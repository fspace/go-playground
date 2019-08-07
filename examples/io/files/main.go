package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func workingDir() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd err: %v", err)
	}
	fmt.Println("current working dir:", wd)

	fmt.Println("starting dir:", wd)

	if err := os.Chdir("/"); err != nil {
		fmt.Println(err)
		return
	}

	if wd, err = os.Getwd(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("final dir:", wd)
}
func filePath() {
	abs, err := filepath.Abs(".")
	if err != nil {
		log.Fatalln("Abs err:", err)
	}
	fmt.Println("abs is :", abs)
	//
	base := filepath.Base(abs)
	fmt.Println("base is :", base)

	clean := filepath.Clean(abs + "/")
	fmt.Println("clean path is :", clean)
	//clean = filepath.Clean("../some-dir")
	//fmt.Println("clean path2 is :", clean)

	parentDir := filepath.Dir(abs)
	fmt.Println("parent dir is :", parentDir)

	// ## func EvalSymlinks(path string) (string, error):
	ext := filepath.Ext("../some-file.go")
	fmt.Println("extension of file is :", ext)

	p := filepath.FromSlash(abs) // win 下啥都不做
	fmt.Println("path is (in windows):", p)

	matches, err := filepath.Glob("*.go")
	if err != nil {
		log.Fatalln("Glob err: ", err)
	}
	fmt.Println("matches count : ", len(matches))
	for i, match := range matches {
		fmt.Println("i:", i, "file :", match)
	}
	// 不推荐使用啦 用 strings.HasPrefix 替代！
	if filepath.HasPrefix("user/register", "user") {
		fmt.Println("it does has prefix user")
	}
	if filepath.IsAbs(abs) {
		fmt.Println("yes it is abs !")
	}
	fmt.Println("new path is ", filepath.Join(parentDir, "sub-dir", "some.txt"))

	m, err := filepath.Match("*.go", "./main.go")
	if err != nil {
		fmt.Println("match err: ", err)
	}
	fmt.Println("match result : ", m)

	relativePath, err := filepath.Rel(abs, abs+"/some-dir") // 共同前缀？
	if err != nil {
		log.Fatalln("filepath.Rel err:", err)
	}
	fmt.Println("relative path is :", relativePath)

	d, f := filepath.Split(abs)
	fmt.Println("dir is ", d, " file is ", f)
	fmt.Println("split list :", filepath.SplitList(abs))
	fmt.Println("volume name :", filepath.VolumeName(abs)) // 仅win下可用

	fmt.Println()
	filepath.Walk(parentDir, func(path string, info os.FileInfo, err error) error {
		fmt.Println("current path is :", path)
		fmt.Println("file info is :", map[string]interface{}{
			"name":    info.Name(),
			"isDir":   info.IsDir(),
			"modTime": info.Mode(),
			"Size":    info.Size(),
		})
		return nil
	})
	fmt.Println()
}

func main() {

	filePath()
	workingDir()
}
