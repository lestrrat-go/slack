package objects

import (
	"github.com/pkg/errors"
)

func BuildAction() *ActionBuilder {
	var b ActionBuilder
	return &b
}

func (b *ActionBuilder) Confirm(v *Confirmation) *ActionBuilder {
	b.confirm = v
	return b
}

func (b *ActionBuilder) DataSource(v string) *ActionBuilder {
	b.dataSource = v
	return b
}

func (b *ActionBuilder) MinQueryLength(v int) *ActionBuilder {
	b.minQueryLength = v
	return b
}

func (b *ActionBuilder) Name(v string) *ActionBuilder {
	b.name = v
	return b
}

func (b *ActionBuilder) OptionGroups(v ...*OptionGroup) *ActionBuilder {
	b.optionGroups = v
	return b
}

func (b *ActionBuilder) Options(v ...*Option) *ActionBuilder {
	b.options = v
	return b
}

func (b *ActionBuilder) SelectedOptions(v ...*Option) *ActionBuilder {
	b.selectedOptions = v
	return b
}

func (b *ActionBuilder) Style(v string) *ActionBuilder {
	b.style = v
	return b
}

func (b *ActionBuilder) Text(v string) *ActionBuilder {
	b.text = v
	return b
}

func (b *ActionBuilder) Type(v string) *ActionBuilder {
	b.typ = v
	return b
}

func (b *ActionBuilder) Value(v string) *ActionBuilder {
	b.value = v
	return b
}

func (b *ActionBuilder) Do() (*Action, error) {
	var v Action
	v.confirm = b.confirm
	v.dataSource = b.dataSource
	v.minQueryLength = b.minQueryLength
	v.name = b.name
	v.optionGroups = b.optionGroups
	v.options = b.options
	v.selectedOptions = b.selectedOptions
	v.style = b.style
	v.text = b.text
	v.typ = b.typ
	v.value = b.value
	return &v, nil
}

func (b *Action) Confirm() *Confirmation {
	return b.confirm
}

func (b *Action) DataSource() string {
	return b.dataSource
}

func (b *Action) MinQueryLength() int {
	return b.minQueryLength
}

func (b *Action) Name() string {
	return b.name
}

func (b *Action) OptionGroups() OptionGroupList {
	return b.optionGroups
}

func (b *Action) Options() OptionList {
	return b.options
}

func (b *Action) SelectedOptions() OptionList {
	return b.selectedOptions
}

func (b *Action) Style() string {
	return b.style
}

func (b *Action) Text() string {
	return b.text
}

func (b *Action) Type() string {
	return b.typ
}

func (b *Action) Value() string {
	return b.value
}

func BuildAttachment() *AttachmentBuilder {
	var b AttachmentBuilder
	return &b
}

func (b *AttachmentBuilder) Actions(v ...*Action) *AttachmentBuilder {
	b.actions = v
	return b
}

func (b *AttachmentBuilder) AttachmentType(v string) *AttachmentBuilder {
	b.attachmentType = v
	return b
}

func (b *AttachmentBuilder) AuthorName(v string) *AttachmentBuilder {
	b.authorName = v
	return b
}

func (b *AttachmentBuilder) AuthorLink(v string) *AttachmentBuilder {
	b.authorLink = v
	return b
}

func (b *AttachmentBuilder) AuthorIcon(v string) *AttachmentBuilder {
	b.authorIcon = v
	return b
}

func (b *AttachmentBuilder) CallbackId(v string) *AttachmentBuilder {
	b.callbackId = v
	return b
}

func (b *AttachmentBuilder) Color(v string) *AttachmentBuilder {
	b.color = v
	return b
}

func (b *AttachmentBuilder) Fallback(v string) *AttachmentBuilder {
	b.fallback = v
	return b
}

func (b *AttachmentBuilder) Fields(v ...*AttachmentField) *AttachmentBuilder {
	b.fields = v
	return b
}

func (b *AttachmentBuilder) Footer(v string) *AttachmentBuilder {
	b.footer = v
	return b
}

func (b *AttachmentBuilder) FooterIcon(v string) *AttachmentBuilder {
	b.footerIcon = v
	return b
}

func (b *AttachmentBuilder) ImageUrl(v string) *AttachmentBuilder {
	b.imageUrl = v
	return b
}

func (b *AttachmentBuilder) ThumbUrl(v string) *AttachmentBuilder {
	b.thumbUrl = v
	return b
}

func (b *AttachmentBuilder) Pretext(v string) *AttachmentBuilder {
	b.pretext = v
	return b
}

func (b *AttachmentBuilder) Text(v string) *AttachmentBuilder {
	b.text = v
	return b
}

func (b *AttachmentBuilder) Timestamp(v Timestamp) *AttachmentBuilder {
	b.timestamp = v
	return b
}

func (b *AttachmentBuilder) Title(v string) *AttachmentBuilder {
	b.title = v
	return b
}

func (b *AttachmentBuilder) TitleLink(v string) *AttachmentBuilder {
	b.titleLink = v
	return b
}

func (b *AttachmentBuilder) Do() (*Attachment, error) {
	var v Attachment
	v.actions = b.actions
	v.attachmentType = b.attachmentType
	v.authorName = b.authorName
	v.authorLink = b.authorLink
	v.authorIcon = b.authorIcon
	v.callbackId = b.callbackId
	v.color = b.color
	v.fallback = b.fallback
	v.fields = b.fields
	v.footer = b.footer
	v.footerIcon = b.footerIcon
	v.imageUrl = b.imageUrl
	v.thumbUrl = b.thumbUrl
	v.pretext = b.pretext
	v.text = b.text
	v.timestamp = b.timestamp
	v.title = b.title
	v.titleLink = b.titleLink
	return &v, nil
}

func (b *Attachment) Actions() ActionList {
	return b.actions
}

func (b *Attachment) AttachmentType() string {
	return b.attachmentType
}

func (b *Attachment) AuthorName() string {
	return b.authorName
}

func (b *Attachment) AuthorLink() string {
	return b.authorLink
}

func (b *Attachment) AuthorIcon() string {
	return b.authorIcon
}

func (b *Attachment) CallbackId() string {
	return b.callbackId
}

func (b *Attachment) Color() string {
	return b.color
}

func (b *Attachment) Fallback() string {
	return b.fallback
}

func (b *Attachment) Fields() AttachmentFieldList {
	return b.fields
}

func (b *Attachment) Footer() string {
	return b.footer
}

func (b *Attachment) FooterIcon() string {
	return b.footerIcon
}

func (b *Attachment) ImageUrl() string {
	return b.imageUrl
}

func (b *Attachment) ThumbUrl() string {
	return b.thumbUrl
}

func (b *Attachment) Pretext() string {
	return b.pretext
}

func (b *Attachment) Text() string {
	return b.text
}

func (b *Attachment) Timestamp() Timestamp {
	return b.timestamp
}

func (b *Attachment) Title() string {
	return b.title
}

func (b *Attachment) TitleLink() string {
	return b.titleLink
}

func BuildAttachmentField() *AttachmentFieldBuilder {
	var b AttachmentFieldBuilder
	return &b
}

func (b *AttachmentFieldBuilder) Title(v string) *AttachmentFieldBuilder {
	b.title = v
	return b
}

func (b *AttachmentFieldBuilder) Value(v string) *AttachmentFieldBuilder {
	b.value = v
	return b
}

func (b *AttachmentFieldBuilder) Short(v string) *AttachmentFieldBuilder {
	b.short = v
	return b
}

