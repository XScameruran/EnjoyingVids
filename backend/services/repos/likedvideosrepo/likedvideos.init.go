package likedvideosrepo

func (lv *LikedVideosRepository) InitializeLikedVideos() (bool, error) {
	_, err := lv.DB.Exec(`CREATE TABLE IF NOT EXISTS LikedVideos (
     	LikedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     	VideoID TEXT NOT NULL,
     	UserID TEXT NOT NULL,

     	PRIMARY KEY (VideoID, UserID),

     	FOREIGN KEY (VideoID) REFERENCES Videos(ID) ON DELETE CASCADE,
     	FOREIGN KEY (UserID) REFERENCES Users(ID) ON DELETE CASCADE
	);`)

	if err != nil {
		return false, err
	}
	return true, nil
}
