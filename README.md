# enogu-archive-viewer
開発中です

## 使い方
- 初回  
`enogu-no6.exe -email xxxxx -pass xxxxx`  
実行したフォルダに 
archives_list.csv が作成されます
- 二回目以降  
`enogu-no6.exe -email xxxxx -pass xxxxx -u`  
もしくは archives_list.csv を削除してから  
`enogu-no6.exe -email xxxxx -pass xxxxx`  
archives_list.csv が再作成されます

## 実装済機能
- 公式HP上の動画アーカイブページをすべてDLする
- DLしたアーカイブページからアーカイブのタイトル、日付、画像URL、youtube URLを抽出しcsvファイルに出力する

## 実装予定機能
- サーバ起動することでブラウザにてアーカイブの一覧表示、検索を行えるようにする
