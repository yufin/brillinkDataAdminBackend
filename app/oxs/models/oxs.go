/**
* @Author: Akiraka
* @Date: 2022/8/17 11:52
 */

package models

// ResponseErr 通用错误返回结构体
type ResponseErr struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