func (b *AttachmentFieldBuilder) Do() (*AttachmentField, error) {
	var v AttachmentField
	v.title = b.title
	v.value = b.value
	v.short = b.short
	return &v, nil
}

func (b *AttachmentField) Title() string {
	return b.title
}

func (b *AttachmentField) Value() string {
	return b.value
}

func (b *AttachmentField) Short() string {
	return b.short
}

func BuildBot() *BotBuilder {
	var b BotBuilder
	return &b
}

func (b *BotBuilder) Id(v string) *BotBuilder {
	b.id = v
	return b
}

func (b *BotBuilder) AppId(v string) *BotBuilder {
	b.appId = v
	return b
}

func (b *BotBuilder) Deleted(v bool) *BotBuilder {
	b.deleted = v
	return b
}

func (b *BotBuilder) Name(v string) *BotBuilder {
	b.name = v
	return b
}

func (b *BotBuilder) Icons(v Icons) *BotBuilder {
	b.icons = v
	return b
}

func (b *BotBuilder) Do() (*Bot, error) {
	var v Bot
	v.id = b.id
	v.appId = b.appId
	v.deleted = b.deleted
	v.name = b.name
	v.icons = b.icons
	return &v, nil
}

func (b *Bot) Id() string {
	return b.id
}

func (b *Bot) AppId() string {
	return b.appId
}

func (b *Bot) Deleted() bool {
	return b.deleted
}

func (b *Bot) Name() string {
	return b.name
}

func (b *Bot) Icons() Icons {
	return b.icons
}

func BuildChannel() *ChannelBuilder {
	var b ChannelBuilder
	return &b
}

func (b *ChannelBuilder) Id(v string) *ChannelBuilder {
	b.id = v
	return b
}

func (b *ChannelBuilder) Created(v EpochTime) *ChannelBuilder {
	b.created = v
	return b
}

func (b *ChannelBuilder) IsOpen(v bool) *ChannelBuilder {
	b.isOpen = v
	return b
}

func (b *ChannelBuilder) LastRead(v string) *ChannelBuilder {
	b.lastRead = v
	return b
}

func (b *ChannelBuilder) Latest(v *Message) *ChannelBuilder {
	b.latest = v
	return b
}

func (b *ChannelBuilder) UnreadCount(v int) *ChannelBuilder {
	b.unreadCount = v
	return b
}

func (b *ChannelBuilder) UnreadCountDisplay(v int) *ChannelBuilder {
	b.unreadCountDisplay = v
	return b
}

func (b *ChannelBuilder) Creator(v string) *ChannelBuilder {
	b.creator = v
	return b
}

func (b *ChannelBuilder) IsArchived(v bool) *ChannelBuilder {
	b.isArchived = v
	return b
}

func (b *ChannelBuilder) IsGroup(v bool) *ChannelBuilder {
	b.isGroup = v
	return b
}

func (b *ChannelBuilder) IsMpim(v bool) *ChannelBuilder {
	b.isMpim = v
	return b
}

func (b *ChannelBuilder) Members(v ...string) *ChannelBuilder {
	b.members = v
	return b
}

func (b *ChannelBuilder) Name(v string) *ChannelBuilder {
	b.name = v
	return b
}

func (b *ChannelBuilder) NameNormalized(v string) *ChannelBuilder {
	b.nameNormalized = v
	return b
}

func (b *ChannelBuilder) NumMembers(v int) *ChannelBuilder {
	b.numMembers = v
	return b
}

func (b *ChannelBuilder) PreviousNames(v ...string) *ChannelBuilder {
	b.previousNames = v
	return b
}

func (b *ChannelBuilder) Purpose(v Purpose) *ChannelBuilder {
	b.purpose = v
	return b
}

func (b *ChannelBuilder) Topic(v Topic) *ChannelBuilder {
	b.topic = v
	return b
}

func (b *ChannelBuilder) IsChannel(v bool) *ChannelBuilder {
	b.isChannel = v
	return b
}

func (b *ChannelBuilder) IsGeneral(v bool) *ChannelBuilder {
	b.isGeneral = v
	return b
}

func (b *ChannelBuilder) IsMember(v bool) *ChannelBuilder {
	b.isMember = v
	return b
}

func (b *ChannelBuilder) IsOrgShared(v bool) *ChannelBuilder {
	b.isOrgShared = v
	return b
}

func (b *ChannelBuilder) IsShared(v bool) *ChannelBuilder {
	b.isShared = v
	return b
}

func (b *ChannelBuilder) Do() (*Channel, error) {
	var v Channel
	v.id = b.id
	v.created = b.created
	v.isOpen = b.isOpen
	v.lastRead = b.lastRead
	v.latest = b.latest
	v.unreadCount = b.unreadCount
	v.unreadCountDisplay = b.unreadCountDisplay
	v.creator = b.creator
	v.isArchived = b.isArchived
	v.isGroup = b.isGroup
	v.isMpim = b.isMpim
	v.members = b.members
	v.name = b.name
	v.nameNormalized = b.nameNormalized
	v.numMembers = b.numMembers
	v.previousNames = b.previousNames
	v.purpose = b.purpose
	v.topic = b.topic
	v.isChannel = b.isChannel
	v.isGeneral = b.isGeneral
	v.isMember = b.isMember
	v.isOrgShared = b.isOrgShared
	v.isShared = b.isShared
	return &v, nil
}

func (b *Channel) Id() string {
	return b.id
}

func (b *Channel) Created() EpochTime {
	return b.created
}

func (b *Channel) IsOpen() bool {
	return b.isOpen
}

func (b *Channel) LastRead() string {
	return b.lastRead
}

func (b *Channel) Latest() *Message {
	return b.latest
}

func (b *Channel) UnreadCount() int {
	return b.unreadCount
}

func (b *Channel) UnreadCountDisplay() int {
	return b.unreadCountDisplay
}

func (b *Channel) Creator() string {
	return b.creator
}

func (b *Channel) IsArchived() bool {
	return b.isArchived
}

func (b *Channel) IsGroup() bool {
	return b.isGroup
}

func (b *Channel) IsMpim() bool {
	return b.isMpim
}

func (b *Channel) Members() []string {
	return b.members
}

func (b *Channel) Name() string {
	return b.name
}

func (b *Channel) NameNormalized() string {
	return b.nameNormalized
}

func (b *Channel) NumMembers() int {
	return b.numMembers
}

func (b *Channel) PreviousNames() []string {
	return b.previousNames
}

func (b *Channel) Purpose() Purpose {
	return b.purpose
}

func (b *Channel) Topic() Topic {
	return b.topic
}

func (b *Channel) IsChannel() bool {
	return b.isChannel
}

func (b *Channel) IsGeneral() bool {
	return b.isGeneral
}

func (b *Channel) IsMember() bool {
	return b.isMember
}

func (b *Channel) IsOrgShared() bool {
	return b.isOrgShared
}

func (b *Channel) IsShared() bool {
	return b.isShared
}

func BuildConfirmation() *ConfirmationBuilder {
	var b ConfirmationBuilder
	return &b
}

func (b *ConfirmationBuilder) Title(v string) *ConfirmationBuilder {
	b.title = v
	return b
}

func (b *ConfirmationBuilder) Text(v string) *ConfirmationBuilder {
	b.text = v
	return b
}

func (b *ConfirmationBuilder) OkText(v string) *ConfirmationBuilder {
	b.okText = v
	return b
}

