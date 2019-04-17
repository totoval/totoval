package models

import (
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/totoval/framework/helpers/debug"

	"github.com/totoval/framework/helpers/ptr"

	"github.com/totoval/framework/model/helper"

	"github.com/totoval/framework/helpers/m"
	"github.com/totoval/framework/model"
)

const AFFILIATION_CODE_LENGTH uint = 6

type UserAffiliation struct {
	UserID    *uint      `gorm:"column:user_id;primary_key;type:int unsigned"`
	Code      *string    `gorm:"column:uaff_code;type:varchar(32);not null"`
	FromCode  *string    `gorm:"column:uaff_from_code;type:varchar(32)"`
	Root      *uint      `gorm:"column:uaff_root_id;type:int unsigned"`
	Parent    *uint      `gorm:"column:uaff_parent_id;type:int unsigned"`
	Left      *uint      `gorm:"column:uaff_left_id;type:int unsigned"`
	Right     *uint      `gorm:"column:uaff_right_id;type:int unsigned"`
	Level     *uint      `gorm:"column:uaff_level;type:int unsigned"`
	CreatedAt *time.Time `gorm:"column:user_created_at"`
	UpdatedAt time.Time  `gorm:"column:user_updated_at"`
	DeletedAt *time.Time `gorm:"column:user_deleted_at"`
	model.BaseModel
}

func (uaff *UserAffiliation) TableName() string {
	return uaff.SetTableName("user_affiliations")
}

func (uaff *UserAffiliation) Default() interface{} {
	return UserAffiliation{}
}

func (uaff *UserAffiliation) InsertRoot() {
	// 'uaff_user_id' => $root_user_id,
	//     'uaff_root_id' => $root_user_id,
	//     'uaff_level' => 1,
	//     'uaff_left' => 1,
	//     'uaff_right' => 2
}

func (uaff *UserAffiliation) generateCode(user *User) (string, error) {
	if float64(*user.ID) > math.Pow(16, float64(AFFILIATION_CODE_LENGTH)) {
		return "", errors.New("userID excceed the max affiliation length")
	}

	affiliationLen := fmt.Sprintf("%d", AFFILIATION_CODE_LENGTH)
	return fmt.Sprintf("%0"+affiliationLen+"x", *user.ID), nil
}

func (uaff *UserAffiliation) InsertNode(user *User, fromCode ...string) error {
	if user.ID == nil {
		return errors.New("user must be existed")
	}

	var fromCodePtr *string
	if len(fromCode) > 0 {
		fromCodePtr = &fromCode[0]
	}

	// insert tree node
	m.Transaction(func(TransactionHelper *helper.Helper) {
		// define affiliation to be inserting
		code, err := uaff.generateCode(user)
		if err != nil {
			panic(err)
		}
		current := UserAffiliation{
			UserID:   user.ID,
			Code:     &code,
			FromCode: fromCodePtr,
		}

		// new tree
		if current.FromCode == nil {
			// no parent
			current.Root = current.UserID
			current.Parent = nil
			current.Level = ptr.Uint(1)
			current.Left = ptr.Uint(1)
			current.Right = ptr.Uint(2)
			if err := TransactionHelper.Create(&current); err != nil {
				panic(err)
			}
			return
		}

		// existed tree
		// lock table
		//@todo switch the DB sql by DB driver
		TransactionHelper.DB().Raw("LOCK TABLES ? WRITE", uaff.TableName())
		// unlock table
		defer TransactionHelper.DB().Raw("UNLOCK TABLES")

		// 1. get parent node
		parent := UserAffiliation{
			Code: fromCodePtr,
		}
		if !TransactionHelper.Exist(&parent, false) {
			panic(errors.New("affiliation code not found"))
		}

		//@todo 2. this.level = parent.level + 1, this.root = parent.root
		// current.Root = parent.Root
		// current.Parent = parent.UserID
		// current.Level = ptr.Uint(*parent.Level + 1)

		//@todo 3. update other nodes
		//@todo 3.1 update left: if other.root == parent.root && other.id != this.id && other.left >= parent.right, other.left += 2
		if err := TransactionHelper.Q([]model.Filter{
			{Key: "uaff_root_id", Op: "=", Val: parent.Root},
			{Key: "user_id", Op: "!=", Val: current.UserID},
			{Key: "uaff_left_id", Op: ">=", Val: parent.Right},
		}, []model.Sort{}, 0, false).Model(UserAffiliation{}).UpdateColumn("uaff_left_id", gorm.Expr("uaff_left_id + ?", 2)).Error; err != nil {
			panic(err)
		}
		//@todo 3.2 update right: if other.root == parent.root && other.id != this.id && other.right  >= parent.right, other.right += 2 ,  here must consider should parent.right+2
		if err := TransactionHelper.Q([]model.Filter{
			{Key: "uaff_root_id", Op: "=", Val: parent.Root},
			{Key: "user_id", Op: "!=", Val: current.UserID},
			{Key: "uaff_right_id", Op: ">=", Val: parent.Right},
		}, []model.Sort{}, 0, false).Model(UserAffiliation{}).UpdateColumn("uaff_right_id", gorm.Expr("uaff_right_id + ?", 2)).Error; err != nil {
			panic(err)
		}

		current.Root = parent.Root
		current.Parent = parent.UserID
		current.Level = ptr.Uint(*parent.Level + 1)

		current.Left = parent.Right
		current.Right = ptr.Uint(*parent.Right + 1)
		if err := TransactionHelper.Create(&current); err != nil {
			panic(err)
		}
		return

	}, 1)
	return nil

	//@todo 3. update other nodes
	//@todo 3.1 update left: if other.root == parent.root && other.id != this.id && other.left >= parent.right, other.left += 2
	//@todo 3.2 update right: if other.root == parent.root && other.id != this.id && other.right  >= parent.right, other.right += 2 ,  here must consider should parent.right+2

	//@todo 4. this.left = parent.right - 2, this.right = parent.right -1

	//@todo lock table
}

