# Slot Game Engine

This repository contains the code for a basic slot game engine. The Makefile provides a set of commands to help in building, running, testing, and generating code for this project.

## Prerequisites

- Ensure you have [Go](https://golang.org/) installed.
- Clone this repository to your local machine.

## Makefile Commands

### Build

To compile the source code and create a binary named `slotgame`, use the following command:

```bash
make build
```

After successful compilation, you'll find the binary `slotgame` in your directory.

### Run

To compile and run the slot game, use:

```bash
make run
```

### Clean

To remove any build artifacts and the generated binary:

```bash
make clean
```

### Test

To run tests across the project with race detection and coverage:

```bash
make test
```

### Generate

To generate mocks:

```bash
make generate
```

### Contributing

Feel free to open issues and pull requests. Every form of contribution is much appreciated!

# Flexible Game Engine Implementation

The application is designed with flexibility in mind. Through the utilization of the `GameEngine` interface, alternative game engine implementations can be easily integrated.

**How It Works:**

- The interface `GameEngine` declares the method `Play(bet float64) ([][]string, float64, error)`. This interface-based design is beneficial as any struct that satisfies the methods declared can be used wherever that interface is expected.

  For instance, an `AdvancedGameEngine` could be created with a different play logic:

  ```go
  type AdvancedGameEngine struct {
      // ... other fields
  }

  func (age AdvancedGameEngine) Play(bet float64) ([][]string, float64, error) {
      // ... implement advanced play logic
      return
  }
  ```

  With this, the `AdvancedGameEngine` can be utilized in places where the `GameEngine` interface is expected.

- The architecture is not just limited to game engine flexibility. The `PayTable` interface is another avenue for customization. It enables the game engine to cater to various winning conditions and pay table designs.

  Different pay table implementations can be integrated. A basic pay table might use fixed patterns and multipliers, but alternative designs, like a `ProgressivePayTable`, could be introduced, potentially leveraging evolving multipliers or more intricate conditions.

- Adjustments in the pay table data, even within a singular implementation, can drastically change the winning outcomes. This highlights the adaptability of the system, catering to different game dynamics without the need for overarching structural changes.

- The reel design, governed by the `Reel` interface, provides another layer of extensibility. It's possible to integrate reels with different symbol arrangements, varying probabilities, or even unique mechanisms like cascading or sticky symbols. The modular structure supports a diverse range of game experiences.

In essence, the modular and interface-based architecture of the application ensures adaptability to various game mechanics, winning conditions, and reel designs, thus allowing for a rich and diverse gaming experience.

# Potential Bug Scenarios and Concerns:
### Pseudo-Randomness of rand:

If not seeded, the Go's rand package will produce the same sequence of numbers every time the application is started. This can be exploitative if someone detects the pattern.