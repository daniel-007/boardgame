/*

moves is a convenience package that implements composable Moves to make it
easy to implement common logic. The Base move type is a very simple move that
implements the basic stubs necessary for your straightforward moves to have
minimal boilerplate. Although it's technically optional, a lot of the magic
features throughout the framework depend on some if its base logic, so it's
recommended to always embed it anonymously in your move struct (or embed a
struct that embeds it).

You interact with and configure various move types by implementing interfaces.
Those interfaes are defined in the interfaces subpackage, to make this
package's design more clear.

There are many move types defined. Some are designed to be used directly with
minimal modification; others are powerful move types that are designed to be
sub-classed.

Automatic MoveConfig Generation

Technically every move needs to be installed on your GameManager by creating a
boardgame.MoveConfig. In practice writing those from scratch is verbose
and error-prone. auto is a sub-package that generates those configurations for
you automatically, and is strongly recommended to use for all of your moves.
All moves in this package are designed to work well with auto.Config.

This package defines a number of interfaces.CustomConfigurationOptions
options, which are designed to be passed as arguments to auto.Config. They all
begin with "With" and are used to set various configuration options that
auto.Config will use.

Read more--including in-depth worked examples and discussions of how to use
auto.Config idiomatically--in the package doc for auto.

Configure Move Helpers

Your Game Delegate's ConfigureMoves() []boardgame.MoveConfig is where the
action happens for installing moves. In practice you can do whatever you want
in there as long as you return a list of MoveConfigs. In practice you often
use auto.Config (see section above). If you have a very simple game type you
might not need to do anythign special.

If, however, your game type is complicated enough to need the notion of
phases, then you'll probably want to use some of the convenience methods for
installing moves: Combine, Add, AddForPhase, and AddOrderedForPhase. These
methods make sure that enough information is stored for the Legal() methods of
those moves to know when the move is legal. Technically they're just
convenience wrappers (each describes the straightforward things it's doing),
but in practice they're the best way to do it. See the tutorial in the main
package for more.

Base Move

Implementing a Move requires a lot of stub methods to implement the
boardgame.Move interface, but also a lot of logic, especially to support
Phases correctly. moves.Base is a move that all moves should embed somewhere
in their hierarchy. It is very important to always call your superclasse's
Legal(), because moves.Base.Legal contains important logic to implement phases
and ordered moves within phases.

Base also includes a number of methods needed for moves to work well with
auto.Config.

FixUp Move

FixUp moves are simple embedding of move.Base, but they default to having
IsFixUp generated by auto.Config be true instead of false. This is useful so
you don't forget to pass WithFixUp(true) yourself in auto.Config.

FixUpMulti is the same as FixUp, but also has a AllowMultipleInProgression
that returns true, meaning that the ordered move logic within phases will
allow multiple of this move type to apply in a row.

Default Component Move

DefaultComponent is a move type that, in DefaultsForState, searches through
all of the components in the stack provided with WithSourceStack, and testing
the Legal() method of each component. It sets the first one that returns nil
to m.ComponentIndex. Its Legal() returns whether there is a valid component
specified, and what its Legal returns. You provide your own Apply().

It's useful for fixup moves that need to apply actions to components in a
given stack when certain conditions are met--for example, crowning a token
that makes it to the opposite end of a board in checkers.

The componentValues.Legal() takes a legalType. This is the way you can use
multiple DefaultComponent moves for the same type of component. If you only
have one then you can skip passing WithLegalType, and just default to 0. If
you do have multiple legalTypes, the idiomatic way is to have those be members
of an Enum for that purpose.

Current Player Move

These moves are for moves that are only legal to be made by the current
player. Their Legal() will verify that it is the proposer's turn.

Move Deal and Collect Component Moves

Generally when moving components from one place to another it makes sense to
move one component at a time, so that each component is animated separately.
However, this is a pain to implement, because it requires implementing a move
that knows how many times to apply itself in a row, which is fincky and error-
prone.

There is a collection of 9 moves that all do basically the same thing for
moving components, one at a time, from stack to stack. Move-type moves move
components between two specific stacks, often both on your GameState. Deal and
Collect type moves move components between a stack in GameState and a stack in
each Player's PlayerState. Deal-type moves move components from the game stack
to the player stack, and Collect-type moves move components from each player
to the GameState.

All of these moves define a way to define the source and destination stack.
For Move-type moves, you define SourceStack() and DestinationStack(). For Deal
and Collect-type moves, you implement GameStack() and PlayerStack().

All moves in this collection implement TargetCount() int, and all of them
default to 1. Override this if you want a different number of components
checked for in the end condition.

In practice you'll often use WithTargetCount, WithGameStack, and friends as
configuration to auto.Config instead of overriding those yourself. In fact, in
many cases configuartion options are powerful enough to allow you to use these
moves types on their own directly in your game. See the documentation in
auto.Config for more examples.

Each of Move, Deal, and Collect have three variants based on the end
condition. Note that Move-type moves have only two stacks, but Deal and
Collect type moves operate on n pairs of stacks, where n is the number of
players in the game. In general for Deal and Collect type moves, the condition
is met when all pairs of stacks meet the end condition.

{Move,Deal,Collect}CountComponents simply apply that many moves without regard
to the number of components in the source or destination stacks. Move names
that end in CountReached operate until the destination stacks all have
TargetCount or more items. Move names that end in CountLeft operate until the
source stacks all have TargetCount or fewer items in them.

ApplyUntil ApplyUntilCount and ApplyCountTimes

These moves are what the MoveComponents moves are based off of and are
designed to be subclassed. They apply the move in question until some
condition is reached.

RoundRobin and RoundRobinNumRounds

Round Robin moves are like ApplyUntilCount and friends, except they go around
and operate on each player in succession. RoundRobinNumRounds goes around each
player until NumRounds() cycles have completed. The base RoundRobin goes
around until the PlayerCondition has been met for each player. These are the
most complicated moves in the set; if you subclass one directly you're most
likely to subclass RoundRobinNumRounds.

FinishTurn

FinishTurn is a move that is designed to be used as a fix-up move during
normal phases of play in your game. It checks whether the current player's
turn is done (based on criteria you specify) and if so advances to the next
player, resetting state as appropriate. In practice you often can use this
move directly in your game without even passing any WithOPTION configuration
to auto.Config.

StartPhase

The StartPhase move is designed to set your game's phase to the next phase.
It's generally used as the last move in an ordered phase, for example, the
last move in your game's SetUp phase. This move can also generally be used
directly in your game, by using the WithPhaseToStart configuration option in
auto.Config.

ShuffleStack

Shuffle stack is a simple move that just shuffles the stack denoted by
SourceStack.

*/
package moves
