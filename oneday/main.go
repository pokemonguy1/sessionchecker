package main

import (
	"fmt"
	"sort"
	"time"

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

func CreateExcel(sessions []Session) {

	sort.Slice(sessions[:], func(i, j int) bool {
		return sessions[i].hall_id < sessions[j].hall_id
	})

	sort.Slice(sessions[:], func(i, j int) bool {
		return sessions[i].startTime.Before(sessions[j].startTime)
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

		row := fmt.Sprintf("%d", i+num) //DECLARATION OF VARIABLES
		startrow := fmt.Sprintf("%d", i+num-1)
		daterow := fmt.Sprintf("%d", num-3)

		starttime := fmt.Sprintf("%d : %d", sessions[i].startTime.Hour(), sessions[i].startTime.Minute())
		endtime := fmt.Sprintf("%d : %d", sessions[i].endTime.Hour(), sessions[i].endTime.Minute())

		day := fmt.Sprintf("%d", sessions[i].startTime.Day())
		month := fmt.Sprintf("%v", sessions[i].startTime.Month())

		if bigjump == true {
			file.SetCellValue("Sheet1", "A"+daterow, "Репертуар на "+day+month)
			bigjump = false
		}

		var duration string

		if sessions[i].endTime.Minute() < sessions[i].startTime.Minute() {
			// TODO отдельную  обработчик времени
			hour := sessions[i].endTime.Hour() - sessions[i].startTime.Hour()
			minute := sessions[i].endTime.Minute() - sessions[i].startTime.Minute()

			duration = fmt.Sprintf("%d : %d", hour, (hour*60+minute)%60)

		} else {
			hour := sessions[i].endTime.Hour() - sessions[i].startTime.Hour()
			minute := sessions[i].endTime.Minute() - sessions[i].startTime.Minute()

			duration = fmt.Sprintf("%d : %d", hour, minute)
		}

		if jump == true {
			file.SetCellValue("Sheet1", "B"+startrow, "Зал"+fmt.Sprintf(" %d", sessions[i].hall_id)) //NEW STARTING ROW CAUSED BY INCREASE OF HALL ID
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

		file.SetCellValue("Sheet1", "B"+row, sessions[i].hall_id) //INFO ROW OF SESSION
		file.SetCellValue("Sheet1", "C"+row, starttime)
		file.SetCellValue("Sheet1", "D"+row, endtime)
		file.SetCellValue("Sheet1", "E"+row, duration)
		file.SetCellValue("Sheet1", "F"+row, sessions[i].interval)
		file.SetCellValue("Sheet1", "G"+row, sessions[i].name)
		file.SetCellValue("Sheet1", "H"+row, sessions[i].priceAdult)
		file.SetCellValue("Sheet1", "I"+row, sessions[i].priceStudent)
		file.SetCellValue("Sheet1", "J"+row, sessions[i].priceChildren)
		fmt.Println(row, "info row!")

		if len(sessions)-1 == i {
			file.SetCellValue("Sheet1", "B"+row, sessions[i].hall_id) //INFO ROW OF SESSION & END OF CYCLE
			file.SetCellValue("Sheet1", "C"+row, starttime)
			file.SetCellValue("Sheet1", "D"+row, endtime)
			file.SetCellValue("Sheet1", "E"+row, duration)
			file.SetCellValue("Sheet1", "F"+row, sessions[i].interval)
			file.SetCellValue("Sheet1", "G"+row, sessions[i].name)
			file.SetCellValue("Sheet1", "H"+row, sessions[i].priceAdult)
			file.SetCellValue("Sheet1", "I"+row, sessions[i].priceStudent)
			file.SetCellValue("Sheet1", "J"+row, sessions[i].priceChildren)
			fmt.Println(row)
			break
		}

		if sessions[i+1].startTime.Day() > sessions[i].startTime.Day() {
			// TODO  отдельную и  должна быть который  с новый
			fmt.Println(row, "do big jump!")
			bigjump = true
			num = num + 10
		}

		if sessions[i+1].hall_id > sessions[i].hall_id {

			file.SetCellValue("Sheet1", "B"+row, sessions[i].hall_id)
			file.SetCellValue("Sheet1", "C"+row, starttime)
			file.SetCellValue("Sheet1", "D"+row, endtime)
			file.SetCellValue("Sheet1", "E"+row, duration)
			file.SetCellValue("Sheet1", "F"+row, sessions[i].interval)
			file.SetCellValue("Sheet1", "G"+row, sessions[i].name)
			file.SetCellValue("Sheet1", "H"+row, sessions[i].priceAdult)
			file.SetCellValue("Sheet1", "I"+row, sessions[i].priceStudent)
			file.SetCellValue("Sheet1", "J"+row, sessions[i].priceChildren)

			jump = true
			fmt.Println("go jump!")
			num = num + 4
		}

	}

	file.SaveAs("programming is interesting.xlsx")
}

func main() {
	Sessions := []Session{
		{
			name:       "superman",
			hall_id:    1,
			startTime:  time.Date(2022, time.August, 15, 16, 40, 0, 0, time.UTC),
			endTime:    time.Date(2022, time.August, 15, 17, 20, 0, 0, time.UTC),
			interval:   10,
			session_id: "123",
		},
		{
			name:       "thor",
			hall_id:    1,
			startTime:  time.Date(2022, time.August, 15, 15, 40, 0, 0, time.UTC),
			endTime:    time.Date(2022, time.August, 15, 17, 20, 0, 0, time.UTC),
			interval:   10,
			session_id: "sss1",
		},
		{
			name:       "joker",
			hall_id:    3,
			startTime:  time.Date(2022, time.August, 15, 19, 40, 0, 0, time.UTC),
			endTime:    time.Date(2022, time.August, 15, 20, 30, 0, 0, time.UTC),
			interval:   10,
			session_id: "1s33",
		},
		{
			name:       "jokeryy",
			hall_id:    3,
			startTime:  time.Date(2022, time.August, 15, 19, 40, 0, 0, time.UTC),
			endTime:    time.Date(2022, time.August, 15, 20, 30, 0, 0, time.UTC),
			interval:   10,
			session_id: "1s33",
		},
		{
			name:       "batman",
			hall_id:    1,
			startTime:  time.Date(2022, time.August, 15, 17, 40, 0, 0, time.UTC),
			endTime:    time.Date(2022, time.August, 15, 18, 20, 0, 0, time.UTC),
			interval:   10,
			session_id: "1s3",
		},
		{
			name:       "cock",
			hall_id:    4,
			startTime:  time.Date(2022, time.August, 15, 17, 40, 0, 0, time.UTC),
			endTime:    time.Date(2022, time.August, 15, 19, 20, 0, 0, time.UTC),
			interval:   10,
			session_id: "1nick3",
		},
		{
			name:       "cockyy",
			hall_id:    2,
			startTime:  time.Date(2022, time.August, 15, 17, 40, 0, 0, time.UTC),
			endTime:    time.Date(2022, time.August, 15, 19, 20, 0, 0, time.UTC),
			interval:   10,
			session_id: "1nick3",
		},
		{
			name:       "what's up?",
			hall_id:    2,
			startTime:  time.Date(2022, time.August, 16, 17, 40, 0, 0, time.UTC),
			endTime:    time.Date(2022, time.August, 16, 19, 20, 0, 0, time.UTC),
			interval:   10,
			session_id: "1nick3",
		},
		{
			name:       "welcome to goodland",
			hall_id:    2,
			startTime:  time.Date(2022, time.August, 17, 17, 40, 0, 0, time.UTC),
			endTime:    time.Date(2022, time.August, 17, 19, 20, 0, 0, time.UTC),
			interval:   10,
			session_id: "1nick3",
		},
	}

	CreateExcel(Sessions)

}
