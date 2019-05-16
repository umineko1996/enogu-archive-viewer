package no6

import (
	"bufio"
	"net/http"
	"strings"
	"testing"
)

func TestExtractCsrf(t *testing.T) {
	t.Run("TestExtractCsrf_OK", func(t *testing.T) {
		// 準備
		res, err := http.ReadResponse(
			bufio.NewReader(
				strings.NewReader(newPageResponse)),
			nil)
		if err != nil {
			t.Fatalf("extractCsrf関数の検証に失敗しました。 err: %s", err.Error())
		}
		// 実行
		param, token, err := extractCsrf(res.Body)

		// 検証
		if err != nil {
			t.Errorf("extractCsrf関数の検証に失敗しました。 err: %s", err.Error())
		}
		if got, want := param, "authenticity_token"; got != want {
			t.Errorf("extractCsrf関数の検証に失敗しました。 paramの値が不正です。 got %s, want %s", got, want)
		}
		if got, want := token, "G0mObqxcxuShfhX0QW4iemKf2k9X3K6lOqh1ntcEFNCy/dkn8eOYXblQekRukR1GZ+iHAcFtrcOs1aogax3cPw=="; got != want {
			t.Errorf("extractCsrf関数の検証に失敗しました。 tokenの値が不正です。 got %s, want %s", got, want)
		}
	})

	t.Run("TestExtractCsrf_NG_token無し", func(t *testing.T) {
		// 準備
		source := strings.NewReader(`<meta name="csrf-param" content="authenticity_token" />`)

		// 実行
		param, token, err := extractCsrf(source)

		// 検証
		if got, want := err.Error(), "response body has not csrf"; got != want {
			t.Errorf("extractCsrf関数の検証に失敗しました。 errの値が不正です。 got %s, want %s", got, want)
		}
		if got, want := param, ""; got != want {
			t.Errorf("extractCsrf関数の検証に失敗しました。 paramの値が不正です。 got %s, want %s", got, want)
		}
		if got, want := token, ""; got != want {
			t.Errorf("extractCsrf関数の検証に失敗しました。 tokenの値が不正です。 got %s, want %s", got, want)
		}
	})

	t.Run("TestExtractCsrf_NG_param無し", func(t *testing.T) {
		// 準備
		source := strings.NewReader(`<meta name="csrf-token" content="HWmEZ6Fk1OnVC+/36Yjnvs+mc5WU5Tk1O+82EfNuX4O03dMu/NuKUM0lgEfGd9iCytEu2wJUOlOtkumvT3eXbA==" />`)

		// 実行
		param, token, err := extractCsrf(source)

		// 検証
		if got, want := err.Error(), "response body has not csrf"; got != want {
			t.Errorf("extractCsrf関数の検証に失敗しました。 errの値が不正です。 got %s, want %s", got, want)
		}
		if got, want := param, ""; got != want {
			t.Errorf("extractCsrf関数の検証に失敗しました。 paramの値が不正です。 got %s, want %s", got, want)
		}
		if got, want := token, ""; got != want {
			t.Errorf("extractCsrf関数の検証に失敗しました。 tokenの値が不正です。 got %s, want %s", got, want)
		}
	})

	t.Run("TestExtractCsrf_NG_param無し_token無し", func(t *testing.T) {
		// 準備
		source := strings.NewReader(`invalid source`)

		// 実行
		param, token, err := extractCsrf(source)

		// 検証
		if got, want := err.Error(), "response body has not csrf"; got != want {
			t.Errorf("extractCsrf関数の検証に失敗しました。 errの値が不正です。 got %s, want %s", got, want)
		}
		if got, want := param, ""; got != want {
			t.Errorf("extractCsrf関数の検証に失敗しました。 paramの値が不正です。 got %s, want %s", got, want)
		}
		if got, want := token, ""; got != want {
			t.Errorf("extractCsrf関数の検証に失敗しました。 tokenの値が不正です。 got %s, want %s", got, want)
		}
	})
}

func TestExtractLastPage(t *testing.T) {
	t.Run("extractLastPage_OK", func(t *testing.T) {
		// 準備
		res, err := http.ReadResponse(
			bufio.NewReader(
				strings.NewReader(archivesPageRespons)),
			nil)
		if err != nil {
			t.Fatalf("extractLastPage関数の検証に失敗しました。 err: %s", err.Error())
		}

		// 実行
		page, err := extractLastPage(res.Body)

		// 検証
		if err != nil {
			t.Errorf("extractLastPage関数の検証に失敗しました。 err: %s", err.Error())
		}
		if got, want := page, 44; got != want {
			t.Errorf("extractLastPage関数の検証に失敗しました。 pageの値が不正です。 got %d, want %d", got, want)
		}
	})
}

