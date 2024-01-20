package parser

import (
	"github.com/jobin-nelson/writing_an_interpreter_in_go/ast"
	"github.com/jobin-nelson/writing_an_interpreter_in_go/lexer"
	"github.com/jobin-nelson/writing_an_interpreter_in_go/token"
)

type Parser struct {
  l *lexer.Lexer
  curToken token.Token
  peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
  p := &Parser{ l: l}

  p.nextToken()
  p.nextToken()
  return p
}

func (p *Parser) nextToken() {
  p.curToken = p.peekToken
  p.peekToken = p.l.NextToken()
}

func (p * Parser) ParseProgram() *ast.Program {
  return nil
}
