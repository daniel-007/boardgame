package boardgame

import (
	"errors"
	"github.com/jkomoros/boardgame/enum"
	"log"
	"testing"
)

//Place to define testing structs and helpers that are useful throughout

type testAgent struct{}

func (t *testAgent) Name() string {
	return "Test"
}

func (t *testAgent) DisplayName() string {
	return "Robby the Robot"
}

func (t *testAgent) SetUpForGame(game *Game, player PlayerIndex) (state []byte) {
	return nil
}

func (t *testAgent) ProposeMove(game *Game, player PlayerIndex, agentState []byte) (move Move, newAgentState []byte) {

	state := game.CurrentState()

	gameState, _ := concreteStates(state)

	if gameState.CurrentPlayer != player {
		return nil, nil
	}

	move = game.MoveByName("Test")

	if move == nil {
		log.Println("Couldn't find move Test")
		return nil, nil
	}

	move.(*testMove).TargetPlayerIndex = player

	return move, nil
}

//testingComponent is a very basic thing that fufills the Component interface.
type testingComponent struct {
	BaseComponentValues
	String  string
	Integer int
}

type testingComponentDynamic struct {
	BaseSubState
	BaseComponentValues
	IntVar int
	Stack  Stack
	Enum   enum.Val
}

const testGameName = "test_game"

func (t *testingComponent) Reader() PropertyReader {
	return getDefaultReader(t)
}

func (t *testingComponentDynamic) Reader() PropertyReader {
	return getDefaultReader(t)
}

func (t *testingComponentDynamic) ReadSetter() PropertyReadSetter {
	return getDefaultReadSetter(t)
}

func (t *testingComponentDynamic) ReadSetConfigurer() PropertyReadSetConfigurer {
	return getDefaultReadSetConfigurer(t)
}

func (t *testingComponentDynamic) SetState(state State) {
	t.state = state
}

func (t *testingComponentDynamic) State() State {
	return t.state
}

func (t *testingComponentDynamic) SetImmutableState(state ImmutableState) {
	t.immutableState = state
}

func (t *testingComponentDynamic) ImmutableState() ImmutableState {
	return t.immutableState
}

//Every game should do such a convenience method. state might be nil.
func concreteStates(state ImmutableState) (*testGameState, []*testPlayerState) {

	if state == nil {
		return nil, nil
	}

	players := make([]*testPlayerState, len(state.ImmutablePlayerStates()))

	for i, player := range state.ImmutablePlayerStates() {
		players[i] = player.(*testPlayerState)
	}

	game, ok := state.ImmutableGameState().(*testGameState)

	if !ok {
		return nil, nil
	}

	return game, players
}

type testGameState struct {
	state              State
	immutableState     ImmutableState
	Phase              enum.TreeVal `enum:"phase"`
	CurrentPlayer      PlayerIndex
	DrawDeck           Stack `sanitize:"len"`
	Timer              Timer
	MyIntSlice         []int
	MyStringSlice      []string
	MyBoolSlice        []bool
	MyPlayerIndexSlice []PlayerIndex
	MyEnumValue        enum.Val
	MyEnumConst        enum.ImmutableVal
	DownSizeStack      SizedStack  `sizedstack:"test, ConstantStackSize"`
	OtherStack         SizedStack  `sizedstack:"test,2"`
	MyMergedStack      MergedStack `concatenate:"DownSizeStack,OtherStack"`
	MyBoard            Board       `stack:"test" board:"3"`
}

func (t *testGameState) Reader() PropertyReader {
	return getDefaultReader(t)
}

func (t *testGameState) ReadSetter() PropertyReadSetter {
	return getDefaultReadSetter(t)
}

func (t *testGameState) ReadSetConfigurer() PropertyReadSetConfigurer {
	return getDefaultReadSetConfigurer(t)
}

func (t *testGameState) SetState(state State) {
	t.state = state
}

func (t *testGameState) State() State {
	return t.state
}

func (t *testGameState) SetImmutableState(state ImmutableState) {
	t.immutableState = state
}

func (t *testGameState) ImmutableState() ImmutableState {
	return t.immutableState
}

type testPlayerState struct {
	state          State
	immutableState ImmutableState
	//Note: PlayerIndex is stored ehre, but not a normal property or
	//serialized, because it's really just a convenience method because it's
	//implied by its position in the State.Users array.
	playerIndex       PlayerIndex
	Score             int
	MovesLeftThisTurn int
	Hand              SizedStack `sizedstack:"test,2"`
	IsFoo             bool
	EnumVal           enum.Val `enum:"color"`
}

func (t *testPlayerState) GameScore() int {
	return t.Score
}

