package main

import (
	"fmt"
	"sort"
	"time"

	"github.com/pokemonguy1/sessionchecker/model"
	formatter "github.com/pokemonguy1/sessionchecker/timeFormatter"
	"github.com/xuri/excelize/v2"
)

type Session struct {
	name          string
	hall_id       int
	startTime     time.Time
	endTime       time.Time
	session_id    string
	status        string
	interval      int
	priceAdult    int
	priceStudent  int
	priceChildren int
}

func CreateExcel(sessions []model.Seance) {

	sort.Slice(sessions[:], func(i, j int) bool {
		return sessions[i].Reference.Hall < sessions[j].Reference.Hall
	})

	sort.Slice(sessions[:], func(i, j int) bool {
		istartTime := formatter.FormatTime(sessions[i].Timeframe.Start)
		jstartTime := formatter.FormatTime(sessions[j].Timeframe.Start)
		return istartTime.Before(jstartTime)
	})

	file := excelize.NewFile()

	num := 4

	if err := file.MergeCell("Sheet1", "A1", "J1"); err != nil {
		fmt.Println(err)
		return
	}

	jump := true
	bigjump := true

	for i := range sessions {

		istartTime := formatter.FormatTime(sessions[i].Timeframe.Start)
		iendTime := formatter.FormatTime(sessions[i].Timeframe.End)
		ip1startTime := formatter.FormatTime(sessions[i+1].Timeframe.Start)

		row := fmt.Sprintf("%d", i+num) //DECLARATION OF VARIABLES
		startrow := fmt.Sprintf("%d", i+num-1)
		daterow := fmt.Sprintf("%d", num-3)

		starttime := fmt.Sprintf("%d : %d", istartTime.Hour(), istartTime.Minute())
		endtime := fmt.Sprintf("%d : %d", iendTime.Hour(), iendTime.Minute())

		day := fmt.Sprintf("%d", istartTime.Day())
		month := fmt.Sprintf("%v", istartTime.Month())

		if bigjump == true {
			file.SetCellValue("Sheet1", "A"+daterow, "Репертуар на "+day+month)
			bigjump = false
		}

		var duration string

		if iendTime.Minute() < istartTime.Minute() {
			// TODO отдельную  обработчик времени
			hour := iendTime.Hour() - istartTime.Hour()
			minute := iendTime.Minute() - istartTime.Minute()

			duration = fmt.Sprintf("%d : %d", hour, (hour*60+minute)%60)

		} else {
			hour := iendTime.Hour() - istartTime.Hour()
			minute := iendTime.Minute() - istartTime.Minute()

			duration = fmt.Sprintf("%d : %d", hour, minute)
		}

		if jump == true {
			file.SetCellValue("Sheet1", "B"+startrow, "Зал"+fmt.Sprintf(" %d", sessions[i].Reference.Hall)) //NEW STARTING ROW CAUSED BY INCREASE OF HALL ID
			file.SetCellValue("Sheet1", "C"+startrow, "Начало")
			file.SetCellValue("Sheet1", "D"+startrow, "Конец")
			file.SetCellValue("Sheet1", "E"+startrow, "Длительность")
			file.SetCellValue("Sheet1", "F"+startrow, "Разрыв")
			file.SetCellValue("Sheet1", "G"+startrow, "Название")
			file.SetCellValue("Sheet1", "H"+startrow, "Цена Взрослый")
			file.SetCellValue("Sheet1", "I"+startrow, "Цена Студенческий")
			file.SetCellValue("Sheet1", "J"+startrow, "Цена Детский")
			fmt.Println(startrow, "starting row!")

			jump = false
		}

		file.SetCellValue("Sheet1", "B"+row, sessions[i].Reference.Hall) //INFO ROW OF SESSION
		file.SetCellValue("Sheet1", "C"+row, starttime)
		file.SetCellValue("Sheet1", "D"+row, endtime)
		file.SetCellValue("Sheet1", "E"+row, duration)
		file.SetCellValue("Sheet1", "F"+row, sessions[i].Timeframe.Interval)
		file.SetCellValue("Sheet1", "G"+row, sessions[i].Name)
		// file.SetCellValue("Sheet1", "H"+row, sessions[i].priceAdult)
		// file.SetCellValue("Sheet1", "I"+row, sessions[i].priceStudent)
		// file.SetCellValue("Sheet1", "J"+row, sessions[i].priceChildren)
		fmt.Println(row, "info row!")

		if len(sessions)-1 == i {
			file.SetCellValue("Sheet1", "B"+row, sessions[i].Reference.Hall) //INFO ROW OF SESSION & END OF CYCLE
			file.SetCellValue("Sheet1", "C"+row, starttime)
			file.SetCellValue("Sheet1", "D"+row, endtime)
			file.SetCellValue("Sheet1", "E"+row, duration)
			file.SetCellValue("Sheet1", "F"+row, sessions[i].Timeframe.Interval)
			file.SetCellValue("Sheet1", "G"+row, sessions[i].Name)
			// file.SetCellValue("Sheet1", "H"+row, sessions[i].priceAdult)
			// file.SetCellValue("Sheet1", "I"+row, sessions[i].priceStudent)
			// file.SetCellValue("Sheet1", "J"+row, sessions[i].priceChildren)
			fmt.Println(row)
			break
		}

		if ip1startTime.Day() > istartTime.Day() {
			// TODO  отдельную и  должна быть который  с новый
			fmt.Println(row, "do big jump!")
			bigjump = true
			num = num + 10
		}

		if sessions[i+1].Reference.Hall > sessions[i].Reference.Hall {

			file.SetCellValue("Sheet1", "B"+row, sessions[i].Reference.Hall) //INFO ROW OF SESSION
			file.SetCellValue("Sheet1", "C"+row, starttime)
			file.SetCellValue("Sheet1", "D"+row, endtime)
			file.SetCellValue("Sheet1", "E"+row, duration)
			file.SetCellValue("Sheet1", "F"+row, sessions[i].Timeframe.Interval)
			file.SetCellValue("Sheet1", "G"+row, sessions[i].Name)
			// file.SetCellValue("Sheet1", "H"+row, sessions[i].priceAdult)
			// file.SetCellValue("Sheet1", "I"+row, sessions[i].priceStudent)
			// file.SetCellValue("Sheet1", "J"+row, sessions[i].priceChildren)

			jump = true
			fmt.Println("go jump!")
			num = num + 4
		}

	}

	file.SaveAs("schedule.xlsx")
}

