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

type SendInfo struct {
	Packs       []PackageInfo
	Name        string // サーバ名
	PackManType string // パッケージマネージャの種類(apt|pacman)
	Arch        string // サーバのOSのArch
}

func SendSrv(info SendInfo, url string) {
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
		url+name,
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
