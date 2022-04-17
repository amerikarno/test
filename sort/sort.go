package main

import (
	"fmt"
	"sort"
	"time"
)

type Info struct {
	DateTime	string
	DateTimeUTC	time.Time
}

type DateInfo struct {
	Date	string
	Info    []Info
}

type Data struct {
	Data []DateInfo
}

func main() {
	data := []string{
		"20 Jan 2022 15:14:23",
		"22 Jan 2022 12:14:23",
		"21 Jan 2022 14:14:23",
		"21 Jan 2022 11:14:23",
		"20 Jan 2022 10:14:23",
	}

	layout := "2 Jan 2006 15:04:05"
	info := Info{}
	dateInfo := DateInfo{}
	datas := Data{}
	mDate := make(map[string]string)
	mInfo := make(map[string][]Info)
	
	for _, d := range data {
		date := d[:11]
		dtUTC, _ := time.Parse(layout, d)
		info = Info{
			DateTime: d,
			DateTimeUTC: dtUTC,
		}
		_, exist := mDate[date]
		if !exist {
			mDate[date] = date
	}
	_, exist = mInfo[date]
	if exist {
		mInfo[date] = append(mInfo[date], info)
		} else {
			mInfo[date] = append(mInfo[date], info)
		}
	}

	for _, date := range mDate {
		dateInfo = DateInfo{
			Date: date,
			Info: mInfo[date],
		}

		sort.Slice(dateInfo.Info, func(i, j int) bool { 
			d1 := dateInfo.Info[i].DateTimeUTC
			d2 := dateInfo.Info[j].DateTimeUTC
			return d1.Before(d2)
		})
		datas.Data = append(datas.Data, dateInfo)
	}

	fmt.Println(datas)
	// fmt.Println(mDate)
	// fmt.Println(mInfo)
}
