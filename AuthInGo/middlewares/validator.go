package middlewares

import (
	"AuthInGo/dto"
	"AuthInGo/utils"
	"context"
	"net/http"
)

func LoginRequestValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload dto.LoginUserRequestDto

		if jsonErr := utils.ReadJsonBody(r, &payload); jsonErr != nil {
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Something went wrong while logging in", jsonErr)
			return
		}

		if validationErr := utils.Validator.Struct(payload); validationErr != nil {
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid Input Data", validationErr)
			return
		}

		ctx := context.WithValue(
			r.Context(),
			"LoginUserPayload",
			payload,
		)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func CreateUserRequestValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload dto.CreateUserRequestDto

		if jsonErr := utils.ReadJsonBody(r, &payload); jsonErr != nil {
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Something went wrong while creating user", jsonErr)
			return
		}

		if validationErr := utils.Validator.Struct(payload); validationErr != nil {
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid Input Data", validationErr)
			return
		}

		ctx := context.WithValue(
			r.Context(),
			"CreateUserPayload",
			payload,
		)

		next.ServeHTTP(w, r.WithContext(ctx))
	})

}

func CreateRoleRequestValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload dto.CreateRoleRequestDto

		// Read and decode the JSON body into the payload
		if err := utils.ReadJsonBody(r, &payload); err != nil {
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid request body", err)
			return
		}

		// Validate the payload using the Validator instance
		if err := utils.Validator.Struct(payload); err != nil {
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Validation failed", err)
			return
		}

		ctx := context.WithValue(r.Context(), "payload", payload)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func UpdateRoleRequestValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload dto.UpdateRoleRequestDto

		// Read and decode the JSON body into the payload
		if err := utils.ReadJsonBody(r, &payload); err != nil {
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid request body", err)
			return
		}

		// Validate the payload using the Validator instance
		if err := utils.Validator.Struct(payload); err != nil {
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Validation failed", err)
			return
		}

		ctx := context.WithValue(r.Context(), "payload", payload)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func AssignPermissionRequestValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload dto.AssignPermissionRequestDto

		// Read and decode the JSON body into the payload
		if err := utils.ReadJsonBody(r, &payload); err != nil {
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid request body", err)
			return
		}

		// Validate the payload using the Validator instance
		if err := utils.Validator.Struct(payload); err != nil {
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Validation failed", err)
			return
		}

		ctx := context.WithValue(r.Context(), "payload", payload)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func RemovePermissionRequestValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload dto.RemovePermissionRequestDto

		// Read and decode the JSON body into the payload
		if err := utils.ReadJsonBody(r, &payload); err != nil {
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid request body", err)
			return
		}

		// Validate the payload using the Validator instance
		if err := utils.Validator.Struct(payload); err != nil {
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Validation failed", err)
			return
		}

		ctx := context.WithValue(r.Context(), "payload", payload)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
