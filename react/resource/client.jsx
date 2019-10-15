function App() {
  const [videos, setVideos] = React.useState([{url: `first.com`}, {url: `second.com`}]);
  // ここは基本値の受け渡し（更新）だけの予定
  // videos が関数から渡される値、prevが以前の値
  const updateVideos = React.useCallback((videos) => setVideos((prev) => [...prev, ...videos]), [setVideos]);

  return (
    <div>
      <p>Hello World!</p>
      <SearchVideo updateVideoList={updateVideos}/>
      <VideoBoxList videos = {videos}/>
    </div>
  );
}

const target = document.querySelector('#app');
ReactDOM.render(<App/>, target);

function SearchVideo(props) {
  const inputRef = React.useRef(null);

    // コールバック関数でフィールドのアウトプットを行うためのpropsを更新する
  const createVideoBox = React.useCallback( () => {
    if(inputRef.current) {
      // TODO ここに通信処理入れる
      getVideos(inputRef.current.value)
      .then((newVideos) => props.updateVideoList(newVideos))
    }
  }, [inputRef.current, props.updateVideoList])

  // inputのフィールドをリターンする
  return (
    <div>
      <div><input className="search-form" ref={inputRef}></input></div>
      <div><button onClick={createVideoBox} className="search">Search</button></div>
    </div>
  );
}

function VideoBoxList(props) {
  const videoBoxList = props.videos.map((v) => (
    // TODO keyは動画で一意の値に
    <VideoBox key={v.url} video = {v}/>
  ));
  return <div> {videoBoxList} </div>;
}

function VideoBox(props) {
  // TODO ここでvideos変数の整形をする
  return <li>{props.video.url}</li>;
}

const server = `localhost:6060`

function getVideos(word) {
  return new Promise(function(resolve) {
    fetch(`http://${server}/search?w=${encodeURIComponent(word)}`)
    .then((response) => {
      if (response.ok) {
        return response.json()
      }
      throw new Error(response.status)
    })
    .then((jsonData) => resolve(convVideos(jsonData)))
    .catch((error) => console.error(error));
  })
}

function convVideos(jsonData) {
  const videos = jsonData.videos.map((v) => {
  const video = {
      // TODO 必要な情報を入れる
      url: v.url
    }
    return video;
  });

  return  videos;
}
