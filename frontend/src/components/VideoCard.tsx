import type { VideoResponse } from "../api/video"

type VideoProp = {
     video : VideoResponse | null
}

const VideoBox = ({ video } : VideoProp) : React.JSX.Element => {
     return (
          <div>
               <img className={`video_thumbnail ${video?.id}`} src={`http://localhost:8080/thumbnails/${video?.id}`} alt="thumbnail" />
               <h2 className={`video_title ${video?.id}`}>{video?.name}</h2>
          </div>
     )
}

export default VideoBox