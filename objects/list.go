package objects

func (l *ActionList) Append(a *Action) *ActionList {
	*l = append(*l, a)
	return l
}

func (l *AttachmentList) Append(a *Attachment) *AttachmentList {
	*l = append(*l, a)
	return l
}

func (l *AttachmentFieldList) Append(a *AttachmentField) *AttachmentFieldList {
	*l = append(*l, a)
	return l
}
