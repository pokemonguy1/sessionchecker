package timeFormatter

import (
	"fmt"
	"sort"
	"strconv"
	"time"
)

type Session struct {
	name       string
	hall_id    int
	startTime  time.Time
	endTime    time.Time
	session_id string
	status     string
}

type UTC struct {
	timezone, utcZ string
	timeZ          int
}

func formatTime(rawtime string) time.Time {
	timeUTC := map[string]string{
		"04:00": "UTC+4",
		"05:00": "UTC+5",
		"06:00": "UTC+6",
	}

	UTCval := map[string]int{
		"UTC+4": 4 * 60 * 60,
		"UTC+5": 5 * 60 * 60,
		"UTC+6": 6 * 60 * 60,
	}

	year := string(rawtime[0:4]) // 2022-04-01T19:30:00+06:00
	month := string(rawtime[5:7])
	day := string(rawtime[8:10])

	hour := string(rawtime[11:13])
	mins := string(rawtime[14:16])
	secs := string(rawtime[17:19])

	tz := string(rawtime[20:25])

	yearINT, err := strconv.Atoi(year)
	monthINT, err := strconv.Atoi(month)
	dayINT, err := strconv.Atoi(day)
	hourINT, err := strconv.Atoi(hour)
	minsINT, err := strconv.Atoi(mins)
	secsINT, err := strconv.Atoi(secs)

	if err != nil {
		fmt.Errorf("err")
	}

	value_1 := timeUTC[tz]
	value_2 := UTCval[value_1]

	formattedTime := time.Date(yearINT, time.Month(monthINT), dayINT, hourINT, minsINT, secsINT, 0, time.FixedZone(value_1, value_2))

	return formattedTime
}

type sessionConfict struct { //a slice which contains ids of conflicting sessions
	ids []string
}

type sessionError struct {
	ids []string
}

func CheckSession(sessions []Session) []sessionError {
	sort.Slice(sessions[:], func(i, j int) bool {
		return sessions[i].startTime.Before(sessions[j].startTime)
	})

	sort.Slice(sessions[:], func(i, j int) bool {
		return sessions[i].hall_id < sessions[j].hall_id
	})

	var (
		result []sessionError
	)

	for i, session := range sessions {

		if len(sessions) == i {
			break
		}

		if sessions[i].endTime.Before(session.startTime) == true {
			result = append(result, sessionError{
				ids: []string{sessions[i].session_id},
			})
			fmt.Printf("err: session with id '%s' is ended before it started.  \n", sessions[i].session_id)
			sessions[i].status = "fail"

		} else {
			sessions[i].status = "success"
		}

	}

	return result

}

func CheckSessions(sessions []Session) []sessionConfict {
	sort.Slice(sessions[:], func(i, j int) bool {
		return sessions[i].startTime.Before(sessions[j].startTime)
	})

	sort.Slice(sessions[:], func(i, j int) bool {
		return sessions[i].hall_id < sessions[j].hall_id
	})

	var (
		result []sessionConfict
	)

	for i, session := range sessions {
		fmt.Printf("Name: '%s', id: '%s', Hall: %d, Starting time: %v, Ending time: %v \n", session.name, session.session_id, session.hall_id, session.startTime, session.endTime)

		if len(sessions)-1 == i {
			break
		}

		date := sessions[i+1].startTime.Year() == sessions[i].endTime.Year() &&
			int(sessions[i+1].startTime.Month()) == int(sessions[i].endTime.Month()) &&
			sessions[i+1].startTime.Day() == sessions[i].endTime.Day()

		if date &&
			sessions[i+1].startTime.Hour() < sessions[i].endTime.Hour() &&
			sessions[i+1].hall_id == sessions[i].hall_id {
			result = append(result, sessionConfict{
				ids: []string{sessions[i+1].session_id, sessions[i].session_id}, //adding conflicting sessions' ids into slice
			})
			fmt.Printf("err: two sessions with ids '%s' and '%s' conflicting.  \n", sessions[i+1].session_id, sessions[i].session_id)
			sessions[i+1].status = "fail"
			sessions[i].status = "fail"
		} else if date &&
			sessions[i+1].startTime.Hour() == sessions[i].endTime.Hour() &&
			sessions[i+1].startTime.Minute() <= sessions[i].endTime.Minute() &&
			sessions[i+1].hall_id == sessions[i].hall_id {

			result = append(result, sessionConfict{
				ids: []string{sessions[i+1].session_id, sessions[i].session_id},
			})
			fmt.Printf("err: two sessions with ids '%s' and '%s' conflicting. \n", sessions[i+1].session_id, sessions[i].session_id)
			sessions[i+1].status = "fail"
			sessions[i].status = "fail"
		} else {
			sessions[i+1].status = "success"
			sessions[i].status = "success"
		}

	}

	return result
}

func main() {
	Sessions := []Session{
		{
			name:       "superman",
			hall_id:    1,
			startTime:  time.Date(2022, time.August, 15, 17, 35, 0, 0, time.UTC),
			endTime:    time.Date(2022, time.August, 15, 18, 35, 0, 0, time.UTC),
			session_id: "123",
		},
		{
			name:       "thor",
			hall_id:    2,
			startTime:  time.Date(2022, time.August, 15, 16, 30, 0, 0, time.UTC),
			endTime:    time.Date(2022, time.August, 15, 17, 30, 0, 0, time.UTC),
			session_id: "sss1",
		},
		{
			name:       "joker",
			hall_id:    3,
			startTime:  time.Date(2022, time.August, 15, 20, 40, 0, 0, time.UTC),
			endTime:    time.Date(2022, time.August, 15, 19, 30, 0, 0, time.UTC),
			session_id: "1s33",
		},
		{
			name:       "batman",
			hall_id:    1,
			startTime:  time.Date(2022, time.August, 15, 17, 40, 0, 0, time.UTC),
			endTime:    time.Date(2022, time.August, 15, 18, 20, 0, 0, time.UTC),
			session_id: "1s3",
		},
		{
			name:       "cock",
			hall_id:    4,
			startTime:  time.Date(2022, time.August, 15, 19, 40, 0, 0, time.UTC),
			endTime:    time.Date(2022, time.August, 15, 18, 20, 0, 0, time.UTC),
			session_id: "1nick3",
		},
	}

	result := CheckSessions(Sessions)
	result2 := CheckSession(Sessions)

	fmt.Printf("Conflicting [failed] sessions' ids: %v\n", result)
	fmt.Printf("Error [failed] sessions' ids: %v\n", result2)
}
