// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package assets

import (
	"github.com/rakyll/statik/fs"
	"github.com/ysicing/ext/logger/zlog"
	"net/http"
)

func EmbedFS() http.FileSystem {
	efs, err := fs.NewWithNamespace(Gexe)
	if err != nil {
		zlog.Fatal("err: %v", err)
	}
	return efs
}
