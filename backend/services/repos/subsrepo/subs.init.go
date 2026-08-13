package subsrepo

func (s *SubsRepository) InitializeSubs() (bool, error) {
	_, err := s.DB.Exec(`CREATE TABLE IF NOT EXISTS Subs (
     	UserID TEXT NOT NULL,
     	TargetID TEXT NOT NULL,
     	PRIMARY KEY (UserID, TargetID),
     	FOREIGN KEY (UserId) REFERENCES Users(ID) ON DELETE CASCADE,
     	FOREIGN KEY (TargetID) REFERENCES Users(ID) ON DELETE CASCADE,
     	CHECK (UserID <> TargetID)
	);`)

	if err != nil {
		return false, err
	}
	return true, nil
}
