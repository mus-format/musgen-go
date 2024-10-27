package basegen

type TypeMetadataBuilder interface {
	BuildTypeMetadata() *Metadata
}

type BoolMetadata struct {
	Validator string
}

func (m BoolMetadata) BuildTypeMetadata() *Metadata {
	return &Metadata{}
}

type NumMetadata struct {
	Encoding  NumEncoding
	Validator string
}

func (m NumMetadata) BuildTypeMetadata() *Metadata {
	return &Metadata{Encoding: m.Encoding, Validator: m.Validator}
}

type StringMetadata struct {
	LenEncoding  NumEncoding
	LenValidator string
	Validator    string
}

func (m StringMetadata) BuildTypeMetadata() *Metadata {
	return &Metadata{
		LenEncoding:  m.LenEncoding,
		LenValidator: m.LenValidator,
		Validator:    m.Validator,
	}
}

type SliceMetadata struct {
	LenEncoding  NumEncoding
	LenValidator string
	Elem         TypeMetadataBuilder
	Validator    string
}

func (m SliceMetadata) BuildTypeMetadata() *Metadata {
	tm := &Metadata{
		LenEncoding:  m.LenEncoding,
		LenValidator: m.LenValidator,
		Validator:    m.Validator,
	}
	if m.Elem != nil {
		tm.Elem = m.Elem.BuildTypeMetadata()
	}
	return tm
}

type MapMetadata struct {
	LenEncoding  NumEncoding
	LenValidator string
	Key          TypeMetadataBuilder
	Elem         TypeMetadataBuilder
	Validator    string
}

func (m MapMetadata) BuildTypeMetadata() *Metadata {
	tm := &Metadata{
		LenEncoding:  m.LenEncoding,
		LenValidator: m.LenValidator,
		Validator:    m.Validator,
	}
	if m.Key != nil {
		tm.Key = m.Key.BuildTypeMetadata()
	}
	if m.Elem != nil {
		tm.Elem = m.Elem.BuildTypeMetadata()
	}
	return tm
}

type PtrMetadata struct {
	Validator string
}

func (m PtrMetadata) BuildTypeMetadata() *Metadata {
	return &Metadata{
		Validator: m.Validator,
	}
}
