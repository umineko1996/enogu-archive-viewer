package no6

import (
	"bufio"
	"net/http"
	"strings"
	"testing"
)

func TestExtractArchiveItems(t *testing.T) {
	t.Run("TestExtractArchiveItems_OK", func(t *testing.T) {
		// 準備
		res, err := http.ReadResponse(
			bufio.NewReader(
				strings.NewReader(archivesPageRespons)),
			nil)
		if err != nil {
			t.Fatalf("extractArchiveItems関数の検証に失敗しました。 err: %s", err.Error())
		}
		// 実行
		archiveItems, err := extractArchiveItems(res.Body)

		// 検証
		if err != nil {
			t.Fatalf("extractArchiveItems関数の検証に失敗しました。 err: %s", err.Error())
		}
		got := string(archiveItems[0])
		want := `<li class='archive-items_item'>
<a target="_blank" href="https://youtu.be/8o9vWxRNjTg"><div class='archive-item_thumbnail'>
<img src="https://i.ytimg.com/vi/8o9vWxRNjTg/hqdefault.jpg" />
</div>
<h3 class='archive-item_title'>
体験型企画プロジェクト『えのぐに逢いに恋‼︎』遂に始動します
</h3>
<div class='archive-item_published-date'>
2019.04.29
</div>
</a></li>`
		if got != want {
			t.Errorf("extractArchiveItems関数の検証に失敗しました。 got %s, want %s", got, want)
		}
	})
}

func TestExtractArchiveInfo(t *testing.T) {
	t.Run("extractArchiveInfo_OK", func(t *testing.T) {
		// 準備
		source := `<li class='archive-items_item'>
<a target="_blank" href="https://youtu.be/8o9vWxRNjTg"><div class='archive-item_thumbnail'>
<img src="https://i.ytimg.com/vi/8o9vWxRNjTg/hqdefault.jpg" />
</div>
<h3 class='archive-item_title'>
体験型企画プロジェクト『えのぐに逢いに恋‼︎』遂に始動します
</h3>
<div class='archive-item_published-date'>
2019.04.29
</div>
</a></li>`

		// 実行
		archive, err := extractArchiveInfo(source)

		// 検証
		if err != nil {
			t.Fatalf("extractArchiveInfo関数の検証に失敗しました。 err: %s", err.Error())
		}
		if got, want := archive.Title, "体験型企画プロジェクト『えのぐに逢いに恋‼︎』遂に始動します"; got != want {
			t.Errorf("extractArchiveInfo関数の検証に失敗しました。 titleの値が不正です。 got %s, want %s", got, want)
		}
		if got, want := archive.PublishedDate, "2019.04.29"; got != want {
			t.Errorf("extractArchiveInfo関数の検証に失敗しました。 publishedDateの値が不正です。 got %s, want %s", got, want)
		}
		if got, want := archive.URL, "https://youtu.be/8o9vWxRNjTg"; got != want {
			t.Errorf("extractArchiveInfo関数の検証に失敗しました。 urlの値が不正です。 got %s, want %s", got, want)
		}
		if got, want := archive.Thumbnail, "https://i.ytimg.com/vi/8o9vWxRNjTg/hqdefault.jpg"; got != want {
			t.Errorf("extractArchiveInfo関数の検証に失敗しました。 thumbnailの値が不正です。 got %s, want %s", got, want)
		}
		if got, want := archive.ID, "8o9vWxRNjTg"; got != want {
			t.Errorf("extractArchiveInfo関数の検証に失敗しました。 idの値が不正です。 got %s, want %s", got, want)
		}
	})

	t.Run("extractArchiveInfo_OK_htmlEncode有り", func(t *testing.T) {
		// 準備
		source := `<li class='archive-items_item'>
<a target="_blank" href="https://youtu.be/Pbf9dWQcI48"><div class='archive-item_thumbnail'>
<img src="https://i.ytimg.com/vi/Pbf9dWQcI48/maxresdefault.jpg" />
</div>
<h3 class='archive-item_title'>
【ASMR】睡眠のおともに - バイノーラルマイクで&quot;しりとり&quot;してみた
</h3>
<div class='archive-item_published-date'>
2019.04.23
</div>
</a></li>`

		// 実行
		archive, err := extractArchiveInfo(source)

		// 検証
		if err != nil {
			t.Fatalf("extractArchiveInfo関数の検証に失敗しました。 err: %s", err.Error())
		}
		if got, want := archive.Title, "【ASMR】睡眠のおともに - バイノーラルマイクで\"しりとり\"してみた"; got != want {
			t.Errorf("extractArchiveInfo関数の検証に失敗しました。 titleの値が不正です。 got %s, want %s", got, want)
		}
		if got, want := archive.PublishedDate, "2019.04.23"; got != want {
			t.Errorf("extractArchiveInfo関数の検証に失敗しました。 publishedDateの値が不正です。 got %s, want %s", got, want)
		}
		if got, want := archive.URL, "https://youtu.be/Pbf9dWQcI48"; got != want {
			t.Errorf("extractArchiveInfo関数の検証に失敗しました。 urlの値が不正です。 got %s, want %s", got, want)
		}
		if got, want := archive.Thumbnail, "https://i.ytimg.com/vi/Pbf9dWQcI48/maxresdefault.jpg"; got != want {
			t.Errorf("extractArchiveInfo関数の検証に失敗しました。 thumbnailの値が不正です。 got %s, want %s", got, want)
		}
		if got, want := archive.ID, "Pbf9dWQcI48"; got != want {
			t.Errorf("extractArchiveInfo関数の検証に失敗しました。 idの値が不正です。 got %s, want %s", got, want)
		}
	})

}

var archivesPageRespons = `HTTP/2.0 200 OK
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
