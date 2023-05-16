package str

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
	"text/template"
)

// SplitSpace 以空格分割字符串-正则方式
func SplitSpace(s string) []string {
	compile, _ := regexp.Compile(`\s+`)
	return compile.Split(s, -1)
}

// Join 组合字符串
func Join(s []string) string {
	return strings.Join(s, "")
}

// Sf 格式化拼接字符串
func Sf(format string, val ...interface{}) string {
	return fmt.Sprintf(format, val...)
}

// Format 格式化字符串
func Format(tpl string, data interface{}) string {
	var buf bytes.Buffer
	tmpl, _ := template.New("test").Parse(tpl)
	_ = tmpl.Execute(&buf, data)
	return buf.String()
}
