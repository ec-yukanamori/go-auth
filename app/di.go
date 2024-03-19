package main

func assembleTokenUsecase() tokenUsecase {
	r := newTokenRepository(rds)
	return newtokenUsecase(r)
}

func assembleTokenHandler() *tokenHandler {
	u := assembleTokenUsecase()
	return newTokenHandler(u)
}
