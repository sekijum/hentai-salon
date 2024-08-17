package service

import (
	"context"
	"errors"
	"server/domain/model"
	"server/infrastructure/datasource"
	"server/infrastructure/ent"
	"server/presentation/request"
	"server/presentation/response"
)

type ThreadApplicationService struct {
	client                  *ent.Client
	threadDatasource        *datasource.ThreadDatasource
	threadCommentDatasource *datasource.ThreadCommentDatasource
	tagDatasource           *datasource.TagDatasource
}

func NewThreadApplicationService(
	client *ent.Client,
	threadDatasource *datasource.ThreadDatasource,
	threadCommentDatasource *datasource.ThreadCommentDatasource,
	tagDatasource *datasource.TagDatasource,
) *ThreadApplicationService {
	return &ThreadApplicationService{
		client:                  client,
		threadDatasource:        threadDatasource,
		threadCommentDatasource: threadCommentDatasource,
		tagDatasource:           tagDatasource,
	}
}

type ThreadApplicationServiceFindAllParams struct {
	Ctx    context.Context
	Qs     request.ThreadFindAllRequest
	UserID int
}

func (svc *ThreadApplicationService) FindAll(params ThreadApplicationServiceFindAllParams) ([]*response.ThreadResponse, error) {
	var threadList []*model.Thread
	var err error

	switch params.Qs.Filter {
	case "popularity":
		threadList, err = svc.threadDatasource.FindByPopularity(datasource.ThreadDatasourceFindByPopularityParams{
			Ctx:    params.Ctx,
			Limit:  params.Qs.Limit,
			Offset: params.Qs.Offset,
		})
		if err != nil {
			return nil, err
		}
	case "tags":
		if len(params.Qs.TagNameList) > 0 {

			threadList, err = svc.threadDatasource.FindAll(datasource.ThreadDatasourceFindAllParams{
				Ctx:         params.Ctx,
				Limit:       params.Qs.Limit,
				Offset:      params.Qs.Offset,
				TagNameList: params.Qs.TagNameList,
			})
			if err != nil {
				return nil, err
			}
		}
	case "related-by-history":
		if len(params.Qs.ThreadIDs) == 0 {
			return nil, nil
		}
		var tagIDs, err = svc.tagDatasource.FindAllIDs(datasource.TagDatasourceFindAllIDsParams{
			Ctx:       params.Ctx,
			ThreadIDs: params.Qs.ThreadIDs,
		})
		if err != nil {
			return nil, err
		}

		threadList, err = svc.threadDatasource.FindByRelatedTag(datasource.ThreadDatasourceFindByRelatedTagParams{
			Ctx:       params.Ctx,
			ThreadIDs: params.Qs.ThreadIDs,
			Limit:     params.Qs.Limit,
			Offset:    params.Qs.Offset,
			TagIDs:    tagIDs,
		})
		if err != nil {
			return nil, err
		}
	case "related-by-thread":
		if params.Qs.ThreadID == 0 {
			return nil, nil
		}
		var tagIDs, err = svc.tagDatasource.FindAllIDs(datasource.TagDatasourceFindAllIDsParams{
			Ctx:       params.Ctx,
			ThreadIDs: []int{params.Qs.ThreadID},
		})
		if err != nil {
			return nil, err
		}

		threadList, err = svc.threadDatasource.FindByRelatedTag(datasource.ThreadDatasourceFindByRelatedTagParams{
			Ctx:       params.Ctx,
			ThreadIDs: params.Qs.ThreadIDs,
			Limit:     params.Qs.Limit,
			Offset:    params.Qs.Offset,
			TagIDs:    tagIDs,
		})
		if err != nil {
			return nil, err
		}
	case "keyword":
		if params.Qs.Keyword == "" {
			return nil, nil
		}
		threadList, err = svc.threadDatasource.FindAll(datasource.ThreadDatasourceFindAllParams{
			Ctx:     params.Ctx,
			Keyword: params.Qs.Keyword,
			Limit:   params.Qs.Limit,
			Offset:  params.Qs.Offset,
		})
		if err != nil {
			return nil, err
		}
	case "history":
		if len(params.Qs.ThreadIDs) == 0 {
			return nil, nil
		}
		threadList, err = svc.threadDatasource.FindAll(datasource.ThreadDatasourceFindAllParams{
			Ctx:       params.Ctx,
			ThreadIDs: params.Qs.ThreadIDs,
		})
		if err != nil {
			return nil, err
		}
		threadMap := make(map[int]*model.Thread)
		for _, thread_i := range threadList {
			threadMap[thread_i.EntThread.ID] = thread_i
		}

		var sortedthreadList []*model.Thread
		for _, id_i := range params.Qs.ThreadIDs {
			if thread, ok := threadMap[id_i]; ok {
				sortedthreadList = append(sortedthreadList, thread)
			}
		}

		threadList = sortedthreadList
	case "board":
		if params.Qs.BoardID == 0 {
			return nil, errors.New("BoardIDが必要です")
		}
		threadList, err = svc.threadDatasource.FindAll(datasource.ThreadDatasourceFindAllParams{
			Ctx:     params.Ctx,
			BoardID: params.Qs.BoardID,
			Limit:   params.Qs.Limit,
			Offset:  params.Qs.Offset,
		})
		if err != nil {
			return nil, err
		}
	case "owner":
		if params.UserID == 0 {
			return nil, errors.New("UserIDが必要です")
		}
		threadList, err = svc.threadDatasource.FindAll(datasource.ThreadDatasourceFindAllParams{
			Ctx:    params.Ctx,
			UserID: params.UserID,
			Limit:  params.Qs.Limit,
			Offset: params.Qs.Offset,
		})
		if err != nil {
			return nil, err
		}
	case "newest":
		threadList, err = svc.threadDatasource.FindAll(datasource.ThreadDatasourceFindAllParams{
			Ctx:    params.Ctx,
			Limit:  params.Qs.Limit,
			Offset: params.Qs.Offset,
			Order:  "desc",
		})
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("無効なfilterです")
	}

	var dto []*response.ThreadResponse
	for _, thread_i := range threadList {
		commentCount := len(thread_i.EntThread.Edges.Comments)
		response := response.NewThreadResponse(response.NewThreadResponseParams{
			Thread:       thread_i,
			CommentCount: &commentCount,
			IncludeBoard: true,
		})
		dto = append(dto, response)
	}

	return dto, nil
}

