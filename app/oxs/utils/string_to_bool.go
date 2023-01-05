/**
* @Author: Akiraka
* @Date: 2022/10/17 11:19
 */

package utils

import (
	"strconv"
)

// StringToBool 字符串转bool
func StringToBool(value string) bool {
	res, _ := strconv.ParseBool(value)
	return res
}
