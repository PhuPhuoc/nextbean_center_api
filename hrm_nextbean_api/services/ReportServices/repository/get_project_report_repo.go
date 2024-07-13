package repository

import (
	"bytes"
	"fmt"

	"github.com/xuri/excelize/v2"
)

func (store *reportStore) GetProjectReportInOJT(oid string) ([]byte, error) {
	f := excelize.NewFile()
	sheetName := "FALL24 Internship Report"
	f.SetSheetName("Sheet1", sheetName)
	// data is 2-dimensional arrays
	data := [][]interface{}{
		{"Semester: Fall2024     University: FPT     Start time: 2024-05-29    End time: 2024-08-12"},
		{"Intern's infomation", nil, "Project - Task info", nil, nil, "Working Time", nil},
		{"Student Name", "Student Code", "Projects Participated", "Total Task", "Total Tasks Completed", "Total Estimated Time", "Total Actual Time"},
		{"Tran Phu Phuoc", "SE173438", 3, 20, 20, 150, 150},
		{"Nguyen Bao Ngoc Anh Thu", "SE173439", 3, 15, 14, 130, 180},
		{"Lam Chi Bao", "SE173401", 3, 21, 20, 150, 160},
		{"Tang Hoang Danh", "SE173402", 3, 19, 19, 160, 140},
		{"Ho Dac Nhan Tam", "SE173403", 3, 20, 18, 150, 160},
		{"Bui Van Tien", "SE173404", 3, 10, 10, 100, 103},
		{"Huynh Khuong Ninh", "SE173405", 3, 20, 19, 150, 130},
	}

	for i, row := range data {
		startCell, err := excelize.JoinCellName("A", i+1)
		if err != nil {
			fmt.Println("error in for loop: ", err)
			return nil, err
		}
		if err := f.SetSheetRow(sheetName, startCell, &row); err != nil {
			fmt.Println("error in set sheet row: ", err)
			return nil, err

		}
	}

	mergeCellRanges := [][]string{{"A1", "G1"}, {"A2", "B2"}, {"C2", "E2"}, {"F2", "G2"}}
	for _, ranges := range mergeCellRanges {
		if err := f.MergeCell(sheetName, ranges[0], ranges[1]); err != nil {
			fmt.Println("error in set sheet row: ", err)
			return nil, err

		}
	}

	style1, err_style1 := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{Horizontal: "center"},
		Fill:      excelize.Fill{Type: "pattern", Color: []string{"#DFEBF6"}, Pattern: 1},
	})
	if err_style1 != nil {
		fmt.Println("error in set sheet style: ", err_style1)
		return nil, err_style1

	}
	if err_set_style1 := f.SetCellStyle(sheetName, "A1", "A1", style1); err_set_style1 != nil {
		fmt.Println("error in set sheet style col A1: ", err_style1)
		return nil, err_set_style1
	}

	style2, err_style2 := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{Horizontal: "center"},
	})
	if err_style2 != nil {
		fmt.Println("error in set sheet style2: ", err_style1)
		return nil, err_style2

	}
	for _, cell := range []string{"A2", "C2", "F2"} {
		if err_style2 := f.SetCellStyle(sheetName, cell, cell, style2); err_style2 != nil {
			fmt.Println("error in set sheet style col 2: ", err_style1)
			return nil, err_style2
		}
	}

	if err := f.SetColWidth(sheetName, "A", "B", 30); err != nil {
		fmt.Println("error in set col width: ", err)
		return nil, err
	}

	if err := f.SetColWidth(sheetName, "C", "E", 25); err != nil {
		fmt.Println("error in set col width: ", err)
		return nil, err

	}

	if err := f.SetColWidth(sheetName, "F", "G", 25); err != nil {
		fmt.Println("error in set col width: ", err)
		return nil, err

	}

	disable := true
	if err := f.AddTable(sheetName, &excelize.Table{
		Range:             "A3:G10",
		Name:              "table",
		StyleName:         "TableStyleMedium2",
		ShowFirstColumn:   false,
		ShowLastColumn:    false,
		ShowRowStripes:    &disable,
		ShowColumnStripes: false,
	}); err != nil {
		fmt.Println("error in set table : ", err)
		return nil, err

	}

	if err := f.SetPanes(sheetName, &excelize.Panes{
		Freeze:      true,
		Split:       false,
		XSplit:      0,
		YSplit:      3,
		TopLeftCell: "A4",
		ActivePane:  "bottomLeft",
	}); err != nil {
		fmt.Println("error in set freeze : ", err)
		return nil, err

	}

	format, err_format := f.NewConditionalStyle(
		&excelize.Style{
			Font: &excelize.Font{Color: "9A0511"},
			Fill: excelize.Fill{
				Type: "pattern", Color: []string{"FEC7CE"}, Pattern: 1,
			},
		},
	)
	if err_format != nil {
		fmt.Println(err_format)
	}
	err_con1 := f.SetConditionalFormat(sheetName, "E4:E10",
		[]excelize.ConditionalFormatOptions{
			{Type: "cell", Criteria: "<", Format: format, Value: "D4:D10"},
		},
	)
	if err_con1 != nil {
		fmt.Println(err_con1)
	}

	format3, err_format3 := f.NewConditionalStyle(
		&excelize.Style{
			Font: &excelize.Font{Color: "09600B"},
			Fill: excelize.Fill{
				Type: "pattern", Color: []string{"C7EECF"}, Pattern: 1,
			},
		},
	)
	if err_format3 != nil {
		fmt.Println(err_format3)
	}

	err_con2 := f.SetConditionalFormat(sheetName, "G4:G10",
		[]excelize.ConditionalFormatOptions{
			{Type: "cell", Criteria: "<=", Format: format3, Value: "F4:F10"},
		},
	)
	if err_con2 != nil {
		fmt.Println(err_con2)
	}
	var buf bytes.Buffer
	if err := f.Write(&buf); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
