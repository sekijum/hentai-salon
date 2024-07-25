// Code generated by ent, DO NOT EDIT.

package ent

import (
	"server/infrastructure/ent/board"
	"server/infrastructure/ent/schema"
	"server/infrastructure/ent/tag"
	"server/infrastructure/ent/thread"
	"server/infrastructure/ent/threadcomment"
	"server/infrastructure/ent/threadcommentattachment"
	"server/infrastructure/ent/user"
	"server/infrastructure/ent/usercommentlike"
	"server/infrastructure/ent/usercommentsubscription"
	"server/infrastructure/ent/userthreadlike"
	"server/infrastructure/ent/userthreadsubscription"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	boardFields := schema.Board{}.Fields()
	_ = boardFields
	// boardDescTitle is the schema descriptor for title field.
	boardDescTitle := boardFields[2].Descriptor()
	// board.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	board.TitleValidator = boardDescTitle.Validators[0].(func(string) error)
	// boardDescDescription is the schema descriptor for description field.
	boardDescDescription := boardFields[3].Descriptor()
	// board.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	board.DescriptionValidator = boardDescDescription.Validators[0].(func(string) error)
	// boardDescStatus is the schema descriptor for status field.
	boardDescStatus := boardFields[5].Descriptor()
	// board.DefaultStatus holds the default value on creation for the status field.
	board.DefaultStatus = boardDescStatus.Default.(int)
	// boardDescCreatedAt is the schema descriptor for created_at field.
	boardDescCreatedAt := boardFields[6].Descriptor()
	// board.DefaultCreatedAt holds the default value on creation for the created_at field.
	board.DefaultCreatedAt = boardDescCreatedAt.Default.(func() time.Time)
	// boardDescUpdatedAt is the schema descriptor for updated_at field.
	boardDescUpdatedAt := boardFields[7].Descriptor()
	// board.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	board.DefaultUpdatedAt = boardDescUpdatedAt.Default.(func() time.Time)
	// board.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	board.UpdateDefaultUpdatedAt = boardDescUpdatedAt.UpdateDefault.(func() time.Time)
	tagFields := schema.Tag{}.Fields()
	_ = tagFields
	// tagDescName is the schema descriptor for name field.
	tagDescName := tagFields[1].Descriptor()
	// tag.NameValidator is a validator for the "name" field. It is called by the builders before save.
	tag.NameValidator = tagDescName.Validators[0].(func(string) error)
	// tagDescCreatedAt is the schema descriptor for created_at field.
	tagDescCreatedAt := tagFields[2].Descriptor()
	// tag.DefaultCreatedAt holds the default value on creation for the created_at field.
	tag.DefaultCreatedAt = tagDescCreatedAt.Default.(func() time.Time)
	threadFields := schema.Thread{}.Fields()
	_ = threadFields
	// threadDescTitle is the schema descriptor for title field.
	threadDescTitle := threadFields[3].Descriptor()
	// thread.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	thread.TitleValidator = threadDescTitle.Validators[0].(func(string) error)
	// threadDescDescription is the schema descriptor for description field.
	threadDescDescription := threadFields[4].Descriptor()
	// thread.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	thread.DescriptionValidator = threadDescDescription.Validators[0].(func(string) error)
	// threadDescIPAddress is the schema descriptor for ip_address field.
	threadDescIPAddress := threadFields[6].Descriptor()
	// thread.IPAddressValidator is a validator for the "ip_address" field. It is called by the builders before save.
	thread.IPAddressValidator = threadDescIPAddress.Validators[0].(func(string) error)
	// threadDescStatus is the schema descriptor for status field.
	threadDescStatus := threadFields[7].Descriptor()
	// thread.DefaultStatus holds the default value on creation for the status field.
	thread.DefaultStatus = threadDescStatus.Default.(int)
	// threadDescCreatedAt is the schema descriptor for created_at field.
	threadDescCreatedAt := threadFields[8].Descriptor()
	// thread.DefaultCreatedAt holds the default value on creation for the created_at field.
	thread.DefaultCreatedAt = threadDescCreatedAt.Default.(func() time.Time)
	// threadDescUpdatedAt is the schema descriptor for updated_at field.
	threadDescUpdatedAt := threadFields[9].Descriptor()
	// thread.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	thread.DefaultUpdatedAt = threadDescUpdatedAt.Default.(func() time.Time)
	// thread.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	thread.UpdateDefaultUpdatedAt = threadDescUpdatedAt.UpdateDefault.(func() time.Time)
	threadcommentFields := schema.ThreadComment{}.Fields()
	_ = threadcommentFields
	// threadcommentDescGuestName is the schema descriptor for guest_name field.
	threadcommentDescGuestName := threadcommentFields[4].Descriptor()
	// threadcomment.GuestNameValidator is a validator for the "guest_name" field. It is called by the builders before save.
	threadcomment.GuestNameValidator = threadcommentDescGuestName.Validators[0].(func(string) error)
	// threadcommentDescIPAddress is the schema descriptor for ip_address field.
	threadcommentDescIPAddress := threadcommentFields[6].Descriptor()
	// threadcomment.IPAddressValidator is a validator for the "ip_address" field. It is called by the builders before save.
	threadcomment.IPAddressValidator = threadcommentDescIPAddress.Validators[0].(func(string) error)
	// threadcommentDescStatus is the schema descriptor for status field.
	threadcommentDescStatus := threadcommentFields[7].Descriptor()
	// threadcomment.DefaultStatus holds the default value on creation for the status field.
	threadcomment.DefaultStatus = threadcommentDescStatus.Default.(int)
	// threadcommentDescCreatedAt is the schema descriptor for created_at field.
	threadcommentDescCreatedAt := threadcommentFields[8].Descriptor()
	// threadcomment.DefaultCreatedAt holds the default value on creation for the created_at field.
	threadcomment.DefaultCreatedAt = threadcommentDescCreatedAt.Default.(func() time.Time)
	// threadcommentDescUpdatedAt is the schema descriptor for updated_at field.
	threadcommentDescUpdatedAt := threadcommentFields[9].Descriptor()
	// threadcomment.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	threadcomment.DefaultUpdatedAt = threadcommentDescUpdatedAt.Default.(func() time.Time)
	// threadcomment.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	threadcomment.UpdateDefaultUpdatedAt = threadcommentDescUpdatedAt.UpdateDefault.(func() time.Time)
	threadcommentattachmentFields := schema.ThreadCommentAttachment{}.Fields()
	_ = threadcommentattachmentFields
	// threadcommentattachmentDescDisplayOrder is the schema descriptor for display_order field.
	threadcommentattachmentDescDisplayOrder := threadcommentattachmentFields[3].Descriptor()
	// threadcommentattachment.DefaultDisplayOrder holds the default value on creation for the display_order field.
	threadcommentattachment.DefaultDisplayOrder = threadcommentattachmentDescDisplayOrder.Default.(int)
	// threadcommentattachmentDescType is the schema descriptor for type field.
	threadcommentattachmentDescType := threadcommentattachmentFields[4].Descriptor()
	// threadcommentattachment.DefaultType holds the default value on creation for the type field.
	threadcommentattachment.DefaultType = threadcommentattachmentDescType.Default.(int)
	// threadcommentattachmentDescCreatedAt is the schema descriptor for created_at field.
	threadcommentattachmentDescCreatedAt := threadcommentattachmentFields[5].Descriptor()
	// threadcommentattachment.DefaultCreatedAt holds the default value on creation for the created_at field.
	threadcommentattachment.DefaultCreatedAt = threadcommentattachmentDescCreatedAt.Default.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[1].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = userDescName.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[2].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescStatus is the schema descriptor for status field.
	userDescStatus := userFields[5].Descriptor()
	// user.DefaultStatus holds the default value on creation for the status field.
	user.DefaultStatus = userDescStatus.Default.(int)
	// userDescRole is the schema descriptor for role field.
	userDescRole := userFields[6].Descriptor()
	// user.DefaultRole holds the default value on creation for the role field.
	user.DefaultRole = userDescRole.Default.(int)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[7].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[8].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	usercommentlikeFields := schema.UserCommentLike{}.Fields()
	_ = usercommentlikeFields
	// usercommentlikeDescLikedAt is the schema descriptor for liked_at field.
	usercommentlikeDescLikedAt := usercommentlikeFields[2].Descriptor()
	// usercommentlike.DefaultLikedAt holds the default value on creation for the liked_at field.
	usercommentlike.DefaultLikedAt = usercommentlikeDescLikedAt.Default.(func() time.Time)
	usercommentsubscriptionFields := schema.UserCommentSubscription{}.Fields()
	_ = usercommentsubscriptionFields
	// usercommentsubscriptionDescIsNotified is the schema descriptor for is_notified field.
	usercommentsubscriptionDescIsNotified := usercommentsubscriptionFields[2].Descriptor()
	// usercommentsubscription.DefaultIsNotified holds the default value on creation for the is_notified field.
	usercommentsubscription.DefaultIsNotified = usercommentsubscriptionDescIsNotified.Default.(bool)
	// usercommentsubscriptionDescIsChecked is the schema descriptor for is_checked field.
	usercommentsubscriptionDescIsChecked := usercommentsubscriptionFields[3].Descriptor()
	// usercommentsubscription.DefaultIsChecked holds the default value on creation for the is_checked field.
	usercommentsubscription.DefaultIsChecked = usercommentsubscriptionDescIsChecked.Default.(bool)
	// usercommentsubscriptionDescSubscribedAt is the schema descriptor for subscribed_at field.
	usercommentsubscriptionDescSubscribedAt := usercommentsubscriptionFields[4].Descriptor()
	// usercommentsubscription.DefaultSubscribedAt holds the default value on creation for the subscribed_at field.
	usercommentsubscription.DefaultSubscribedAt = usercommentsubscriptionDescSubscribedAt.Default.(func() time.Time)
	userthreadlikeFields := schema.UserThreadLike{}.Fields()
	_ = userthreadlikeFields
	// userthreadlikeDescLikedAt is the schema descriptor for liked_at field.
	userthreadlikeDescLikedAt := userthreadlikeFields[2].Descriptor()
	// userthreadlike.DefaultLikedAt holds the default value on creation for the liked_at field.
	userthreadlike.DefaultLikedAt = userthreadlikeDescLikedAt.Default.(func() time.Time)
	userthreadsubscriptionFields := schema.UserThreadSubscription{}.Fields()
	_ = userthreadsubscriptionFields
	// userthreadsubscriptionDescIsNotified is the schema descriptor for is_notified field.
	userthreadsubscriptionDescIsNotified := userthreadsubscriptionFields[2].Descriptor()
	// userthreadsubscription.DefaultIsNotified holds the default value on creation for the is_notified field.
	userthreadsubscription.DefaultIsNotified = userthreadsubscriptionDescIsNotified.Default.(bool)
	// userthreadsubscriptionDescIsChecked is the schema descriptor for is_checked field.
	userthreadsubscriptionDescIsChecked := userthreadsubscriptionFields[3].Descriptor()
	// userthreadsubscription.DefaultIsChecked holds the default value on creation for the is_checked field.
	userthreadsubscription.DefaultIsChecked = userthreadsubscriptionDescIsChecked.Default.(bool)
	// userthreadsubscriptionDescSubscribedAt is the schema descriptor for subscribed_at field.
	userthreadsubscriptionDescSubscribedAt := userthreadsubscriptionFields[4].Descriptor()
	// userthreadsubscription.DefaultSubscribedAt holds the default value on creation for the subscribed_at field.
	userthreadsubscription.DefaultSubscribedAt = userthreadsubscriptionDescSubscribedAt.Default.(func() time.Time)
}
