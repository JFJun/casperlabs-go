package deploy

import cl "github.com/JFJun/casperlabs-go/clvalue"

type ExecutableDeployItem struct {
	moduleBytes                   ModuleBytes
	storedContractByHash          StoredContractByHash
	storedContractByName          StoredContractByName
	storedVersionedContractByHash StoredVersionedContractByHash
	storedVersionedContractByName StoredVersionedContractByName
	transfer                      Transfer
}

type ModuleBytes struct {
	tag         int
	moduleBytes []byte
	args        RuntimeArgs
}

type RuntimeArgs struct {
	args map[string]cl.CLTypedAndToBytes
}

type StoredContractByHash struct {
	tag        int
	hash       []byte
	entryPoint string
	args       RuntimeArgs
}

type StoredContractByName struct {
	tag        int
	hash       []byte
	entryPoint string
	args       RuntimeArgs
}

type StoredVersionedContractByHash struct {
	tag        int
	hash       []byte
	version    int
	entryPoint string
	args       RuntimeArgs
}

type StoredVersionedContractByName struct {
	tag        int
	hash       []byte
	version    int
	entryPoint string
	args       RuntimeArgs
}

type Transfer struct {
	tag  int
	args RuntimeArgs
}

func NewModuleBytes(moduleBytes []byte, args RuntimeArgs) *ModuleBytes {
	return &ModuleBytes{
		moduleBytes: moduleBytes,
		args:        args,
	}
}
