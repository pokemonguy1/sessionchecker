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

// TODO на один день 3-4 зала
func CreateExcel(sessions []Session) {

	sort.Slice(sessions[:], func(i, j int) bool {
		return sessions[i].startTime.Before(sessions[j].startTime)
	})

	sort.Slice(sessions[:], func(i, j int) bool {
		return sessions[i].hall_id < sessions[j].hall_id
	})

	file := excelize.NewFile()

	num := 4
	/// TODO создать массив из букв альфавита и пробегаться через него
	/// TODO возможно функция Next() (row, col)

	rowMaker := func(row string, hid int, startT string, endT string, dur string, interval int, name string, priceA float64, priceS float64, priceC float64) {
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

			if err := file.MergeCell("Sheet1", "A"+daterow, "J"+daterow); err != nil {
				fmt.Println(err)
				return
			}

			file.SetCellValue("Sheet1", "A"+daterow, "Репертуар на "+day+" "+month)
			bigjump = false
			fmt.Println(daterow, "NEW LIST", "NUM: ", num)
		}

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

		if jump == true {

			file.SetCellValue("Sheet1", "B"+startrow, "Зал"+fmt.Sprintf(" %d", sessions[i].hall_id)) //NEW STARTING ROW CAUSED BY INCREASE OF HALL ID
			file.SetCellValue("Sheet1", "C"+startrow, "Начало")
			file.SetCellValue("Sheet1", "D"+startrow, "Конец")
			file.SetCellValue("Sheet1", "E"+startrow, "Длительность")
			file.SetCellValue("Sheet1", "F"+startrow, "Разрыв (мин)")
			file.SetCellValue("Sheet1", "G"+startrow, "Название")
			file.SetCellValue("Sheet1", "H"+startrow, "Цена Взрослый")
			file.SetCellValue("Sheet1", "I"+startrow, "Цена Студенческий")
			file.SetCellValue("Sheet1", "J"+startrow, "Цена Детский")
			fmt.Println(startrow, "starting row!", "NUM: ", num)

			jump = false
		}

		if len(sessions)-1 == i {
			rowMaker(row, sessions[i].hall_id, starttime, endtime, duration, sessions[i].interval, sessions[i].name, sessions[i].priceAdult, sessions[i].priceStudent, sessions[i].priceChildren)
			fmt.Println(row)
			break
		}

		if sessions[i+1].startTime.Day() == sessions[i].startTime.Day() {
			rowMaker(row, sessions[i].hall_id, starttime, endtime, duration, sessions[i].interval, sessions[i].name, sessions[i].priceAdult, sessions[i].priceStudent, sessions[i].priceChildren)
		}

		fmt.Println(row, "info row!")

		if sessions[i+1].startTime.Day() > sessions[i].startTime.Day() {
			fmt.Println(row, "do big jump!")

			file.SetCellValue("Sheet1", "B"+startrow, "Зал"+fmt.Sprintf(" %d", sessions[i].hall_id)) //NEW STARTING ROW CAUSED BY INCREASE OF HALL ID
			file.SetCellValue("Sheet1", "C"+startrow, "Начало")
			file.SetCellValue("Sheet1", "D"+startrow, "Конец")
			file.SetCellValue("Sheet1", "E"+startrow, "Длительность")
			file.SetCellValue("Sheet1", "F"+startrow, "Разрыв (мин)")
			file.SetCellValue("Sheet1", "G"+startrow, "Название")
			file.SetCellValue("Sheet1", "H"+startrow, "Цена Взрослый")
			file.SetCellValue("Sheet1", "I"+startrow, "Цена Студенческий")
			file.SetCellValue("Sheet1", "J"+startrow, "Цена Детский")
			fmt.Println(startrow, "starting row!", "NUM: ", num)

			bigjump = true
			num = num + 6
		}

		if sessions[i+1].hall_id > sessions[i].hall_id {
			rowMaker(row, sessions[i].hall_id, starttime, endtime, duration, sessions[i].interval, sessions[i].name, sessions[i].priceAdult, sessions[i].priceStudent, sessions[i].priceChildren)
			jump = true
			fmt.Println(row, "go jump!", "NUM: ", num)
			num = num + 4
			fmt.Println("number + 4: ", num)
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
