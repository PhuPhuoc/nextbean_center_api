package repository

import (
	"bytes"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/PhuPhuoc/hrm_nextbean_api/services/ReportServices/model"
	"github.com/xuri/excelize/v2"
)

func (store *reportStore) GetProjectReportInOJT(oid string) ([]byte, error) {
	f := excelize.NewFile()
	sheetName := "Internship Report"
	f.SetSheetName("Sheet1", sheetName)
	if err := SetupExcel(store, oid, f, sheetName); err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	if err := f.Write(&buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func formatDate(input string) (string, error) {
	t, err := time.Parse(time.RFC3339, input)
	if err != nil {
		return "", err
	}
	return t.Format("02-01-2006"), nil
}

func getOJTDetails(store *reportStore, oid string) (*model.ReportOJT, error) {
	ojt := new(model.ReportOJT)
	rawsql := `select o.semester, o.university, o.start_at, o.end_at, (select count(*) from intern where ojt_id=o.id) as total_intern from ojt o where o.id=?`
	if err := store.db.QueryRow(rawsql, oid).Scan(&ojt.Semester, &ojt.University, &ojt.StartAt, &ojt.EndAt, &ojt.TotalIntern); err != nil {
		return ojt, err
	}
	ojt.StartAt, _ = formatDate(ojt.StartAt)
	ojt.EndAt, _ = formatDate(ojt.EndAt)
	return ojt, nil
}

func SetupExcel(store *reportStore, oid string, f *excelize.File, sheetName string) error {
	interns, count, err_get_intern := getInternInfo(store, oid)
	if err_get_intern != nil {
		return err_get_intern
	}
	internCount := count + 3

	ojt, err_query_ojt := getOJTDetails(store, oid)
	if err_query_ojt != nil {
		return err_query_ojt
	}
	if err := addHeader(f, sheetName, ojt); err != nil {
		return err
	}
	if err := addInternData(f, sheetName, interns); err != nil {
		return err
	}
	if err := addStyles(f, sheetName, internCount); err != nil {
		return err
	}
	if err := addConditionalFormatting(f, sheetName, internCount); err != nil {
		return err
	}
	if err := addNotes(f, sheetName, internCount); err != nil {
		return err
	}
	return nil
}

func getInternInfoDuumyData(oid string) ([]model.ReportIntern, int, error) {
	data := []model.ReportIntern{
		{Name: "Tran Phu Phuoc", Code: "SE173438", ProjectsParticipated: 3, TotalTasks: 20, EstimatedTime: 150, ActualTime: 160, TotalDaysWorkOffline: 30, TotalTimesWorkOffline: 150},
		{Name: "Nguyen Bao Ngoc Anh Thu", Code: "SE173438", ProjectsParticipated: 3, TotalTasks: 20, EstimatedTime: 150, ActualTime: 100, TotalDaysWorkOffline: 29, TotalTimesWorkOffline: 160},
		{Name: "Lam Chi Bao", Code: "SE173438", ProjectsParticipated: 3, TotalTasks: 18, EstimatedTime: 150, ActualTime: 140, TotalDaysWorkOffline: 30, TotalTimesWorkOffline: 150},
		{Name: "Tang Hoang Danh", Code: "SE173438", ProjectsParticipated: 3, TotalTasks: 21, EstimatedTime: 150, ActualTime: 160, TotalDaysWorkOffline: 28, TotalTimesWorkOffline: 170},
		{Name: "Ho Dac Nhan Tam", Code: "SE173438", ProjectsParticipated: 3, TotalTasks: 20, EstimatedTime: 150, ActualTime: 170, TotalDaysWorkOffline: 31, TotalTimesWorkOffline: 160},
		{Name: "Bui Van Tien", Code: "SE173438", ProjectsParticipated: 3, TotalTasks: 15, EstimatedTime: 130, ActualTime: 120, TotalDaysWorkOffline: 20, TotalTimesWorkOffline: 100},
		{Name: "Huynh Khuong Ninh", Code: "SE173438", ProjectsParticipated: 3, TotalTasks: 19, EstimatedTime: 150, ActualTime: 160, TotalDaysWorkOffline: 25, TotalTimesWorkOffline: 140},
		{Name: "Tran Huynh Nhat Anh", Code: "SE173439", ProjectsParticipated: 3, TotalTasks: 20, EstimatedTime: 150, ActualTime: 140, TotalDaysWorkOffline: 35, TotalTimesWorkOffline: 140},
		{Name: "Pham Tien Dat", Code: "SE173430", ProjectsParticipated: 3, TotalTasks: 29, EstimatedTime: 190, ActualTime: 160, TotalDaysWorkOffline: 35, TotalTimesWorkOffline: 120},
		{Name: "Pham Phuc Nghi", Code: "SE173431", ProjectsParticipated: 3, TotalTasks: 20, EstimatedTime: 150, ActualTime: 140, TotalDaysWorkOffline: 25, TotalTimesWorkOffline: 140},
	}
	return data, len(data), nil
}

func rawsqlGetIntern() string {
	var query strings.Builder
	query.WriteString(`select acc.user_name, i.student_code,`)
	query.WriteString(`(select count(*) from project_intern pri where pri.intern_id=i.id and pri.status='in_progress') as total_project,`)
	query.WriteString(`(select count(*) from task t where t.assigned_to=i.id and t.status='completed') as total_task,`)
	query.WriteString(`(select sum(t.estimated_effort) from task t where t.assigned_to=i.id and t.status='completed') as total_est_time,`)
	query.WriteString(`(select sum(t.actual_effort) from task t where t.assigned_to=i.id and t.status='completed') as total_act_time,`)
	query.WriteString(`(select count(*) from timetable ti where ti.intern_id=i.id and ti.status_attendance='present') as total_day_work_offline,`)
	query.WriteString(`(select sum(TIMESTAMPDIFF(hour, ti.act_clockin, ti.act_clockout)) from timetable ti where ti.intern_id = i.id and ti.act_clockin is not null and ti.act_clockout is not null and ti.status_attendance='present') AS total_working_time`)
	query.WriteString(` `)
	query.WriteString(`from intern i join account acc on i.account_id=acc.id where i.ojt_id=?`)
	return query.String()
}

func getInternInfo(store *reportStore, oid string) ([]model.ReportIntern, int, error) {
	data := []model.ReportIntern{}
	rawsql := rawsqlGetIntern()
	rows, err_query := store.db.Query(rawsql, oid)
	if err_query != nil {
		return data, 0, err_query
	}
	defer rows.Close()

	var (
		est_time  sql.NullInt64
		act_time  sql.NullInt64
		work_time sql.NullInt64
	)

	for rows.Next() {
		intern := new(model.ReportIntern)
		if err_scan := rows.Scan(&intern.Name, &intern.Code, &intern.ProjectsParticipated, &intern.TotalTasks, &est_time, &act_time, &intern.TotalDaysWorkOffline, &work_time); err_scan != nil {
			return data, 0, err_scan
		}

		if est_time.Valid {
			intern.EstimatedTime = int(est_time.Int64)
		} else {
			intern.EstimatedTime = 0
		}
		if act_time.Valid {
			intern.ActualTime = int(act_time.Int64)
		} else {
			intern.ActualTime = 0
		}
		if work_time.Valid {
			intern.TotalTimesWorkOffline = int(work_time.Int64)
		} else {
			intern.TotalTimesWorkOffline = 0
		}
		data = append(data, *intern)
	}

	return data, len(data), nil
}

func addHeader(f *excelize.File, sheetName string, ojt *model.ReportOJT) error {
	header := fmt.Sprintf("Semester:   %v          University:   %v          Start time:   %v          End time:   %v          Total intern:   %v", ojt.Semester, ojt.University, ojt.StartAt, ojt.EndAt, ojt.TotalIntern)
	data := [][]interface{}{
		{header},
		{"Intern's information", nil, "Project - Task info", nil, "Working Time Task", nil, "Offline working time", nil},
		{"Student Name", "Student Code", "Projects Participated", "Total Task", "Total Estimated Time (hours)", "Total Actual Time (hours)", "Total Days In Office (days)", "Total working time in Office (hours)"},
	}

	for i, row := range data {
		startCell, err := excelize.JoinCellName("A", i+1)
		if err != nil {
			return fmt.Errorf("error in addHeader: %v", err)
		}
		if err := f.SetSheetRow(sheetName, startCell, &row); err != nil {
			return fmt.Errorf("error in addHeader: %v", err)
		}
	}

	mergeCellRanges := [][]string{{"A1", "H1"}, {"A2", "B2"}, {"C2", "D2"}, {"E2", "F2"}, {"G2", "H2"}}
	for _, ranges := range mergeCellRanges {
		if err := f.MergeCell(sheetName, ranges[0], ranges[1]); err != nil {
			return fmt.Errorf("error in addHeader: %v", err)
		}
	}
	return nil
}

func addInternData(f *excelize.File, sheetName string, interns []model.ReportIntern) error {
	for i, intern := range interns {
		row := []interface{}{
			intern.Name,
			intern.Code,
			intern.ProjectsParticipated,
			intern.TotalTasks,
			intern.EstimatedTime,
			intern.ActualTime,
			intern.TotalDaysWorkOffline,
			intern.TotalTimesWorkOffline,
		}
		if err := f.SetSheetRow(sheetName, "A"+fmt.Sprint(4+i), &row); err != nil {
			return fmt.Errorf("error in addInternData: %v", err)
		}
	}
	return nil
}

func addStyles(f *excelize.File, sheetName string, internCount int) error {
	style1, err := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{Horizontal: "center"},
		Fill:      excelize.Fill{Type: "pattern", Color: []string{"#DFEBF6"}, Pattern: 1},
	})
	if err != nil {
		return fmt.Errorf("error in addStyles: %v", err)
	}
	if err := f.SetCellStyle(sheetName, "A1", "A1", style1); err != nil {
		return fmt.Errorf("error in addStyles: %v", err)
	}

	style2, err := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{Horizontal: "center"},
	})
	if err != nil {
		return fmt.Errorf("error in addStyles: %v", err)
	}
	for _, cell := range []string{"A2", "C2", "E2", "G2"} {
		if err := f.SetCellStyle(sheetName, cell, cell, style2); err != nil {
			return fmt.Errorf("error in addStyles: %v", err)
		}
	}

	if err := f.SetColWidth(sheetName, "A", "A", 30); err != nil {
		return fmt.Errorf("error in addStyles: %v", err)
	}

	if err := f.SetColWidth(sheetName, "B", "B", 15); err != nil {
		return fmt.Errorf("error in addStyles: %v", err)
	}

	if err := f.SetColWidth(sheetName, "C", "C", 25); err != nil {
		return fmt.Errorf("error in addStyles: %v", err)
	}

	if err := f.SetColWidth(sheetName, "D", "D", 15); err != nil {
		return fmt.Errorf("error in addStyles: %v", err)
	}

	if err := f.SetColWidth(sheetName, "E", "E", 30); err != nil {
		return fmt.Errorf("error in addStyles: %v", err)
	}

	if err := f.SetColWidth(sheetName, "F", "F", 25); err != nil {
		return fmt.Errorf("error in addStyles: %v", err)
	}

	if err := f.SetColWidth(sheetName, "G", "G", 30); err != nil {
		return fmt.Errorf("error in addStyles: %v", err)
	}

	if err := f.SetColWidth(sheetName, "H", "H", 35); err != nil {
		return fmt.Errorf("error in addStyles: %v", err)
	}

	disable := true
	if err := f.AddTable(sheetName, &excelize.Table{
		Range:             "A3:H" + fmt.Sprint(internCount),
		Name:              "table",
		StyleName:         "TableStyleMedium2",
		ShowFirstColumn:   false,
		ShowLastColumn:    false,
		ShowRowStripes:    &disable,
		ShowColumnStripes: false,
	}); err != nil {
		return fmt.Errorf("error in addStyles: %v", err)
	}

	if err := f.SetPanes(sheetName, &excelize.Panes{
		Freeze:      true,
		Split:       false,
		XSplit:      0,
		YSplit:      3,
		TopLeftCell: "A4",
		ActivePane:  "bottomLeft",
	}); err != nil {
		return fmt.Errorf("error in addStyles: %v", err)
	}
	return nil
}

