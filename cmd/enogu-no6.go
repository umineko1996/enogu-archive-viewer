package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/umineko1996/enogu-archive-viewer/no6"
)

var (
	email    string
	password string
	isUpdate bool
)

func parseFlag() {
	flag.StringVar(&email, "email", "", "ログインメールアドレス")
	flag.StringVar(&password, "pass", "", "ログインパスワード")
	flag.BoolVar(&isUpdate, "u", false, "ファイル更新フラグ")
	flag.Parse()
}

func main() {
	parseFlag()

	config := no6.Config{
		Email:    email,
		Password: password,
	}
	if isUpdate {
		if err := os.Remove(no6.ArchivesListFile); err != nil &&
			err.Error() != fmt.Sprintf("remove %s: The system cannot find the file specified.", no6.ArchivesListFile) {

			fmt.Printf("ファイルの削除に失敗しました。 err: %s\n", err.Error())
			os.Exit(1)
		}
	}

	if _, err := os.Stat(no6.ArchivesListFile); err == nil {
		fmt.Printf("既に%sが作成されています。\n最新のデータを使用し再作成するには -u オプションを使用してください。\n", no6.ArchivesListFile)
		os.Exit(1)
	}
	client, err := no6.NewClient(config)
	if err != nil {
		fmt.Printf("ログイン処理に失敗しました。 err: %s\n", err.Error())
		os.Exit(1)
	}
	if err := client.GetALLArchivesPage(); err != nil {
		fmt.Printf("アーカイブページの取得処理に失敗しました。 err: %s\n", err.Error())
		os.Exit(1)
	}

	if err := no6.MakeArchivesList(); err != nil {
		fmt.Printf("csvファイルの作成に失敗しました。 err: %s\n", err.Error())
		os.Exit(1)
	}
}
