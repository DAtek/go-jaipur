package core

const NotEnoughCardsToSellError = JaipurError("Player doesn't have enough cards to sell")
const NotEnoughCardsOnTableError = JaipurError("Not enough cards on the table")
const RoundEndedError = JaipurError("Round ended")
const GameEndedError = JaipurError("Round ended")
