package storage

import "github.com/DrSKPunk/JobBoard/internal/models"

var Users = []models.User{}
var NextUserID = 1
var Tokens = map[string]int{}
