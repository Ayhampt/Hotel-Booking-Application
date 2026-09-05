package controllers

import (
	"AuthInGo/dto"
	"AuthInGo/services"
	"AuthInGo/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type RoleController struct {
	RoleService services.RoleService
}

func NewRoleController(_roleService services.RoleService) *RoleController {
	return &RoleController{
		RoleService: _roleService,
	}
}

func (rc *RoleController) AssignRoleToUser(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")
	roleId := chi.URLParam(r, "roleId")
	if userId == "" || roleId == "" {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "User ID and Role ID are required", nil)
		return
	}
	roleIdInt, err := strconv.ParseInt(roleId, 10, 64)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid roleId", err)
		return
	}
	userIdInt, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid userId", err)
		return
	}
	err = rc.RoleService.AssignRoleToUser(userIdInt, roleIdInt)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Failed to assign role to user", err)
		return
	}
	utils.WriteJsonSuccessResponse(w, http.StatusOK, "Role assigned to user successfully", nil)
}

func (rc *RoleController) GetRoleById(w http.ResponseWriter, r *http.Request) {
	roleId := chi.URLParam(r, "id")
	if roleId == "" {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Role ID is required", nil)
		return
	}
	roleIdInt, err := strconv.ParseInt(roleId, 10, 64)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid roleId", err)
		return
	}
	role, err := rc.RoleService.GetRoleById(roleIdInt)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Failed to fetch role", err)
		return
	}
	if role == nil {
		utils.WriteJsonErrorResponse(w, http.StatusNotFound, "Role not found", nil)
		return
	}
	utils.WriteJsonSuccessResponse(w, http.StatusOK, "Role fetched successfully", role)
}

func (rc *RoleController) GetAllRoles(w http.ResponseWriter, r *http.Request) {
	roles, err := rc.RoleService.GetAllRoles()
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Failed to fetch roles", err)
		return
	}
	utils.WriteJsonSuccessResponse(w, http.StatusOK, "Roles fetched successfully", roles)
}

func (rc *RoleController) CreateRole(w http.ResponseWriter, r *http.Request) {
	payload := r.Context().Value("payload").(dto.CreateRoleRequestDto)
	role, err := rc.RoleService.CreateRole(payload.Name, payload.Description)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Failed to create role", err)
		return
	}
	utils.WriteJsonSuccessResponse(w, http.StatusOK, "Role created successfully", role)
}

func (rc *RoleController) UpdateRole(w http.ResponseWriter, r *http.Request) {
	roleId := chi.URLParam(r, "id")
	if roleId == "" {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Role ID is required", nil)
		return
	}
	roleIdInt, err := strconv.ParseInt(roleId, 10, 64)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid roleId", err)
		return
	}
	payload := r.Context().Value("payload").(dto.UpdateRoleRequestDto)
	role, err := rc.RoleService.UpdateRole(roleIdInt, payload.Name, payload.Description)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Failed to update role", err)
		return
	}
	utils.WriteJsonSuccessResponse(w, http.StatusOK, "Role updated successfully", role)
}

func (rc *RoleController) DeleteRole(w http.ResponseWriter, r *http.Request) {
	roleId := chi.URLParam(r, "id")
	if roleId == "" {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Role ID is required", nil)
		return
	}
	roleIdInt, err := strconv.ParseInt(roleId, 10, 64)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid roleId", err)
		return
	}
	err = rc.RoleService.DeleteRoleById(roleIdInt)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Failed to delete role", err)
		return
	}
	utils.WriteJsonSuccessResponse(w, http.StatusOK, "Role deleted successfully", nil)
}

func (rc *RoleController) GetRolePermissions(w http.ResponseWriter, r *http.Request) {
	roleId := chi.URLParam(r, "id")
	if roleId == "" {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Role ID is required", nil)
		return
	}
	roleIdInt, err := strconv.ParseInt(roleId, 10, 64)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid roleId", err)
		return
	}
	rolePermissions, err := rc.RoleService.GetRolePermissions(roleIdInt)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Failed to fetch role permissions", err)
		return
	}
	utils.WriteJsonSuccessResponse(w, http.StatusOK, "Role permissions fetched successfully", rolePermissions)
}

func (rc *RoleController) AssignPermissionToRole(w http.ResponseWriter, r *http.Request) {
	roleId := chi.URLParam(r, "id")
	if roleId == "" {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Role ID is required", fmt.Errorf("missing role Id"))
		return
	}
	roleIdInt, err := strconv.ParseInt(roleId, 10, 64)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid roleId", err)
		return
	}
	payload := r.Context().Value("payload").(dto.AssignPermissionRequestDto)
	rolePermissions, err := rc.RoleService.AddPermissionToRole(roleIdInt, payload.PermissionId)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Failed to assign permission to role", err)
	}
	utils.WriteJsonSuccessResponse(w, http.StatusCreated, "Permission assigned to role successfully", rolePermissions)
}

func (rc *RoleController) RemovePermissionFromRole(w http.ResponseWriter, r *http.Request) {
	roleId := chi.URLParam(r, "id")
	if roleId == "" {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Role ID is required", fmt.Errorf("missing role ID"))
		return
	}

	roleIdInt, err := strconv.ParseInt(roleId, 10, 64)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid role ID", err)
		return
	}

	payload := r.Context().Value("payload").(dto.RemovePermissionRequestDto)

	err = rc.RoleService.RemovePermissionFromRole(roleIdInt, payload.PermissionId)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Failed to remove permission from role", err)
		return
	}

	utils.WriteJsonSuccessResponse(w, http.StatusOK, "Permission removed from role successfully", nil)
}

func (rc *RoleController) GetAllRolePermissions(w http.ResponseWriter, r *http.Request) {
	rolePermissions, err := rc.RoleService.GetAllRolePermissions()
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Failed to fetch all role permissions", err)
		return
	}

	utils.WriteJsonSuccessResponse(w, http.StatusOK, "All role permissions fetched successfully", rolePermissions)
}
