package objects

func (l *ActionList) Append(v *Action) *ActionList {
	*l = append(*l, v)
	return l
}

func (l *AttachmentList) Append(v *Attachment) *AttachmentList {
	*l = append(*l, v)
	return l
}

func (l *AttachmentFieldList) Append(v *AttachmentField) *AttachmentFieldList {
	*l = append(*l, v)
	return l
}

func (l *ChannelList) Append(v *Channel) *ChannelList {
	*l = append(*l, v)
	return l
}

func (l *ConfirmationList) Append(v *Confirmation) *ConfirmationList {
	*l = append(*l, v)
	return l
}

func (l *ConversationList) Append(v *Conversation) *ConversationList {
	*l = append(*l, v)
	return l
}

func (l *MessageList) Append(v *Message) *MessageList {
	*l = append(*l, v)
	return l
}

func (l *OptionList) Append(v *Option) *OptionList {
	*l = append(*l, v)
	return l
}

func (l *OptionGroupList) Append(v *OptionGroup) *OptionGroupList {
	*l = append(*l, v)
	return l
}

func (l *ReactionList) Append(v *Reaction) *ReactionList {
	*l = append(*l, v)
	return l
}

func (l *UserProfileList) Append(v *UserProfile) *UserProfileList {
	*l = append(*l, v)
	return l
}
