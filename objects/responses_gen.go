package objects

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

func BuildChannelsHistoryResponse() *ChannelsHistoryResponseBuilder {
	var b ChannelsHistoryResponseBuilder
	return &b
}

func (b *ChannelsHistoryResponseBuilder) HasMore(v bool) *ChannelsHistoryResponseBuilder {
	b.hasMore = v
	return b
}

func (b *ChannelsHistoryResponseBuilder) Latest(v string) *ChannelsHistoryResponseBuilder {
	b.latest = v
	return b
}

func (b *ChannelsHistoryResponseBuilder) Messages(v ...*Message) *ChannelsHistoryResponseBuilder {
	b.messages = v
	return b
}

func (b *ChannelsHistoryResponseBuilder) Do() (*ChannelsHistoryResponse, error) {
	var v ChannelsHistoryResponse
	v.hasMore = b.hasMore
	v.latest = b.latest
	v.messages = b.messages
	return &v, nil
}

func (b *ChannelsHistoryResponse) HasMore() bool {
	return b.hasMore
}

func (b *ChannelsHistoryResponse) Latest() string {
	return b.latest
}

func (b *ChannelsHistoryResponse) Messages() MessageList {
	return b.messages
}

func BuildChatResponse() *ChatResponseBuilder {
	var b ChatResponseBuilder
	return &b
}

func (b *ChatResponseBuilder) Channel(v string) *ChatResponseBuilder {
	b.channel = v
	return b
}

func (b *ChatResponseBuilder) Timestamp(v string) *ChatResponseBuilder {
	b.ts = v
	return b
}

func (b *ChatResponseBuilder) Message(v interface{}) *ChatResponseBuilder {
	b.message = v
	return b
}

func (b *ChatResponseBuilder) Do() (*ChatResponse, error) {
	var v ChatResponse
	v.channel = b.channel
	v.ts = b.ts
	v.message = b.message
	return &v, nil
}

func (b *ChatResponse) Channel() string {
	return b.channel
}

func (b *ChatResponse) Timestamp() string {
	return b.ts
}

func (b *ChatResponse) Message() interface{} {
	return b.message
}

func BuildEphemeralResponse() *EphemeralResponseBuilder {
	var b EphemeralResponseBuilder
	return &b
}

func (b *EphemeralResponseBuilder) MessageTs(v string) *EphemeralResponseBuilder {
	b.messageTs = v
	return b
}

func (b *EphemeralResponseBuilder) Do() (*EphemeralResponse, error) {
	var v EphemeralResponse
	v.messageTs = b.messageTs
	return &v, nil
}

func (b *EphemeralResponse) MessageTs() string {
	return b.messageTs
}

func BuildGenericResponse() *GenericResponseBuilder {
	var b GenericResponseBuilder
	return &b
}

func (b *GenericResponseBuilder) OK(v bool) *GenericResponseBuilder {
	b.ok = v
	return b
}

func (b *GenericResponseBuilder) ReplyTo(v int) *GenericResponseBuilder {
	b.replyTo = v
	return b
}

func (b *GenericResponseBuilder) Error(v *ErrorResponse) *GenericResponseBuilder {
	b.error = v
	return b
}

func (b *GenericResponseBuilder) Timestamp(v string) *GenericResponseBuilder {
	b.ts = v
	return b
}

func (b *GenericResponseBuilder) Do() (*GenericResponse, error) {
	var v GenericResponse
	v.ok = b.ok
	v.replyTo = b.replyTo
	v.error = b.error
	v.ts = b.ts
	return &v, nil
}

func (b *GenericResponse) OK() bool {
	return b.ok
}

func (b *GenericResponse) ReplyTo() int {
	return b.replyTo
}

func (b *GenericResponse) Error() *ErrorResponse {
	return b.error
}

func (b *GenericResponse) Timestamp() string {
	return b.ts
}

