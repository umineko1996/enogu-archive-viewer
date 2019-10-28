function VideoNav(props) {
        const createVideoBox = React.useCallback((page) => {
                getVideos(searchWords, page)
                .then((newVideos) => props.updateVideoList(newVideos))
        }, [props.updateVideoList])

        //console.log(props)
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
