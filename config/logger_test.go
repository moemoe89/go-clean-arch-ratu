//
//  Simple REST API
//
//  Copyright © 2020. All rights reserved.
//

package config_test

import (
	"simple-go-clean-arch/config"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLog(t *testing.T) {
	log := config.InitLog()

	assert.NotNil(t, log)
}