package units

import (
	"bytes"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/text/gregex"
	"text/template"
)

//通过文本模板进行变量替换
func StringLiteralTemplate(str string, param interface{}) string {
	t, err := template.New("test").Parse(str)
	if err != nil {
		glog.Fatal("Parse string literal template error:", err)
	}
	buf := new(bytes.Buffer) //读写方法的可变大小的字节缓冲
	err = t.Execute(buf, param)
	if err != nil {
		glog.Fatal("Execute string literal template error:", err)
		return ""
	}
	return buf.String()
}

//通过正则表达式进行数据过滤,过虑掉指定的内容
func NormFormat(str, filter string) string {
	if filter != "" {
		tmpTxt, err := gregex.ReplaceString(filter, "", str)
		if err != nil {
			glog.Error(err)
		}
		return tmpTxt
	}
	return str
}
