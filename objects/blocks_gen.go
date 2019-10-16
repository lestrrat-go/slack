package objects

func BuildActionsBlock(elements ...BlockElement) *ActionsBlockBuilder {
	var b ActionsBlockBuilder
	b.elements = elements
	return &b
}

func (b *ActionsBlockBuilder) BlockId(v string) *ActionsBlockBuilder {
	b.blockId = v
	return b
}

func (b *ActionsBlockBuilder) Do() (*ActionsBlock, error) {
	var v ActionsBlock
	v.elements = b.elements
	v.blockId = b.blockId
	return &v, nil
}

func (b *ActionsBlock) Elements() []BlockElement {
	return b.elements
}

func (b *ActionsBlock) BlockId() string {
	return b.blockId
}

func (b ActionsBlock) Type() BlockType {
	return ActionsBlockType
}

func BuildContextBlock(elements ...interface{}) *ContextBlockBuilder {
	var b ContextBlockBuilder
	b.elements = elements
	return &b
}

func (b *ContextBlockBuilder) BlockId(v string) *ContextBlockBuilder {
	b.blockId = v
	return b
}

func (b *ContextBlockBuilder) Do() (*ContextBlock, error) {
	var v ContextBlock
	v.elements = b.elements
	v.blockId = b.blockId
	return &v, nil
}

func (b *ContextBlock) Elements() []interface{} {
	return b.elements
}

func (b *ContextBlock) BlockId() string {
	return b.blockId
}

func (b ContextBlock) Type() BlockType {
	return ContextBlockType
}

func BuildDividerBlock() *DividerBlockBuilder {
	var b DividerBlockBuilder
	return &b
}

func (b *DividerBlockBuilder) BlockId(v string) *DividerBlockBuilder {
	b.blockId = v
	return b
}

func (b *DividerBlockBuilder) Do() (*DividerBlock, error) {
	var v DividerBlock
	v.blockId = b.blockId
	return &v, nil
}

func (b *DividerBlock) BlockId() string {
	return b.blockId
}

func (b DividerBlock) Type() BlockType {
	return DividerBlockType
}

func BuildFileBlock(externalId string) *FileBlockBuilder {
	var b FileBlockBuilder
	b.externalId = externalId
	return &b
}

func (b *FileBlockBuilder) Source(v string) *FileBlockBuilder {
	b.source = v
	return b
}

func (b *FileBlockBuilder) BlockId(v string) *FileBlockBuilder {
	b.blockId = v
	return b
}

func (b *FileBlockBuilder) Do() (*FileBlock, error) {
	var v FileBlock
	v.externalId = b.externalId
	v.source = b.source
	v.blockId = b.blockId
	return &v, nil
}

func (b *FileBlock) ExternalId() string {
	return b.externalId
}

func (b *FileBlock) Source() string {
	return b.source
}

func (b *FileBlock) BlockId() string {
	return b.blockId
}

func (b FileBlock) Type() BlockType {
	return FileBlockType
}

func BuildImageBlock(imageUrl string, altText string) *ImageBlockBuilder {
	var b ImageBlockBuilder
	b.imageUrl = imageUrl
	b.altText = altText
	return &b
}

func (b *ImageBlockBuilder) Title(v string) *ImageBlockBuilder {
	b.title = v
	return b
}

func (b *ImageBlockBuilder) BlockId(v string) *ImageBlockBuilder {
	b.blockId = v
	return b
}

func (b *ImageBlockBuilder) Do() (*ImageBlock, error) {
	var v ImageBlock
	v.imageUrl = b.imageUrl
	v.altText = b.altText
	v.title = b.title
	v.blockId = b.blockId
	return &v, nil
}

func (b *ImageBlock) ImageUrl() string {
	return b.imageUrl
}

func (b *ImageBlock) AltText() string {
	return b.altText
}

func (b *ImageBlock) Title() string {
	return b.title
}

func (b *ImageBlock) BlockId() string {
	return b.blockId
}

func (b ImageBlock) Type() BlockType {
	return ImageBlockType
}

func BuildInputBlock(label string) *InputBlockBuilder {
	var b InputBlockBuilder
	b.label = label
	return &b
}

func (b *InputBlockBuilder) Element(v interface{}) *InputBlockBuilder {
	b.element = v
	return b
}

func (b *InputBlockBuilder) Hint(v *Text) *InputBlockBuilder {
	b.hint = v
	return b
}

func (b *InputBlockBuilder) Optional(v bool) *InputBlockBuilder {
	b.optional = v
	return b
}

func (b *InputBlockBuilder) Do() (*InputBlock, error) {
	var v InputBlock
	v.label = b.label
	v.element = b.element
	v.hint = b.hint
	v.optional = b.optional
	return &v, nil
}

func (b *InputBlock) Label() string {
	return b.label
}

func (b *InputBlock) Element() interface{} {
	return b.element
}

func (b *InputBlock) Hint() *Text {
	return b.hint
}

func (b *InputBlock) Optional() bool {
	return b.optional
}

func (b InputBlock) Type() BlockType {
	return InputBlockType
}

func BuildSectionBlock(text *Text) *SectionBlockBuilder {
	var b SectionBlockBuilder
	b.text = text
	return &b
}

func (b *SectionBlockBuilder) Fields(v []*Text) *SectionBlockBuilder {
	b.fields = v
	return b
}

func (b *SectionBlockBuilder) BlockId(v string) *SectionBlockBuilder {
	b.blockId = v
	return b
}

func (b *SectionBlockBuilder) Accessory(v BlockElement) *SectionBlockBuilder {
	b.accessory = v
	return b
}

func (b *SectionBlockBuilder) Do() (*SectionBlock, error) {
	var v SectionBlock
	v.text = b.text
	v.fields = b.fields
	v.blockId = b.blockId
	v.accessory = b.accessory
	return &v, nil
}

func (b *SectionBlock) Text() *Text {
	return b.text
}

func (b *SectionBlock) Fields() []*Text {
	return b.fields
}

func (b *SectionBlock) BlockId() string {
	return b.blockId
}

func (b *SectionBlock) Accessory() BlockElement {
	return b.accessory
}

func (b SectionBlock) Type() BlockType {
	return SectionBlockType
}
