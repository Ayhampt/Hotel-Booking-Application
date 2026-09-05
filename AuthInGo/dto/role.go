package dto

type CreateRoleRequestDto struct {
	Name        string `json:"name" validate:"required,min=2,max=50"`
	Description string `json:"description" validate:"required,min=5,max=200"`
}

type UpdateRoleRequestDto struct {
	Name        string `json:"name" validate:"required,min=2,max=50"`
	Description string `json:"description" validate:"required,min=5,max=200"`
}

type AssignPermissionRequestDto struct {
	PermissionId int64 `json:"permission_id" validate:"required"`
}

type RemovePermissionRequestDto struct {
	PermissionId int64 `json:"permission_id" validate:"required"`
}
