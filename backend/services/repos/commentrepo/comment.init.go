package commentrepo

func (c *CommentRepository) InitializeComments() (bool, error) {
	_, err := c.DB.Exec(`CREATE TABLE IF NOT EXISTS Comments (
     	ID BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
     	AuthorID TEXT NOT NULL,
     	VideoID TEXT NOT NULL,
     	CommentText VARCHAR(256) NOT NULL,
     	Date DATE DEFAULT CURRENT_DATE,
     	ParentID BIGINT,
     	Likes INT DEFAULT 0,
     	FOREIGN KEY (AuthorID) REFERENCES Users(ID),
     	FOREIGN KEY (VideoID) REFERENCES Videos(ID),
     	FOREIGN KEY (ParentID) REFERENCES Comments(ID)
	)`)

	if err != nil {
		return false, err
	}
	return true, nil
}