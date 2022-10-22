package async

type Handler interface {
	OnUserDelete(userID uint)
	OnPartyDelete(partyID uint)
}
