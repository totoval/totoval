package models

import (
	"github.com/totoval/framework/helpers/zone"
	"github.com/totoval/framework/request"

	"github.com/totoval/framework/helpers/m"
	"github.com/totoval/framework/helpers/pb"
	"github.com/totoval/framework/model"
	"github.com/totoval/framework/queue"
	message "github.com/totoval/framework/queue/protocol_buffers"
)

type FailedQueue struct {
	ID *uint `gorm:"column:failed_queue_id;primary_key;auto_increment"`

	Hash      *string        `gorm:"column:failed_queue_hash;type:varchar(100);unique_index;not null"`
	Topic     *string        `gorm:"column:failed_queue_topic_name;type:varchar(100);not null"`
	Channel   *string        `gorm:"column:failed_queue_channel_name;type:varchar(100);not null"`
	DataProto *[]byte        `gorm:"column:failed_queue_data;type:varbinary"`
	PushedAt  *zone.Time     `gorm:"column:failed_queue_pushed_at;not null"`
	Delay     *zone.Duration `gorm:"column:failed_queue_delay;type:bigint;not null"`
	Retries   *uint32        `gorm:"column:failed_queue_retries;type:integer unsigned;not null"`
	Tried     *uint32        `gorm:"column:failed_queue_tried;type:integer unsigned;not null"`
	Err       *string        `gorm:"column:failed_queue_err;size:65535"`

	CreatedAt *zone.Time `gorm:"column:failed_queue_created_at"`
	UpdatedAt zone.Time  `gorm:"column:failed_queue_updated_at"`
	DeletedAt *zone.Time `gorm:"column:failed_queue_deleted_at"`
	model.BaseModel
}

func (fq *FailedQueue) TableName() string {
	return fq.SetTableName("failed_queue")
}
func (fq *FailedQueue) RetryTopic() string {
	return *fq.Topic
}

func (fq *FailedQueue) RetryChannel() string {
	return *fq.Channel
}

func (fq *FailedQueue) RetryRetries() uint32 {
	return *fq.Retries
}

func (fq *FailedQueue) RetryDelay() zone.Duration {
	return *fq.Delay
}

func (fq *FailedQueue) RetryParamProtoBytes() []byte {
	return *fq.DataProto
}
func (fq *FailedQueue) RetryHash() string {
	return *fq.Hash
}

func (fq *FailedQueue) Default() interface{} {
	return FailedQueue{}
}

func (fq *FailedQueue) FailedToDatabase(topicName string, channelName string, msg *message.Message, errStr string) error {
	_fq := FailedQueue{
		Topic:     &topicName,
		Channel:   &channelName,
		Hash:      &msg.Hash,
		DataProto: &msg.Param,
		PushedAt:  pb.TimestampConvert(msg.PushedAt),
		Delay:     pb.DurationConvert(msg.Delay),
		Retries:   &msg.Retries,
		Tried:     &msg.Tried,
		Err:       &errStr,
	}
	return m.H().Create(&_fq)
}

//type retryError byte
//
//const (
//	RETRY_NOT_FOUND retryError = iota
//	RETRY_DELETE_FAILED
//	RETRY_PUSH_FAILED
//)
//
//func (r retryError) Error() string {
//	switch r {
//	case RETRY_NOT_FOUND:
//		return "retry queue not found"
//	case RETRY_DELETE_FAILED:
//		return "retry queue delete failed"
//	case RETRY_PUSH_FAILED:
//	default:
//		return "retry queue repush failed"
//	}
//	return "retry error parse error"
//}
func (fq *FailedQueue) FailedQueueById(id uint) (failedQueuerPtr queue.FailedQueuer, err error) {
	_fq := FailedQueue{ID: &id}
	if err := m.H().First(&_fq, false); err != nil {
		return nil, err
	}
	return &_fq, nil
}
func (fq *FailedQueue) DeleteQueueById(id uint) error {
	_fq := FailedQueue{ID: &id}
	if err := m.H().Delete(&_fq, true); err != nil {
		return err
	}
	return nil
}

func (fq *FailedQueue) DeleteAll() error {
	//@todo need to be down for queue flush
	panic("need implements")
}

func (fq *FailedQueue) ObjArr(filterArr []model.Filter, sortArr []model.Sort, limit int, withTrashed bool) (interface{}, error) {
	var outArr []FailedQueue
	if err := m.H().Q(filterArr, sortArr, limit, withTrashed).Find(&outArr).Error; err != nil {
		return nil, err
	}
	return outArr, nil
}
func (fq *FailedQueue) ObjArrPaginate(c *request.Context, perPage uint, filterArr []model.Filter, sortArr []model.Sort, limit int, withTrashed bool) (pagination model.Pagination, err error) {
	var outArr []FailedQueue
	filter := model.Model(*m.H().Q(filterArr, sortArr, limit, withTrashed))
	return filter.Paginate(&outArr, c, perPage)
}
