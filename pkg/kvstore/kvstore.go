//
// Copyright 2016 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
package kvstore

import (
	"encoding/json"
	"time"

	"github.com/cilium/cilium/common/types"

	"github.com/op/go-logging"
)

var (
	log = logging.MustGetLogger("cilium-net")
)

type KVClient interface {
	LockPath(path string) (KVLocker, error)
	GetValue(k string) (json.RawMessage, error)
	SetValue(k string, v interface{}) error
	InitializeFreeID(path string, firstID uint32) error
	GetMaxID(key string, firstID uint32) (uint32, error)
	SetMaxID(key string, firstID, maxID uint32) error

	GASNewSecLabelID(baseKeyPath string, baseID uint32, secCtxLabels *types.SecCtxLabel) error
	GASNewServiceL4ID(basePath string, baseID uint32, sl4 *types.ServiceL4ID) error

	DeleteTree(path string) error

	GetWatcher(key string, timeSleep time.Duration) <-chan []uint32

	Status() (string, error)
}

type KVLocker interface {
	Unlock() error
}