type ThreadApplicationServiceFindByIDParams struct {
	Ctx      context.Context
	UserID   *int
	ThreadID int
	Qs       request.ThreadFindByIdRequest
}

func (svc *ThreadApplicationService) FindByID(params ThreadApplicationServiceFindByIDParams) (*response.ThreadResponse, error) {
	thread, err := svc.threadDatasource.FindById(datasource.ThreadDatasourceFindByIDParams{
		Ctx:      params.Ctx,
		Order:    params.Qs.Order,
		Limit:    params.Qs.Limit,
		Offset:   params.Qs.Offset,
		ThreadID: params.ThreadID,
	})
	if err != nil {
		return nil, err
	}

	commentCount, err := svc.threadCommentDatasource.GetCommentCount(datasource.ThreadDatasourceGetCommentCountParams{
		Ctx:      params.Ctx,
		ThreadID: &params.ThreadID,
	})
	if err != nil {
		return nil, err
	}

	dto := response.NewThreadResponse(response.NewThreadResponseParams{
		Thread:             thread,
		CommentCount:       &commentCount,
		Limit:              params.Qs.Limit,
		Offset:             params.Qs.Offset,
		UserID:             params.UserID,
		IncludeComments:    true,
		IncludeAttachments: true,
		IncludeBoard:       true,
		IncludeTagNameList: true,
	})
	return dto, nil
}

type ThreadApplicationServiceCreateParams struct {
	Ctx      context.Context
	UserID   int
	ClientIP string
	Body     request.ThreadCreateRequest
}

