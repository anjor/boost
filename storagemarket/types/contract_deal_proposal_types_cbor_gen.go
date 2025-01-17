// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package types

import (
	"fmt"
	"io"
	"math"
	"sort"

	abi "github.com/filecoin-project/go-state-types/abi"
	cid "github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

var _ = xerrors.Errorf
var _ = cid.Undef
var _ = math.E
var _ = sort.Sort

var lengthBufContractDealProposal = []byte{140}

func (t *ContractDealProposal) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufContractDealProposal); err != nil {
		return err
	}

	// t.PieceCID (cid.Cid) (struct)

	if err := cbg.WriteCid(cw, t.PieceCID); err != nil {
		return xerrors.Errorf("failed to write cid field t.PieceCID: %w", err)
	}

	// t.PieceSize (abi.PaddedPieceSize) (uint64)

	if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.PieceSize)); err != nil {
		return err
	}

	// t.VerifiedDeal (bool) (bool)
	if err := cbg.WriteBool(w, t.VerifiedDeal); err != nil {
		return err
	}

	// t.Client (address.Address) (struct)
	if err := t.Client.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.Label (types.DealLabel) (struct)
	if err := t.Label.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.StartEpoch (abi.ChainEpoch) (int64)
	if t.StartEpoch >= 0 {
		if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.StartEpoch)); err != nil {
			return err
		}
	} else {
		if err := cw.WriteMajorTypeHeader(cbg.MajNegativeInt, uint64(-t.StartEpoch-1)); err != nil {
			return err
		}
	}

	// t.EndEpoch (abi.ChainEpoch) (int64)
	if t.EndEpoch >= 0 {
		if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.EndEpoch)); err != nil {
			return err
		}
	} else {
		if err := cw.WriteMajorTypeHeader(cbg.MajNegativeInt, uint64(-t.EndEpoch-1)); err != nil {
			return err
		}
	}

	// t.StoragePricePerEpoch (big.Int) (struct)
	if err := t.StoragePricePerEpoch.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.ProviderCollateral (big.Int) (struct)
	if err := t.ProviderCollateral.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.ClientCollateral (big.Int) (struct)
	if err := t.ClientCollateral.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.Version (string) (string)
	if len(t.Version) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.Version was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(t.Version))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.Version)); err != nil {
		return err
	}

	// t.Params ([]uint8) (slice)
	if len(t.Params) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Params was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajByteString, uint64(len(t.Params))); err != nil {
		return err
	}

	if _, err := cw.Write(t.Params[:]); err != nil {
		return err
	}
	return nil
}

func (t *ContractDealProposal) UnmarshalCBOR(r io.Reader) (err error) {
	*t = ContractDealProposal{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 12 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.PieceCID (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(cr)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.PieceCID: %w", err)
		}

		t.PieceCID = c

	}
	// t.PieceSize (abi.PaddedPieceSize) (uint64)

	{

		maj, extra, err = cr.ReadHeader()
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.PieceSize = abi.PaddedPieceSize(extra)

	}
	// t.VerifiedDeal (bool) (bool)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}
	if maj != cbg.MajOther {
		return fmt.Errorf("booleans must be major type 7")
	}
	switch extra {
	case 20:
		t.VerifiedDeal = false
	case 21:
		t.VerifiedDeal = true
	default:
		return fmt.Errorf("booleans are either major type 7, value 20 or 21 (got %d)", extra)
	}
	// t.Client (address.Address) (struct)

	{

		if err := t.Client.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.Client: %w", err)
		}

	}
	// t.Label (types.DealLabel) (struct)

	{

		if err := t.Label.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.Label: %w", err)
		}

	}
	// t.StartEpoch (abi.ChainEpoch) (int64)
	{
		maj, extra, err := cr.ReadHeader()
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative oveflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.StartEpoch = abi.ChainEpoch(extraI)
	}
	// t.EndEpoch (abi.ChainEpoch) (int64)
	{
		maj, extra, err := cr.ReadHeader()
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative oveflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.EndEpoch = abi.ChainEpoch(extraI)
	}
	// t.StoragePricePerEpoch (big.Int) (struct)

	{

		if err := t.StoragePricePerEpoch.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.StoragePricePerEpoch: %w", err)
		}

	}
	// t.ProviderCollateral (big.Int) (struct)

	{

		if err := t.ProviderCollateral.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.ProviderCollateral: %w", err)
		}

	}
	// t.ClientCollateral (big.Int) (struct)

	{

		if err := t.ClientCollateral.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.ClientCollateral: %w", err)
		}

	}
	// t.Version (string) (string)

	{
		sval, err := cbg.ReadString(cr)
		if err != nil {
			return err
		}

		t.Version = string(sval)
	}
	// t.Params ([]uint8) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.Params: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.Params = make([]uint8, extra)
	}

	if _, err := io.ReadFull(cr, t.Params[:]); err != nil {
		return err
	}
	return nil
}

var lengthBufParamsVersion1 = []byte{131}

func (t *ContractParamsVersion1) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufParamsVersion1); err != nil {
		return err
	}

	// t.LocationRef (string) (string)
	if len(t.LocationRef) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.LocationRef was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(t.LocationRef))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.LocationRef)); err != nil {
		return err
	}

	// t.CarSize (uint64) (uint64)

	if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.CarSize)); err != nil {
		return err
	}

	// t.SkipIpniAnnounce (bool) (bool)
	if err := cbg.WriteBool(w, t.SkipIpniAnnounce); err != nil {
		return err
	}
	return nil
}

func (t *ContractParamsVersion1) UnmarshalCBOR(r io.Reader) (err error) {
	*t = ContractParamsVersion1{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 3 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.LocationRef (string) (string)

	{
		sval, err := cbg.ReadString(cr)
		if err != nil {
			return err
		}

		t.LocationRef = string(sval)
	}
	// t.CarSize (uint64) (uint64)

	{

		maj, extra, err = cr.ReadHeader()
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.CarSize = uint64(extra)

	}
	// t.SkipIpniAnnounce (bool) (bool)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}
	if maj != cbg.MajOther {
		return fmt.Errorf("booleans must be major type 7")
	}
	switch extra {
	case 20:
		t.SkipIpniAnnounce = false
	case 21:
		t.SkipIpniAnnounce = true
	default:
		return fmt.Errorf("booleans are either major type 7, value 20 or 21 (got %d)", extra)
	}
	return nil
}