func (uaff *UserAffiliation) Tree(rootID uint) ([]UserAffiliation, error) {
	root := UserAffiliation{
		UserID: &rootID,
	}
	if err := m.H().First(&root, false); err != nil {
		return []UserAffiliation{}, err
	}

	nodes, err := uaff.ObjArr([]model.Filter{
		{Key: "uaff_root_id", Op: "=", Val: *root.UserID},
		{Key: "uaff_left_id", Op: ">", Val: *root.Left},
		{Key: "uaff_right_id", Op: "<", Val: *root.Right},
	}, []model.Sort{
		{Key: "uaff_left_id", Direction: model.ASC},
	}, 0, false)

	if err != nil {
		return []UserAffiliation{}, err
	}

	return nodes.([]UserAffiliation), nil
}

func (uaff *UserAffiliation) TreeByParent(parentID uint) ([]UserAffiliation, error) {
	parent := UserAffiliation{
		UserID: &parentID,
	}
	if err := m.H().First(&parent, false); err != nil {
		return []UserAffiliation{}, err
	}

	nodes, err := uaff.ObjArr([]model.Filter{
		{Key: "uaff_root_id", Op: "=", Val: *parent.Root},
		{Key: "uaff_left_id", Op: ">", Val: *parent.Left},
		{Key: "uaff_right_id", Op: "<", Val: *parent.Right},
	}, []model.Sort{
		{Key: "uaff_left_id", Direction: model.ASC},
	}, 0, false)

	if err != nil {
		return []UserAffiliation{}, err
	}

	return nodes.([]UserAffiliation), nil
}

type Tree struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Value    string `json:"value"`
	Children []Tree `json:"children"`
}

func (uaff *UserAffiliation) All() Tree {
	rootNodes, err := uaff.ObjArr([]model.Filter{
		{Key: "uaff_parent_id", Op: "is_null"},
	}, []model.Sort{
		{Key: "uaff_left_id", Direction: model.ASC},
	}, 0, false)
	if err != nil {
		return Tree{}
	}

	r := Tree{ID: 0, Children: []Tree{}, Name: "root", Value: "root value"}
	for _, val := range rootNodes.([]UserAffiliation) {
		current := Tree{ID: *val.UserID, Children: []Tree{}, Name: *val.Code, Value: *val.Code}
		nodes, err := uaff.Tree(*val.UserID)
		if err != nil {
			debug.Dump(err)
			continue
		}

		current.Children = r.recursiveCombineTree(current, 1, nodes)

		r.Children = append(r.Children, current)
	}

	return r
}

func (t *Tree) recursiveCombineTree(current Tree, level uint, nodes []UserAffiliation) []Tree {
	for _, uaff := range nodes {
		if *uaff.Level < level+1 {
			continue
		}

		if *uaff.Level > level+1 {
			continue
		}

		if current.ID == *uaff.Parent {
			_current := Tree{
				ID:       *uaff.UserID,
				Children: []Tree{},
				Name:     *uaff.Code, Value: *uaff.Code,
			}
			_current.Children = t.recursiveCombineTree(_current, level+1, nodes)

			current.Children = append(current.Children, _current)
			continue
		}
	}

	return current.Children
}

func (uaff *UserAffiliation) ObjArr(filterArr []model.Filter, sortArr []model.Sort, limit int, withTrashed bool) (interface{}, error) {
	var outArr []UserAffiliation
	if err := m.H().Q(filterArr, sortArr, limit, withTrashed).Find(&outArr).Error; err != nil {
		return nil, err
	}
	return outArr, nil
}
func (uaff *UserAffiliation) ObjArrPaginate(c *gin.Context, perPage uint, filterArr []model.Filter, sortArr []model.Sort, limit int, withTrashed bool) (pagination model.Pagination, err error) {
	var outArr []UserAffiliation
	filter := model.Model(*m.H().Q(filterArr, sortArr, limit, withTrashed))
	return filter.Paginate(&outArr, c, perPage)
}