func (b *ConfirmationBuilder) DismissText(v string) *ConfirmationBuilder {
	b.dismissText = v
	return b
}

func (b *ConfirmationBuilder) Do() (*Confirmation, error) {
	var v Confirmation
	v.title = b.title
	v.text = b.text
	v.okText = b.okText
	v.dismissText = b.dismissText
	return &v, nil
}

func (b *Confirmation) Title() string {
	return b.title
}

func (b *Confirmation) Text() string {
	return b.text
}

func (b *Confirmation) OkText() string {
	return b.okText
}

func (b *Confirmation) DismissText() string {
	return b.dismissText
}

func BuildConversation() *ConversationBuilder {
	var b ConversationBuilder
	return &b
}

func (b *ConversationBuilder) Id(v string) *ConversationBuilder {
	b.id = v
	return b
}

func (b *ConversationBuilder) Created(v EpochTime) *ConversationBuilder {
	b.created = v
	return b
}

func (b *ConversationBuilder) IsOpen(v bool) *ConversationBuilder {
	b.isOpen = v
	return b
}

func (b *ConversationBuilder) LastRead(v string) *ConversationBuilder {
	b.lastRead = v
	return b
}

func (b *ConversationBuilder) Latest(v *Message) *ConversationBuilder {
	b.latest = v
	return b
}

func (b *ConversationBuilder) UnreadCount(v int) *ConversationBuilder {
	b.unreadCount = v
	return b
}

func (b *ConversationBuilder) UnreadCountDisplay(v int) *ConversationBuilder {
	b.unreadCountDisplay = v
	return b
}

func (b *ConversationBuilder) Do() (*Conversation, error) {
	var v Conversation
	v.id = b.id
	v.created = b.created
	v.isOpen = b.isOpen
	v.lastRead = b.lastRead
	v.latest = b.latest
	v.unreadCount = b.unreadCount
	v.unreadCountDisplay = b.unreadCountDisplay
	return &v, nil
}

func (b *Conversation) Id() string {
	return b.id
}

func (b *Conversation) Created() EpochTime {
	return b.created
}

func (b *Conversation) IsOpen() bool {
	return b.isOpen
}

func (b *Conversation) LastRead() string {
	return b.lastRead
}

func (b *Conversation) Latest() *Message {
	return b.latest
}

func (b *Conversation) UnreadCount() int {
	return b.unreadCount
}

func (b *Conversation) UnreadCountDisplay() int {
	return b.unreadCountDisplay
}

func BuildDialog() *DialogBuilder {
	var b DialogBuilder
	return &b
}

func (b *DialogBuilder) CallbackId(v string) *DialogBuilder {
	b.callbackId = v
	return b
}

func (b *DialogBuilder) Title(v string) *DialogBuilder {
	b.title = v
	return b
}

func (b *DialogBuilder) SubmitLabel(v string) *DialogBuilder {
	b.submitLabel = v
	return b
}

func (b *DialogBuilder) Elements(v ...*DialogElement) *DialogBuilder {
	b.elements = v
	return b
}

func (b *DialogBuilder) Do() (*Dialog, error) {
	var v Dialog
	v.callbackId = b.callbackId
	v.title = b.title
	v.submitLabel = b.submitLabel
	v.elements = b.elements
	return &v, nil
}

func (b *Dialog) CallbackId() string {
	return b.callbackId
}

func (b *Dialog) Title() string {
	return b.title
}

func (b *Dialog) SubmitLabel() string {
	return b.submitLabel
}

func (b *Dialog) Elements() DialogElementList {
	return b.elements
}

func BuildDialogElement() *DialogElementBuilder {
	var b DialogElementBuilder
	return &b
}

func (b *DialogElementBuilder) Label(v string) *DialogElementBuilder {
	b.label = v
	return b
}

func (b *DialogElementBuilder) Type(v string) *DialogElementBuilder {
	b.typ = v
	return b
}

func (b *DialogElementBuilder) Name(v string) *DialogElementBuilder {
	b.name = v
	return b
}

func (b *DialogElementBuilder) Hint(v string) *DialogElementBuilder {
	b.hint = v
	return b
}

func (b *DialogElementBuilder) MaxLength(v string) *DialogElementBuilder {
	b.maxLength = v
	return b
}

func (b *DialogElementBuilder) MinLength(v string) *DialogElementBuilder {
	b.minLength = v
	return b
}

func (b *DialogElementBuilder) Optional(v bool) *DialogElementBuilder {
	b.optional = v
	return b
}

func (b *DialogElementBuilder) Placeholder(v string) *DialogElementBuilder {
	b.placeholder = v
	return b
}

func (b *DialogElementBuilder) Subtype(v string) *DialogElementBuilder {
	b.subtype = v
	return b
}

func (b *DialogElementBuilder) Value(v string) *DialogElementBuilder {
	b.value = v
	return b
}

func (b *DialogElementBuilder) Do() (*DialogElement, error) {
	var v DialogElement
	v.label = b.label
	v.typ = b.typ
	v.name = b.name
	v.hint = b.hint
	v.maxLength = b.maxLength
	v.minLength = b.minLength
	v.optional = b.optional
	v.placeholder = b.placeholder
	v.subtype = b.subtype
	v.value = b.value
	return &v, nil
}

func (b *DialogElement) Label() string {
	return b.label
}

func (b *DialogElement) Type() string {
	return b.typ
}

func (b *DialogElement) Name() string {
	return b.name
}

func (b *DialogElement) Hint() string {
	return b.hint
}

func (b *DialogElement) MaxLength() string {
	return b.maxLength
}

func (b *DialogElement) MinLength() string {
	return b.minLength
}

func (b *DialogElement) Optional() bool {
	return b.optional
}

func (b *DialogElement) Placeholder() string {
	return b.placeholder
}

func (b *DialogElement) Subtype() string {
	return b.subtype
}

func (b *DialogElement) Value() string {
	return b.value
}

func BuildEdited() *EditedBuilder {
	var b EditedBuilder
	return &b
}

func (b *EditedBuilder) Ts(v string) *EditedBuilder {
	b.ts = v
	return b
}

func (b *EditedBuilder) User(v string) *EditedBuilder {
	b.user = v
	return b
}

func (b *EditedBuilder) Do() (*Edited, error) {
	var v Edited
	v.ts = b.ts
	v.user = b.user
	return &v, nil
}

func (b *Edited) Ts() string {
	return b.ts
}

func (b *Edited) User() string {
	return b.user
}

func BuildFile() *FileBuilder {
	var b FileBuilder
	return &b
}

func (b *FileBuilder) Id(v string) *FileBuilder {
	b.id = v
	return b
}

func (b *FileBuilder) Name(v string) *FileBuilder {
	b.name = v
	return b
}

func (b *FileBuilder) User(v string) *FileBuilder {
	b.user = v
	return b
}

func (b *FileBuilder) Created(v int) *FileBuilder {
	b.created = v
	return b
}

func (b *FileBuilder) Timestamp(v int) *FileBuilder {
	b.timestamp = v
	return b
}

func (b *FileBuilder) Updated(v int) *FileBuilder {
	b.updated = v
	return b
}

func (b *FileBuilder) MimeType(v string) *FileBuilder {
	b.mimetype = v
	return b
}

func (b *FileBuilder) FileType(v string) *FileBuilder {
	b.filetype = v
	return b
}

func (b *FileBuilder) PrettyType(v string) *FileBuilder {
	b.prettyType = v
	return b
}

