package delivery

import "immersive-dashboard/features/user"

type UserResponse struct {
	ID         uint          `json:"id"`
	Full_Name  string        `json:"full_name"`
	Email      string        `json:"email"`
	Teams      user.CoreTeam `json:"teams"`
	Role       string        `json:"role"`
	Status     string        `json:"status"`
	Permission string        `json:"permission"`
}

func fromCore(dataCore user.CoreUser) UserResponse {
	return UserResponse{
		ID:        dataCore.ID,
		Full_Name: dataCore.Full_Name,
		Email:     dataCore.Email,
		Teams:     dataCore.Teams,
		Role:      dataCore.Role,
		Status:    dataCore.Status,
	}
}

func fromCoreList(dataCore []user.CoreUser) []UserResponse {
	var dataResponse []UserResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}