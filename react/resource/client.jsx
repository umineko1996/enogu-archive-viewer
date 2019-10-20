function App() {
  const [videos, setVideos] = React.useState([]);
  // ここは基本値の受け渡し（更新）だけの予定
  // videos が関数から渡される値、prevが以前の値
  const updateVideos = React.useCallback((videos) => setVideos(() => videos), [setVideos]);

  return (
    <div>
      <SearchVideo updateVideoList={updateVideos}/>
      <h2 className='content-title content-title_archive'>動画アーカイブ</h2>
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
    <VideoBox key={v.id} video = {v}/>
  ));
  return (
    <div class='contents'>
      <ul className='archive-items'> {videoBoxList} </ul>
    </div>
  );
}

function VideoBox(props) {
  // TODO ここでvideos変数の整形をする
  let v = props.video
  return (
    <li className='archive-items_item'><a target="_blank" href={v.url}>
      <div className='archive-item_thumbnail'><img src={v.thumbnail} /></div>
      <h3 className='archive-item_title'>{v.title}</h3>
      <div className='archive-item_published-date'>{v.date}</div>
    </a></li>
  );
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
  //console.log(jsonData)
  const videos = jsonData.Videos.map((v) => {
    //console.log(v)
  const video = {
      title: v.Title,
      url: v.URL,
      thumbnail: v.Thumbnail,
      id: v.ID,
      date: v.PublishedDate
    }
    return video;
  });

  return  videos;
}
