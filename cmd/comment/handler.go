package main

import (
	"context"
	"dousheng/cmd/comment/kafkaop"
	"dousheng/dal/db"
	douyin_comment "dousheng/kitex_gen/douyin_comment"
	"strconv"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// CommentAction implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentAction(ctx context.Context, req *douyin_comment.CommentActionRequest) (resp *douyin_comment.CommentActionResponse, err error) {
	// TODO: Your code here...
	var commentText string
	var commentId int64
	if req.CommentText != nil {
		commentText = *req.CommentText
	}
	if req.CommentId != nil {
		commentId = *req.CommentId
	}

	err = kafkaop.ProduceComment(ctx, req.UserId, req.VideoId, req.ActionType, commentText, strconv.FormatInt(commentId, 10))
	if err != nil {
		return nil, err
	}

	resp = douyin_comment.NewCommentActionResponse()
	return resp, nil
}

// CommentList implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentList(ctx context.Context, req *douyin_comment.CommentListRequest) (resp *douyin_comment.CommentListResponse, err error) {
	// TODO: Your code here...
	comments := db.GetComments(req.VideoId)
	resp = douyin_comment.NewCommentListResponse()
	resp.CommentList = comments
	return
}
