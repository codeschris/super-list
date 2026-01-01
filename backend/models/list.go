package models

import "time"

type ShoppingList struct {
    ID        string    `json:"id"`
    UserID    string    `json:"user_id"`
    Name      string    `json:"name"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type CreateListInput struct {
    Name string `json:"name" binding:"required,min=1,max=255"`
}

type UpdateListInput struct {
    Name string `json:"name" binding:"required,min=1,max=255"`
}