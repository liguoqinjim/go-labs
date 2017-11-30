/**
 * Author:        Tony.Shao
 * Email:         xiocode@gmail.com
 * Github:        github.com/xiocode
 * File:          rpn.go
 * Description:   https://gist.github.com/achun/5730664
 */

//ReversePolishnotation，RPN
//逆波兰表达式相关
//lexer 是一个中缀数学四则计算表达式词法解析器
package rpn

import (
	"errors"
)

type Token int

func (p Token) String() string {
	return tokens[p]
}

const eof = -1

const (
	TokenNil        Token = iota
	paren_beg             //paren_beg
	TokenLeftParen        //"("
	TokenRightParen       //")"
	paren_end

	literal_beg //literal_beg
	TokenNumber //0..9
	TokenDot    //"."
	literal_end

	operator_beg //operator_beg
	TokenAdd     //"+"
	TokenSub     //"-"
	TokenDiv     //"/"
	TokenMul     //"*"
	operator_end

	//TokenAssign //"="
	//词义token
	TokenInt      //整数
	TokenDecimal  //十进制小数
	TokenFloat    //浮点数
	TokenOperator //+-*/
	TokenValue
	TokenExp //表达式
	TokenWhiteSpace
	TokenError
)

var tokens = [...]string{
	TokenNil:        "Nil",
	TokenError:      "Error",
	TokenDot:        "Dot",
	TokenNumber:     "Number",
	TokenDecimal:    "Decimal",
	TokenAdd:        "Add",
	TokenSub:        "Sub",
	TokenMul:        "Mul",
	TokenDiv:        "Div",
	TokenWhiteSpace: "WhiteSpace",
	TokenLeftParen:  "LeftParen",
	TokenRightParen: "RightParen",
}

// token变化关系,设置上一个token允许的类型
var lexerwish = [...][]Token{
	TokenNumber:     []Token{TokenNil, TokenNumber, TokenDecimal, TokenAdd, TokenSub, TokenMul, TokenDiv, TokenLeftParen},
	TokenDot:        []Token{TokenNil, TokenNumber, TokenAdd, TokenSub, TokenMul, TokenDiv, TokenLeftParen},
	TokenDecimal:    []Token{TokenDot, TokenAdd, TokenSub, TokenMul, TokenDiv, TokenLeftParen},
	TokenAdd:        []Token{TokenNumber, TokenDecimal, TokenRightParen},
	TokenSub:        []Token{TokenNumber, TokenDecimal, TokenRightParen},
	TokenMul:        []Token{TokenNumber, TokenDecimal, TokenRightParen},
	TokenDiv:        []Token{TokenNumber, TokenDecimal, TokenRightParen},
	TokenLeftParen:  []Token{TokenNil, TokenAdd, TokenSub, TokenMul, TokenDiv, TokenLeftParen},
	TokenRightParen: []Token{TokenNumber, TokenDecimal, TokenRightParen},
}

type stateFn func() stateFn

type lexer struct {
	closed            bool
	input             string
	pos               int    //pos
	runer             rune   //当前pos的字符
	token             Token  //上一个已经确定的 Token
	value             []rune //上一个token 对应的value
	t                 Token  //当前可变动的 token
	v                 []rune //当前可变动的 token 对应的value
	leftParenCounter  int    // (计数器
	rightParenCounter int    // )计数器
	errMsg            string //错误信息
	cha               chan bool
}

//根据输入文本新建词法解析实例
func NewLexer(input string) (p *lexer) {
	p = new(lexer)
	p.cha = make(chan bool)
	p.input = input
	go p.run()
	return
}

func (p *lexer) run() {
	fn := p.walk
	for pos, r := range p.input {
		p.pos = pos
		p.runer = r
		if p.closed {
			print("closed ", pos, string(r))
			return
		}
		fn = fn()
		if fn == nil {
			print("fn == nil ")
			p.Close()
			p.token = TokenError
			if p.errMsg == "" {
				p.error("lexer error")
			}
			return
		}
	}
	p.save()
	p.Close()
	if p.leftParenCounter != p.rightParenCounter {
		p.error("rightParen Not paired")
	} else {
		p.token = TokenNil
		p.v = []rune{}
	}
}

func (p *lexer) Error() error {
	if len(p.errMsg) == 0 {
		return nil
	}
	return errors.New(p.errMsg)
}

//单步词法解析
//返回: 词法 token 序号和字面值
//  以下两种情况解析过程将自动终止
//  TokenError 表示发生了词法错误
//  TokenNil   表示解析结束
func (p *lexer) Next() (t Token, v []rune) {
	<-p.cha
	//println("Next()", p.pos)
	t = p.token
	v = p.value
	<-p.cha
	return
}