func BuildOAuthAccessResponse() *OAuthAccessResponseBuilder {
	var b OAuthAccessResponseBuilder
	return &b
}

func (b *OAuthAccessResponseBuilder) AccessToken(v string) *OAuthAccessResponseBuilder {
	b.accessToken = v
	return b
}

func (b *OAuthAccessResponseBuilder) Scope(v string) *OAuthAccessResponseBuilder {
	b.scope = v
	return b
}

func (b *OAuthAccessResponseBuilder) Do() (*OAuthAccessResponse, error) {
	var v OAuthAccessResponse
	v.accessToken = b.accessToken
	v.scope = b.scope
	return &v, nil
}

func (b *OAuthAccessResponse) AccessToken() string {
	return b.accessToken
}

func (b *OAuthAccessResponse) Scope() string {
	return b.scope
}

func BuildPermalinkResponse() *PermalinkResponseBuilder {
	var b PermalinkResponseBuilder
	return &b
}

func (b *PermalinkResponseBuilder) Channel(v string) *PermalinkResponseBuilder {
	b.channel = v
	return b
}

func (b *PermalinkResponseBuilder) Permalink(v string) *PermalinkResponseBuilder {
	b.permalink = v
	return b
}

func (b *PermalinkResponseBuilder) Do() (*PermalinkResponse, error) {
	var v PermalinkResponse
	v.channel = b.channel
	v.permalink = b.permalink
	return &v, nil
}

func (b *PermalinkResponse) Channel() string {
	return b.channel
}

func (b *PermalinkResponse) Permalink() string {
	return b.permalink
}

func BuildReactionsGetResponse() *ReactionsGetResponseBuilder {
	var b ReactionsGetResponseBuilder
	return &b
}

func (b *ReactionsGetResponseBuilder) Channel(v string) *ReactionsGetResponseBuilder {
	b.channel = v
	return b
}

func (b *ReactionsGetResponseBuilder) Comment(v string) *ReactionsGetResponseBuilder {
	b.comment = v
	return b
}

func (b *ReactionsGetResponseBuilder) File(v *File) *ReactionsGetResponseBuilder {
	b.file = v
	return b
}

func (b *ReactionsGetResponseBuilder) Message(v *Message) *ReactionsGetResponseBuilder {
	b.message = v
	return b
}

func (b *ReactionsGetResponseBuilder) Type(v string) *ReactionsGetResponseBuilder {
	b.typ = v
	return b
}

func (b *ReactionsGetResponseBuilder) Do() (*ReactionsGetResponse, error) {
	var v ReactionsGetResponse
	v.channel = b.channel
	v.comment = b.comment
	v.file = b.file
	v.message = b.message
	v.typ = b.typ
	return &v, nil
}

func (b *ReactionsGetResponse) Channel() string {
	return b.channel
}

func (b *ReactionsGetResponse) Comment() string {
	return b.comment
}

func (b *ReactionsGetResponse) File() *File {
	return b.file
}

func (b *ReactionsGetResponse) Message() *Message {
	return b.message
}

func (b *ReactionsGetResponse) Type() string {
	return b.typ
}

func BuildReactionsListResponse() *ReactionsListResponseBuilder {
	var b ReactionsListResponseBuilder
	return &b
}

func (b *ReactionsListResponseBuilder) Items(v ...*ReactionsGetResponse) *ReactionsListResponseBuilder {
	b.items = v
	return b
}

func (b *ReactionsListResponseBuilder) Paging(v *Paging) *ReactionsListResponseBuilder {
	b.paging = v
	return b
}

func (b *ReactionsListResponseBuilder) Do() (*ReactionsListResponse, error) {
	var v ReactionsListResponse
	v.items = b.items
	v.paging = b.paging
	return &v, nil
}

func (b *ReactionsListResponse) Items() ReactionsGetResponseList {
	return b.items
}

func (b *ReactionsListResponse) Paging() *Paging {
	return b.paging
}
