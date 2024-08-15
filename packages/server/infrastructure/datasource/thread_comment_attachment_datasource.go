package datasource

import (
	"context"
	"server/domain/model"
	"server/infrastructure/ent"
	"server/infrastructure/ent/tag"
	"server/infrastructure/ent/thread"
	"server/infrastructure/ent/threadcomment"
	"server/infrastructure/ent/threadcommentattachment"

	"entgo.io/ent/dialect/sql"
)

type ThreadCommentAttachmentDatasource struct {
	client *ent.Client
}

func NewThreadCommentAttachmentDatasource(client *ent.Client) *ThreadCommentAttachmentDatasource {
	return &ThreadCommentAttachmentDatasource{client: client}
}

type ThreadCommentAttachmentDatasourceFindAllByHistoyParams struct {
	Ctx    context.Context
	TagIDs []int
}

func (ds *ThreadCommentAttachmentDatasource) FindAllByHistory(params ThreadCommentAttachmentDatasourceFindAllByHistoyParams) ([]*model.ThreadCommentAttachment, error) {
	entCommentAttachmentList, err := ds.client.
		ThreadCommentAttachment.
		Query().
		Where(
			threadcommentattachment.HasCommentWith(
				threadcomment.HasThreadWith(
					thread.HasTagsWith(
						tag.IDIn(params.TagIDs...),
					),
				),
			),
		).
		WithComment(func(q *ent.ThreadCommentQuery) {
			q.WithAuthor().
				WithThread()
		}).
		Order(sql.OrderByRand()).
		Limit(100).
		All(params.Ctx)

	if err != nil {
		return nil, err
	}

	var attachmentList []*model.ThreadCommentAttachment
	for _, entCommentAttachment_i := range entCommentAttachmentList {
		attachmentList = append(attachmentList, model.NewThreadCommentAttachment(model.NewThreadCommentAttachmentParams{EntAttachment: entCommentAttachment_i}))
	}

	return attachmentList, nil
}
