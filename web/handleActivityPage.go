package web

import (
	"fmt"
	"forum/comments"
	"forum/posts"
	"forum/users"
	"net/http"
	"strconv"
)

type ActivityPage struct {
	Posts             []posts.HomepagePosts
	CommentsWithPosts []posts.ActPage
	LikedPosts        []posts.ActPage
	DislikedPosts     []posts.ActPage
	LikedComments     []posts.ActPage
	DislikedComments  []posts.ActPage
	Comments          []posts.Post
	Username          string
	LoggedIn          bool
	UserID            int
	Nbool             bool
	Notification      int
	CommentNote       []Notify
	LikeNote          []Notify
	DisLikeNote       []Notify
}

func (s *myServer) EditActComment() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			r.ParseForm()

			fmt.Println("--------", r.Form["editpost"])
		}

		Tpl.ExecuteTemplate(w, "editactpost.html", PostPageData{LoggedIn: users.AlreadyLoggedIn(r), Username: users.CurrentUser, UserID: UserIdint})
	}
}

func (s *myServer) DeleteActComment() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		for _, v := range r.Form {
			for _, id := range v {
				idInt, _ := strconv.Atoi(id)
				comments.DeleteComment(s.Db, idInt)
			}
		}
		stringGID := strconv.Itoa(GuserId)
		http.Redirect(w, r, "/activitypage?userid="+stringGID, http.StatusSeeOther)
	}
}

func (s *myServer) DeleteActPost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		for _, v := range r.Form {
			for _, id := range v {
				idInt, _ := strconv.Atoi(id)
				posts.DeletePost(s.Db, idInt)
			}
		}

		stringGID := strconv.Itoa(GuserId)

		http.Redirect(w, r, "/activitypage?userid="+stringGID, http.StatusSeeOther)

	}
}

func (s *myServer) ActivityPage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var data ActivityPage

		data.Notification = (len(CommentNotify(s.Db)) + len(LikesNotify(s.Db)) + len(DisLikesNotify(s.Db)))

		if data.Notification > 0 {
			data.Nbool = true
		}

		data.CommentNote = CommentNotify(s.Db)

		data.LikeNote = LikesNotify(s.Db)

		data.DisLikeNote = DisLikesNotify(s.Db)

		data.Posts = posts.UsersPostsHomepageData(s.Db, GuserId)

		data.CommentsWithPosts = posts.ActivityComments(s.Db, GuserId)

		data.LikedPosts = posts.ActivityPostLikes(s.Db, GuserId)

		data.DislikedPosts = posts.ActivityPostDislikes(s.Db, GuserId)

		data.LikedComments = posts.ActivityCommentLikes(s.Db, GuserId)

		data.DislikedComments = posts.ActivityCommentDislikes(s.Db, GuserId)
		data.Username = users.CurrentUser
		data.LoggedIn = users.AlreadyLoggedIn(r)
		data.UserID = GuserId
		SuserID := strconv.Itoa(GuserId)

		if string(r.URL.RawQuery[len(r.URL.RawQuery)-1]) != SuserID {
			http.Error(w, "Incorrect user request made!", http.StatusBadRequest)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		Tpl.ExecuteTemplate(w, "activitypage.html", data)

		func() {
			ResetCommentNotified(s.Db)
		}()

		func() {
			ResetLikesNotified(s.Db)
		}()

		func() {
			ResetDisLikesNotified(s.Db)
		}()

	}
}