//终止解析过程
//因内部使用了 gorountine.如果需要中途终止解析过程，Close()提供结束 gorountine 的方法
func (p *lexer) Close() {
	if p.closed {
		return
	}
	p.closed = true
	close(p.cha)
}

//保存 p.t ,返回 p.walk 继续执行
func (p *lexer) save() {
	if p.closed || p.t == TokenNil {
		return
	}
	p.token = p.t
	p.value = p.v
	p.cha <- true
	p.cha <- true
	p.v = []rune{}
}

//设置 p.t
func (p *lexer) emit(token Token) {
	p.t = token
}

// append(p.v, p.runer) 返回 p.walk 继续执行
func (p *lexer) append() {
	p.v = append(p.v, p.runer)
}

//设置错误信息
func (p *lexer) error(str string) {
	p.token = TokenError
	p.errMsg = "lexer error: " + str
}

func (p *lexer) walk() stateFn {
	r := p.runer
	var fn stateFn
	var token Token

	if r >= '0' && r <= '9' {
		token = TokenNumber
		fn = p.number
	} else {
		switch r {
		case '\t', '\n', ' ':
			token = TokenWhiteSpace
			fn = p.whiteSpace
		case '.':
			token = TokenDot
			fn = p.dot
		case '+':
			token = TokenAdd
			fn = p.add
		case '-':
			token = TokenSub
			fn = p.sub
		case '*':
			token = TokenMul
			fn = p.mul
		case '/':
			token = TokenDiv
			fn = p.div
		case '(':
			token = TokenLeftParen
			fn = p.leftParen
		case ')':
			token = TokenRightParen
			fn = p.rightParen
		}
	}

	if fn == nil {
		p.error("unknow char " + " \"" + string(r) + "\"")
		return nil
	}
	for _, i := range lexerwish[token] {
		if p.t == i {
			return fn()
		}
	}
	p.error("token " + token.String() + " \"" + string(r) + "\"")
	return nil
}

func (p *lexer) eof() stateFn {
	return nil
}
func (p *lexer) number() stateFn {
	switch p.t {
	default:
		p.error("number")
		return nil
	case TokenLeftParen:
		p.save()
		p.emit(TokenNumber)
	case TokenNil, TokenNumber:
		p.emit(TokenNumber)
	case TokenDot, TokenDecimal:
		p.emit(TokenDecimal)
	case TokenAdd, TokenSub, TokenMul, TokenDiv:
		p.save()
		p.emit(TokenNumber)
	}
	p.append()
	return p.walk
}
func (p *lexer) dot() stateFn {
	switch p.t {
	default:
		p.error("dot")
		return nil
	case TokenNil:
	case TokenNumber:
		p.emit(TokenDecimal)
	}
	p.append()
	return p.walk
}
func (p *lexer) add() stateFn {
	p.save()
	p.emit(TokenAdd)
	p.append()
	return p.walk
}
func (p *lexer) sub() stateFn {
	p.save()
	p.emit(TokenSub)
	p.append()
	return p.walk
}
func (p *lexer) mul() stateFn {
	p.save()
	p.emit(TokenMul)
	p.append()
	return p.walk
}
func (p *lexer) div() stateFn {
	p.save()
	p.emit(TokenDiv)
	p.append()
	return p.walk
}

func (p *lexer) leftParen() stateFn {
	p.leftParenCounter++
	p.save()
	p.emit(TokenLeftParen)
	p.append()
	return p.walk
}
func (p *lexer) rightParen() stateFn {
	if p.rightParenCounter >= p.leftParenCounter {
		p.error("rightParen Not paired")
		return nil
	}
	p.rightParenCounter++
	p.save()
	p.emit(TokenRightParen)
	p.append()
	return p.walk
}

// 待完善
func (p *lexer) whiteSpace() stateFn {
	p.error("whiteSpace")
	return nil
}

type node struct {
	token Token
	v     []rune
	as    Token // 分类token
}

//ast树
type tree struct {
	node   *node
	left   *tree
	right  *tree
	parent *tree
	errMsg string
}

