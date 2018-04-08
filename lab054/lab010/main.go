package main

import (
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
)

func main() {
	data, err := ioutil.ReadFile("data.json")
	if err != nil {
		log.Fatalf("readFile error:%v", err)
	}

	var commentResult CommentResult
	err = gjson.Unmarshal(data, &commentResult)
	if err != nil {
		log.Fatalf("gjson error:%v", err)
	} else {
		log.Printf("%+v", commentResult)
	}
}

type CommentResult struct {
	IsMusician  bool `json:"isMusician"`
	UserID      int  `json:"userId"`
	TopComments []struct {
		User struct {
			LocationInfo interface{} `json:"locationInfo"`
			VipType      int         `json:"vipType"`
			Nickname     string      `json:"nickname"`
			UserID       int         `json:"userId"`
			UserType     int         `json:"userType"`
			ExpertTags   interface{} `json:"expertTags"`
			AuthStatus   int         `json:"authStatus"`
			RemarkName   interface{} `json:"remarkName"`
			AvatarURL    string      `json:"avatarUrl"`
			Experts      interface{} `json:"experts"`
		} `json:"user"`
		TopCommentID int `json:"topCommentId"`
		BeReplied    []struct {
			User struct {
				LocationInfo interface{} `json:"locationInfo"`
				VipType      int         `json:"vipType"`
				Nickname     string      `json:"nickname"`
				UserID       int         `json:"userId"`
				UserType     int         `json:"userType"`
				ExpertTags   interface{} `json:"expertTags"`
				AuthStatus   int         `json:"authStatus"`
				RemarkName   interface{} `json:"remarkName"`
				AvatarURL    string      `json:"avatarUrl"`
				Experts      interface{} `json:"experts"`
			} `json:"user"`
			Content string `json:"content"`
			Status  string `json:"status"`
		} `json:"beReplied"`
		PendantData interface{} `json:"pendantData"`
		CommentID   int         `json:"commentId"`
		L           int         `json:"l"`
		Time        int64       `json:"time"`
		Content     string      `json:"content"`
	} `json:"topComments"`
	MoreHot     bool `json:"moreHot"`
	HotComments []struct {
		User struct {
			LocationInfo interface{} `json:"locationInfo"`
			VipType      int         `json:"vipType"`
			Nickname     string      `json:"nickname"`
			UserID       int         `json:"userId"`
			UserType     int         `json:"userType"`
			ExpertTags   interface{} `json:"expertTags"`
			AuthStatus   int         `json:"authStatus"`
			RemarkName   interface{} `json:"remarkName"`
			AvatarURL    string      `json:"avatarUrl"`
			Experts      interface{} `json:"experts"`
		} `json:"user"`
		BeReplied []struct {
			User struct {
				LocationInfo interface{} `json:"locationInfo"`
				VipType      int         `json:"vipType"`
				Nickname     string      `json:"nickname"`
				UserID       int         `json:"userId"`
				UserType     int         `json:"userType"`
				ExpertTags   interface{} `json:"expertTags"`
				AuthStatus   int         `json:"authStatus"`
				RemarkName   interface{} `json:"remarkName"`
				AvatarURL    string      `json:"avatarUrl"`
				Experts      interface{} `json:"experts"`
			} `json:"user"`
			Content string `json:"content"`
			Status  string `json:"status"`
		} `json:"beReplied"`
		PendantData interface{} `json:"pendantData"`
		Liked       bool        `json:"liked"`
		CommentID   int         `json:"commentId"`
		LikedCount  int         `json:"likedCount"`
		Time        int64       `json:"time"`
		Content     string      `json:"content"`
	} `json:"hotComments"`
	Code     int `json:"code"`
	Comments []struct {
		User struct {
			LocationInfo interface{} `json:"locationInfo"`
			VipType      int         `json:"vipType"`
			Nickname     string      `json:"nickname"`
			UserID       int         `json:"userId"`
			UserType     int         `json:"userType"`
			ExpertTags   interface{} `json:"expertTags"`
			AuthStatus   int         `json:"authStatus"`
			RemarkName   interface{} `json:"remarkName"`
			AvatarURL    string      `json:"avatarUrl"`
			Experts      interface{} `json:"experts"`
		} `json:"user"`
		BeReplied []struct {
			User struct {
				LocationInfo interface{} `json:"locationInfo"`
				VipType      int         `json:"vipType"`
				Nickname     string      `json:"nickname"`
				UserID       int         `json:"userId"`
				UserType     int         `json:"userType"`
				ExpertTags   interface{} `json:"expertTags"`
				AuthStatus   int         `json:"authStatus"`
				RemarkName   interface{} `json:"remarkName"`
				AvatarURL    string      `json:"avatarUrl"`
				Experts      interface{} `json:"experts"`
			} `json:"user"`
			Content string `json:"content"`
			Status  string `json:"status"`
		} `json:"beReplied"`
		PendantData        interface{} `json:"pendantData"`
		Liked              bool        `json:"liked"`
		CommentID          int         `json:"commentId"`
		LikedCount         int         `json:"likedCount"`
		Time               int64       `json:"time"`
		Content            string      `json:"content"`
		IsRemoveHotComment bool        `json:"isRemoveHotComment"`
	} `json:"comments"`
	Total int  `json:"total"`
	More  bool `json:"more"`
}
