package core

import (
	"fmt"

	b3 "github.com/liguoqinjim/behavior3go"
	"github.com/liguoqinjim/behavior3go/config"
)

/**
 * The BehaviorTree class, as the name implies, represents the Behavior Tree
 * structure.
 *
 * There are two ways to construct a Behavior Tree: by manually setting the
 * root node, or by loading it from a data structure (which can be loaded
 * from a JSON). Both methods are shown in the examples below and better
 * explained in the user guide.
 *
 * The tick method must be called periodically, in order to send the tick
 * signal to all nodes in the tree, starting from the root. The method
 * `BehaviorTree.tick` receives a target object and a blackboard as
 * parameters. The target object can be anything: a game agent, a system, a
 * DOM object, etc. This target is not used by any piece of Behavior3JS,
 * i.e., the target object will only be used by custom nodes.
 *
 * The blackboard is obligatory and must be an instance of `Blackboard`. This
 * requirement is necessary due to the fact that neither `BehaviorTree` or
 * any node will store the execution variables in its own object (e.g., the
 * BT does not store the target, information about opened nodes or number of
 * times the tree was called). But because of this, you only need a single
 * tree instance to control multiple (maybe hundreds) objects.
**/
type BehaviorTree struct { //这里很重要，也就是说，行为树只要有一个实例就可以了，就可以控制多个使用相同行为树的对象
	id          string
	title       string
	description string
	properties  map[string]interface{}

	root IBaseNode // The reference to the root node. Must be an instance of `b3.BaseNode`.

	debug    interface{} //The reference to the debug instance.
	dumpInfo *config.BTTreeCfg
}

func NewBeTree() *BehaviorTree {
	tree := &BehaviorTree{}
	tree.Initialize()
	return tree
}

func (this *BehaviorTree) Initialize() {
	this.id = b3.CreateUUID()
	this.title = "The behavior tree"
	this.description = "Default description"
	this.properties = make(map[string]interface{})
	this.root = nil
	this.debug = nil
}

func (this *BehaviorTree) GetID() string {
	return this.id
}

func (this *BehaviorTree) SetDebug(debug interface{}) {
	this.debug = debug
}

func (this *BehaviorTree) Load(data *config.BTTreeCfg, maps *b3.RegisterStructMaps, extMaps *b3.RegisterStructMaps) { //This method loads a Behavior Tree from a data structure
	this.title = data.Title             // || this.title;
	this.description = data.Description // || this.description;
	this.properties = data.Properties   // || this.properties;
	this.dumpInfo = data
	nodes := make(map[string]IBaseNode)

	// Create the node list (without connection between them)
	for id, s := range data.Nodes {
		spec := &s
		var node IBaseNode
		if extMaps != nil && extMaps.CheckElem(spec.Name) {
			// Look for the name in custom nodes
			if tnode, err := extMaps.New(spec.Name); err == nil {
				node = tnode.(IBaseNode)
			}
		} else {
			if tnode, err2 := maps.New(spec.Name); err2 == nil {
				node = tnode.(IBaseNode)
			} else {
				fmt.Println("new ", spec.Name, " err:", err2)
			}
		}

		if node == nil {
			// Invalid node name
			panic("BehaviorTree.load: Invalid node name:" + spec.Name + ",title:" + spec.Title)
		}

		node.Ctor()
		node.Initialize(spec)
		node.SetBaseNodeWorker(node.(IBaseWorker)) //node是interface，可以转到类型IBaseWorker
		nodes[id] = node
	}

	// Connect the nodes
	for id, spec := range data.Nodes {
		node := nodes[id]

		if node.GetCategory() == b3.COMPOSITE && spec.Children != nil {
			for i := 0; i < len(spec.Children); i++ {
				var cid = spec.Children[i]
				comp := node.(IComposite)
				comp.AddChild(nodes[cid])
			}
		} else if node.GetCategory() == b3.DECORATOR && len(spec.Child) > 0 {
			dec := node.(IDecorator)
			dec.SetChild(nodes[spec.Child])
		}
	}

	this.root = nodes[data.Root]
}

/**
 * This method dump the current BT into a data structure.
 *
 * Note: This method does not record the current node parameters. Thus,
 * it may not be compatible with load for now.
 *
 * @method dump
 * @return {Object} A data object representing this tree.
**/
func (this *BehaviorTree) dump() *config.BTTreeCfg {
	return this.dumpInfo
}

/**
 * Propagates the tick signal through the tree, starting from the root.
 *
 * This method receives a target object of any type (Object, Array,
 * DOMElement, whatever) and a `Blackboard` instance. The target object has
 * no use at all for all Behavior3JS components, but surely is important
 * for custom nodes. The blackboard instance is used by the tree and nodes
 * to store execution variables (e.g., last node running) and is obligatory
 * to be a `Blackboard` instance (or an object with the same interface).
 *
 * Internally, this method creates a Tick object, which will store the
 * target and the blackboard objects.
 *
 * Note: BehaviorTree stores a list of open nodes from last tick, if these
 * nodes weren't called after the current tick, this method will close them
 * automatically.
 *
 * @method tick
 * @param {Object} target A target object.
 * @param {Blackboard} blackboard An instance of blackboard object.
 * @return {Constant} The tick signal state.
**/
func (this *BehaviorTree) Tick(target interface{}, blackboard *Blackboard) b3.Status {
	if blackboard == nil {
		panic("The blackboard parameter is obligatory and must be an instance of b3.Blackboard")
	}

	/* CREATE A TICK OBJECT */
	var tick = NewTick()
	tick.debug = this.debug
	tick.Target = target
	tick.Blackboard = blackboard
	tick.tree = this

	/* TICK NODE */
	var state = this.root._execute(tick)

	/* CLOSE NODES FROM LAST TICK, IF NEEDED */
	var lastOpenNodes = blackboard._getTreeData(this.id).OpenNodes
	var currOpenNodes []IBaseNode
	copy(tick._openNodes, currOpenNodes)

	// does not close if it is still open in this tick
	var start = 0
	for i := 0; i < b3.MinInt(len(lastOpenNodes), len(currOpenNodes)); i++ {
		start = i + 1
		if lastOpenNodes[i] != currOpenNodes[i] {
			break
		}
	}

	// close the nodes
	for i := len(lastOpenNodes) - 1; i >= start; i-- {
		lastOpenNodes[i]._close(tick)
	}

	/* POPULATE BLACKBOARD */
	blackboard._getTreeData(this.id).OpenNodes = currOpenNodes
	blackboard.SetTree("nodeCount", tick._nodeCount, this.id)

	return state
}

func (this *BehaviorTree) Print() {
	printNode(this.root, 0)
}

func printNode(root IBaseNode, blk int) { //打印树
	//fmt.Println("new node:", root.Name, " children:", len(root.Children), " child:", root.Child)
	for i := 0; i < blk; i++ {
		fmt.Print(" ") //缩进
	}

	//fmt.Println("|—<", root.Name, ">") //打印"|—<id>"形式
	fmt.Print("|—", root.GetTitle())

	if root.GetCategory() == b3.DECORATOR {
		dec := root.(IDecorator)
		if dec.GetChild() != nil {
			//fmt.Print("=>")
			printNode(dec.GetChild(), blk+3)
		}
	}

	fmt.Println("")
	if root.GetCategory() == b3.COMPOSITE {
		comp := root.(IComposite)
		if comp.GetChildCount() > 0 {
			for i := 0; i < comp.GetChildCount(); i++ {
				printNode(comp.GetChild(i), blk+3)
			}
		}
	}
}
