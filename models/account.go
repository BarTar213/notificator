package models

type Account struct {
	ID    int    `header:"X-Account-Id" binding:"required"`
	Login string `header:"X-Account" binding:"required"`
}
