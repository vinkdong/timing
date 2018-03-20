package common

import "testing"

func TestCheckExistTemplate(t *testing.T) {
	origin := "this is a test  vk.list{1..2} and this is vk.random{1..3} "
	RegisterTemplate(origin)
}