func (t *testPlayerState) PlayerIndex() PlayerIndex {
	return t.playerIndex
}

func (t *testPlayerState) ReadSetter() PropertyReadSetter {
	return getDefaultReadSetter(t)
}

func (t *testPlayerState) Reader() PropertyReader {
	return getDefaultReader(t)
}

func (t *testPlayerState) ReadSetConfigurer() PropertyReadSetConfigurer {
	return getDefaultReadSetConfigurer(t)
}

func (t *testPlayerState) SetState(state State) {
	t.state = state
}

func (t *testPlayerState) State() State {
	return t.state
}

func (t *testPlayerState) SetImmutableState(state ImmutableState) {
	t.immutableState = state
}

func (t *testPlayerState) ImmutableState() ImmutableState {
	return t.immutableState
}

type testMoveInvalidPlayerIndex struct {
	baseFixUpMove
	//This move is a dangerous one and also a fix-up. So make it so by default
	//it doesn't apply.
	CurrentlyLegal bool
}

var testMoveInvalidPlayerIndexConfig = MoveConfig{
	Name: "Invalid PlayerIndex",
	Constructor: func() Move {
		return new(testMoveInvalidPlayerIndex)
	},
}

func (t *testMoveInvalidPlayerIndex) HelpText() string {
	return "Set one of the PlayerIndex properties to an invalid number, so we can verify that ApplyMove catches it."
}

func (t *testMoveInvalidPlayerIndex) Reader() PropertyReader {
	return getDefaultReader(t)
}

func (t *testMoveInvalidPlayerIndex) ReadSetter() PropertyReadSetter {
	return getDefaultReadSetter(t)
}

func (t *testMoveInvalidPlayerIndex) ReadSetConfigurer() PropertyReadSetConfigurer {
	return getDefaultReadSetConfigurer(t)
}

func (t *testMoveInvalidPlayerIndex) Legal(state ImmutableState, propopser PlayerIndex) error {

	if !t.CurrentlyLegal {
		return errors.New("Move not currently legal")
	}

	return nil
}

func (t *testMoveInvalidPlayerIndex) Apply(state State) error {
	game, players := concreteStates(state)

	game.CurrentPlayer = PlayerIndex(len(players))

	return nil
}

type testMoveMakeIllegalPhase struct {
	baseMove
}

var testMoveMakeIllegalPhaseConfig = MoveConfig{
	Name: "Make Illegal Phase",
	Constructor: func() Move {
		return new(testMoveMakeIllegalPhase)
	},
}

func (t *testMoveMakeIllegalPhase) HelpText() string {
	return "Sets to illegal phase which should fail to apply"
}

func (t *testMoveMakeIllegalPhase) Reader() PropertyReader {
	return getDefaultReader(t)
}

func (t *testMoveMakeIllegalPhase) ReadSetter() PropertyReadSetter {
	return getDefaultReadSetter(t)
}

func (t *testMoveMakeIllegalPhase) ReadSetConfigurer() PropertyReadSetConfigurer {
	return getDefaultReadSetConfigurer(t)
}

func (t *testMoveMakeIllegalPhase) Legal(state ImmutableState, player PlayerIndex) error {
	return nil
}

func (t *testMoveMakeIllegalPhase) Apply(state State) error {

	g := state.GameState().(*testGameState)

	//phaseNormal is illegal because it is a non-leaf enum val move
	g.Phase.SetValue(phaseNormal)

	return nil

}

type testMoveIncrementCardInHand struct {
	baseMove
	TargetPlayerIndex PlayerIndex
}

var testMoveIncrementCardInHandConfig = MoveConfig{
	Name: "Increment IntValue of Card in Hand",
	Constructor: func() Move {
		return new(testMoveIncrementCardInHand)
	},
}

func (t *testMoveIncrementCardInHand) HelpText() string {
	return "Increments the IntValue of the card in the hand"
}

func (t *testMoveIncrementCardInHand) DefaultsForState(state ImmutableState) {
	t.TargetPlayerIndex = state.CurrentPlayer().PlayerIndex()
}

func (t *testMoveIncrementCardInHand) Reader() PropertyReader {
	return getDefaultReader(t)
}

func (t *testMoveIncrementCardInHand) ReadSetter() PropertyReadSetter {
	return getDefaultReadSetter(t)
}

func (t *testMoveIncrementCardInHand) ReadSetConfigurer() PropertyReadSetConfigurer {
	return getDefaultReadSetConfigurer(t)
}