func addConditionalFormatting(f *excelize.File, sheetName string, internCount int) error {
	styles := []struct {
		color    string
		fill     string
		pattern  int
		criteria string
		rangeStr string
		value    string
	}{
		{"9A0511", "FEC7CE", 1, "<", "G4:G", "7"},
		{"09600B", "C7EECF", 1, "<=", "F4:F", "E4:E1"},
		{"9A0511", "FFFFE0", 1, ">", "H4:H", "60"},
		{"9A0511", "FFDAB9", 1, ">=", "D4:D", "7"},
	}

	for _, s := range styles {
		format, err := f.NewConditionalStyle(&excelize.Style{
			Font: &excelize.Font{Color: s.color},
			Fill: excelize.Fill{
				Type: "pattern", Color: []string{s.fill}, Pattern: s.pattern,
			},
		})
		if err != nil {
			fmt.Println(err)
			continue
		}

		if err := f.SetConditionalFormat(sheetName, s.rangeStr+fmt.Sprint(internCount),
			[]excelize.ConditionalFormatOptions{
				{Type: "cell", Criteria: s.criteria, Format: format, Value: s.value},
			},
		); err != nil {
			return fmt.Errorf("error in addConditionalFormatting: %v", err)
		}
	}
	return nil
}

func addNotes(f *excelize.File, sheetName string, internCount int) error {
	notes := [][]interface{}{
		{"Note:"},
		{"", "Total Task >= 7 tasks"},
		{"", "Total Actual Time <= Total Estimated Time"},
		{"", "Total Days In Office < 7 days"},
		{"", "Total working time in Office > 60 hours"},
	}

	startRow := internCount + 3 // Assuming data rows end at row 10

	for i, row := range notes {
		cell, err := excelize.JoinCellName("A", startRow+i)
		if err != nil {
			return fmt.Errorf("error in addNotes: %v", err)
		}
		if err := f.SetSheetRow(sheetName, cell, &row); err != nil {
			return fmt.Errorf("error in addNotes: %v", err)
		}
	}

	// Define the colors for the empty cells
	colorStyles := []string{"FFDAB9", "C7EECF", "FEC7CE", "FFFFE0"}
	for i := 1; i < len(notes); i++ {
		cell, err := excelize.JoinCellName("A", startRow+i)
		if err != nil {
			return fmt.Errorf("error in addNotes: %v", err)
		}
		style, err := f.NewStyle(&excelize.Style{
			Fill: excelize.Fill{
				Type:    "pattern",
				Color:   []string{colorStyles[i-1]},
				Pattern: 1,
			},
		})
		if err != nil {
			return fmt.Errorf("error in addNotes: %v", err)
		}
		if err := f.SetCellStyle(sheetName, cell, cell, style); err != nil {
			return fmt.Errorf("error in addNotes: %v", err)
		}
	}

	noteStyle, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{Bold: true, Color: "EE0000"},
	})
	if err != nil {
		return fmt.Errorf("error in addNotes: %v", err)
	}

	if err := f.SetCellStyle(sheetName, "A"+fmt.Sprint(startRow), "A"+fmt.Sprint(startRow), noteStyle); err != nil {
		return fmt.Errorf("error in addNotes: %v", err)
	}
	return nil
}
