package commentrepo

import (
	"fmt"
	"time"
)

type CommentsResponse struct {
	Id int64 `json:"id"`
	AuthorID string `json:"authorID"`
	VideoID string `json:"videoID"`
	CommentText string `json:"commentText"`
	Date time.Time `json:"date"`
	ParentID int64 `json:"parentID"`
	Likes int	`json:"likes"`
}

func (c *CommentRepository) GetComments(videoID string) ([]CommentsResponse, error) {
	comments := make([]CommentsResponse, 0)
	res, err := c.DB.Query(`
		SELECT 
		ID,
		AuthorID,
     	VideoID,
     	CommentText,
     	Date,
     	ParentID,
     	Likes
		FROM comments 
		WHERE videoid = $1;
	`, videoID)

	if err != nil {
		return nil, err
	}
	fmt.Println("Did through query comments")
	defer res.Close()
	for res.Next() {
		var comment CommentsResponse;

		err := res.Scan(&comment.Id, &comment.AuthorID, &comment.VideoID, &comment.CommentText, &comment.Date, &comment.ParentID, &comment.Likes)

		if err != nil {
			fmt.Println("err:", err)
		}
		comments = append(comments, comment)
	}
	fmt.Println("Did through for")
	if err := res.Err(); err != nil {
		return nil, err
	}
	fmt.Println("About to return data")
	return comments, nil
}