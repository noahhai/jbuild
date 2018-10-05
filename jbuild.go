package jbuild

type Jmap map[string]interface{}

type MergeOptions struct {
	ErrorOnKeyConflict bool
}

func (n *Jmap) Add(val string, path ...string) {
	if len(path) < 1 {
		return
	}
	key := path[len(path)-1]
	path = path[0 : len(path)-1]
	n.AddMap(Jmap{key: val}, path...)
}

func (n *Jmap) AddMap(node Jmap, path ...string) {
	currNode := n
	for _, k := range path {
		currVal := *currNode
		untyped, ok := currVal[k]
		var nextNode Jmap
		if !ok || untyped == nil {
			nextNode = Jmap{}
		}
		switch v := currVal[k].(type) {
		case Jmap:
			nextNode = v
		default:
			nextNode = Jmap{}
		}
		currVal[k] = nextNode
		currNode = &nextNode
	}

	for k, v := range node {
		(*currNode)[k] = v
	}
}

// // TODO
// func (n *Jmap) Merge(node Jmap, opt *MergeOptions) {

// }
