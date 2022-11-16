package selfpkg

import "strings"

func Trans2cnForSignUp(en string) string {
	zn := ""
	if strings.Contains(en, "required") {
		zn += "请求参数输入不能为空 "
	}
	if strings.Contains(en, "eqfield") {
		zn += "密码和确认密码不一致"
	}

	return zn
}
