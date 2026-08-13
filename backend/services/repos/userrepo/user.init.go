package userrepo

func (u *UserRepository) InitializeUser() (bool, error) {
	_, err := u.DB.Exec(`CREATE TABLE IF NOT EXISTS Users (
     	ID TEXT NOT NULL PRIMARY KEY,
     	UserName VARCHAR(16) UNIQUE,
     	CreationDate TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     	Bio VARCHAR(255),
     	Subs INT NOT NULL DEFAULT 0
	);`)

	if err != nil {
		return false, err
	}

	return true, nil
}