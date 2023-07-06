package main

import (
	"fmt"
	"os"

	"github.com/eventials/go-tus"

	"github.com/astaxie/beego/httplib"
)

func main() {

	//打印参数数量
	fmt.Println("命令行参数数量:", len(os.Args))
	//只打印某个参数
	fmt.Println("第二个参数:", os.Args[1])
	//遍历参数并打印
	for k, v := range os.Args {
		fmt.Printf("args[%v]=[%v]\n", k, v)
	}

	flag := os.Args[1]
	if flag == "1" {
		ddxc()
	} else if flag == "0" {
		pt()
	}
}

func pt() {
	var obj interface{}
	req := httplib.Post("http://172.22.90.1:8080/group1/upload")
	req.PostFile("file", "fileserver.exe")
	req.Param("output", "json")
	req.Param("scene", "")
	req.Param("path", "")
	req.ToJSON(&obj)
	fmt.Print(obj)
}

func ddxc() {
	f, err := os.Open("file.jar")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	// create the tus client.
	client, err := tus.NewClient("http://172.22.90.1:8080/group1/big/upload", nil)
	fmt.Println(err)
	// create an upload from a file.
	upload, err := tus.NewUploadFromFile(f)
	fmt.Println(err)
	// create the uploader.
	uploader, err := client.CreateUpload(upload)
	fmt.Println(err)
	// start the uploading process.
	fmt.Println(uploader.Upload())

}
