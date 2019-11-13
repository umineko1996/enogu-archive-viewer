package no6

import (
	"bufio"
	"errors"
	"html"
	"io"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/gocarina/gocsv"
)

var (
	// ArchivesListFile 変数はアーカイブ情報をまとめたcsvファイルです
	ArchivesListFile = "archives_list.csv"
	utf8BOM          = []byte{0xEF, 0xBB, 0xBF}
)

type archiveInfo struct {
	Title         string `csv:"タイトル"`
	PublishedDate string `csv:"公開日"`
	URL           string `csv:"youtube"`
	Thumbnail     string `csv:"サムネ"`
	ID            string `csv:"id"`
}

func (a *archiveInfo) checkField() error {
	if a == nil {
		return errors.New("archiveInfo: receiver is nil")
	}
	switch {
	case a.Title == "":
		return errors.New("archiveInfo: title field is space")
	case a.PublishedDate == "":
		return errors.New("archiveInfo: publishedDate field is space")
	case a.URL == "":
		return errors.New("archiveInfo: url field is space")
	case a.Thumbnail == "":
		return errors.New("archiveInfo: thumbnail field is space")
	}
	return nil
}

// MakeArchivesList 関数はlocalにあるアーカイブページのデータからアーカイブの情報を抽出したcsvファイルを生成します
// 抽出が終わったファイルは削除されます
func MakeArchivesList(archivesInfo []*archiveInfo) error {
	f, err := os.Create(ArchivesListFile)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.Write(utf8BOM); err != nil {
		return err
	}

	if err := gocsv.MarshalFile(&archivesInfo, f); err != nil {
		return err
	}

	return nil
}

func extractArchivesInfo(r io.Reader) ([]*archiveInfo, error) {
	arcItems, err := extractArchiveItems(r)
	if err != nil {
		return nil, err
	}

	var archivesInfo []*archiveInfo
	for _, arcItem := range arcItems {
		arcInf, err := extractArchiveInfo(arcItem)
		if err != nil {
			return nil, err
		}
		archivesInfo = append(archivesInfo, arcInf)
	}

	return archivesInfo, nil
}

func extractArchiveInfo(archiveItem string) (*archiveInfo, error) {
	archive := new(archiveInfo)

	lines := strings.Split(archiveItem, "\n")
	for i, line := range lines {
		switch {
		case strings.Contains(line, `<a target="_blank"`):
			archive.URL = extractArchiveURL(line)
		case strings.Contains(line, `<img src=`):
			archive.Thumbnail = extractImgURL(line)
		case strings.Contains(line, `archive-item_title`):
			if len(lines)-1 == i {
				break
			}
			archive.Title = html.UnescapeString(lines[i+1])
		case strings.Contains(line, `archive-item_published-date`):
			if len(lines)-1 == i {
				break
			}
			archive.PublishedDate = lines[i+1]
		}
	}

	if err := archive.checkField(); err != nil {
		return nil, err
	}

	archive.ID = path.Base(archive.URL)

	return archive, nil
}

var (
	archiveURLRegexp = regexp.MustCompile(`href="https:.*"`)
	imgURLRegexp     = regexp.MustCompile(`https:\/\/i\.ytimg\.com\/.*jpg`)
)

func extractArchiveItems(r io.Reader) ([]string, error) {
	var archiveItems []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, `<li class='archive-items_item'>`) {
			archiveItem := line
			for scanner.Scan() {
				line = scanner.Text()
				archiveItem += "\n" + line
				if strings.Contains(line, `</li>`) {
					break
				}
			}
			archiveItems = append(archiveItems, archiveItem)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if len(archiveItems) == 0 {
		return nil, errors.New("can not find archive items")
	}

	return archiveItems, nil
}

func extractArchiveURL(s string) string {
	r1 := archiveURLRegexp.Find([]byte(s))
	r2 := strings.Split(string(r1), string('"'))
	return r2[1]
}

func extractImgURL(s string) string {
	r := imgURLRegexp.Find([]byte(s))
	return string(r)
}
