function App() {
  const [videos, setVideos] = React.useState([]);
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

var searchWords = ""

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

function VideoBoxList(props) {

  const nav = VideoNav(props);
  const videoBoxList = props.videos.map((v) => (
    <VideoBox key={v.id} video={v}/>
  ));

  return (
    <div className='contents'>
      <div>{nav}</div>
      <ul className='archive-items'> {videoBoxList} </ul>
    </div>
  );
}


function VideoNav(props) {
  const createVideoBox = React.useCallback((page) => {
    getVideos(searchWords, page)
    .then((newVideos) => props.updateVideoList(newVideos))
  }, [props.updateVideoList])

  console.log(props)
  const v = props.videos

  if (v.length == 0) {
    return
  }

  const spansFirstPage = navFirstPage(createVideoBox, v.page)
  const spans = navPages(createVideoBox, v.page, v.lastPage)
  const spansLastPage = navLastPage(createVideoBox, v.page, v.lastPage)
  return (
    <nav className='pagination'>
      {spansFirstPage}
      {spans}
      {spansLastPage}
    </nav>
  )
}

function navPages(createVideoBox, page, last) {
  let spans = []
  if (last <= 5) {
    for (var i = 1; i < last + 1; i++) {
      if (i == page) {
        spans = [...spans, createCurrentNav(createVideoBox, i)]
      } else {
        spans = [...spans, createNav(createVideoBox, i)]
      }
    }
  } else if (page < 3) {
    for (var i = 1; i < 6; i++) {
      if (i == page) {
        spans = [...spans, createCurrentNav(createVideoBox, i)]
      } else {
        spans = [...spans, createNav(createVideoBox, i)]
      }
    }
  } else if (page > last - 2){
    for (var i = last - 4; i < last + 1; i++) {
      if (i == page) {
        spans = [...spans, createCurrentNav(createVideoBox, i)]
      } else {
        spans = [...spans, createNav(createVideoBox, i)]
      }
    }
  } else {
    for (var i = page - 2; i < page + 3; i++) {
      if (i == page) {
        spans = [...spans, createCurrentNav(createVideoBox, i)]
      } else {
        spans = [...spans, createNav(createVideoBox, i)]
      }
    }
  }
  console.log(spans)
  return spans
}

function createNav(createVideoBox, page) {
  return (<span key={page} className='page' onClick={()=>(createVideoBox(page))}>{page}</span>)
}

function createCurrentNav(createVideoBox, page) {
  return (<span key={page} className='page current'>{page}</span>)
}

function navFirstPage(createVideoBox, page) {
  if (page == 1) {
    return (
      <span className='current first page'></span>
    )
  }

  return (
    <span className='first page' onClick={()=>(createVideoBox(1))}></span>
  )
}

function navLastPage(createVideoBox, page, last) {
  if (page == last) {
    return (
      <span className='current last page'></span>
    )
  }

  return (
      <span className='page last' onClick={()=>(createVideoBox(last))}></span>
  )
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

function getVideos(word, page) {
  return new Promise(function(resolve) {
    fetch(`/search?w=${encodeURIComponent(word)}&page=${page}`)
    .then((response) => {
      if (response.ok) {
        return response.json();
      }
      throw new Error(response.status);
    })
    .then((jsonData) => resolve(convVideos(jsonData)))
    .catch((error) => console.error(error));
  })
}

function convVideos(jsonData) {
  //console.log(jsonData);
  const videos = jsonData.Videos.map((v) => {
  //console.log(v);
  const video = {
      title: v.Title,
      url: v.URL,
      thumbnail: v.Thumbnail,
      id: v.ID,
      date: v.PublishedDate,
    }
    return video;
  });
  //videos.total = jsonData.Total;
  videos.page = jsonData.Page;
  videos.lastPage = jsonData.LastPage;
  return  videos;
}
