package models

type Migration struct {
	BaseModel
	ID        uint   `gorm:"column:id;primary_key"`
	Migration string `gorm:"column:migration;type:varchar(255)"`
	Batch     int    `gorm:"column:batch;"`
}

func (*Migration) GetObjArr()         {} //@todo     public function getObjArr(?array $filter_arr = [], ?array $sort_arr = null, ?int $limit = null, bool $with_trashed = false): Collection;
func (*Migration) GetObjArrPaginate() {} //@todo     public function getObjArrPaginate(int $per_page, ?array $filter_arr = [], ?array $sort_arr = null, bool $with_trashed = false): LengthAwarePaginator;