func (b *FileBuilder) Mode(v string) *FileBuilder {
	b.mode = v
	return b
}

func (b *FileBuilder) Editable(v bool) *FileBuilder {
	b.editable = v
	return b
}

func (b *FileBuilder) IsExternal(v bool) *FileBuilder {
	b.isExternal = v
	return b
}

func (b *FileBuilder) ExternalType(v string) *FileBuilder {
	b.externalType = v
	return b
}

func (b *FileBuilder) Size(v int) *FileBuilder {
	b.size = v
	return b
}

func (b *FileBuilder) Url(v string) *FileBuilder {
	b.url = v
	return b
}

func (b *FileBuilder) UrlDownload(v string) *FileBuilder {
	b.urlDownload = v
	return b
}

func (b *FileBuilder) UrlPrivate(v string) *FileBuilder {
	b.urlPrivate = v
	return b
}

func (b *FileBuilder) UrlPrivateDownload(v string) *FileBuilder {
	b.urlPrivateDownload = v
	return b
}

func (b *FileBuilder) ImageExifRotation(v string) *FileBuilder {
	b.imageExifRotation = v
	return b
}

func (b *FileBuilder) OriginalWidth(v int) *FileBuilder {
	b.originalW = v
	return b
}

func (b *FileBuilder) OriginalHeight(v int) *FileBuilder {
	b.originalH = v
	return b
}

func (b *FileBuilder) Thumb64(v string) *FileBuilder {
	b.thumb64 = v
	return b
}

func (b *FileBuilder) Thumb80(v string) *FileBuilder {
	b.thumb80 = v
	return b
}

func (b *FileBuilder) Thumb160(v string) *FileBuilder {
	b.thumb160 = v
	return b
}

func (b *FileBuilder) Thumb360(v string) *FileBuilder {
	b.thumb360 = v
	return b
}

func (b *FileBuilder) Thumb360Gif(v string) *FileBuilder {
	b.thumb360Gif = v
	return b
}

func (b *FileBuilder) Thumb360W(v int) *FileBuilder {
	b.thumb360W = v
	return b
}

func (b *FileBuilder) Thumb360H(v int) *FileBuilder {
	b.thumb360H = v
	return b
}

func (b *FileBuilder) Thumb480(v string) *FileBuilder {
	b.thumb480 = v
	return b
}

func (b *FileBuilder) Thumb480W(v int) *FileBuilder {
	b.thumb480W = v
	return b
}

func (b *FileBuilder) Thumb480H(v int) *FileBuilder {
	b.thumb480H = v
	return b
}

func (b *FileBuilder) Thumb720(v string) *FileBuilder {
	b.thumb720 = v
	return b
}

func (b *FileBuilder) Thumb720W(v int) *FileBuilder {
	b.thumb720W = v
	return b
}

func (b *FileBuilder) Thumb720H(v int) *FileBuilder {
	b.thumb720H = v
	return b
}

func (b *FileBuilder) Thumb960(v string) *FileBuilder {
	b.thumb960 = v
	return b
}

func (b *FileBuilder) Thumb960W(v int) *FileBuilder {
	b.thumb960W = v
	return b
}

func (b *FileBuilder) Thumb960H(v int) *FileBuilder {
	b.thumb960H = v
	return b
}

func (b *FileBuilder) Thumb1024(v string) *FileBuilder {
	b.thumb1024 = v
	return b
}

func (b *FileBuilder) Thumb1024W(v int) *FileBuilder {
	b.thumb1024W = v
	return b
}

func (b *FileBuilder) Thumb1024H(v int) *FileBuilder {
	b.thumb1024H = v
	return b
}

func (b *FileBuilder) Permalink(v string) *FileBuilder {
	b.permalink = v
	return b
}

func (b *FileBuilder) PermalinkPublic(v string) *FileBuilder {
	b.permalinkPublic = v
	return b
}

func (b *FileBuilder) EditLink(v string) *FileBuilder {
	b.editLink = v
	return b
}

func (b *FileBuilder) Preview(v string) *FileBuilder {
	b.preview = v
	return b
}

func (b *FileBuilder) PreviewHighlight(v string) *FileBuilder {
	b.previewHighlight = v
	return b
}

func (b *FileBuilder) Lines(v int) *FileBuilder {
	b.lines = v
	return b
}

func (b *FileBuilder) LinesMore(v int) *FileBuilder {
	b.linesMore = v
	return b
}

func (b *FileBuilder) IsPublic(v bool) *FileBuilder {
	b.isPublic = v
	return b
}

func (b *FileBuilder) PublicUrlShared(v bool) *FileBuilder {
	b.publicUrlShared = v
	return b
}

func (b *FileBuilder) Channels(v ...string) *FileBuilder {
	b.channels = v
	return b
}

func (b *FileBuilder) Groups(v ...string) *FileBuilder {
	b.groups = v
	return b
}

func (b *FileBuilder) IMs(v ...string) *FileBuilder {
	b.ims = v
	return b
}

func (b *FileBuilder) InitialComment(v Comment) *FileBuilder {
	b.initialComment = v
	return b
}

func (b *FileBuilder) CommentCount(v int) *FileBuilder {
	b.commentCount = v
	return b
}

func (b *FileBuilder) NumStars(v int) *FileBuilder {
	b.numStars = v
	return b
}

func (b *FileBuilder) IsStarred(v bool) *FileBuilder {
	b.isStarred = v
	return b
}

func (b *FileBuilder) Title(v string) *FileBuilder {
	b.title = v
	return b
}

func (b *FileBuilder) Reactions(v ...*Reaction) *FileBuilder {
	b.reactions = v
	return b
}

func (b *FileBuilder) Do() (*File, error) {
	var v File
	v.id = b.id
	v.name = b.name
	v.user = b.user
	v.created = b.created
	v.timestamp = b.timestamp
	v.updated = b.updated
	v.mimetype = b.mimetype
	v.filetype = b.filetype
	v.prettyType = b.prettyType
	v.mode = b.mode
	v.editable = b.editable
	v.isExternal = b.isExternal
	v.externalType = b.externalType
	v.size = b.size
	v.url = b.url
	v.urlDownload = b.urlDownload
	v.urlPrivate = b.urlPrivate
	v.urlPrivateDownload = b.urlPrivateDownload
	v.imageExifRotation = b.imageExifRotation
	v.originalW = b.originalW
	v.originalH = b.originalH
	v.thumb64 = b.thumb64
	v.thumb80 = b.thumb80
	v.thumb160 = b.thumb160
	v.thumb360 = b.thumb360
	v.thumb360Gif = b.thumb360Gif
	v.thumb360W = b.thumb360W
	v.thumb360H = b.thumb360H
	v.thumb480 = b.thumb480
	v.thumb480W = b.thumb480W
	v.thumb480H = b.thumb480H
	v.thumb720 = b.thumb720
	v.thumb720W = b.thumb720W
	v.thumb720H = b.thumb720H
	v.thumb960 = b.thumb960
	v.thumb960W = b.thumb960W
	v.thumb960H = b.thumb960H
	v.thumb1024 = b.thumb1024
	v.thumb1024W = b.thumb1024W
	v.thumb1024H = b.thumb1024H
	v.permalink = b.permalink
	v.permalinkPublic = b.permalinkPublic
	v.editLink = b.editLink
	v.preview = b.preview
	v.previewHighlight = b.previewHighlight
	v.lines = b.lines
	v.linesMore = b.linesMore
	v.isPublic = b.isPublic
	v.publicUrlShared = b.publicUrlShared
	v.channels = b.channels
	v.groups = b.groups
	v.ims = b.ims
	v.initialComment = b.initialComment
	v.commentCount = b.commentCount
	v.numStars = b.numStars
	v.isStarred = b.isStarred
	v.title = b.title
	v.reactions = b.reactions
	return &v, nil
}

