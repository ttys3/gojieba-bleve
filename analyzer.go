package jbleve

import (
	"errors"

	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

func analyzerConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.Analyzer, error) {
	tokenizerName, ok := config["tokenizer"].(string)
	if !ok {
		return nil, errors.New("must specify tokenizer")
	}
	tokenizer, err := cache.TokenizerNamed(tokenizerName)
	if err != nil {
		return nil, err
	}

	jbtk, ok := tokenizer.(*JiebaTokenizer)
	if !ok {
		return nil, errors.New("tokenizer must be of type jieba")
	}
	alz := &JiebaAnalyzer{
		Tokenizer: jbtk,
	}
	return alz, nil
}

func init() {
	registry.RegisterAnalyzer("gojieba", analyzerConstructor)
}

// JiebaAnalyzer from analysis.DefaultAnalyzer
type JiebaAnalyzer struct {
	CharFilters  []analysis.CharFilter
	Tokenizer    *JiebaTokenizer
	TokenFilters []analysis.TokenFilter
}

func (a *JiebaAnalyzer) Analyze(input []byte) analysis.TokenStream {
	if a.CharFilters != nil {
		for _, cf := range a.CharFilters {
			input = cf.Filter(input)
		}
	}
	tokens := a.Tokenizer.Tokenize(input)
	if a.TokenFilters != nil {
		for _, tf := range a.TokenFilters {
			tokens = tf.Filter(tokens)
		}
	}
	return tokens
}

func (a *JiebaAnalyzer) Free() {
	if a.Tokenizer != nil {
		a.Tokenizer.Free()
	} else {
		panic("JiebaAnalyzer.Tokenizer is nil, this should not happen")
	}
}
