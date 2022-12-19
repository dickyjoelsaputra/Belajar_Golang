package main

import (
	"fmt"
	"time"
)

// Time, Parsing Time, & Format Time
func main() {
	var time1 = time.Now()
    fmt.Printf("time1 %v\n", time1)
    // time1 2015-09-01 17:59:31.73600891 +0700 WIB

    var time2 = time.Date(2011, 12, 24, 10, 20, 0, 0, time.UTC)
    fmt.Printf("time2 %v\n", time2)
    // time2 2011-12-24 10:20:00 +0000 UTC
    // time.Date(tahun, bulan, tanggal, jam, menit, detik, nanodetik, timezone)

    fmt.Println("time3 ", time1.Year() , time1.Local() , time1.Weekday())

    // Parsing dari string ke time.Time
    fmt.Println("===================================")
    var layoutFormat, value string
    var date time.Time

    // layout format nya
    layoutFormat = "2006-01-02 15:04:05"
    // data asli nta
    value = "2015-09-02 08:04:00"
    // time parse memiliki 2 parameter
    date, _ = time.Parse(layoutFormat, value)

    fmt.Println(value, "\t->", date.String())
    // 2015-09-02 08:04:00 +0000 UTC

    layoutFormat = "02/Jan/2006 MST"
    value = "02/Aug/2015 WIB"
    date, _ = time.Parse(layoutFormat, value)
    fmt.Println(value, "\t\t->", date.String())
    // 2015-09-02 00:00:00 +0700 WIB

    // Predefined Layout Format Untuk Keperluan Parsing Time

    data := "02 Sep 15 08:00 WIB"
    var dates, _ = time.Parse(time.RFC822, data)
    fmt.Println(dates.String())
    // 2015-09-02 08:00:00 +0700 WIB

    fmt.Println("===================================")
    //  Format dari time.Time ke string
    var datee, err = time.Parse(time.RFC822, "02 Sep 15 08:00 WIB")
    // handling error
    if err != nil {
        fmt.Println("error", err.Error())
        return
    }
    // 	02 Jan 06 15:04 MST
    fmt.Println("date : " , datee)

    var dateS1 = datee.Format("Monday 02, January 2006 15:04 MST")
    fmt.Println("dateS1", dateS1)
    // Wednesday 02, September 2015 08:00 WIB

    var dateS2 = datee.Format(time.RFC3339)
    fmt.Println("dateS2", dateS2)
    // 2015-09-02T08:00:00+07:00
}