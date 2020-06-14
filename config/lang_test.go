//
//  Simple REST API
//
//  Copyright © 2020. All rights reserved.
//

package config_test

import (
	"github.com/moemoe89/go-clean-arch-ratu/config"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLang(t *testing.T) {
	_, err := config.InitLang()

	assert.Nil(t, err)
}
