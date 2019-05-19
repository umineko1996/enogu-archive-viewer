# enogu-archive-viewer
開発中です

## 実装済機能
- 公式HP上の動画アーカイブページをすべてDLする
- DLしたアーカイブページからアーカイブのタイトル、日付、画像URL、youtube URLを抽出しcsvファイルに出力する

## 使い方
- 初回  
`enogu-no6.exe -email xxxxx -pass xxxxx`  
- 二回目以降  
`enogu-no6.exe -email xxxxx -pass xxxxx -u`  
もしくはarchives_list.csvを削除してから実行



## 実装予定機能
- サーバ起動することでブラウザにてアーカイブの一覧表示、検索を行えるようにする
