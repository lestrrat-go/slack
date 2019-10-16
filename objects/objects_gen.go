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

func BuildAuthTestResponse() *AuthTestResponseBuilder {
	var b AuthTestResponseBuilder
	return &b
}

func (b *AuthTestResponseBuilder) Url(v string) *AuthTestResponseBuilder {
	b.url = v
	return b
}

func (b *AuthTestResponseBuilder) Team(v string) *AuthTestResponseBuilder {
	b.team = v
	return b
}

func (b *AuthTestResponseBuilder) User(v string) *AuthTestResponseBuilder {
	b.user = v
	return b
}

func (b *AuthTestResponseBuilder) TeamId(v string) *AuthTestResponseBuilder {
	b.teamId = v
	return b
}

func (b *AuthTestResponseBuilder) UserId(v string) *AuthTestResponseBuilder {
	b.userId = v
	return b
}

func (b *AuthTestResponseBuilder) Do() (*AuthTestResponse, error) {
	var v AuthTestResponse
	v.url = b.url
	v.team = b.team
	v.user = b.user
	v.teamId = b.teamId
	v.userId = b.userId
	return &v, nil
}

func (b *AuthTestResponse) Url() string {
	return b.url
}

func (b *AuthTestResponse) Team() string {
	return b.team
}

func (b *AuthTestResponse) User() string {
	return b.user
}

func (b *AuthTestResponse) TeamId() string {
	return b.teamId
}

func (b *AuthTestResponse) UserId() string {
	return b.userId
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