func (b *File) Id() string {
	return b.id
}

func (b *File) Name() string {
	return b.name
}

func (b *File) User() string {
	return b.user
}

func (b *File) Created() int {
	return b.created
}

func (b *File) Timestamp() int {
	return b.timestamp
}

func (b *File) Updated() int {
	return b.updated
}

func (b *File) MimeType() string {
	return b.mimetype
}

func (b *File) FileType() string {
	return b.filetype
}

func (b *File) PrettyType() string {
	return b.prettyType
}

func (b *File) Mode() string {
	return b.mode
}

func (b *File) Editable() bool {
	return b.editable
}

func (b *File) IsExternal() bool {
	return b.isExternal
}

func (b *File) ExternalType() string {
	return b.externalType
}

func (b *File) Size() int {
	return b.size
}

func (b *File) Url() string {
	return b.url
}

func (b *File) UrlDownload() string {
	return b.urlDownload
}

func (b *File) UrlPrivate() string {
	return b.urlPrivate
}

func (b *File) UrlPrivateDownload() string {
	return b.urlPrivateDownload
}

func (b *File) ImageExifRotation() string {
	return b.imageExifRotation
}

func (b *File) OriginalWidth() int {
	return b.originalW
}

func (b *File) OriginalHeight() int {
	return b.originalH
}

func (b *File) Thumb64() string {
	return b.thumb64
}

func (b *File) Thumb80() string {
	return b.thumb80
}

func (b *File) Thumb160() string {
	return b.thumb160
}

func (b *File) Thumb360() string {
	return b.thumb360
}

func (b *File) Thumb360Gif() string {
	return b.thumb360Gif
}

func (b *File) Thumb360W() int {
	return b.thumb360W
}

func (b *File) Thumb360H() int {
	return b.thumb360H
}

func (b *File) Thumb480() string {
	return b.thumb480
}

func (b *File) Thumb480W() int {
	return b.thumb480W
}

func (b *File) Thumb480H() int {
	return b.thumb480H
}

func (b *File) Thumb720() string {
	return b.thumb720
}

func (b *File) Thumb720W() int {
	return b.thumb720W
}

func (b *File) Thumb720H() int {
	return b.thumb720H
}

func (b *File) Thumb960() string {
	return b.thumb960
}

func (b *File) Thumb960W() int {
	return b.thumb960W
}

func (b *File) Thumb960H() int {
	return b.thumb960H
}

func (b *File) Thumb1024() string {
	return b.thumb1024
}

func (b *File) Thumb1024W() int {
	return b.thumb1024W
}

func (b *File) Thumb1024H() int {
	return b.thumb1024H
}

func (b *File) Permalink() string {
	return b.permalink
}

func (b *File) PermalinkPublic() string {
	return b.permalinkPublic
}

func (b *File) EditLink() string {
	return b.editLink
}

func (b *File) Preview() string {
	return b.preview
}

func (b *File) PreviewHighlight() string {
	return b.previewHighlight
}

func (b *File) Lines() int {
	return b.lines
}

func (b *File) LinesMore() int {
	return b.linesMore
}

func (b *File) IsPublic() bool {
	return b.isPublic
}

func (b *File) PublicUrlShared() bool {
	return b.publicUrlShared
}

func (b *File) Channels() []string {
	return b.channels
}

func (b *File) Groups() []string {
	return b.groups
}

func (b *File) IMs() []string {
	return b.ims
}

func (b *File) InitialComment() Comment {
	return b.initialComment
}

func (b *File) CommentCount() int {
	return b.commentCount
}

func (b *File) NumStars() int {
	return b.numStars
}

func (b *File) IsStarred() bool {
	return b.isStarred
}

func (b *File) Title() string {
	return b.title
}

func (b *File) Reactions() ReactionList {
	return b.reactions
}

func BuildGroup() *GroupBuilder {
	var b GroupBuilder
	return &b
}

func (b *GroupBuilder) Id(v string) *GroupBuilder {
	b.id = v
	return b
}

func (b *GroupBuilder) Created(v EpochTime) *GroupBuilder {
	b.created = v
	return b
}

func (b *GroupBuilder) IsOpen(v bool) *GroupBuilder {
	b.isOpen = v
	return b
}

func (b *GroupBuilder) LastRead(v string) *GroupBuilder {
	b.lastRead = v
	return b
}

func (b *GroupBuilder) Latest(v *Message) *GroupBuilder {
	b.latest = v
	return b
}

func (b *GroupBuilder) UnreadCount(v int) *GroupBuilder {
	b.unreadCount = v
	return b
}

func (b *GroupBuilder) UnreadCountDisplay(v int) *GroupBuilder {
	b.unreadCountDisplay = v
	return b
}

func (b *GroupBuilder) Creator(v string) *GroupBuilder {
	b.creator = v
	return b
}

func (b *GroupBuilder) IsArchived(v bool) *GroupBuilder {
	b.isArchived = v
	return b
}

func (b *GroupBuilder) IsGroup(v bool) *GroupBuilder {
	b.isGroup = v
	return b
}

func (b *GroupBuilder) IsMpim(v bool) *GroupBuilder {
	b.isMpim = v
	return b
}

func (b *GroupBuilder) Members(v ...string) *GroupBuilder {
	b.members = v
	return b
}

func (b *GroupBuilder) Name(v string) *GroupBuilder {
	b.name = v
	return b
}

func (b *GroupBuilder) NameNormalized(v string) *GroupBuilder {
	b.nameNormalized = v
	return b
}

func (b *GroupBuilder) NumMembers(v int) *GroupBuilder {
	b.numMembers = v
	return b
}

func (b *GroupBuilder) PreviousNames(v ...string) *GroupBuilder {
	b.previousNames = v
	return b
}

func (b *GroupBuilder) Purpose(v Purpose) *GroupBuilder {
	b.purpose = v
	return b
}

func (b *GroupBuilder) Topic(v Topic) *GroupBuilder {
	b.topic = v
	return b
}

func (b *GroupBuilder) Do() (*Group, error) {
	var v Group
	v.id = b.id
	v.created = b.created
	v.isOpen = b.isOpen
	v.lastRead = b.lastRead
	v.latest = b.latest
	v.unreadCount = b.unreadCount
	v.unreadCountDisplay = b.unreadCountDisplay
	v.creator = b.creator
	v.isArchived = b.isArchived
	v.isGroup = b.isGroup
	v.isMpim = b.isMpim
	v.members = b.members
	v.name = b.name
	v.nameNormalized = b.nameNormalized
	v.numMembers = b.numMembers
	v.previousNames = b.previousNames
	v.purpose = b.purpose
	v.topic = b.topic
	return &v, nil
}

func (b *Group) Id() string {
	return b.id
}

func (b *Group) Created() EpochTime {
	return b.created
}

func (b *Group) IsOpen() bool {
	return b.isOpen
}

func (b *Group) LastRead() string {
	return b.lastRead
}

func (b *Group) Latest() *Message {
	return b.latest
}

func (b *Group) UnreadCount() int {
	return b.unreadCount
}

func (b *Group) UnreadCountDisplay() int {
	return b.unreadCountDisplay
}

