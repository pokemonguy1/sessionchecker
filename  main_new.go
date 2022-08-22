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
	priceAdult    float64
	priceStudent  float64
	priceChildren float64
}

func CreateExcel(sessions []Session) {

	sort.Slice(sessions[:], func(i, j int) bool {
		return sessions[i].startTime.Before(sessions[j].startTime)
	})

	sort.Slice(sessions[:], func(i, j int) bool {
		return sessions[i].hall_id < sessions[j].hall_id
	})

	file := excelize.NewFile()

	if err := file.MergeCell("Sheet1", "A1", "J1"); err != nil {
		fmt.Println(err)
		return
	}

	CreateRow := func(row string, hid int, startT string, endT string, dur string, interval int, name string, priceA float64, priceS float64, priceC float64) {
		file.SetCellValue("Sheet1", "B"+row, hid) //INFO ROW OF SESSION
		file.SetCellValue("Sheet1", "C"+row, startT)
		file.SetCellValue("Sheet1", "D"+row, endT)
		file.SetCellValue("Sheet1", "E"+row, dur)
		file.SetCellValue("Sheet1", "F"+row, interval)
		file.SetCellValue("Sheet1", "G"+row, name)
		file.SetCellValue("Sheet1", "H"+row, priceA)
		file.SetCellValue("Sheet1", "I"+row, priceS)
		file.SetCellValue("Sheet1", "J"+row, priceC)
	}

	doRowJump := func(erowN int) {

		erow := fmt.Sprintf("%d", erowN)

		file.SetCellValue("Sheet1", "B"+erow, "") //INFO ROW OF SESSION
		file.SetCellValue("Sheet1", "C"+erow, "")
		file.SetCellValue("Sheet1", "D"+erow, "")
		file.SetCellValue("Sheet1", "E"+erow, "")
		file.SetCellValue("Sheet1", "F"+erow, "")
		file.SetCellValue("Sheet1", "G"+erow, "")
		file.SetCellValue("Sheet1", "H"+erow, "")
		file.SetCellValue("Sheet1", "I"+erow, "")
		file.SetCellValue("Sheet1", "J"+erow, "")
	}

	num := 5
	jump := true
	bigjump := true

	for i := range sessions {

		daterow := fmt.Sprintf("%d", i+num)
		startrow := fmt.Sprintf("%d", i+num-1)
		row := fmt.Sprintf("%d", i+num-2)

		starttime := fmt.Sprintf("%d : %d", sessions[i].startTime.Hour(), sessions[i].startTime.Minute())
		endtime := fmt.Sprintf("%d : %d", sessions[i].endTime.Hour(), sessions[i].endTime.Minute())

		var duration string

		if sessions[i].endTime.Minute() < sessions[i].startTime.Minute() {

			hour := sessions[i].endTime.Hour() - sessions[i].startTime.Hour()
			minute := sessions[i].endTime.Minute() - sessions[i].startTime.Minute()

			duration = fmt.Sprintf("%d : %d", hour, (hour*60+minute)%60)

		} else {
			hour := sessions[i].endTime.Hour() - sessions[i].startTime.Hour()
			minute := sessions[i].endTime.Minute() - sessions[i].startTime.Minute()

			duration = fmt.Sprintf("%d : %d", hour, minute)
		}

		if bigjump == true {

			fmt.Println("about to jump...")

			run := num

			for i := 1; i < 5; i++ {
				doRowJump(i + run)
				fmt.Println(run)
				run++
			}

			num = run

			fmt.Println("jumped! ", num)
			fmt.Println(row, "date row")

			file.SetCellValue("Sheet1", "A"+daterow, "Repertoire for [date]")

		}

		if jump == true || bigjump == true {

			file.SetCellValue("Sheet1", "B"+startrow, "Зал"+fmt.Sprintf(" %d", sessions[i].hall_id)) //NEW STARTING ROW CAUSED BY INCREASE OF HALL ID
			file.SetCellValue("Sheet1", "C"+startrow, "Начало")
			file.SetCellValue("Sheet1", "D"+startrow, "Конец")
			file.SetCellValue("Sheet1", "E"+startrow, "Длительность")
			file.SetCellValue("Sheet1", "F"+startrow, "Разрыв (мин)")
			file.SetCellValue("Sheet1", "G"+startrow, "Название")
			file.SetCellValue("Sheet1", "H"+startrow, "Цена Взрослый")
			file.SetCellValue("Sheet1", "I"+startrow, "Цена Студенческий")
			file.SetCellValue("Sheet1", "J"+startrow, "Цена Детский")
			fmt.Println(startrow, "starting row!")

			bigjump = false
			jump = false
		}

		CreateRow(row, sessions[i].hall_id, starttime, endtime, duration, sessions[i].interval, sessions[i].name, sessions[i].priceAdult, sessions[i].priceStudent, sessions[i].priceChildren)
		fmt.Println(row, "info row!")

		if len(sessions)-1 == i {
			CreateRow(row, sessions[i].hall_id, starttime, endtime, duration, sessions[i].interval, sessions[i].name, sessions[i].priceAdult, sessions[i].priceStudent, sessions[i].priceChildren)

			break
		}

		if sessions[i+1].hall_id > sessions[i].hall_id {

			CreateRow(row, sessions[i].hall_id, starttime, endtime, duration, sessions[i].interval, sessions[i].name, sessions[i].priceAdult, sessions[i].priceStudent, sessions[i].priceChildren)

			fmt.Println(row, "jump!", " ", num)

			num = num + 5

			fmt.Println("		!Number: ", num)

			jump = true
		}

	}

	//}

	file.SetCellValue("Sheet1", "F2", "Репертуар на 15 августа")
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
	}

	CreateExcel(Sessions)

}
