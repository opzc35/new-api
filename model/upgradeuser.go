//升级用户等级

import (
    "github.com/QuantumNous/new-api/common"
)

// UpdateUserGroup 将用户的 group 字段更新为指定值，并更新缓存（如果有缓存逻辑）
func UpdateUserGroup(userId int, newGroup string) error {
    if err := DB.Model(&User{}).Where("id = ?", userId).Update("`group`", newGroup).Error; err != nil {
        common.SysLog("failed to update user group: " + err.Error())
        return err
    }
    return nil
}