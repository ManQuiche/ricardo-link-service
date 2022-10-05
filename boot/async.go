package boot

func ListenEvents() {
	_, _ = natsEncConn.Subscribe(natsUserDeleted, asyncHandler.OnUserDelete)
	_, _ = natsEncConn.Subscribe(natsPartyDeleted, asyncHandler.OnPartyDelete)
}