func (b *Group) Creator() string {
	return b.creator
}

func (b *Group) IsArchived() bool {
	return b.isArchived
}

func (b *Group) IsGroup() bool {
	return b.isGroup
}

func (b *Group) IsMpim() bool {
	return b.isMpim
}

func (b *Group) Members() []string {
	return b.members
}

func (b *Group) Name() string {
	return b.name
}

func (b *Group) NameNormalized() string {
	return b.nameNormalized
}

func (b *Group) NumMembers() int {
	return b.numMembers
}

func (b *Group) PreviousNames() []string {
	return b.previousNames
}

func (b *Group) Purpose() Purpose {
	return b.purpose
}

func (b *Group) Topic() Topic {
	return b.topic
}

func BuildIcons() *IconsBuilder {
	var b IconsBuilder
	return &b
}

func (b *IconsBuilder) Image36(v string) *IconsBuilder {
	b.image36 = v
	return b
}

func (b *IconsBuilder) Image48(v string) *IconsBuilder {
	b.image48 = v
	return b
}

func (b *IconsBuilder) Image72(v string) *IconsBuilder {
	b.image72 = v
	return b
}

func (b *IconsBuilder) Do() (*Icons, error) {
	var v Icons
	v.image36 = b.image36
	v.image48 = b.image48
	v.image72 = b.image72
	return &v, nil
}

func (b *Icons) Image36() string {
	return b.image36
}

func (b *Icons) Image48() string {
	return b.image48
}

func (b *Icons) Image72() string {
	return b.image72
}

func BuildMessage() *MessageBuilder {
	var b MessageBuilder
	return &b
}

func (b *MessageBuilder) Attachments(v ...*Attachment) *MessageBuilder {
	b.attachments = v
	return b
}

func (b *MessageBuilder) Channel(v string) *MessageBuilder {
	b.channel = v
	return b
}

func (b *MessageBuilder) Edited(v *Edited) *MessageBuilder {
	b.edited = v
	return b
}

func (b *MessageBuilder) IsStarred(v bool) *MessageBuilder {
	b.isStarred = v
	return b
}

func (b *MessageBuilder) PinnedTo(v ...string) *MessageBuilder {
	b.pinnedTo = v
	return b
}

func (b *MessageBuilder) Text(v string) *MessageBuilder {
	b.text = v
	return b
}

func (b *MessageBuilder) Timestamp(v string) *MessageBuilder {
	b.ts = v
	return b
}

func (b *MessageBuilder) Type(v string) *MessageBuilder {
	b.typ = v
	return b
}

func (b *MessageBuilder) User(v string) *MessageBuilder {
	b.user = v
	return b
}

func (b *MessageBuilder) Subtype(v string) *MessageBuilder {
	b.subtype = v
	return b
}

func (b *MessageBuilder) Hidden(v bool) *MessageBuilder {
	b.hidden = v
	return b
}

func (b *MessageBuilder) DeletedTimestamkp(v string) *MessageBuilder {
	b.deletedTs = v
	return b
}

func (b *MessageBuilder) EventTimestamkp(v string) *MessageBuilder {
	b.eventTs = v
	return b
}

func (b *MessageBuilder) BotId(v string) *MessageBuilder {
	b.botId = v
	return b
}

func (b *MessageBuilder) Username(v string) *MessageBuilder {
	b.username = v
	return b
}

func (b *MessageBuilder) Icons(v *Icons) *MessageBuilder {
	b.icons = v
	return b
}

func (b *MessageBuilder) Inviter(v string) *MessageBuilder {
	b.inviter = v
	return b
}

func (b *MessageBuilder) Topic(v string) *MessageBuilder {
	b.topic = v
	return b
}

func (b *MessageBuilder) Purpose(v string) *MessageBuilder {
	b.purpose = v
	return b
}

func (b *MessageBuilder) Name(v string) *MessageBuilder {
	b.name = v
	return b
}

func (b *MessageBuilder) OldName(v string) *MessageBuilder {
	b.oldName = v
	return b
}

func (b *MessageBuilder) Members(v ...string) *MessageBuilder {
	b.members = v
	return b
}

func (b *MessageBuilder) Upload(v bool) *MessageBuilder {
	b.upload = v
	return b
}

func (b *MessageBuilder) Comment(v *Comment) *MessageBuilder {
	b.comment = v
	return b
}

func (b *MessageBuilder) ItemType(v string) *MessageBuilder {
	b.itemType = v
	return b
}

func (b *MessageBuilder) ReplyTo(v int) *MessageBuilder {
	b.replyTo = v
	return b
}

func (b *MessageBuilder) Team(v string) *MessageBuilder {
	b.team = v
	return b
}

func (b *MessageBuilder) Reactions(v ...*Reaction) *MessageBuilder {
	b.reactions = v
	return b
}

func (b *MessageBuilder) Blocks(v ...Block) *MessageBuilder {
	b.blocks = v
	return b
}

func (b *MessageBuilder) Do() (*Message, error) {
	var v Message
	v.attachments = b.attachments
	v.channel = b.channel
	v.edited = b.edited
	v.isStarred = b.isStarred
	v.pinnedTo = b.pinnedTo
	v.text = b.text
	v.ts = b.ts
	v.typ = b.typ
	v.user = b.user
	v.subtype = b.subtype
	v.hidden = b.hidden
	v.deletedTs = b.deletedTs
	v.eventTs = b.eventTs
	v.botId = b.botId
	v.username = b.username
	v.icons = b.icons
	v.inviter = b.inviter
	v.topic = b.topic
	v.purpose = b.purpose
	v.name = b.name
	v.oldName = b.oldName
	v.members = b.members
	v.upload = b.upload
	v.comment = b.comment
	v.itemType = b.itemType
	v.replyTo = b.replyTo
	v.team = b.team
	v.reactions = b.reactions
	v.blocks = b.blocks
	return &v, nil
}

func (b *Message) Attachments() AttachmentList {
	return b.attachments
}

func (b *Message) Channel() string {
	return b.channel
}

func (b *Message) Edited() *Edited {
	return b.edited
}

func (b *Message) IsStarred() bool {
	return b.isStarred
}

func (b *Message) PinnedTo() []string {
	return b.pinnedTo
}

func (b *Message) Text() string {
	return b.text
}

func (b *Message) Timestamp() string {
	return b.ts
}

func (b *Message) Type() string {
	return b.typ
}

func (b *Message) User() string {
	return b.user
}

func (b *Message) Subtype() string {
	return b.subtype
}

func (b *Message) Hidden() bool {
	return b.hidden
}

func (b *Message) DeletedTimestamkp() string {
	return b.deletedTs
}

func (b *Message) EventTimestamkp() string {
	return b.eventTs
}

func (b *Message) BotId() string {
	return b.botId
}

func (b *Message) Username() string {
	return b.username
}

func (b *Message) Icons() *Icons {
	return b.icons
}

func (b *Message) Inviter() string {
	return b.inviter
}

func (b *Message) Topic() string {
	return b.topic
}

func (b *Message) Purpose() string {
	return b.purpose
}

func (b *Message) Name() string {
	return b.name
}

func (b *Message) OldName() string {
	return b.oldName
}

func (b *Message) Members() []string {
	return b.members
}

func (b *Message) Upload() bool {
	return b.upload
}

func (b *Message) Comment() *Comment {
	return b.comment
}

func (b *Message) ItemType() string {
	return b.itemType
}

func (b *Message) ReplyTo() int {
	return b.replyTo
}