func (t *testMoveIncrementCardInHand) Legal(state ImmutableState, proposer PlayerIndex) error {
	game, players := concreteStates(state)

	player := players[game.CurrentPlayer]

	if !t.TargetPlayerIndex.Equivalent(proposer) {
		return errors.New("The proposer is not the target Player index")
	}

	if !t.TargetPlayerIndex.Equivalent(game.CurrentPlayer) {
		return errors.New("The target player index is not the current player")
	}

	if player.Hand.NumComponents() == 0 {
		return errors.New("The current player does not have any components in their hand")
	}

	if game.DrawDeck.NumComponents() < 1 {
		return errors.New("There aren't enough cards left over in the draw deck")
	}

	return nil
}

func (t *testMoveIncrementCardInHand) Apply(state State) error {
	game, players := concreteStates(state)

	player := players[game.CurrentPlayer]

	for _, component := range player.Hand.Components() {
		if component == nil {
			continue
		}

		values := component.DynamicValues()

		if values == nil {
			return errors.New("DynamicValues was nil")
		}

		easyValues := values.(*testingComponentDynamic)

		easyValues.IntVar += 3
		game.DrawDeck.ComponentAt(game.DrawDeck.Len() - 1).MoveToFirstSlot(easyValues.Stack)

		return nil

	}

	return errors.New("Didn't find a component in hand")
}

type testMoveDrawCard struct {
	baseMove
	TargetPlayerIndex PlayerIndex
}

var testMoveDrawCardConfig = MoveConfig{
	Name: "Draw Card",
	Constructor: func() Move {
		return new(testMoveDrawCard)
	},
}

func (t *testMoveDrawCard) HelpText() string {
	return "Draws one card from draw deck into player's hand"
}

func (t *testMoveDrawCard) DefaultsForState(state ImmutableState) {
	t.TargetPlayerIndex = state.CurrentPlayer().PlayerIndex()
}

func (t *testMoveDrawCard) Reader() PropertyReader {
	return getDefaultReader(t)
}

func (t *testMoveDrawCard) ReadSetter() PropertyReadSetter {
	return getDefaultReadSetter(t)
}

func (t *testMoveDrawCard) ReadSetConfigurer() PropertyReadSetConfigurer {
	return getDefaultReadSetConfigurer(t)
}

func (t *testMoveDrawCard) Legal(state ImmutableState, proposer PlayerIndex) error {
	game, players := concreteStates(state)

	player := players[game.CurrentPlayer]

	if !t.TargetPlayerIndex.Equivalent(proposer) {
		return errors.New("The proposer is not equivalent to targetplayerindex")
	}

	if !t.TargetPlayerIndex.Equivalent(game.CurrentPlayer) {
		return errors.New("The target player is not the current player")
	}

	if player.Hand.SlotsRemaining() == 0 {
		return errors.New("The current player does not have enough slots in their hand")
	}

	if game.DrawDeck.NumComponents() == 0 {
		return errors.New("there are no cards to draw")
	}

	return nil
}

func (t *testMoveDrawCard) Apply(state State) error {
	game, players := concreteStates(state)

	player := players[game.CurrentPlayer]

	if err := game.DrawDeck.ComponentAt(0).MoveToFirstSlot(player.Hand); err != nil {
		return errors.New("couldn't move component from draw deck to hand: " + err.Error())
	}

	return nil
}

type testMoveAdvanceCurentPlayer struct {
	baseFixUpMove
}

var testMoveAdvanceCurrentPlayerConfig = MoveConfig{
	Name: "Advance Current Player",
	Constructor: func() Move {
		return new(testMoveAdvanceCurentPlayer)
	},
}

func (t *testMoveAdvanceCurentPlayer) HelpText() string {
	return "Advances to the next player when the current player has no more legal moves they can make this turn."
}

func (t *testMoveAdvanceCurentPlayer) Reader() PropertyReader {
	return getDefaultReader(t)
}

func (t *testMoveAdvanceCurentPlayer) ReadSetter() PropertyReadSetter {
	return getDefaultReadSetter(t)
}

func (t *testMoveAdvanceCurentPlayer) ReadSetConfigurer() PropertyReadSetConfigurer {
	return getDefaultReadSetConfigurer(t)
}

func (t *testMoveAdvanceCurentPlayer) Legal(state ImmutableState, proposer PlayerIndex) error {

	game, players := concreteStates(state)

	player := players[game.CurrentPlayer]

	if player.MovesLeftThisTurn > 0 {
		return errors.New("The current player still has moves left this turn.")
	}

	return nil
}

func (t *testMoveAdvanceCurentPlayer) Apply(state State) error {

	game, players := concreteStates(state)

	//Make sure we're leaving it at 0
	players[game.CurrentPlayer].MovesLeftThisTurn = 0

	game.CurrentPlayer++

	if !game.CurrentPlayer.Valid(state) {
		//Must have looped back around to 0
		game.CurrentPlayer = 0
	}

	players[game.CurrentPlayer].MovesLeftThisTurn = 1

	return nil
}

