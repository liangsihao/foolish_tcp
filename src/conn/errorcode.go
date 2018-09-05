package conn

import "errors"

const (
	Marshall_Success        = 0
	Marshall_Flag_fail      = 10000
	Marshall_Len_Less_fail  = 10001
	Marshall_Len_Zero_fail  = 10002
	Marshall_RLen_Zero_fail = 10003
	Marshall_Uid_Zero_fail  = 10004
	Marshall_Sid_Zero_fail  = 10005
	Marshall_Code_Zero_fail = 10006
	Marshall_err_fail       = 10007
	Buff_Len_Zero       = 10007
)

var (
	Marshall_Flag_Err = errors.New("is not standrd flag")
	Marshall_Len_Zero_Err = errors.New("len can not be zero")
	Marshall_Len_Less_Err = errors.New("len is not enough")
	Buff_Len_Zero_Err = errors.New("len is zero")
)