func (b *Message) Team() string {
	return b.team
}

func (b *Message) Reactions() ReactionList {
	return b.reactions
}

func (b *Message) Blocks() BlockList {
	return b.blocks
}

func BuildOption() *OptionBuilder {
	var b OptionBuilder
	return &b
}

func (b *OptionBuilder) Text(v string) *OptionBuilder {
	b.text = v
	return b
}

func (b *OptionBuilder) Value(v string) *OptionBuilder {
	b.value = v
	return b
}

func (b *OptionBuilder) Description(v string) *OptionBuilder {
	b.description = v
	return b
}

func (b *OptionBuilder) Do() (*Option, error) {
	var v Option
	v.text = b.text
	v.value = b.value
	v.description = b.description
	return &v, nil
}

func (b *Option) Text() string {
	return b.text
}

func (b *Option) Value() string {
	return b.value
}

func (b *Option) Description() string {
	return b.description
}

func BuildOptionGroup() *OptionGroupBuilder {
	var b OptionGroupBuilder
	return &b
}

func (b *OptionGroupBuilder) Text(v string) *OptionGroupBuilder {
	b.text = v
	return b
}

func (b *OptionGroupBuilder) Options(v ...*Option) *OptionGroupBuilder {
	b.options = v
	return b
}

func (b *OptionGroupBuilder) Do() (*OptionGroup, error) {
	var v OptionGroup
	v.text = b.text
	v.options = b.options
	return &v, nil
}

func (b *OptionGroup) Text() string {
	return b.text
}

func (b *OptionGroup) Options() OptionList {
	return b.options
}

func BuildPurpose() *PurposeBuilder {
	var b PurposeBuilder
	return &b
}

func (b *PurposeBuilder) Value(v string) *PurposeBuilder {
	b.value = v
	return b
}

func (b *PurposeBuilder) Creator(v string) *PurposeBuilder {
	b.creator = v
	return b
}

func (b *PurposeBuilder) LastSet(v EpochTime) *PurposeBuilder {
	b.lastSet = v
	return b
}

func (b *PurposeBuilder) Do() (*Purpose, error) {
	var v Purpose
	v.value = b.value
	v.creator = b.creator
	v.lastSet = b.lastSet
	return &v, nil
}

func (b *Purpose) Value() string {
	return b.value
}

func (b *Purpose) Creator() string {
	return b.creator
}

func (b *Purpose) LastSet() EpochTime {
	return b.lastSet
}

func BuildReaction() *ReactionBuilder {
	var b ReactionBuilder
	return &b
}

func (b *ReactionBuilder) Count(v int) *ReactionBuilder {
	b.count = v
	return b
}

func (b *ReactionBuilder) Name(v string) *ReactionBuilder {
	b.name = v
	return b
}

func (b *ReactionBuilder) Users(v ...string) *ReactionBuilder {
	b.users = v
	return b
}

func (b *ReactionBuilder) Do() (*Reaction, error) {
	var v Reaction
	v.count = b.count
	v.name = b.name
	v.users = b.users
	return &v, nil
}

func (b *Reaction) Count() int {
	return b.count
}

func (b *Reaction) Name() string {
	return b.name
}

func (b *Reaction) Users() []string {
	return b.users
}

func BuildReminder() *ReminderBuilder {
	var b ReminderBuilder
	return &b
}

func (b *ReminderBuilder) Id(v string) *ReminderBuilder {
	b.id = v
	return b
}

func (b *ReminderBuilder) Creator(v string) *ReminderBuilder {
	b.creator = v
	return b
}

func (b *ReminderBuilder) User(v string) *ReminderBuilder {
	b.user = v
	return b
}

func (b *ReminderBuilder) Text(v string) *ReminderBuilder {
	b.text = v
	return b
}

func (b *ReminderBuilder) Recurring(v bool) *ReminderBuilder {
	b.recurring = v
	return b
}

func (b *ReminderBuilder) Time(v EpochTime) *ReminderBuilder {
	b.time = v
	return b
}

func (b *ReminderBuilder) CompleteTimestamp(v EpochTime) *ReminderBuilder {
	b.completeTimestamp = v
	return b
}

func (b *ReminderBuilder) Do() (*Reminder, error) {
	var v Reminder
	v.id = b.id
	v.creator = b.creator
	v.user = b.user
	v.text = b.text
	v.recurring = b.recurring
	v.time = b.time
	v.completeTimestamp = b.completeTimestamp
	return &v, nil
}

func (b *Reminder) Id() string {
	return b.id
}

func (b *Reminder) Creator() string {
	return b.creator
}

func (b *Reminder) User() string {
	return b.user
}

func (b *Reminder) Text() string {
	return b.text
}

func (b *Reminder) Recurring() bool {
	return b.recurring
}

func (b *Reminder) Time() EpochTime {
	return b.time
}

func (b *Reminder) CompleteTimestamp() EpochTime {
	return b.completeTimestamp
}

func BuildTeam() *TeamBuilder {
	var b TeamBuilder
	return &b
}

func (b *TeamBuilder) Id(v string) *TeamBuilder {
	b.id = v
	return b
}

func (b *TeamBuilder) Name(v string) *TeamBuilder {
	b.name = v
	return b
}

func (b *TeamBuilder) Domain(v string) *TeamBuilder {
	b.domain = v
	return b
}

func (b *TeamBuilder) EmailDomain(v string) *TeamBuilder {
	b.emailDomain = v
	return b
}

func (b *TeamBuilder) EnterpriseId(v string) *TeamBuilder {
	b.enterpriseId = v
	return b
}

func (b *TeamBuilder) EnterpriseName(v string) *TeamBuilder {
	b.enterpriseName = v
	return b
}

func (b *TeamBuilder) Icon(v map[string]interface{}) *TeamBuilder {
	b.icon = v
	return b
}

func (b *TeamBuilder) MsgEditWindowMins(v int) *TeamBuilder {
	b.msgEditWindowMins = v
	return b
}

func (b *TeamBuilder) OverStorageLimit(v bool) *TeamBuilder {
	b.overStorageLimit = v
	return b
}

func (b *TeamBuilder) Prefs(v interface{}) *TeamBuilder {
	b.prefs = v
	return b
}

func (b *TeamBuilder) Plan(v string) *TeamBuilder {
	b.plan = v
	return b
}

func (b *TeamBuilder) Do() (*Team, error) {
	var v Team
	v.id = b.id
	v.name = b.name
	v.domain = b.domain
	v.emailDomain = b.emailDomain
	v.enterpriseId = b.enterpriseId
	v.enterpriseName = b.enterpriseName
	v.icon = b.icon
	v.msgEditWindowMins = b.msgEditWindowMins
	v.overStorageLimit = b.overStorageLimit
	v.prefs = b.prefs
	v.plan = b.plan
	return &v, nil
}

func (b *Team) Id() string {
	return b.id
}

func (b *Team) Name() string {
	return b.name
}

func (b *Team) Domain() string {
	return b.domain
}

func (b *Team) EmailDomain() string {
	return b.emailDomain
}

func (b *Team) EnterpriseId() string {
	return b.enterpriseId
}

func (b *Team) EnterpriseName() string {
	return b.enterpriseName
}

func (b *Team) Icon() map[string]interface{} {
	return b.icon
}

func (b *Team) MsgEditWindowMins() int {
	return b.msgEditWindowMins
}

func (b *Team) OverStorageLimit() bool {
	return b.overStorageLimit
}

func (b *Team) Prefs() interface{} {
	return b.prefs
}

