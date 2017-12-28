package main

import "strings"

// borrowed from https://github.com/hashicorp/serf/blob/master/command/agent/flag_slice_value.go
// ↑ リンク切れ
// https://github.com/hashicorp/serf/blob/555e0dcbb180ecbd03431adc28226bb3192558bc/cmd/serf/command/agent/flag_slice_value.go

// AppendSliceValueはflag.Valueインターフェイスを実装し、同じ変数への複数の呼び出しでリストを追加できます。
type AppendSliceValue []string

func (s *AppendSliceValue) String() string {
	return strings.Join(*s, ",")
}

func (s *AppendSliceValue) Set(value string) error {
	if *s == nil {
		*s = make([]string, 0, 1)
	}

	*s = append(*s, value)
	return nil
}
