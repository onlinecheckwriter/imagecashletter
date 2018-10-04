// Copyright 2018 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

import (
	"bytes"
	"strings"
	"testing"
)

// TestX9Write writes an x9 File
func TestX9Write(t *testing.T) {
	file := NewFile().SetHeader(mockFileHeader())

	// Create CheckDetail
	cd := mockCheckDetail()
	cd.AddCheckDetailAddendumA(mockCheckDetailAddendumA())
	cd.AddCheckDetailAddendumB(mockCheckDetailAddendumB())
	cd.AddCheckDetailAddendumC(mockCheckDetailAddendumC())
	cd.AddImageViewDetail(mockImageViewDetail())
	cd.AddImageViewData(mockImageViewData())
	cd.AddImageViewAnalysis(mockImageViewAnalysis())
	bundle := NewBundle(mockBundleHeader())
	bundle.AddCheckDetail(cd)

	// CheckDetail 2
	cdTwo := mockCheckDetail()
	cdTwo.AddCheckDetailAddendumA(mockCheckDetailAddendumA())
	cdTwo.AddCheckDetailAddendumB(mockCheckDetailAddendumB())
	cdTwo.AddCheckDetailAddendumC(mockCheckDetailAddendumC())
	cdTwo.AddImageViewDetail(mockImageViewDetail())
	cdTwo.AddImageViewData(mockImageViewData())
	cdTwo.AddImageViewAnalysis(mockImageViewAnalysis())
	bundle.AddCheckDetail(cdTwo)

	// Create ReturnDetail
	rd := mockReturnDetail()
	rd.AddReturnDetailAddendumA(mockReturnDetailAddendumA())
	rd.AddReturnDetailAddendumB(mockReturnDetailAddendumB())
	rd.AddReturnDetailAddendumC(mockReturnDetailAddendumC())
	rd.AddReturnDetailAddendumD(mockReturnDetailAddendumD())
	rd.AddImageViewDetail(mockImageViewDetail())
	rd.AddImageViewData(mockImageViewData())
	rd.AddImageViewAnalysis(mockImageViewAnalysis())
	returnBundle := NewBundle(mockBundleHeader())
	returnBundle.BundleHeader.BundleSequenceNumber = "2"
	returnBundle.AddReturnDetail(rd)

	rdTwo := mockReturnDetail()
	rdTwo.AddReturnDetailAddendumA(mockReturnDetailAddendumA())
	rdTwo.AddReturnDetailAddendumB(mockReturnDetailAddendumB())
	rdTwo.AddReturnDetailAddendumC(mockReturnDetailAddendumC())
	rdTwo.AddReturnDetailAddendumD(mockReturnDetailAddendumD())
	rdTwo.AddImageViewDetail(mockImageViewDetail())
	rdTwo.AddImageViewData(mockImageViewData())
	rdTwo.AddImageViewAnalysis(mockImageViewAnalysis())
	returnBundle.AddReturnDetail(rdTwo)

	// Create CashLetter
	cl := NewCashLetter(mockCashLetterHeader())
	cl.AddBundle(bundle)
	cl.AddBundle(returnBundle)
	cl.Create()
	file.AddCashLetter(cl)

	clTwo := NewCashLetter(mockCashLetterHeader())
	clTwo.CashLetterHeader.CashLetterID = "A2"
	clTwo.AddBundle(bundle)
	clTwo.AddBundle(returnBundle)
	clTwo.Create()
	file.AddCashLetter(clTwo)

	// Create file
	if err := file.Create(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	if err := file.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	b := &bytes.Buffer{}
	f := NewWriter(b)

	if err := f.Write(file); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	/*
		// We want to write the file to an io.Writer
		w := NewWriter(os.Stdout)
		if err := w.Write(file); err != nil {
			log.Fatalf("Unexpected error: %s\n", err)
		}
		w.Flush()*/

	r := NewReader(strings.NewReader(b.String()))
	_, err := r.Read()
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	if err = r.File.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
}