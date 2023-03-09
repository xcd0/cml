package main

import (
	"flag"
	"log"
	"os"

	"github.com/xcd0/cml"
)

// 名前は後で変える。
// .git/hooks/commit-msgから `cml $1` のように呼ばれることを想定する。
// 第一引数で与えられた文字列をコミットメッセージがかかれたテキストファイルのパスだと見做して読み込み、
// コミットメッセージの書式をチェックし、チェック結果を標準出力に出力して、
// 書式に問題がなければ正常終了、問題があれば異常終了する。

// linterとしての設定は特定のパスに配置した設定ファイルを使用する。
// 特定のパスは環境変数から変更できる。

// main関数は使いまわし。
func main() {
	log.SetFlags(log.Ltime | log.Lshortfile)
	//log.SetFlags(0)
	err := cml.Run(os.Args[1:], os.Stdout, os.Stderr)
	if err != nil && err != flag.ErrHelp {
		log.Println(err)
		exitCode := 1
		if ecoder, ok := err.(interface{ ExitCode() int }); ok {
			exitCode = ecoder.ExitCode()
		}
		os.Exit(exitCode)
	}
}
