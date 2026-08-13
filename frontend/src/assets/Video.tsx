import { useEffect, useState } from "react"
import { useParams, type Params } from "react-router"
import { type VideoResponse, GetVideo } from "../api/video"
import VideoBox from "../components/VideoCard"

const Video = () : React.JSX.Element => {
     const [video, setVideo] = useState<VideoResponse | null>(null)
     const { id } : Params = useParams()
     useEffect(() => {
          console.log("/Video/")
     }, [])
     useEffect(() => {
          async function loadVideo(id : string | undefined) {
               const videores = await GetVideo(id)
               console.log(videores)
               setVideo(videores)
          }
          loadVideo(id)
     }, [id])
     return (
          <>
               <VideoBox video={video} />
          </>
     )
}

export default Video