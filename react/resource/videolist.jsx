function VideoBoxList(props) {
        const nav = VideoNav(props);
        //console.log(props)
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
