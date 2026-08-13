export type VideoResponse = 
     {
          id : string,
          name : string,
          date : string,
          description : string,
          likes : number,
          dislikes : number,
          status : string,
     }
export async function GetVideo(id : string | undefined) : Promise<VideoResponse> {
     const response = await fetch(`http://127.0.0.1:8080/Videos/${id}/`);
     const data = await response.json();
     if (!response.ok) {
          throw new Error(`HTTP ${response.status}`)
     }

     return data;
}