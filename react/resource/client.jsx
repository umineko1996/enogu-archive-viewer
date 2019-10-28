var searchWords = "all"

function App() {
  const [videos, setVideos] = React.useState(
  () => {
      getVideos(searchWords, 1)
      .then((newVideos) => setVideos(newVideos))
      return []
  });
  // ここは基本値の受け渡し（更新）だけの予定
  // videos が関数から渡される値、prevが以前の値
  const updateVideos = React.useCallback((newVideos) => setVideos(() => newVideos), [setVideos]);

  return (
    <div>
      <SearchVideo updateVideoList={updateVideos}/>
      <h2 className='content-title content-title_archive'>動画アーカイブ</h2>
      <VideoBoxList videos = {videos} updateVideoList={updateVideos}/>
    </div>
  );
}

const target = document.querySelector('#app');
ReactDOM.render(<App/>, target);

// form無効化用
function handleSubmit(e) {
  e.preventDefault();
}

function SearchVideo(props) {
  const page = 1
  const inputRef = React.useRef(null);

  // コールバック関数でフィールドのアウトプットを行うためのpropsを更新する
  const createVideoBox = React.useCallback(() => {
    if(inputRef.current) {
      searchWords = inputRef.current.value
      getVideos(searchWords, page)
      .then((newVideos) => props.updateVideoList(newVideos))
    }
  }, [inputRef.current, props.updateVideoList])

  // inputのフィールドをリターンする
  return (
    <div className="bordered-box">
      <h2>検索</h2>
      <form onSubmit={handleSubmit}>
      <div className='form-group'>
        <label htmlFor="search-word">検索ワード</label>
        <input className="form-control" required="required" placeholder="栗原桜子" ref={inputRef}></input>
      </div>
        <input onClick={createVideoBox} type="submit" name="search" value="表示" data-disable-with="表示" />
        </form>
    </div>
  );
}

