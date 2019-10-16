package objects

import (
	"github.com/pkg/errors"
)

func BuildText(typ string, text string) *TextBuilder {
	var b TextBuilder
	b.typ = typ
	b.text = text
	return &b
}

func (b *TextBuilder) Emoji(v bool) *TextBuilder {
	b.emoji = v
	return b
}

func (b *TextBuilder) Verbatim(v bool) *TextBuilder {
	b.verbatim = v
	return b
}

func (b *TextBuilder) Do() (*Text, error) {
	if err := b.Validate(); err != nil {
		return nil, errors.Wrap(err, `validation for Text failed`)
	}

	var v Text
	v.typ = b.typ
	v.text = b.text
	v.emoji = b.emoji
	v.verbatim = b.verbatim
	return &v, nil
}

func (b *Text) Type() string {
	return b.typ
}

func (b *Text) Text() string {
	return b.text
}

func (b *Text) Emoji() bool {
	return b.emoji
}

func (b *Text) Verbatim() bool {
	return b.verbatim
}

func BuildUserProfile() *UserProfileBuilder {
	var b UserProfileBuilder
	return &b
}

func (b *UserProfileBuilder) AlwaysActive(v bool) *UserProfileBuilder {
	b.alwaysActive = v
	return b
}

func (b *UserProfileBuilder) AvatarHash(v string) *UserProfileBuilder {
	b.avatarHash = v
	return b
}

func (b *UserProfileBuilder) Email(v string) *UserProfileBuilder {
	b.email = v
	return b
}

func (b *UserProfileBuilder) FirstName(v string) *UserProfileBuilder {
	b.firstName = v
	return b
}

func (b *UserProfileBuilder) Image24(v string) *UserProfileBuilder {
	b.image24 = v
	return b
}

func (b *UserProfileBuilder) Image32(v string) *UserProfileBuilder {
	b.image32 = v
	return b
}

func (b *UserProfileBuilder) Image48(v string) *UserProfileBuilder {
	b.image48 = v
	return b
}

func (b *UserProfileBuilder) Image72(v string) *UserProfileBuilder {
	b.image72 = v
	return b
}

func (b *UserProfileBuilder) Image192(v string) *UserProfileBuilder {
	b.image192 = v
	return b
}

func (b *UserProfileBuilder) Image512(v string) *UserProfileBuilder {
	b.image512 = v
	return b
}

func (b *UserProfileBuilder) LastName(v string) *UserProfileBuilder {
	b.lastName = v
	return b
}

func (b *UserProfileBuilder) RealName(v string) *UserProfileBuilder {
	b.realName = v
	return b
}

func (b *UserProfileBuilder) RealNameNormalized(v string) *UserProfileBuilder {
	b.realNameNormalized = v
	return b
}

func (b *UserProfileBuilder) StatusText(v string) *UserProfileBuilder {
	b.statusText = v
	return b
}

func (b *UserProfileBuilder) StatusEmoji(v string) *UserProfileBuilder {
	b.statusEmoji = v
	return b
}

func (b *UserProfileBuilder) Do() (*UserProfile, error) {

	var v UserProfile
	v.alwaysActive = b.alwaysActive
	v.avatarHash = b.avatarHash
	v.email = b.email
	v.firstName = b.firstName
	v.image24 = b.image24
	v.image32 = b.image32
	v.image48 = b.image48
	v.image72 = b.image72
	v.image192 = b.image192
	v.image512 = b.image512
	v.lastName = b.lastName
	v.realName = b.realName
	v.realNameNormalized = b.realNameNormalized
	v.statusText = b.statusText
	v.statusEmoji = b.statusEmoji
	return &v, nil
}

func (b *UserProfile) AlwaysActive() bool {
	return b.alwaysActive
}

func (b *UserProfile) AvatarHash() string {
	return b.avatarHash
}

func (b *UserProfile) Email() string {
	return b.email
}

func (b *UserProfile) FirstName() string {
	return b.firstName
}

func (b *UserProfile) Image24() string {
	return b.image24
}

func (b *UserProfile) Image32() string {
	return b.image32
}

func (b *UserProfile) Image48() string {
	return b.image48
}

func (b *UserProfile) Image72() string {
	return b.image72
}

func (b *UserProfile) Image192() string {
	return b.image192
}

func (b *UserProfile) Image512() string {
	return b.image512
}

func (b *UserProfile) LastName() string {
	return b.lastName
}

func (b *UserProfile) RealName() string {
	return b.realName
}

func (b *UserProfile) RealNameNormalized() string {
	return b.realNameNormalized
}

func (b *UserProfile) StatusText() string {
	return b.statusText
}

func (b *UserProfile) StatusEmoji() string {
	return b.statusEmoji
}