type testMove struct {
	baseMove
	AString           string
	ScoreIncrement    int
	TargetPlayerIndex PlayerIndex
	ABool             bool
}

var testMoveConfig = MoveConfig{
	Name: "Test",
	Constructor: func() Move {
		return new(testMove)
	},
}

func (t *testMove) HelpText() string {
	return "Advances the score of the current player by the specified amount."
}

func (t *testMove) DefaultsForState(state ImmutableState) {
	game, _ := concreteStates(state)

	if game == nil {
		blob, _ := DefaultMarshalJSON(state)
		log.Println(string(blob))
		return
	}

	t.TargetPlayerIndex = game.CurrentPlayer
	t.ScoreIncrement = 3
}

func (t *testMove) Reader() PropertyReader {
	return getDefaultReader(t)
}

func (t *testMove) ReadSetter() PropertyReadSetter {
	return getDefaultReadSetter(t)
}

func (t *testMove) ReadSetConfigurer() PropertyReadSetConfigurer {
	return getDefaultReadSetConfigurer(t)
}

func (t *testMove) Legal(state ImmutableState, proposer PlayerIndex) error {

	game, _ := concreteStates(state)

	if !t.TargetPlayerIndex.Equivalent(proposer) {
		return errors.New("The target player index is not equivalent to the proposer")
	}

	if !t.TargetPlayerIndex.Equivalent(game.CurrentPlayer) {
		return errors.New("The target player is not hte current player")
	}

	if game.CurrentPlayer != t.TargetPlayerIndex {
		return errors.New("The current player is not the same as the target player")
	}

	return nil

}

func (t *testMove) Apply(state State) error {

	game, players := concreteStates(state)

	players[game.CurrentPlayer].Score += t.ScoreIncrement

	players[game.CurrentPlayer].MovesLeftThisTurn -= 1

	return nil
}

type testAlwaysLegalMove struct {
	baseFixUpMove
}

var testAlwaysLegalMoveConfig = MoveConfig{
	Name: "Test Always Legal Move",
	Constructor: func() Move {
		return new(testAlwaysLegalMove)
	},
}

func (t *testAlwaysLegalMove) HelpText() string {
	return "A move that is always legal"
}

func (t *testAlwaysLegalMove) Reader() PropertyReader {
	return getDefaultReader(t)
}

func (t *testAlwaysLegalMove) ReadSetter() PropertyReadSetter {
	return getDefaultReadSetter(t)
}

func (t *testAlwaysLegalMove) ReadSetConfigurer() PropertyReadSetConfigurer {
	return getDefaultReadSetConfigurer(t)
}

func (t *testAlwaysLegalMove) Legal(state ImmutableState, proposer PlayerIndex) error {

	return nil

}

func (t *testAlwaysLegalMove) Apply(state State) error {

	//This move doesn't do anything

	return nil
}

type illegalMove struct {
	baseMove
	Stack Stack
}

func (i *illegalMove) Reader() PropertyReader {
	return getDefaultReader(i)
}

func (i *illegalMove) ReadSetter() PropertyReadSetter {
	return getDefaultReadSetter(i)
}

func (i *illegalMove) ReadSetConfigurer() PropertyReadSetConfigurer {
	return getDefaultReadSetConfigurer(i)
}

func (i *illegalMove) Legal(state ImmutableState, proposer PlayerIndex) error {
	return nil
}

func (i *illegalMove) Apply(state State) error {
	return nil
}

var testIllegalMoveConfig = MoveConfig{
	Name: "Illegal Move",
	Constructor: func() Move {
		return new(illegalMove)
	},
}

func (i *illegalMove) HelpText() string {
	return "Move that is illegal because it has an illegal property type on it"
}

//testingComponentValues is designed to be run on a stack.ComponentValues() of
//a stack of testingComponents, in order to convert them all to the specified
//underlying struct.
func testingComponentValues(in []Reader) []*testingComponent {
	result := make([]*testingComponent, len(in))
	for i := 0; i < len(in); i++ {
		if in[i] == nil {
			result[i] = nil
			continue
		}
		result[i] = in[i].(*testingComponent)
	}
	return result
}

func makeTestGameIdsStable(game *Game) {
	//having the same fixed salt helps make the test predictable regarding
	//component ids.
	game.secretSalt = "FAKESALTFORTESTING"
	game.id = "FAKEIDFORTESTING"
}

//testGame returns a Game that has not yet had SetUp() called.
func testGame(t *testing.T) *Game {

	manager := newTestGameManger(t)

	game := manager.NewGame()

	return game
}
