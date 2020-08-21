package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lru_Read(t *testing.T) {

	type initCacheArg struct { // helps to build the initCacheArg
		key             int
		value           int
		callReadOrWrite bool // call Read with key if true else call Write with key and value
	}
	tests := []struct {
		name          string
		initCacheArgs []initCacheArg
		key           int
		returnValue   int
	}{
		{
			name:          "Empty cache",
			initCacheArgs: []initCacheArg{},
			key:           1,
			returnValue:   -1,
		},
		{
			name: "Cache Full and key not present",
			initCacheArgs: []initCacheArg{
				{
					key:             7,
					value:           7,
					callReadOrWrite: false,
				},
				{
					key:             0,
					value:           0,
					callReadOrWrite: false,
				},
				{
					key:             1,
					value:           1,
					callReadOrWrite: false,
				},
			},
			key:         2,
			returnValue: -1,
		},
		{
			name: "Cache Full and key present",
			initCacheArgs: []initCacheArg{
				{
					key:             7,
					value:           7,
					callReadOrWrite: false,
				},
				{
					key:             0,
					value:           0,
					callReadOrWrite: false,
				},
				{
					key:             1,
					value:           1,
					callReadOrWrite: false,
				},
			},
			key:         0,
			returnValue: 0,
		},
		{
			name: "Cache Full, read an element to change lru, and write",
			initCacheArgs: []initCacheArg{
				{
					key:             7,
					value:           7,
					callReadOrWrite: false,
				},
				{
					key:             0,
					value:           0,
					callReadOrWrite: false,
				},
				{
					key:             1,
					value:           1,
					callReadOrWrite: false,
				},
				{
					key:             7,
					value:           7,
					callReadOrWrite: true,
				},
				{
					key:             1,
					value:           1,
					callReadOrWrite: false,
				},
				{
					key:             2,
					value:           2,
					callReadOrWrite: false,
				},
			},
			key:         0,
			returnValue: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var c Cache
			c = NewLRUCache()
			// initialize cache
			for _, ctxt := range tt.initCacheArgs {
				if ctxt.callReadOrWrite {
					c.Read(ctxt.key)
				} else {
					c.Write(ctxt.key, ctxt.value)
				}
			}
			assert.Equal(t, tt.returnValue, c.Read(tt.key), "should be same")
		})
	}
}

func Test_lru_Write(t *testing.T) {

	type initCacheArg struct { // helps to build the initCacheArg
		key             int
		value           int
		callReadOrWrite bool // call Read with key if true else call Write with key and value
	}
	tests := []struct {
		name          string
		initCacheArgs []initCacheArg
		key           int
		value         int
	}{
		{
			name: "Write Success",
			initCacheArgs: []initCacheArg{
				{
					key:             7,
					value:           7,
					callReadOrWrite: false,
				},
				{
					key:             0,
					value:           0,
					callReadOrWrite: false,
				},
				{
					key:             1,
					value:           1,
					callReadOrWrite: false,
				},
			},
			key: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var c Cache
			c = NewLRUCache()
			// initialize cache
			for _, ctxt := range tt.initCacheArgs {
				if ctxt.callReadOrWrite {
					c.Read(ctxt.key)
				} else {
					c.Write(ctxt.key, ctxt.value)
				}
			}
			c.Write(tt.key, tt.value)
		})
	}
}