func (svc *ThreadApplicationService) Create(params ThreadApplicationServiceCreateParams) (*response.ThreadResponse, error) {
	threadList, err := svc.threadDatasource.FindByTitle(datasource.ThreadDatasourceFindByTitleParams{
		Ctx:   params.Ctx,
		Title: params.Body.Title,
	})
	if err != nil {
		return nil, err
	}
	if len(threadList) > 0 {
		return nil, errors.New("スレタイが重複しています")
	}

	tx, err := svc.client.Tx(params.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	tagList, err := svc.tagDatasource.CreateManyTx(datasource.TagDatasourceCreateManyTxParams{
		Ctx:         params.Ctx,
		Tx:          tx,
		TagNameList: params.Body.TagNameList,
	})
	if err != nil {
		return nil, err
	}

	var tagIDs []int
	for _, tag_i := range tagList {
		tagIDs = append(tagIDs, tag_i.EntTag.ID)
	}

	thread := model.NewThread(model.NewThreadParams{
		EntThread: &ent.Thread{
			Title:        params.Body.Title,
			BoardID:      params.Body.BoardId,
			UserID:       params.UserID,
			Description:  params.Body.Description,
			ThumbnailURL: params.Body.ThumbnailURL,
			IPAddress:    params.ClientIP,
		},
		OptionList: []func(*model.Thread){
			model.WithThreadStatus(model.ThreadStatusOpen),
		},
	})

	thread, err = svc.threadDatasource.CreateTx(datasource.ThreadDatasourceCreateTxParams{
		Ctx:    params.Ctx,
		Tx:     tx,
		Thread: thread,
		TagIDs: tagIDs,
	})
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	dto := response.NewThreadResponse(response.NewThreadResponseParams{Thread: thread})

	return dto, nil
}

type ThreadApplicationServiceUpdateParams struct {
	Ctx      context.Context
	ThreadID int
	Body     request.ThreadUpdateRequest
	UserID   int
}

func (svc *ThreadApplicationService) Update(params ThreadApplicationServiceUpdateParams) (*response.ThreadResponse, error) {
	thread, err := svc.threadDatasource.FindById(datasource.ThreadDatasourceFindByIDParams{
		Ctx:      params.Ctx,
		ThreadID: params.ThreadID,
	})
	if err != nil {
		return nil, err
	}

	if thread.EntThread.UserID != params.UserID {
		return nil, errors.New("編集権限がありません。")
	}

	tx, err := svc.client.Tx(params.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	tagList, err := svc.tagDatasource.CreateManyTx(datasource.TagDatasourceCreateManyTxParams{
		Ctx:         params.Ctx,
		Tx:          tx,
		TagNameList: params.Body.TagNameList,
	})
	if err != nil {
		return nil, err
	}

	var tagIDs []int
	for _, tag_i := range tagList {
		tagIDs = append(tagIDs, tag_i.EntTag.ID)
	}

	thread = model.NewThread(model.NewThreadParams{
		EntThread: &ent.Thread{
			ID:           params.ThreadID,
			Description:  params.Body.Description,
			ThumbnailURL: params.Body.ThumbnailURL,
		},
	})

	thread, err = svc.threadDatasource.UpdateTx(datasource.ThreadDatasourceUpdateTxParams{
		Ctx:    params.Ctx,
		Tx:     tx,
		Thread: thread,
		TagIDs: tagIDs,
	})
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	dto := response.NewThreadResponse(response.NewThreadResponseParams{Thread: thread})

	return dto, nil
}

type ThreadApplicationServiceLikeParams struct {
	Ctx      context.Context
	UserID   int
	ThreadID int
}

func (svc *ThreadApplicationService) Like(params ThreadApplicationServiceLikeParams) error {
	err := svc.threadDatasource.Like(datasource.ThreadDatasourceLikeParams{
		Ctx:      params.Ctx,
		UserID:   params.UserID,
		ThreadID: params.ThreadID,
	})
	if err != nil {
		return err
	}

	return nil
}

type ThreadApplicationServiceUnlikeParams struct {
	Ctx      context.Context
	UserID   int
	ThreadID int
}

func (svc *ThreadApplicationService) Unlike(params ThreadApplicationServiceUnlikeParams) error {
	_, err := svc.threadDatasource.Unlike(datasource.ThreadDatasourceUnlikeParams{
		Ctx:      params.Ctx,
		UserID:   params.UserID,
		ThreadID: params.ThreadID,
	})
	if err != nil {
		return err
	}

	return nil
}

type ThreadApplicationServiceDeleteParams struct {
	Ctx      context.Context
	ThreadID int
}

func (svc *ThreadApplicationService) Delete(params ThreadApplicationServiceDeleteParams) error {
	err := svc.threadDatasource.Delete(datasource.ThreadDatasourceDeleteParams{
		Ctx:      params.Ctx,
		ThreadID: params.ThreadID,
	})
	if err != nil {
		return err
	}

	return nil
}
