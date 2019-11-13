package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/umineko1996/enogu-archive-viewer/no6"
)

const (
	success = 0
	failed  = 1
)

func main() {
	os.Exit(run())
}

func run() int {

	email, password, isUpdate, isServer, err := getArgs()
	if err != nil {
		fmt.Println(err)
		return failed
	}

	config := no6.Config{
		Email:    email,
		Password: password,
	}

	if isServer {
		no6.Listen()
	}

	if isUpdate {
		if err := os.Remove(no6.ArchivesListFile); err != nil && !os.IsNotExist(err) {
			fmt.Printf("ファイルの削除に失敗しました。 err: %s\n", err.Error())
			return failed
		}
	}

	if _, err := os.Stat(no6.ArchivesListFile); err == nil {
		fmt.Printf("既に%sが作成されています。\n最新のデータを使用し再作成するには -u オプションを使用してください。\n", no6.ArchivesListFile)

		return failed
	}
	client, err := no6.NewClient(config)
	if err != nil {
		fmt.Printf("ログイン処理に失敗しました。 err: %s\n", err.Error())
		return failed
	}
	archivesInfo, err := client.GetArchivesInfoALL()
	if err != nil {
		fmt.Printf("アーカイブページの取得処理に失敗しました。 err: %s\n", err.Error())
		return failed
	}

	if err := no6.MakeArchivesList(archivesInfo); err != nil {
		fmt.Printf("csvファイルの作成に失敗しました。 err: %s\n", err.Error())
		return failed
	}

	return success
}

func getArgs() (email, password string, isUpdate, isServer bool, err error) {
	flag.StringVar(&email, "email", "", "ログインメールアドレス")
	flag.StringVar(&password, "pass", "", "ログインパスワード")
	flag.BoolVar(&isUpdate, "u", false, "ファイル更新フラグ")
	flag.BoolVar(&isServer, "http", false, "サーバー起動フラグ")
	flag.Parse()

	return
}
