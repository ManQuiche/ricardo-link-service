package boot

func ListenEvents() {
	_, _ = natsEncConn.Subscribe(natsPartyDeleted, asyncHandler.OnPartyDelete)
}
