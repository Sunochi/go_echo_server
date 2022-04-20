package csv

import (
	"io"
	"log"
	"mime/multipart"
	"os"
	"strings"
	"time"

	"github.com/jszwec/csvutil"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

// SJIS形式のCSVをUTF8形式に変換して構造体にいれる。
// TODO: 文字コードの判別を追加する - UTF8の場合は変換処理を通さないべき
func ReadFile(fh *multipart.FileHeader, s interface{}) {
	csvFile, err := fh.Open()
	if err != nil {
		log.Println("csv file open error:", err)
	}
	defer csvFile.Close()

	csv, err := io.ReadAll(csvFile)
	log.Printf("%s", csv)
	if err != nil {
		panic(err)
	}
	// SJIS形式 => UTF8形式
	b, err := io.ReadAll(transform.NewReader(strings.NewReader(string(csv)), japanese.ShiftJIS.NewDecoder()))
	if err != nil {
		log.Println("csv trans error:", err)
	}

	err = csvutil.Unmarshal(b, s)
	if err != nil {
		log.Println("csv file unmarshal error:", err)
	}
}

// 構造体を受け取りSJIS形式のCSVとして出力する。
func WriteFile(filePrefix string, s interface{}) (string, string, error) {
	b, err := csvutil.Marshal(s)
	if err != nil {
		log.Println("struct to csv error:", err)
		return "", "", err
	}
	// UTF8 => SJIS形式
	sjisByte, _, err := transform.Bytes(japanese.ShiftJIS.NewEncoder(), b)
	if err != nil {
		log.Println("csv trans error:", err)
		return "", "", err
	}

	fileName := filePrefix + "_" + time.Now().Format("20060102150405.csv")
	filePath := "csv/" + fileName
	err = os.WriteFile(filePath, sjisByte, 0666)
	if err != nil {
		log.Println("file write error:", err)
		return "", "", err
	}
	return filePath, fileName, nil
}
