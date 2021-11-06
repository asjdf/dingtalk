package dingtalk

import (
	"errors"
	"net/http"
	"strings"
)

import (
	"github.com/zhaoyunxing92/dingtalk/v2/constant"
	"github.com/zhaoyunxing92/dingtalk/v2/model"
	"github.com/zhaoyunxing92/dingtalk/v2/request"
	"github.com/zhaoyunxing92/dingtalk/v2/response"
)

//GetRoleUserList:获取指定角色的员工列表
//roleId:角色Id
//offset:支持分页查询，与size参数同时设置时才生效，此参数代表偏移量，偏移量从0开始
//size:支持分页查询，与offset参数同时设置时才生效，此参数代表分页大小，默认值20，最大值200。
//
//{
//    "errcode": 0,
//    "result": {
//        "hasMore": false,
//        "list": [
//            {
//                "name": "xx",
//                "userid": "xx"
//            },
//            {
//                "name": "xx",
//                "userid": "xx"
//            }
//        ]
//    },
//    "request_id": "y5y017a37yhd"
//}
func (ding *dingTalk) GetRoleUserList(roleId, offset, size int) (apps model.RoleUserListResponse, err error) {
	if size > 200 || size < 0 {
		size = 200
	}
	form := map[string]int{
		"role_id": roleId,
		"offset":  offset,
		"size":    size,
	}
	err = ding.Request(http.MethodPost, constant.GetRoleUserListKey, nil, form, &apps)
	return apps, err
}

//GetRoleGroup:获取角色组
//groupId:组id
func (ding *dingTalk) GetRoleGroup(groupId int) (apps model.RoleGroupResponse, err error) {

	form := map[string]int{
		"group_id": groupId,
	}
	err = ding.Request(http.MethodPost, constant.GetRoleGroupKey, nil, form, &apps)
	return apps, err
}

//GetRoleDetail:获取角色详情
//todo：如果你传的是角色组id，那么会返回角色组的信息
func (ding *dingTalk) GetRoleDetail(roleId int) (apps model.RoleDetailResponse, err error) {

	form := map[string]int{
		"roleId": roleId,
	}
	err = ding.Request(http.MethodPost, constant.GetRoleDetailKey, nil, form, &apps)
	return apps, err
}

//DeleteRole:删除角色
// todo:如果传入的是角色组id则会删除角色组
//roleId:角色id
func (ding *dingTalk) DeleteRole(roleId int) (apps model.Response, err error) {

	form := map[string]int{
		"role_id": roleId,
	}
	err = ding.Request(http.MethodPost, constant.DeleteRoleKey, nil, form, &apps)
	return apps, err
}

//RoleBatchRemoveUser:批量删除员工角色
//roleIds:角色roleId列表，最多可传20个。
//userIds:员工的userId,，最多可传20个。
func (ding *dingTalk) RoleBatchRemoveUser(roleIds []string, userIds []string) (apps model.Response, err error) {

	if len(roleIds) > 20 {
		err = errors.New("一次最多20个角色")
	}

	if len(userIds) > 20 {
		err = errors.New("一次最多20个用户")
	}

	var form = map[string]string{
		"userIds": strings.Join(userIds, ","),
		"roleIds": strings.Join(roleIds, ","),
	}
	err = ding.Request(http.MethodPost, constant.RoleBatchRemoveUserKey, nil, form, &apps)
	return apps, err
}

//RoleUpdateUserManageScope:设定角色成员管理范围
//userId:用户id
//roleId:角色id
//deptIds:部门ID列表数。最多50个，不传则设置范围为所有人
func (ding *dingTalk) RoleUpdateUserManageScope(userId string, roleId int, deptIds []int) (apps model.Response, err error) {

	if deptIds != nil && len(deptIds) > 50 {
		err = errors.New("最多50个部门")
	}

	var form = map[string]interface{}{
		"userid":  userId,
		"role_id": roleId,
	}

	if deptIds != nil {
		form["dept_ids"] = deptIds
	}
	err = ding.Request(http.MethodPost, constant.RoleUpdateUserManageScopeKey, nil, form, &apps)
	return apps, err
}

//CreateRole 创建角色
func (ding *dingTalk) CreateRole(name string, groupId int) (apps response.CreateRole, err error) {

	return apps, ding.Request(http.MethodPost, constant.CreateRoleKey, nil,
		request.NewCreateRole(name, groupId), &apps)
}

//CreateRoleGroup 创建角色组
func (ding *dingTalk) CreateRoleGroup(name string) (apps response.CreateRoleGroup, err error) {

	return apps, ding.Request(http.MethodPost, constant.CreateRoleGroupKey, nil,
		request.NewCreateRoleGroup(name), &apps)
}

//UpdateRole 更新角色
func (ding *dingTalk) UpdateRole(id int, name string) (apps response.Response, err error) {

	return apps, ding.Request(http.MethodPost, constant.UpdateRoleKey, nil,
		request.NewUpdateRole(id, name), &apps)
}

//RoleAddUser  批量增加员工角色
func (ding *dingTalk) RoleAddUser(rs []int, us []string) (apps response.Response, err error) {

	return apps, ding.Request(http.MethodPost, constant.RoleBatchAddUserKey, nil,
		request.NewRoleAddUser(rs, us), &apps)
}

//GetRoleList 获取角色列表
func (ding *dingTalk) GetRoleList(offset, size int) (apps response.RoleList, err error) {

	return apps, ding.Request(http.MethodPost, constant.GetRoleListKey, nil,
		request.NewRoleList(offset, size), &apps)
}
