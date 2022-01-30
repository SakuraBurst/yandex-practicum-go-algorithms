package schedule

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type LessonTime struct {
	startHour        int
	startMinute      int
	endHour          int
	endMinute        int
	lessonTimeString string
}

type AuditorySchedule []LessonTime

func Schedule(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())
	auditorySchedule := make(AuditorySchedule, 0, N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		lessonData := strings.Fields(scanner.Text())
		startLessonData := strings.Split(lessonData[0], ".")
		endLessonData := strings.Split(lessonData[1], ".")
		startHour, _ := strconv.Atoi(startLessonData[0])
		startMinute := 0
		if len(startLessonData) > 1 {
			startMinute, _ = strconv.Atoi(startLessonData[1])
		}
		endHour, _ := strconv.Atoi(endLessonData[0])
		endMinute := 0
		if len(endLessonData) > 1 {
			endMinute, _ = strconv.Atoi(endLessonData[1])
		}
		auditorySchedule = append(auditorySchedule, LessonTime{
			startHour:        startHour,
			startMinute:      startMinute,
			endHour:          endHour,
			endMinute:        endMinute,
			lessonTimeString: scanner.Text(),
		})
	}
	auditorySchedule = merge(auditorySchedule)
	var lastLessonTime *LessonTime
	writer := bufio.NewWriter(w)
	counter := 0
	stringBuilder := strings.Builder{}
	for i, time := range auditorySchedule {
		if lastLessonTime != nil {
			if time.startHour < lastLessonTime.endHour {
				continue
			}
			if time.startHour == lastLessonTime.endHour {
				if time.startMinute < lastLessonTime.endMinute {
					continue
				}
			}
		}
		counter++
		stringBuilder.WriteString(time.lessonTimeString)
		stringBuilder.WriteByte('\n')
		lastLessonTime = &auditorySchedule[i]
	}
	writer.WriteString(strconv.Itoa(counter))
	writer.WriteByte('\n')
	writer.WriteString(stringBuilder.String())
	writer.Flush()
}

func merge(schedule AuditorySchedule) AuditorySchedule {
	scheduleLength := len(schedule)
	if scheduleLength < 2 {
		return schedule
	}
	mid := scheduleLength >> 1
	a := merge(schedule[:mid])
	aLength := len(a)
	b := merge(schedule[mid:])
	bLength := len(b)
	aIndex := 0
	bIndex := 0
	result := make(AuditorySchedule, 0, scheduleLength)
	for aIndex < aLength || bIndex < bLength {
		if aIndex == aLength {
			result = append(result, b[bIndex])
			bIndex++
			continue
		}
		if bIndex == bLength {
			result = append(result, a[aIndex])
			aIndex++
			continue
		}
		if LessonTimeLess(a[aIndex], b[bIndex]) {
			result = append(result, a[aIndex])
			aIndex++
		} else {
			result = append(result, b[bIndex])
			bIndex++
		}
	}
	return result
}

func LessonTimeLess(a, b LessonTime) bool {
	if a.endHour == b.endHour {
		if a.endMinute == b.endMinute {
			return StartLess(a, b)
		}
		return a.endMinute < b.endMinute
	}
	return a.endHour < b.endHour

}

func StartLess(a, b LessonTime) bool {
	if a.startHour == b.startHour {
		return a.startMinute < b.startMinute
	}

	return a.startHour < b.startHour

}

func main() {
	Schedule(os.Stdin, os.Stdout)
}
