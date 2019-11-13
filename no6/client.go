package no6

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Client 構造体は公式HPにアクセスするためのクライアントです
type Client struct {
	client *http.Client
}

// Config 構造体はClient作成用のConfigです
type Config struct {
	Email      string
	Password   string
	HTTPClient *http.Client
}

// 公式HPのURL情報
const (
	protocol     = "https://"
	host         = "enogu-no6.com"
	login        = "/session"
	loginPage    = "/session/new"
	archivesPage = "/archives"
)

// NewClient 関数はえのぐ公式サイトno6に接続するクライアントを作成しログイン処理を実行します
// ログイン情報は引数のconfig構造体に設定してください
func NewClient(config Config) (*Client, error) {
	httpClient := config.HTTPClient
	if httpClient == nil {
		jar, err := cookiejar.New(nil)
		if err != nil {
			return nil, err
		}
		httpClient = &http.Client{Jar: jar}
	}
	client := &Client{
		client: httpClient,
	}
	if err := client.login(config); err != nil {
		return nil, err
	}

	return client, nil
}

func (c *Client) login(config Config) error {
	if config.Email == "" || config.Password == "" {
		return errors.New("email or password is nil")
	}
	param, token, err := c.getCsrf()
	if err != nil {
		return err
	}

	loginURL := protocol + host + login
	values := url.Values{}
	values.Set("utf8", "✓")
	values.Set(param, token)
	values.Add("session[email]", config.Email)
	values.Add("session[password]", config.Password)
	values.Add("commit", "ログイン")
	body := strings.NewReader(values.Encode())

	req, err := http.NewRequest(http.MethodPost, loginURL, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.client.Do(req)
	if err != nil {
		fmt.Println(resp.StatusCode)
		return err
	}
	defer resp.Body.Close()

	if resp.Request.URL.Path != "/home" {
		return errors.New("ログインセッションに失敗しました")
	}

	return nil
}

// GetArchiveseALL 関数は公式HP上のすべてのアーカイブページをローカルに保存します
func (c *Client) GetArchivesInfoALL() (archivesInfo []*archiveInfo, err error) {

	n, err := c.getArchivesLastPage()
	if err != nil {
		return nil, err
	}

	for i := 1; i < n+1; i++ {
		time.Sleep(100 * time.Millisecond)
		arcsInfo, err := c.getArchivesInfo(i)
		if err != nil {
			return nil, err
		}
		archivesInfo = append(archivesInfo, arcsInfo...)
	}

	return archivesInfo, nil
}

func (c *Client) getArchivesLastPage() (int, error) {
	url := protocol + host + archivesPage
	req, err := http.NewRequest(http.MethodGet, url, http.NoBody)
	if err != nil {
		return 0, err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return 0, err
	}
	defer func() {
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}()

	n, err := extractLastPage(resp.Body)
	if err != nil {
		return 0, err
	}
	return n, nil
}

func (c *Client) getArchivesInfo(pageNumber int) ([]*archiveInfo, error) {
	resp, err := c.getArchivesPage(pageNumber)
	if err != nil {
		return nil, err
	}
	defer func() {
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}()

	archivesInfo, err := extractArchivesInfo(resp.Body)
	if err != nil {
		return nil, err
	}

	return archivesInfo, nil
}

func (c *Client) getArchivesPage(n int) (*http.Response, error) {
	url := protocol + host + archivesPage
	req, err := http.NewRequest(http.MethodGet, url, http.NoBody)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = fmt.Sprintf("page=%d", n)
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func extractLastPage(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "page last") {
			if !scanner.Scan() {
				break
			}
			return extractPage(scanner.Text()), nil
		}
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return 0, errors.New("response body has not page last")
}

func (c *Client) getCsrf() (param, token string, err error) {
	url := protocol + host + loginPage
	req, err := http.NewRequest(http.MethodGet, url, http.NoBody)
	if err != nil {
		return "", "", err
	}

	resp, err := c.client.Do(req)
	switch {
	case err != nil:
		return "", "", err
	case resp.Body == nil:
		return "", "", errors.New("response body is nil")
	}
	defer resp.Body.Close()

	return extractCsrf(resp.Body)
}

func extractCsrf(r io.Reader) (param, token string, err error) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case strings.Contains(line, "csrf-param"):
			param = extractContent(line)
		case strings.Contains(line, "csrf-token"):
			token = extractContent(line)
		case param != "" && token != "":
			return param, token, nil
		}
	}
	if err := scanner.Err(); err != nil {
		return "", "", err
	}
	if param == "" || token == "" {
		return "", "", errors.New("response body has not csrf")
	}
	return param, token, nil
}

var (
	contentRegexp = regexp.MustCompile(`content=".+?"`)
	numberRegexp  = regexp.MustCompile(`\d+`)
)

func extractContent(s string) string {
	r1 := contentRegexp.Find([]byte(s))
	r2 := strings.Split(string(r1), string('"'))
	return r2[1]
}

func extractPage(s string) int {
	r1 := numberRegexp.Find([]byte(s))
	n, err := strconv.Atoi(string(r1))
	if err != nil {
		log.Panic(err)
	}
	return n
}
