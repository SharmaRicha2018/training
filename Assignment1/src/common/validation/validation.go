package validation

import (
	"common/utility"
	"sort"
	"strconv"
)

func IsDateValid(date1 string) bool {

	test, err := fmtdate.Parse("YYYY-MM-DD", date1)
	if err != nil {

		return false
	}
	date2 := fmtdate.Format("YYYY-MM-DD", test)

	if compare(date1, date2) != 0 {

		return false
	}

	return true

}

func compare(a, b string) int {

	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return +1
}

func StringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

func SameStringArrays(a, b []string) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	sort.Strings(a)
	sort.Strings(b)

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func GetCommonStringArrays(a, b []string) string {

	var commonString string

	for i := range a {

		for j := range b {
			if a[i] == b[j] {
				if commonString == "" {
					commonString = a[i]

				} else {
					commonString = commonString + "," + a[i]
				}
			}
		}
	}

	return commonString
}

func CreateStringFromStringArray(list []string) string {

	var str string

	for index, each := range list {

		if index == 0 {
			str = each
		} else {
			str = str + "," + each
		}

	}

	return str
}

func CreateStringArrayFromIntArray(list []int) []string {

	var str []string
	var item string

	for _, each := range list {

		item = strconv.Itoa(each)
		str = append(str, item)

	}

	return str
}

func IsItemArrayUnique(itemList []string) bool {
	var hashMap = make(map[string]struct{})

	for _, item := range itemList {
		if _, ok := hashMap[item]; ok {
			return false
		} else {
			hashMap[item] = struct{}{}
		}
	}
	return true
}
