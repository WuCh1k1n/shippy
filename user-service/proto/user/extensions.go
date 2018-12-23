// 我们还需要更改ORM的行为以在创建时生成UUID，而不是尝试生成整数ID。
// UUID是一组随机生成的带连字符的字符串，被用作ID或主键。
// 这比仅使用自动递增ID更安全，因为它可以阻止人们猜测或遍历API端点。
// MongoDB已经使用了这种变体，但我们需要告诉我们的Postgres模型使用UUID。

// 设置与 user.pb.go 的包名相同
package go_micro_srv_user

import (
	"github.com/satori/go.uuid"
	"github.com/jinzhu/gorm"
)

// BeforeCreate - 这将挂钩到GORM的事件生命周期，以便在保存实体之前为Id列生成UUID。
func (model *User) BeforeCreate(scope *gorm.Scope) error {
	
	uuid, err := uuid.NewV4()
	if err != nil {
		return err
	}

	err = scope.SetColumn("Id", uuid.String())
	if err != nil {
		return err
	}

	return nil 
}