package models

import "time"

type ShoppingItem struct {
    ID        string    `json:"id"`
    ListID    string    `json:"list_id"`
    Name      string    `json:"name"`
    Quantity  string    `json:"quantity"`
    IsChecked bool      `json:"is_checked"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type CreateItemInput struct {
    Name     string `json:"name" binding:"required,min=1,max=255"`
    Quantity string `json:"quantity" binding:"max=50"`
}

type UpdateItemInput struct {
    Name      string `json:"name" binding:"omitempty,min=1,max=255"`
    Quantity  string `json:"quantity" binding:"omitempty,max=50"`
    IsChecked *bool  `json:"is_checked"`
}