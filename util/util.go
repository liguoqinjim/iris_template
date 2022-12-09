package util

import "github.com/liguoqinjim/iris_template/consts"

func LikeParam(param string) string {
	return "%" + param + "%"
}

func LikeParamPre(param string) string {
	return "%" + param
}

func LikeParamPost(param string) string {
	return param + "%"
}

// 返回的是offset,limit，是  20,10的格式     不是20,30
func GetPageQueryParams(pageNum, pageSize int) (int, int) {
	if pageNum <= 0 {
		return 0, consts.PageDefault
	} else {
		if pageSize == 0 {
			pageSize = consts.PageSizeDefault
		}

		return (pageNum - 1) * pageSize, pageSize
	}
}