var (
	newPageResponse = `HTTP/2.0 200 OK
date: Tue, 14 May 2019 13:26:48 GMT
content-type: text/html; charset=utf-8
vary: Accept-Encoding
vary: Origin
x-frame-options: SAMEORIGIN
x-xss-protection: 1; mode=block
x-content-type-options: nosniff
x-download-options: noopen
x-permitted-cross-domain-policies: none
referrer-policy: strict-origin-when-cross-origin
etag: W/"06035128e2e63e89f45606d00651a127"
cache-control: max-age=0, private, must-revalidate
set-cookie: _enogu_fanclub_session=ZZE2K5f1CErNApOhm5u3zDS1S3qjXDhhU3HV%2FCkRphIvvlLVsKmg5TSPIfdZ%2FfYIvpQQUsTwwbUyk4iefj7v3GjYP9xCRs1a1E4rmTRmQSe7yHhFmNAa270pxtyMHgnZoSxeW0yhefPj35crn14%3D--ZM7cvX5u2e6GUg4t--owSVr3mGGyDqqfqWn0B7%2FQ%3D%3D; path=/; secure; HttpOnly
x-request-id: fd09e280-7e3b-49fa-a665-e5a1d7f8ed36
x-runtime: 0.008838
x-cloud-trace-context: dda9eba439e8e87f4c730285fa55948b/6096731631813269834;o=1
strict-transport-security: max-age=31536000; includeSubDomains
content-encoding: gzip
via: 1.1 google
X-Firefox-Spdy: h2

<!DOCTYPE html>
<html>
<head>
<meta content='text/html; charset=UTF-8' http-equiv='Content-Type'>
<meta charset='utf-8'>
<meta content='width=device-width, initial-scale=1, shrink-to-fit=no, viewport-fit=cover' name='viewport'>
<title>えのぐ公式ファンクラブ -No,6-</title>
<meta name="csrf-param" content="authenticity_token" />
<meta name="csrf-token" content="G0mObqxcxuShfhX0QW4iemKf2k9X3K6lOqh1ntcEFNCy/dkn8eOYXblQekRukR1GZ+iHAcFtrcOs1aogax3cPw==" />

<link rel="stylesheet" media="all" href="/assets/application-4b46eb12d8c7da1e1586ea0b7d7a6c9a0d0c17fdfdb729ce7897a86aaf985da4.css" data-turbolinks-track="reload" />
<script src="/assets/application-6e438b307bf4ec2b96355941f40b2eecec0f12239862748a7017b5aaab2295a4.js" data-turbolinks-track="reload"></script>
<script async src='https://www.googletagmanager.com/gtag/js?id=UA-104370994-2'></script>
<script>
  window.dataLayer = window.dataLayer || [];
  function gtag(){dataLayer.push(arguments);}
  gtag('js', new Date());

  gtag('config', 'UA-104370994-2');
</script>

</head>
<body class='sessions is-editable' lang='ja'>
<div class='container'>
<h1 class='title-logo'>
<a href="/">えのぐ公式ファンクラブ -No,6-</a>
</h1>
<div class='bordered-box'>
<h2>ログイン</h2>
<form class="mb-3" action="/session" accept-charset="UTF-8" method="post"><input name="utf8" type="hidden" value="&#x2713;" /><input type="hidden" name="authenticity_token" value="dWU2YlOkIeOjimXxpTyOLZdY3ozthfKWUlpN3OhU6dvS9Y95Ld36Nycokl+E7vS7wMRJYoUoaIpAtmvbsbbI4g==" />
<div class='form-group'>
<label for="session_email">メールアドレス</label>
<input class="form-control" required="required" placeholder="mail@example.com" type="email" name="session[email]" id="session_email" />
</div>
<div class='form-group'>
<label for="session_password">パスワード</label>
<input class="form-control" required="required" placeholder="password" type="password" name="session[password]" id="session_password" />
</div>
<input type="submit" name="commit" value="ログイン" data-disable-with="ログイン" />
</form><div class='text-center'>
<small>
<a href="/welcome">アカウントをお持ちでない方</a>
</small>
</div>
<div class='text-center'>
<small>
<a href="/password_resets/new">パスワードをお忘れの方</a>
</small>
</div>
</div>

</div>
</body>
</html>`

	sessionResponse = `HTTP/2.0 302 Found
date: Tue, 14 May 2019 13:46:22 GMT
content-type: text/html; charset=utf-8
x-frame-options: SAMEORIGIN
x-xss-protection: 1; mode=block
x-content-type-options: nosniff
x-download-options: noopen
x-permitted-cross-domain-policies: none
referrer-policy: strict-origin-when-cross-origin
location: https://enogu-no6.com/home
cache-control: no-cache
set-cookie: _enogu_fanclub_session=l7lqgrnlQe0OxVc%2B2KFFnqKkZ1IOw5jFUBWggYR38EkIuZLsm3kn264Is1JZFlCSm4JYAc7WvSusYeu9jvBoyaCo6suK24jIsqd1zOpKAZQTr7xQYZzn5wmtVhUhfgSS8OSFhtyVYKYN4GZ12%2B1y4SMFEAvM5xTT%2F0hddg%3D%3D--XPIihn%2F8J%2FhxlsCT--VOa8MZzxJ7oq5%2FY3OB%2BiNw%3D%3D; path=/; secure; HttpOnly
x-request-id: 72fcd462-1621-4b1d-b431-a483ce4840ec
x-runtime: 0.081221
x-cloud-trace-context: 69dcb9b8478a2547efedbbf3a6a13c97/9685412063886635249;o=1
strict-transport-security: max-age=31536000; includeSubDomains
vary: Origin
via: 1.1 google
X-Firefox-Spdy: h2`

	homePageResponse = `HTTP/2.0 200 OK
date: Tue, 14 May 2019 13:46:22 GMT
content-type: text/html; charset=utf-8
vary: Accept-Encoding
vary: Origin
x-frame-options: SAMEORIGIN
x-xss-protection: 1; mode=block
x-content-type-options: nosniff
x-download-options: noopen
x-permitted-cross-domain-policies: none
referrer-policy: strict-origin-when-cross-origin
etag: W/"46062fad83a65a58d2cc89f09e16579d"
cache-control: max-age=0, private, must-revalidate
set-cookie: _enogu_fanclub_session=KbkQ0mWzUlNYHreuCoogEO%2BWFBK%2BJn3QAqQXdhb%2FAhNhvdcD31VWGkQFW%2FmT7VKL8lQG1n%2FDVOK9IuZbNjFXBAmciVjfGKdrcXJCqk1NyfYTypR0JokO%2BKZKimzUjES7VqsTUSIwpOZFce7cX3JeJLSd7heQgX0qoDmIVQ%3D%3D--EWMgcyarlueYjvC8--DnPecCWRwr%2B2JlBzNR%2BxTg%3D%3D; path=/; secure; HttpOnly
x-request-id: fe5679d6-9e26-41a9-b1de-b9540dbb435d
x-runtime: 0.086584
x-cloud-trace-context: 4dda9cfaa573b5f87d546166afec90b4/5421197360250430304;o=1
strict-transport-security: max-age=31536000; includeSubDomains
content-encoding: gzip
via: 1.1 google
X-Firefox-Spdy: h2

<html>
<head>
<meta content='text/html; charset=UTF-8' http-equiv='Content-Type'>
<meta charset='utf-8'>
<meta content='width=device-width, initial-scale=1, shrink-to-fit=no, viewport-fit=cover' name='viewport'>
<title>えのぐ公式ファンクラブ -No,6-</title>
<meta name="csrf-param" content="authenticity_token" />
<meta name="csrf-token" content="j6LvmrMTA3XZ+gkyws+NujLuPAlrJwIdjB6OMcu9HnsmFrjT7qxdzMHUZoLtMLKGN5lhR/2WAXsaY1GPd6TWlA==" />

<link rel="stylesheet" media="all" href="/assets/application-4b46eb12d8c7da1e1586ea0b7d7a6c9a0d0c17fdfdb729ce7897a86aaf985da4.css" data-turbolinks-track="reload" />
<script src="/assets/application-6e438b307bf4ec2b96355941f40b2eecec0f12239862748a7017b5aaab2295a4.js" data-turbolinks-track="reload"></script>
<script async src='https://www.googletagmanager.com/gtag/js?id=UA-104370994-2'></script>
<script>
  window.dataLayer = window.dataLayer || [];
  function gtag(){dataLayer.push(arguments);}
  gtag('js', new Date());

  gtag('config', 'UA-104370994-2');
</script>

</head>
<body class='pages lemon' lang='ja'>
<div class='container'>
<div class='navigation-bar'>
<button class='drawer-open'>
<img src="/assets/components/drawer/open_button-2c203f8d3250f4e1850d8ff11cd92a288ec462e25ba6d50dd83e2c2148d3c488.svg" />
</button>
<div class='title-logo'>
<a href="/">えのぐ公式ファンクラブ -No,6-</a>
</div>
</div>
<div class='drawer'>
<button class='drawer-close'>
<img src="/assets/components/drawer/close_button-09ecdc04ecc788c8fbf57358d718661ad2fe76856e203297353241f02fbd4e64.svg" />
</button>
<div class='title-logo'>
<a href="/home">えのぐ公式ファンクラブ -No,6-</a>
</div>
<div class='list-group'>
<a class="list-group-item list-group-item-action d-flex align-items-center group1" href="/home"><span class='icon'>
<img src="/assets/components/drawer/ic_drawer_home-1531f8bbb1445531a0398b5b46c98bdb3e2383b45cddf21cbd9e6be5eb540fd5.svg" />
</span>
<span class='title flex-grow-1 pl-2'>
ホーム
</span>
</a><a class="list-group-item list-group-item-action d-flex align-items-center group1" href="/news"><span class='icon'>
<img src="/assets/components/drawer/ic_drawer_news-319d1e4d58f495cbafe3ae01d1549791c4f1d9eb9c5a0a8febec14d531730bc3.svg" />
</span>
<span class='title flex-grow-1 pl-2'>
ニュース
</span>
</a><a class="list-group-item list-group-item-action d-flex align-items-center group1" href="/blogs"><span class='icon'>
<img src="/assets/components/drawer/ic_drawer_blog-4fd91d1e42959fe2537c0ac88753def63e1e147d93990da56971a0215e3efea0.svg" />
</span>
<span class='title flex-grow-1 pl-2'>
ブログ
</span>
</a><a class="list-group-item list-group-item-action d-flex align-items-center group1" href="/gallery/movies"><span class='icon'>
<img src="/assets/components/drawer/ic_drawer_gallery-697063773e7f8149ed4796c65891345a651cabcd82826b3a18ec7d41c23a9ec6.svg" />
</span>
<span class='title flex-grow-1 pl-2'>
ギャラリー
</span>
</a><a class="list-group-item list-group-item-action d-flex align-items-center group1" href="/archives"><span class='icon'>
<img src="/assets/components/drawer/ic_drawer_movie-f7c8e7479c206b3497fb29e7fba9c9706837f70aff841b29c221d14a59b4398f.svg" />
</span>
<span class='title flex-grow-1 pl-2'>
動画アーカイブ
</span>
</a><a class="list-group-item d-flex align-items-center group1 coming-soon" href="/goods"><span class='icon'>
<img src="/assets/components/drawer/ic_drawer_goods-e4880bb7fa2eac0251e342e310c0ecf98fe9cb45ed16a3d3dcbac8f305236a00.svg" />
</span>
<span class='title flex-grow-1 pl-2'>
グッズ
</span>
<span class='coming-soon-badge'>
<img src="/assets/components/drawer/ic_comingsoon_normal-f5de52da61de684a9729e2ea34127c03a06a2fe7253ac574f9f1679989e09d10.svg" />
</span>
</a><a class="list-group-item list-group-item-action d-flex align-items-center group2" href="/letters/new"><span class='icon'>
<img src="/assets/components/drawer/ic_drawer_hutsuota-ad3f9d714c924d36954995f691f30e99f1f4fcdd24122d6535bd89f2cc563100.svg" />
</span>
<span class='title flex-grow-1 pl-2'>
ふつおた
</span>
</a><a class="list-group-item list-group-item-action d-flex align-items-center group2" href="/sponsors/new"><span class='icon'>
<img src="/assets/components/drawer/ic_drawer_teikyo-e03ea7d1bb3df365d0faacaf1d24ecd1d5151a2ce77729ca56601f9c45e136a8.svg" />
</span>
<span class='title flex-grow-1 pl-2'>
提供
</span>
</a><a class="list-group-item list-group-item-action d-flex align-items-center group3" href="/mypage"><span class='icon'>
<img src="/assets/components/drawer/ic_drawer_mypage-471c11e66479661ee2461ab3dcf754eea339402dcf9c421791c2949105b22d12.svg" />
</span>
<span class='title flex-grow-1 pl-2'>
マイページ
</span>
</a></div>
<div class='buttons'>
<a class="button" href="/helps">ヘルプ</a>
<a class="button" href="/privacy_policy">プライバシーポリシー</a>
<a class="button" href="/inquiries/new">お問い合わせ</a>
</div>
</div>

<div class='home' data-navigation-position='130'>
<div class='contents'>
<div class='home_contents news fillin'>
<h2 class='content-title content-title_news'>ニュース</h2>
<div class='contents'>
<ul class='news-items'>
<li class='category-information news-items_item'>
<a href="/news/22"><div class='news-item_published-date'>
2019.04.29
</div>
<div class='news-item_title'>
えのぐに逢いに恋!!プロジェクト始動
</div>
</a></li>
<li class='category-information news-items_item'>
<a href="/news/21"><div class='news-item_published-date'>
2019.04.02
</div>
<div class='news-item_title'>
栗原桜子の一時休養について
</div>
</a></li>
</ul>

<div class='buttons'>
<a class="button button-normal-20" href="/news">ニュースへ</a>
</div>
</div>
</div>
<div class='home_contents blogs fillin'>
<h2 class='content-title content-title_blogs'>ブログ</h2>
<div class='contents'>
<ul class='blog-entries'>
<li class='blog-entry nao'>
<a target="_blank" href="https://ameblo.jp/enogu-nao/entry-12460330783.html"><div class='blog-entry_meta'>
<div class='blog-entry_author'>日向奈央</div>
<div class='blog-entry_published-date'>
2019.05.10
</div>
</div>
<div class='blog-entry_title'>
おもひでぽろぽろぽろぽろ
</div>
<div class='blog-entry_body'>

 みんな〜〜〜！やほやほ！！！えのぐの日向奈央です！！！！！遅くなっちゃったけど、仙台握手会の感想とか！書こうかなって思って！！！！思い返すと、思い出が...
</div>
</a></li>
<li class='anzu blog-entry'>
<a target="_blank" href="https://ameblo.jp/enogu-anzu/entry-12459896476.html"><div class='blog-entry_meta'>
<div class='blog-entry_author'>鈴木あんず</div>
<div class='blog-entry_published-date'>
2019.05.08
</div>
</div>
<div class='blog-entry_title'>
平成最後の仙台。
</div>
<div class='blog-entry_body'>

 こんばんわんず～。VRアイドル えのぐ の鈴木あんずです(*&#39;▽&#39;*)令和だ！！！先月15日から27日まで仙台にて握手会やら、チェキ会やら、ライブをし...
</div>
</a></li>
</ul>

<div class='buttons'>
<a class="button button-normal-20" href="/blogs">ブログへ</a>
</div>
</div>
</div>
<div class='home_contents archives fillin'>
<h2 class='content-title content-title_archive'>動画アーカイブ</h2>
<div class='contents'>
<ul class='archive-items'>
<li class='archive-items_item'>
<a target="_blank" href="https://youtu.be/8o9vWxRNjTg"><div class='archive-item_thumbnail'>
<img src="https://i.ytimg.com/vi/8o9vWxRNjTg/hqdefault.jpg" />
</div>
<h3 class='archive-item_title'>
体験型企画プロジェクト『えのぐに逢いに恋‼︎』遂に始動します
</h3>
<div class='archive-item_published-date'>
2019.04.29
</div>
</a></li>
<li class='archive-items_item'>
<a target="_blank" href="https://youtu.be/0i-xt7Xuirw"><div class='archive-item_thumbnail'>
<img src="https://i.ytimg.com/vi/0i-xt7Xuirw/maxresdefault.jpg" />
</div>
<h3 class='archive-item_title'>
【検証】自分たちの曲ならヘリウム吸ったって余裕でハモれるに決まってら〜〜〜〜！！！
</h3>
<div class='archive-item_published-date'>
2019.04.27
</div>
</a></li>
</ul>

<div class='buttons'>
<a class="button button-normal-20" href="/archives">動画アーカイブへ</a>
</div>
</div>
</div>
<div class='home_contents gallery fillin'>
<h2 class='content-title content-title_gallery'>ギャラリー</h2>
<div class='contents'>
<ul class='gallery-items'>
<li class='gallery-items_item gallery_wallpaper'>
<a href="/gallery/wallpapers/16"><div class='gallery-item_thumbnail'>
<img src="https://enogu-no6.com/rails/active_storage/blobs/eyJfcmFpbHMiOnsibWVzc2FnZSI6IkJBaHBBdklGIiwiZXhwIjpudWxsLCJwdXIiOiJibG9iX2lkIn19--a7e82d70690723cc9e2c8cd324f57353abdda5ba/%E5%A3%81%E7%B4%99_%E4%B8%80%E5%91%A8%E5%B9%B4%E3%83%A9%E3%82%A4%E3%83%96.png" />
</div>
<h3 class='gallery-item_title'>
enogu 1st Anniversary Live ~ARe You Ready?~ 壁紙
</h3>
<div class='gallery-item_published-date'>
2019.03.29
</div>

</a></li>
<li class='gallery-items_item gallery_wallpaper'>
<a href="/gallery/wallpapers/15"><div class='gallery-item_thumbnail'>
<img src="https://enogu-no6.com/rails/active_storage/blobs/eyJfcmFpbHMiOnsibWVzc2FnZSI6IkJBaHBBdFFGIiwiZXhwIjpudWxsLCJwdXIiOiJibG9iX2lkIn19--f513f5719ec2d19cc4b6d093e6d429029438d7cf/%E5%A3%81%E7%B4%99_%E6%A1%9C%E5%AD%90%E8%AA%95%E7%94%9F%E7%A5%AD.jpg" />
</div>
<h3 class='gallery-item_title'>
#桜色に染まる -栗原桜子17歳誕生祭- 壁紙
</h3>
<div class='gallery-item_published-date'>
2019.03.11
</div>

</a></li>
</ul>

<div class='buttons'>
<a class="button button-normal-20" href="/gallery/movies">ギャラリーへ</a>
</div>
</div>
</div>
<div class='home_contents letters'>
<div class='balloon'>
<div class='balloon_title'>
<h2>ふつおた</h2>
</div>
<div class='balloon_body'>
えのぐのメンバーに聞きたいこと、最近あった面白いこと、なんでもお送りください！
</div>
<div class='balloon_buttons'>
<a class="button button-normal-10" href="/letters/new">ふつおたを送る</a>
</div>
</div>
</div>
<div class='home_contents sponsors'>
<div class='balloon'>
<div class='balloon_title'>
<h2>提供</h2>
<div>毎週月曜日 生放送</div>
</div>
<div class='balloon_body'>
生放送の最後にえのぐのメンバーがお名前をお呼びします。呼ばれたいお名前を投稿してください！
</div>
<div class='balloon_buttons'>
<a class="button button-normal-10" href="/sponsors/new">提供リストに名前を送る</a>
</div>
</div>
</div>
</div>
<div class='footer'>
岩本町芸能社 &copy; 2018-2019 iwamotocho geinosha All rights reserved.
</div>
</div>

</div>
</body>
</html>`

	archivesPageRespons = `HTTP/2.0 200 OK
date: Tue, 14 May 2019 13:48:48 GMT
content-type: text/html; charset=utf-8
vary: Accept-Encoding
vary: Origin
x-frame-options: SAMEORIGIN
x-xss-protection: 1; mode=block
x-content-type-options: nosniff
x-download-options: noopen
x-permitted-cross-domain-policies: none
referrer-policy: strict-origin-when-cross-origin
etag: W/"e5d142e41091528bcbe6d16bd46250ba"
cache-control: max-age=0, private, must-revalidate
set-cookie: _enogu_fanclub_session=ikMg%2FCZJWEwmJ%2B8JMgPLytFOWJ7mS4Vj0QxCBCNkr0h2AOKxSDmeLJxhIMuP1Wiv5KxmYPy6EVUwVaIUIRUGsisGIJeT17lkp45yuUzaX45%2FE6lcHtjm1lrYmVDryDAkwM8uZRe8TOb8uadeqojhNhCyDutu%2BniDE3WrvQ%3D%3D--2rGje1LyYBi2qBj9--hXy43AcdZjoxlBOBxiVYLg%3D%3D; path=/; secure; HttpOnly
x-request-id: 71f4ae71-e19a-4336-b82a-dc8f0bc9b0c6
x-runtime: 0.046169
x-cloud-trace-context: 8fb0927ca82c7fb219b0c1260de97ad2/11743544157570766645;o=1
strict-transport-security: max-age=31536000; includeSubDomains
content-encoding: gzip
via: 1.1 google
X-Firefox-Spdy: h2

<!DOCTYPE html>
<html>
<head>
<meta content='text/html; charset=UTF-8' http-equiv='Content-Type'>
<meta charset='utf-8'>
<meta content='width=device-width, initial-scale=1, shrink-to-fit=no, viewport-fit=cover' name='viewport'>
<title>えのぐ公式ファンクラブ -No,6-</title>
<meta name="csrf-param" content="authenticity_token" />
<meta name="csrf-token" content="qcALwO1O1zX7pc0TuKFTgh9evdMRW3OPpPS8kqrsxFIAdFyJsPGJjOOLoqOXXmy+GingnYfqcOkyiWMsFvUMvQ==" />

<link rel="stylesheet" media="all" href="/assets/application-4b46eb12d8c7da1e1586ea0b7d7a6c9a0d0c17fdfdb729ce7897a86aaf985da4.css" data-turbolinks-track="reload" />
<script src="/assets/application-6e438b307bf4ec2b96355941f40b2eecec0f12239862748a7017b5aaab2295a4.js" data-turbolinks-track="reload"></script>
<script async src='https://www.googletagmanager.com/gtag/js?id=UA-104370994-2'></script>
<script>
  window.dataLayer = window.dataLayer || [];
  function gtag(){dataLayer.push(arguments);}
  gtag('js', new Date());

  gtag('config', 'UA-104370994-2');
</script>

</head>
<body class='archives lemon' lang='ja'>
<div class='container'>
<h1 class='title-logo'>
<a href="/">えのぐ公式ファンクラブ -No,6-</a>
</h1>
<div class='navigation-bar'>
<button class='drawer-open'>
<img src="/assets/components/drawer/open_button-2c203f8d3250f4e1850d8ff11cd92a288ec462e25ba6d50dd83e2c2148d3c488.svg" />
</button>
<div class='title-logo'>
<a href="/">えのぐ公式ファンクラブ -No,6-</a>
</div>
</div>
<div class='drawer'>
<button class='drawer-close'>
<img src="/assets/components/drawer/close_button-09ecdc04ecc788c8fbf57358d718661ad2fe76856e203297353241f02fbd4e64.svg" />
</button>
<div class='title-logo'>
<a href="/home">えのぐ公式ファンクラブ -No,6-</a>
</div>
<div class='list-group'>
<a class="list-group-item list-group-item-action d-flex align-items-center group1" href="/home"><span class='icon'>
<img src="/assets/components/drawer/ic_drawer_home-1531f8bbb1445531a0398b5b46c98bdb3e2383b45cddf21cbd9e6be5eb540fd5.svg" />
</span>
<span class='title flex-grow-1 pl-2'>
ホーム
</span>
</a><a class="list-group-item list-group-item-action d-flex align-items-center group1" href="/news"><span class='icon'>
<img src="/assets/components/drawer/ic_drawer_news-319d1e4d58f495cbafe3ae01d1549791c4f1d9eb9c5a0a8febec14d531730bc3.svg" />
</span>
<span class='title flex-grow-1 pl-2'>
ニュース
</span>
</a><a class="list-group-item list-group-item-action d-flex align-items-center group1" href="/blogs"><span class='icon'>
<img src="/assets/components/drawer/ic_drawer_blog-4fd91d1e42959fe2537c0ac88753def63e1e147d93990da56971a0215e3efea0.svg" />
</span>
<span class='title flex-grow-1 pl-2'>
ブログ
</span>
</a><a class="list-group-item list-group-item-action d-flex align-items-center group1" href="/gallery/movies"><span class='icon'>
<img src="/assets/components/drawer/ic_drawer_gallery-697063773e7f8149ed4796c65891345a651cabcd82826b3a18ec7d41c23a9ec6.svg" />
</span>
<span class='title flex-grow-1 pl-2'>
ギャラリー
</span>
</a><a class="list-group-item list-group-item-action d-flex align-items-center group1" href="/archives"><span class='icon'>
<img src="/assets/components/drawer/ic_drawer_movie-f7c8e7479c206b3497fb29e7fba9c9706837f70aff841b29c221d14a59b4398f.svg" />
</span>
<span class='title flex-grow-1 pl-2'>
動画アーカイブ
</span>
</a><a class="list-group-item d-flex align-items-center group1 coming-soon" href="/goods"><span class='icon'>
<img src="/assets/components/drawer/ic_drawer_goods-e4880bb7fa2eac0251e342e310c0ecf98fe9cb45ed16a3d3dcbac8f305236a00.svg" />
</span>
<span class='title flex-grow-1 pl-2'>
グッズ
</span>
<span class='coming-soon-badge'>
<img src="/assets/components/drawer/ic_comingsoon_normal-f5de52da61de684a9729e2ea34127c03a06a2fe7253ac574f9f1679989e09d10.svg" />
</span>
</a><a class="list-group-item list-group-item-action d-flex align-items-center group2" href="/letters/new"><span class='icon'>
<img src="/assets/components/drawer/ic_drawer_hutsuota-ad3f9d714c924d36954995f691f30e99f1f4fcdd24122d6535bd89f2cc563100.svg" />
</span>
<span class='title flex-grow-1 pl-2'>
ふつおた
</span>
</a><a class="list-group-item list-group-item-action d-flex align-items-center group2" href="/sponsors/new"><span class='icon'>
<img src="/assets/components/drawer/ic_drawer_teikyo-e03ea7d1bb3df365d0faacaf1d24ecd1d5151a2ce77729ca56601f9c45e136a8.svg" />
</span>
<span class='title flex-grow-1 pl-2'>
提供
</span>
</a><a class="list-group-item list-group-item-action d-flex align-items-center group3" href="/mypage"><span class='icon'>
<img src="/assets/components/drawer/ic_drawer_mypage-471c11e66479661ee2461ab3dcf754eea339402dcf9c421791c2949105b22d12.svg" />
</span>
<span class='title flex-grow-1 pl-2'>
マイページ
</span>
</a></div>
<div class='buttons'>
<a class="button" href="/helps">ヘルプ</a>
<a class="button" href="/privacy_policy">プライバシーポリシー</a>
<a class="button" href="/inquiries/new">お問い合わせ</a>
</div>
</div>

<h2 class='content-title content-title_archive'>動画アーカイブ</h2>
<div class='contents'>
<p class='lead'>
Youtubeで公開された動画のアーカイブです。
<br>
サムネイルを選択するとYoutubeの画面が開きます。
</p>
<nav class='pagination'>
<span class='current first page'>

</span>

<span class='page current'>
1
</span>

<span class='page'>
<a rel="next" href="/archives?page=2">2</a>
</span>

<span class='page'>
<a href="/archives?page=3">3</a>
</span>

<span class='page'>
<a href="/archives?page=4">4</a>
</span>

<span class='page'>
<a href="/archives?page=5">5</a>
</span>

<span class='page last'>
<a href="/archives?page=44"></a>
</span>

</nav>

<ul class='archive-items'>
<li class='archive-items_item'>
<a target="_blank" href="https://youtu.be/8o9vWxRNjTg"><div class='archive-item_thumbnail'>
<img src="https://i.ytimg.com/vi/8o9vWxRNjTg/hqdefault.jpg" />
</div>
<h3 class='archive-item_title'>
体験型企画プロジェクト『えのぐに逢いに恋‼︎』遂に始動します
</h3>
<div class='archive-item_published-date'>
2019.04.29
</div>
</a></li>
<li class='archive-items_item'>
<a target="_blank" href="https://youtu.be/0i-xt7Xuirw"><div class='archive-item_thumbnail'>
<img src="https://i.ytimg.com/vi/0i-xt7Xuirw/maxresdefault.jpg" />
</div>
<h3 class='archive-item_title'>
【検証】自分たちの曲ならヘリウム吸ったって余裕でハモれるに決まってら〜〜〜〜！！！
</h3>
<div class='archive-item_published-date'>
2019.04.27
</div>
</a></li>
<li class='archive-items_item'>
<a target="_blank" href="https://youtu.be/Y-YOI-bKp9w"><div class='archive-item_thumbnail'>
<img src="https://i.ytimg.com/vi/Y-YOI-bKp9w/hqdefault.jpg" />
</div>
<h3 class='archive-item_title'>
【ゲーム実況】ア゛ア゛ア゛ア゛ァ゛ァ゛ァ゛ッ！！！！！【スプラトゥーン2】
</h3>
<div class='archive-item_published-date'>
2019.04.25
</div>
</a></li>
<li class='archive-items_item'>
<a target="_blank" href="https://youtu.be/Pbf9dWQcI48"><div class='archive-item_thumbnail'>
<img src="https://i.ytimg.com/vi/Pbf9dWQcI48/maxresdefault.jpg" />
</div>
<h3 class='archive-item_title'>
【ASMR】睡眠のおともに - バイノーラルマイクで&quot;しりとり&quot;してみた
</h3>
<div class='archive-item_published-date'>
2019.04.23
</div>
</a></li>
<li class='archive-items_item'>
<a target="_blank" href="https://youtu.be/T0reTlszrdw"><div class='archive-item_thumbnail'>
<img src="https://i.ytimg.com/vi/T0reTlszrdw/hqdefault.jpg" />
</div>
<h3 class='archive-item_title'>
【ゲーム実況】今日は叫びません！ だって…アイドルだからさ【スプラトゥーン2】
</h3>
<div class='archive-item_published-date'>
2019.04.22
</div>
</a></li>
<li class='archive-items_item'>
<a target="_blank" href="https://youtu.be/eO72c2jSYOM"><div class='archive-item_thumbnail'>
<img src="https://i.ytimg.com/vi/eO72c2jSYOM/maxresdefault.jpg" />
</div>
<h3 class='archive-item_title'>
あのね…？　たまき、横浜に行きたいの！
</h3>
<div class='archive-item_published-date'>
2019.04.18
</div>
</a></li>
</ul>

<nav class='pagination'>
<span class='current first page'>

</span>

<span class='page current'>
1
</span>

<span class='page'>
<a rel="next" href="/archives?page=2">2</a>
</span>

<span class='page'>
<a href="/archives?page=3">3</a>
</span>

<span class='page'>
<a href="/archives?page=4">4</a>
</span>

<span class='page'>
<a href="/archives?page=5">5</a>
</span>

<span class='page last'>
<a href="/archives?page=44"></a>
</span>

</nav>

</div>

</div>
</body>
</html>`
)
