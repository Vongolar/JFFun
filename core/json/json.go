/*
 * @Author: Vongola
 * @LastEditTime: 2021-02-19 14:48:33
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: \JFFun\core\json\json.go
 * @Date: 2021-02-19 14:46:41
 * @描述: 文件描述
 */
package json

import (
	"encoding/json"
	"io"
)

func GetExt() []string {
	return []string{`json`}
}

func Encode(w io.Writer, v interface{}) error {
	return json.NewEncoder(w).Encode(v)
}

func Decode(r io.Reader, out interface{}) error {
	return json.NewDecoder(r).Decode(out)
}
