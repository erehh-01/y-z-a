package utils

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

var (
	CCNumRegex   = regexp.MustCompile(`\b\d{16}\b`)
	CCCvvRegex   = regexp.MustCompile(`\b\d{3}\b`)
	CCMonthRegex = regexp.MustCompile(`\b(0[1-9]|1[0-2])\b`)
	CCYearRegex  = regexp.MustCompile(`\b(20\d{2}|2[5-9]\d|2[4-9])\b`)
)

func ParseCC(text string) (CC, error) {

	yearF := CCYearRegex.FindAllString(text, -1)
	monthF := CCMonthRegex.FindAllString(text, -1)
	numF := CCNumRegex.FindAllString(text, -1)
	cvvF := CCCvvRegex.FindAllString(text, -1)

	//fmt.Println("Number:", numF)
	//fmt.Println("CVV:", cvvF)
	//fmt.Println("Year:", yearF)
	//fmt.Println("Month:", monthF)

	var err error
	cc := CC{}
	if len(numF) <= 0 {
		return cc, errors.New("no cc number found in text")
	}

	cc.CCNUM, err = strconv.Atoi(numF[0])
	if err != nil {
		return cc, errors.New("no cc number found")
	}

	if len(cvvF) <= 0 {
		return CC{}, errors.New("no cc cvv found in text")
	}
	cc.CVV = cvvF[0]

	if len(monthF) <= 0 {
		return CC{}, errors.New("no cc month found in text")
	}
	m, err := strconv.Atoi(monthF[0])
	if err != nil {
		return cc, errors.New("no cc month found")
	}
	cc.MONTH = uint(m)

	if len(yearF) <= 0 {
		return CC{}, errors.New("no cc year found in text")
	}

	if len(yearF[0]) == 2 {
		yearF[0] = fmt.Sprint("20" + yearF[0])
	}

	y, err := strconv.Atoi(yearF[0])
	if err != nil {
		return cc, errors.New("no cc year found")
	}
	cc.YEAR = uint(y)

	return cc, nil
}
