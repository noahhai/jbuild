package jbuild

import "fmt"

type Jmap map[string]interface{}

type MergeOptions struct {
	ErrorOnKeyConflict bool
}

func (n Jmap) Add(val interface{}, path ...string) {
	if len(path) < 1 {
		return
	}
	key := path[len(path)-1]
	path = path[0 : len(path)-1]
	n.AddMap(Jmap{key: val}, path...)
}

func (n Jmap) AddMap(node Jmap, path ...string) {
	currNode := n
	for _, k := range path {
		currVal := currNode
		untyped, ok := currVal[k]
		var nextNode Jmap
		if !ok || untyped == nil {
			nextNode = Jmap{}
		} else {
			switch v := currVal[k].(type) {
			case Jmap:
				nextNode = v
			default:
				nextNode = Jmap{}
			}
		}

		currVal[k] = nextNode
		currNode = nextNode
	}

	for k, v := range node {
		currNode[k] = v
	}
}

func (node1 Jmap) Merge(node2 Jmap, opt *MergeOptions) error {
	for k, v := range node2 {
		untyped, ok := node1[k]
		if !ok || untyped == nil {
			node1[k] = v
		} else if opt.ErrorOnKeyConflict {
			return fmt.Errorf("conflict on key '%s'", k)
		} else {
			switch t1 := node1[k].(type) {
			case Jmap:
				switch t2 := node2[k].(type) {
				case Jmap:
					// non-trivial case; need to merge here
					if err := t1.Merge(t2, opt); err != nil {
						return err
					}
				default:
					// node2 wins. this is arbitrary
					node1[k] = v
				}
			default:
				// node2 wins. this is arbitrary
				node1[k] = v
			}
		}
	}
	return nil
}
