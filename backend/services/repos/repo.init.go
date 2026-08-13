package repos

import "fmt"

func (r *RepoInit) Initialize() error {
	if _, err := r.Handler.UserRepo.InitializeUser(); err != nil {
		fmt.Println("Error by userrepo")
		return err
	}

	if _, err := r.Handler.VideoRepo.InitializeVideos(); err != nil {
		fmt.Println("Error by videorepo")
		return err
	}

	if _, err := r.Handler.CommentRepo.InitializeComments(); err != nil {
		fmt.Println("Error by commentrepo")
		return err
	}

	if _, err := r.Handler.SubsRepo.InitializeSubs(); err != nil {
		fmt.Println("Error by subsrepo")
		return err
	}

	if _, err := r.Handler.LikedVideosRepo.InitializeLikedVideos(); err != nil {
		fmt.Println("Error by likedvideosrepo")
		return err
	}

	return nil
}