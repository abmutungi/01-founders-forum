package web

import (
	"fmt"
	"forum/comments"
	"net/http"
	"strconv"
)

var CommentData comments.Comment
var CUserIdint int
var ContentComment string
var CommentId int

//var CPostIdint int

func (s *myServer) CreateCommentHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		userID := r.URL.Query().Get("userid")

		PostIDInt, _ = strconv.Atoi(userID)

		Tpl.ExecuteTemplate(w, "createcomment.html", nil)

	}
}

func (s *myServer) StoreCommentHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		r.ParseForm()
		ContentComment = r.FormValue("content")

		if ContentComment != "" {

			comments.CreateComment(s.Db, GuserId, PostIDInt, ContentComment)

			var commentData comments.Comment
			SPostID := strconv.Itoa(PostIDInt)

			commentData.CommentText = ContentComment
			commentData.PostID = PostIDInt
			commentData.UserID = GuserId

			CommentId = comments.GetCommentID(s.Db)

			fmt.Println("testing method to get comment id", CommentId)
			fmt.Println("testing cd id ************************8", CommentId)
			// fmt.Println("comment data check: ---> ", commentData.CommentText)
			// fmt.Println("comment post id check: ---> ", commentData.PostID)
			// fmt.Println("comment user id check: ---> ", commentData.UserID)

			// fmt.Println("content: ", ContentComment)
			//Tpl.ExecuteTemplate(w, "storecomment.html", commentData.CommentText)
			http.Redirect(w, r, "/showpost/?postid="+SPostID, http.StatusSeeOther)
			//http.Redirect(w, r, "/home", http.StatusSeeOther)
		}
		SPostID := strconv.Itoa(PostIDInt)
		fmt.Println("testing cd id ************************9", CommentData.CommentID)
		http.Redirect(w, r, "/showpost/?postid="+SPostID, http.StatusSeeOther)
	}
}