//要求node已经保障了合法性
//树节点次序由左至右
func (t *tree) push(n *node) *tree {
	switch n.as {
	case TokenValue:
		if t.left == nil {
			t.left = &tree{node: n, parent: t}
		} else if t.right == nil {
			t.right = &tree{node: n, parent: t}
		} else {
			t.errMsg = "tree:not nil node of left and right " + string(n.v)
			return nil
		}
		return t
	case TokenOperator:
		if t.node == nil {
			t.node = n
			return t
		}
		if t.right != nil {
			x := opLeve[n.token] - opLeve[t.node.token]
			if x > 0 {
				t.right = &tree{node: n, left: t.right, parent: t}
				t.right.left.parent = t.right
				return t.right
			}
			if x < 0 && t.parent != nil {
				return t.parent.push(n)
			}
			nt := &tree{node: n, left: t, parent: t.parent}
			if t.parent != nil {
				if t.parent.left == t {
					t.parent.left = nt
				} else {
					t.parent.right = nt
				}
			}
			t.parent = nt
			return nt

		}
		t.errMsg = "tree:expect not nil node of right"
		return nil
	case TokenLeftParen:
		if t.left == nil {
			t.left = &tree{parent: t}
			return t.left
		}
		if t.right == nil {
			t.right = &tree{parent: t}
			return t.right
		}
		t.errMsg = "tree:not nil node of left and right when TokenLeftParen"
		return nil
	case TokenRightParen:
		return t.parent
	}
	t.errMsg = "tree:unknow node as " + n.token.String()
	return nil
}
func (t *tree) root() *tree {
	ret := t
	for ret.parent != nil {
		ret = ret.parent
	}
	return ret
}
func (t *tree) String(deep int) string {
	prefix := ""
	for i := deep; i > 0; i-- {
		prefix += "\t"
	}
	str := ""
	if t.node != nil {
		str += t.node.token.String() + ":" + string(t.node.v) + "\n"
	} else {
		str += "NULL:\n"
	}
	if t.left != nil {
		str += prefix + "\tLeft " + t.left.String(deep+1)
	}
	if t.right != nil {
		str += prefix + "\tRight " + t.right.String(deep+1)
	}
	return str
}

//返回左叶子节点
func (t *tree) leaf(stop *tree) *tree {
	if t == stop {
		return t.parent
	}
	if t.left == nil {
		return t
	}
	return t.left.leaf(stop)
}

//消除有侧多余的空节点
func (t *tree) less() {
	leaf := t
	for leaf != nil {
		if leaf.node != nil && leaf.node.as == TokenValue {
			break
		}
		if leaf.parent == nil {
			leaf = leaf.right
			continue
		}
		if leaf.node == nil {
			if leaf.left == nil {
				leaf.parent.right = leaf.right
				leaf.right.parent = leaf.parent
			} else if leaf.right == nil {
				leaf.parent.right = leaf.left
				leaf.left.parent = leaf.parent
			}
		}
		leaf = leaf.right
	}
	leaf = t
	for leaf != nil {
		if leaf.node != nil && leaf.node.as == TokenValue {
			break
		}
		if leaf.parent == nil {
			leaf = leaf.left
			continue
		}
		if leaf.node == nil {
			if leaf.left == nil {
				leaf.parent.left = leaf.right
				leaf.right.parent = leaf.parent
			} else if leaf.right == nil {
				leaf.parent.left = leaf.left
				leaf.left.parent = leaf.parent
			}
		}
		leaf = leaf.left
	}
}

//遍历树
//fn 为遍历回调函数,他返回一般bool值，表示是否终止遍历
func (t *tree) Walk(fn func(*tree) bool) {
	t.walk(t, fn)
}
func (t *tree) walk(stop *tree, fn func(*tree) bool) bool {
	if t.left != nil {
		if t.left.walk(stop, fn) {
			return true
		}
	}

	if t.right != nil {
		if t.right.walk(stop, fn) {
			return true
		}
	}
	if fn(t) {
		return true
	}
	return t == stop
}

func (t *tree) Rpn() (ret []string) {
	fn := func(tr *tree) bool {
		ret = append(ret, string(tr.node.v))
		return false
	}
	t.Walk(fn)
	return
}

var opLeve = [...]int{
	TokenAdd: 0,
	TokenSub: 0,
	TokenMul: 1,
	TokenDiv: 1,
}

// 解析中缀表达式到tree
func Parse(input string) (*tree, error) {
	l := NewLexer(input)
	t := new(tree)
	token, v := l.Next()
	paren := 0
	for {
		if token == TokenError {
			return t, l.Error()
		}
		if token == TokenNil {
			break
		}
		switch token {
		default:
			l.Close()
			return t, errors.New("tree:unknow token " + token.String())
		case TokenDecimal, TokenNumber:
			t = t.push(&node{token: token, v: v, as: TokenValue})
		case TokenAdd, TokenSub, TokenMul, TokenDiv:
			t = t.push(&node{token: token, v: v, as: TokenOperator})
		case TokenLeftParen:
			paren++
			t = t.push(&node{token: token, v: v, as: TokenLeftParen})
		case TokenRightParen:
			paren--
			t = t.push(&node{token: token, v: v, as: TokenRightParen})
			//消除多余的()配对
			if paren == 0 {
				//println("paren == 0");println(t.String(0));println("less")
				t.less()
				//println(t.String(0));println("===")
			}
		}
		if t == nil {
			l.Close()
			return t, errors.New("tree:AST Error")
		}
		//println(t.root().String(0))
		token, v = l.Next()
	}
	return t.root(), nil
}