func (b *Team) Plan() string {
	return b.plan
}

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

func BuildThreadInfo() *ThreadInfoBuilder {
	var b ThreadInfoBuilder
	return &b
}

func (b *ThreadInfoBuilder) Complete(v bool) *ThreadInfoBuilder {
	b.complete = v
	return b
}

func (b *ThreadInfoBuilder) Count(v int) *ThreadInfoBuilder {
	b.count = v
	return b
}

func (b *ThreadInfoBuilder) Do() (*ThreadInfo, error) {
	var v ThreadInfo
	v.complete = b.complete
	v.count = b.count
	return &v, nil
}

func (b *ThreadInfo) Complete() bool {
	return b.complete
}

func (b *ThreadInfo) Count() int {
	return b.count
}

func BuildTopic() *TopicBuilder {
	var b TopicBuilder
	return &b
}

func (b *TopicBuilder) Value(v string) *TopicBuilder {
	b.value = v
	return b
}

func (b *TopicBuilder) Creator(v string) *TopicBuilder {
	b.creator = v
	return b
}

func (b *TopicBuilder) LastSet(v EpochTime) *TopicBuilder {
	b.lastSet = v
	return b
}

func (b *TopicBuilder) Do() (*Topic, error) {
	var v Topic
	v.value = b.value
	v.creator = b.creator
	v.lastSet = b.lastSet
	return &v, nil
}

func (b *Topic) Value() string {
	return b.value
}

func (b *Topic) Creator() string {
	return b.creator
}

func (b *Topic) LastSet() EpochTime {
	return b.lastSet
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

func BuildUsergroup() *UsergroupBuilder {
	var b UsergroupBuilder
	return &b
}

func (b *UsergroupBuilder) AutoProvision(v bool) *UsergroupBuilder {
	b.autoProvision = v
	return b
}

func (b *UsergroupBuilder) AutoType(v string) *UsergroupBuilder {
	b.autoType = v
	return b
}

func (b *UsergroupBuilder) CreatedBy(v string) *UsergroupBuilder {
	b.createdBy = v
	return b
}

func (b *UsergroupBuilder) DateCreate(v EpochTime) *UsergroupBuilder {
	b.dateCreate = v
	return b
}

func (b *UsergroupBuilder) DateDelete(v EpochTime) *UsergroupBuilder {
	b.dateDelete = v
	return b
}

func (b *UsergroupBuilder) DateUpdate(v EpochTime) *UsergroupBuilder {
	b.dateUpdate = v
	return b
}

func (b *UsergroupBuilder) DeletedBy(v string) *UsergroupBuilder {
	b.deletedBy = v
	return b
}

func (b *UsergroupBuilder) Description(v string) *UsergroupBuilder {
	b.description = v
	return b
}

func (b *UsergroupBuilder) EnterpriseSubteamId(v string) *UsergroupBuilder {
	b.enterpriseSubteamId = v
	return b
}

func (b *UsergroupBuilder) Handle(v string) *UsergroupBuilder {
	b.handle = v
	return b
}

func (b *UsergroupBuilder) Id(v string) *UsergroupBuilder {
	b.id = v
	return b
}

func (b *UsergroupBuilder) IsExternal(v bool) *UsergroupBuilder {
	b.isExternal = v
	return b
}

func (b *UsergroupBuilder) IsSubteam(v bool) *UsergroupBuilder {
	b.isSubteam = v
	return b
}

func (b *UsergroupBuilder) IsUsergroup(v bool) *UsergroupBuilder {
	b.isUsergroup = v
	return b
}

func (b *UsergroupBuilder) Name(v string) *UsergroupBuilder {
	b.name = v
	return b
}

func (b *UsergroupBuilder) Prefs(v *UsergroupPrefs) *UsergroupBuilder {
	b.prefs = v
	return b
}

func (b *UsergroupBuilder) TeamId(v string) *UsergroupBuilder {
	b.teamId = v
	return b
}

func (b *UsergroupBuilder) UpdatedBy(v string) *UsergroupBuilder {
	b.updatedBy = v
	return b
}

func (b *UsergroupBuilder) Users(v ...string) *UsergroupBuilder {
	b.users = v
	return b
}

func (b *UsergroupBuilder) UserCount(v int) *UsergroupBuilder {
	b.userCount = v
	return b
}

func (b *UsergroupBuilder) Do() (*Usergroup, error) {
	var v Usergroup
	v.autoProvision = b.autoProvision
	v.autoType = b.autoType
	v.createdBy = b.createdBy
	v.dateCreate = b.dateCreate
	v.dateDelete = b.dateDelete
	v.dateUpdate = b.dateUpdate
	v.deletedBy = b.deletedBy
	v.description = b.description
	v.enterpriseSubteamId = b.enterpriseSubteamId
	v.handle = b.handle
	v.id = b.id
	v.isExternal = b.isExternal
	v.isSubteam = b.isSubteam
	v.isUsergroup = b.isUsergroup
	v.name = b.name
	v.prefs = b.prefs
	v.teamId = b.teamId
	v.updatedBy = b.updatedBy
	v.users = b.users
	v.userCount = b.userCount
	return &v, nil
}

func (b *Usergroup) AutoProvision() bool {
	return b.autoProvision
}

func (b *Usergroup) AutoType() string {
	return b.autoType
}

func (b *Usergroup) CreatedBy() string {
	return b.createdBy
}

func (b *Usergroup) DateCreate() EpochTime {
	return b.dateCreate
}

func (b *Usergroup) DateDelete() EpochTime {
	return b.dateDelete
}

func (b *Usergroup) DateUpdate() EpochTime {
	return b.dateUpdate
}

func (b *Usergroup) DeletedBy() string {
	return b.deletedBy
}

func (b *Usergroup) Description() string {
	return b.description
}

func (b *Usergroup) EnterpriseSubteamId() string {
	return b.enterpriseSubteamId
}

func (b *Usergroup) Handle() string {
	return b.handle
}

func (b *Usergroup) Id() string {
	return b.id
}

func (b *Usergroup) IsExternal() bool {
	return b.isExternal
}

func (b *Usergroup) IsSubteam() bool {
	return b.isSubteam
}

func (b *Usergroup) IsUsergroup() bool {
	return b.isUsergroup
}

func (b *Usergroup) Name() string {
	return b.name
}

func (b *Usergroup) Prefs() *UsergroupPrefs {
	return b.prefs
}

func (b *Usergroup) TeamId() string {
	return b.teamId
}

func (b *Usergroup) UpdatedBy() string {
	return b.updatedBy
}

func (b *Usergroup) Users() []string {
	return b.users
}

func (b *Usergroup) UserCount() int {
	return b.userCount
}

func BuildUsergroupPrefs() *UsergroupPrefsBuilder {
	var b UsergroupPrefsBuilder
	return &b
}

func (b *UsergroupPrefsBuilder) Channels(v ...string) *UsergroupPrefsBuilder {
	b.channels = v
	return b
}

func (b *UsergroupPrefsBuilder) Groups(v ...string) *UsergroupPrefsBuilder {
	b.groups = v
	return b
}

func (b *UsergroupPrefsBuilder) Do() (*UsergroupPrefs, error) {
	var v UsergroupPrefs
	v.channels = b.channels
	v.groups = b.groups
	return &v, nil
}

func (b *UsergroupPrefs) Channels() []string {
	return b.channels
}

func (b *UsergroupPrefs) Groups() []string {
	return b.groups
}