func main() {

	seancesss := []model.Seance{
		{
			Name:      "superman",
			Reference: &model.SeanceReference{Hall: "1"},
			Timeframe: &model.SeanceTimeframe{Start: "2022-04-01T19:30:00+06:00", End: "2022-04-01T20:30:00+06:00", Interval: 10},
			Uuid:      "123",
		},
	}

	fmt.Println(seancesss)

	// Sessions := []Session{
	// 	{
	// 		name:       "superman",
	// 		hall_id:    1,
	// 		startTime:  time.Date(2022, time.August, 15, 16, 40, 0, 0, time.UTC),
	// 		endTime:    time.Date(2022, time.August, 15, 17, 20, 0, 0, time.UTC),
	// 		interval:   10,
	// 		session_id: "123",
	// 	},
	// 	{
	// 		name:       "thor",
	// 		hall_id:    1,
	// 		startTime:  time.Date(2022, time.August, 15, 15, 40, 0, 0, time.UTC),
	// 		endTime:    time.Date(2022, time.August, 15, 17, 20, 0, 0, time.UTC),
	// 		interval:   10,
	// 		session_id: "sss1",
	// 	},
	// 	{
	// 		name:       "joker",
	// 		hall_id:    3,
	// 		startTime:  time.Date(2022, time.August, 15, 19, 40, 0, 0, time.UTC),
	// 		endTime:    time.Date(2022, time.August, 15, 20, 30, 0, 0, time.UTC),
	// 		interval:   10,
	// 		session_id: "1s33",
	// 	},
	// 	{
	// 		name:       "jokeryy",
	// 		hall_id:    3,
	// 		startTime:  time.Date(2022, time.August, 15, 19, 40, 0, 0, time.UTC),
	// 		endTime:    time.Date(2022, time.August, 15, 20, 30, 0, 0, time.UTC),
	// 		interval:   10,
	// 		session_id: "1s33",
	// 	},
	// 	{
	// 		name:       "batman",
	// 		hall_id:    1,
	// 		startTime:  time.Date(2022, time.August, 15, 17, 40, 0, 0, time.UTC),
	// 		endTime:    time.Date(2022, time.August, 15, 18, 20, 0, 0, time.UTC),
	// 		interval:   10,
	// 		session_id: "1s3",
	// 	},
	// 	{
	// 		name:       "cock",
	// 		hall_id:    4,
	// 		startTime:  time.Date(2022, time.August, 15, 17, 40, 0, 0, time.UTC),
	// 		endTime:    time.Date(2022, time.August, 15, 19, 20, 0, 0, time.UTC),
	// 		interval:   10,
	// 		session_id: "1nick3",
	// 	},
	// 	{
	// 		name:       "cockyy",
	// 		hall_id:    2,
	// 		startTime:  time.Date(2022, time.August, 15, 17, 40, 0, 0, time.UTC),
	// 		endTime:    time.Date(2022, time.August, 15, 19, 20, 0, 0, time.UTC),
	// 		interval:   10,
	// 		session_id: "1nick3",
	// 	},
	// }

	CreateExcel(seancesss)

}
