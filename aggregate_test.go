package lazyrnn

import (
	"testing"

	"github.com/unixpickle/anydiff"
	"github.com/unixpickle/anydiff/anyseq"
	"github.com/unixpickle/anyvec/anyvec64"
)

func TestMean(t *testing.T) {
	const inSize = 3
	c := anyvec64.DefaultCreator{}
	inSeqs := testSeqs(c, inSize)

	count := 0
	for _, seq := range anyseq.SeparateSeqs(inSeqs.Output()) {
		count += len(seq)
	}

	actualFunc := func() anydiff.Res {
		return Mean(Lazify(inSeqs))
	}
	expectedFunc := func() anydiff.Res {
		return anydiff.Scale(anyseq.Sum(inSeqs),
			inSeqs.Creator().MakeNumeric(1/float64(count)))
	}
	testEquivalentRes(t, actualFunc, expectedFunc)
}

func TestSumEach(t *testing.T) {
	const inSize = 3
	c := anyvec64.DefaultCreator{}
	inSeqs := testSeqs(c, inSize)

	actualFunc := func() anydiff.Res {
		return SumEach(Lazify(inSeqs))
	}
	expectedFunc := func() anydiff.Res {
		return anyseq.SumEach(inSeqs)
	}
	testEquivalentRes(t, actualFunc, expectedFunc)
}

func TestSum(t *testing.T) {
	const inSize = 3
	c := anyvec64.DefaultCreator{}
	inSeqs := testSeqs(c, inSize)

	actualFunc := func() anydiff.Res {
		return Sum(Lazify(inSeqs))
	}
	expectedFunc := func() anydiff.Res {
		return anyseq.Sum(inSeqs)
	}
	testEquivalentRes(t, actualFunc, expectedFunc)
}
