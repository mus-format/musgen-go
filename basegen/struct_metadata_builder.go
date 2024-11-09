package basegen

type StructMetadataBuilder interface {
	BuildStructMetadata() []*Metadata
}

type StructMetadata []FieldMetadataBuilder

func (m StructMetadata) BuildStructMetadata() (sl []*Metadata) {
	if m != nil {
		sl = make([]*Metadata, len(m))
		for i := 0; i < len(m); i++ {
			if m[i] != nil {
				fm := (FieldMetadataBuilder)(m[i]).BuildFieldMetadata()
				sl[i] = fm
			}
		}
	}
	return
}

type FieldMetadataBuilder interface {
	BuildFieldMetadata() *Metadata
}

type BoolFieldMetadata struct {
	BoolMetadata
	Ignore bool
}

func (m BoolFieldMetadata) BuildFieldMetadata() (meta *Metadata) {
	meta = m.BoolMetadata.BuildTypeMetadata()
	meta.Ignore = m.Ignore
	return
}

type NumFieldMetadata struct {
	NumMetadata
	Ignore bool
}

func (m NumFieldMetadata) BuildFieldMetadata() (meta *Metadata) {
	meta = m.NumMetadata.BuildTypeMetadata()
	meta.Ignore = m.Ignore
	return
}

type StringFieldMetadata struct {
	StringMetadata
	Ignore bool
}

func (m StringFieldMetadata) BuildFiledMetadata() (meta *Metadata) {
	meta = m.StringMetadata.BuildTypeMetadata()
	meta.Ignore = m.Ignore
	return
}

type SliceFieldMetadata struct {
	SliceMetadata
	Ignore bool
}

func (m SliceFieldMetadata) BuildFieldMetadata() (meta *Metadata) {
	meta = m.SliceMetadata.BuildTypeMetadata()
	meta.Ignore = m.Ignore
	return
}

type ArrayFieldMetadata struct {
	ArrayMetadata
	Ignore bool
}

func (m ArrayFieldMetadata) BuildFieldMetadata() (meta *Metadata) {
	meta = m.ArrayMetadata.BuildTypeMetadata()
	meta.Ignore = m.Ignore
	return
}

type MapFieldMetadata struct {
	MapMetadata
	Ignore bool
}

func (m MapFieldMetadata) BuildFieldMetadata() (meta *Metadata) {
	meta = m.MapMetadata.BuildTypeMetadata()
	meta.Ignore = m.Ignore
	return
}

type PtrFieldMetadata struct {
	PtrMetadata
	Ignore bool
}

func (m PtrFieldMetadata) BuildTypeMetadata() (meta *Metadata) {
	meta = m.PtrMetadata.BuildTypeMetadata()
	meta.Ignore = m.Ignore
	return
}

type CustomTypeFieldMetadata struct {
	Prefix    string
	Ignore    bool
	Validator string
}

func (m CustomTypeFieldMetadata) BuildFieldMetadata() *Metadata {
	return &Metadata{
		Prefix:    m.Prefix,
		Ignore:    m.Ignore,
		Validator: m.Validator,
	}
}
