package repository

import (
	"bytes"
	"fmt"

	"github.com/xuri/excelize/v2"
)

func (store *reportStore) GetProjectReportInOJT(oid string) ([]byte, error) {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	
	sheet1 := "Internship in FALL24"
	index, err := f.NewSheet(sheet1)
	if err != nil {
		return nil, err
	}
	f.DeleteSheet("Sheet1")

	f.SetCellValue(sheet1, "A1", "Hello")
	f.SetCellValue(sheet1, "B1", "World")
	f.SetActiveSheet(index)

	var buf bytes.Buffer
	if err := f.Write(&buf); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
