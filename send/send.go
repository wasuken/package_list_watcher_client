package send

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type PackageInfo struct {
	Name    string
	Version string
}

type OldPackageInfo struct {
	New PackageInfo
	Old PackageInfo
}

type SendInfo struct {
	Packs       []PackageInfo
	OldPacks    []PackageInfo
	Name        string // サーバ名
	PackManType string // パッケージマネージャの種類(apt|pacman)
}

func SendSrv(tp string, info SendInfo) {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	json, err := json.Marshal(info)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(
		"POST",
		"http://127.0.0.1:3000/api/v1/server/"+name,
		bytes.NewBuffer(json),
	)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	var buf []byte
	_, err = resp.Body.Read(buf)
	if err != nil {
		panic(err)
	}

	err = resp.Body.Close()

	if err != nil {
		panic(err)
	}

	fmt.Println(string(buf))
	fmt.Println("finished.")
}
