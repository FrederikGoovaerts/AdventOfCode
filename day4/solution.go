package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func (p passport) hasFields() bool {
	return p.byr != "" &&
		p.iyr != "" &&
		p.eyr != "" &&
		p.hgt != "" &&
		p.hcl != "" &&
		p.ecl != "" &&
		p.pid != ""
}

func (p passport) isValid() bool {
	byrVal, _ := strconv.Atoi(p.byr)
	byrValid := 1920 <= byrVal && byrVal <= 2002

	iyrVal, _ := strconv.Atoi(p.iyr)
	iyrValid := 2010 <= iyrVal && iyrVal <= 2020

	eyrVal, _ := strconv.Atoi(p.eyr)
	eyrValid := 2020 <= eyrVal && eyrVal <= 2030

	hgtUnit := p.hgt[len(p.hgt)-2:]
	var hgtUnitValid bool
	var hgtValValid bool
	if hgtUnit == "cm" {
		hgtUnitValid = true
		cms, err := strconv.Atoi(p.hgt[:len(p.hgt)-2])
		if err == nil {
			hgtValValid = cms >= 150 && cms <= 193
		}
	} else if hgtUnit == "in" {
		hgtUnitValid = true
		ins, err := strconv.Atoi(p.hgt[:len(p.hgt)-2])
		if err == nil {
			hgtValValid = ins >= 59 && ins <= 76
		}
	}

	hgtValid := hgtUnitValid && hgtValValid

	hclMatch, _ := regexp.MatchString("#[0-9a-f]{6}", p.hcl)
	hclValid := len(p.hcl) == 7 && hclMatch

	eclMatch, _ := regexp.MatchString("amb|blu|brn|gry|grn|hzl|oth", p.ecl)
	eclValid := len(p.ecl) == 3 && eclMatch

	pidMatch, _ := regexp.MatchString("[0-9]{9}", p.pid)
	pidValid := len(p.pid) == 9 && pidMatch

	return byrValid && iyrValid && eyrValid && hgtValid && hclValid && eclValid && pidValid
}

func main() {
	dat, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	groups := strings.Split(strings.TrimSpace(string(dat)), "\n\n")
	passports := make([]passport, 0)
	for _, group := range groups {
		fields := strings.Fields(group)
		fieldMap := make(map[string]string, 0)
		for _, field := range fields {
			parts := strings.Split(field, ":")
			fieldMap[parts[0]] = parts[1]
		}
		passports = append(passports, passport{
			byr: fieldMap["byr"],
			iyr: fieldMap["iyr"],
			eyr: fieldMap["eyr"],
			hgt: fieldMap["hgt"],
			hcl: fieldMap["hcl"],
			ecl: fieldMap["ecl"],
			pid: fieldMap["pid"],
			cid: fieldMap["cid"],
		})
	}

	completePassports := 0
	validPassports := 0
	for _, passport := range passports {
		if passport.hasFields() {
			completePassports++
			if passport.isValid() {
				validPassports++
			}
		}
	}
	fmt.Println(completePassports)
	fmt.Println(validPassports)
}
