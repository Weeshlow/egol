# EGOL

## Dependencies
---
- [Go](https://golang.org/) programming language binaries with the `GOPATH` environment variable specified.
- [Go](https://golang.org/) version 1.6, or version 1.5 with the `GO15VENDOREXPERIMENT` environment variable set to `1`.
- [NodeJS](http://nodejs.org/) JavaScript runtime.
- [gulp](http://http://gulpjs.com/) build toolkit (npm install gulp -g).

## Development
---
Clone the repository:

```bash
mkdir $GOPATH/src/github.com/unchartedsoftware
cd $GOPATH/src/github.com/unchartedsoftware
git clone git@github.com:unchartedsoftware/egol.git
```

Install dependencies

```bash
cd egol
make deps
```

Run the server:

```bash
gulp
```

## Priorities
---
### High

#### Frontend
- basic organism rendering
- rendering / interpolation of scene state

#### Backend
- low frequency server simulation loop
- expose hook for AI to plug in
- manage websocket connections and broadcasting state to clients

#### AI
- impl of most basic constraints
- impl of trivial behavior with FSM

### Medium

#### Frontend
- procedural organism generation based on attributes
- unique rendering / representation based on states

#### Backend
- minimize amount of information broadcast to clients
- tune networking code to support more organisms and higher broadcast frequency

#### AI
- impl of more advanced constraints
- impl of more advanced behavior FSM
- balancing constraints / behavior

### Low

#### Frontend
- animate state changes
- implement organism abilities

#### AI
- implement organism specific abilities

## Ideas
---
### Simulation Constraints

- hunger will slowly increase from 0 to 1. Rate at which hunger increases will be determined by weights applied to certain attributes (size?).
- energy, from 0 to 1, will be consumed at fixed rates based on attributes and the current state of the organism.
- if hunger hits 1, and there is remaining energy, add hunger addition * K to energy consumption.
- if energy hits 0, and there is remaining hunder, add energy consumption * K to hunger accumulation.
- if energy reaches 0, and hunger is 1, organism dies.

### Organism attributes:

- family: Type of the organism: {A, B, C, D, or E}. Each type has the following:
	- type {i} will prefer to eat type {i+1} (1.5x multiplier for hunger / energy gain)
	- type {i} will prefer not to eat type {i+2} (0.5 multiplier for hunger / energy gain)
	- possible attribute modifiers?
	- possible abilities unique to families?
- offense: determines ability to attack / feed on another organism.
	- increases size.
	- decreases speed.
- defense: determines ability to defend an attack.
	- increases size.
	- decreases speed.
- agility: ability to dodge an attack
	- increase speed.
	- decrease size.
- intelligence: ability for AI to make better decisions based on perceptions?
- speed: based on agility and size?
- size: based on offense, defense, and speed?
- range: determines attack range
- perception: determines range and sensitivity of perception.
	- ability to detect family, attributes, hunger, etc of other organisms?
- reproductivity: rate at which organism reproduces.
	- number of children? speed? resources consumed?

### Organism States:

- seeking: searching for organisms to consume
- attacking: actively attacking another organism
- defending: actively defending an attacking organism
- consuming: successfully attacked, currently consuming an organism
- pinned: unsuccessfully defended, currently being consumed
- fleeing: fleeing a threat
- rest: no action, recovers energy, reduces hunger accumulation
- reproduce: produce offspring with random mutations to attributes
	- see if net

### Attacking / Consuming:

- target must be within attackers range
- attacker offense + agility vs defender defense + agility to determine if attack is successful
	- if attack is successful, attacker enters consuming state, defender enters a pinned state
	- if attack is unsuccessful, defender moves out of range and enters fleeing state.
- attacker in a consuming state:
	- continues consuming, amount of hunger and energy consumed calculated by offense of attacker vs defense of defender
	- stops consuming and enters another state (being attacked, higher priority target).
- defender in pinned state:
 	- attempts to break pinned state, based on offense + agility of defender vs defense + agility of attacker. If successful moves out of range, and enters fleeing state.

### Reproduction

- produce N children
- children start with some K<1 attribute multiplier fit to some function of t.
	- based on family?
- parents protect children (attribute?)
- children stay near parents(attribute?)

### Perception

- detection range
- detection sensitivity, different types based on value:
	- direction vs position
	- discern organism size / family type, maybe other attributes?
