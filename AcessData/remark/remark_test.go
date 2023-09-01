package remark

import (
	"MicroTikTok/AcessData/mysql"
	"fmt"
	"testing"
)

func TestGetRemarkListByVideoIdAndUserId(t *testing.T) {
	mysql.Init()
	fmt.Println(GetRemarkListByVideoIdAndUserId(1, 1))
}

func TestCreateRemark(t *testing.T) {
	mysql.Init()
	err, _ := CreateRemark(1, 1, "哈哈哈")
	if err != nil {
		return
	}
}

func TestDeleteRemark(t *testing.T) {
	mysql.Init()
	err := DeleteRemark(1)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(GetRemarkListByVideoIdAndUserId(1, 1))
}
