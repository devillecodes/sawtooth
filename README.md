# A simple [Battlesnake](http://play.battlesnake.com) written in Go.

This is a basic implementation of the [Battlesnake API](https://docs.battlesnake.com/snake-api). This project was originally cloned from [starter-snake-go](https://github.com/BattlesnakeOfficial/starter-snake-go) (by [BattlesnakeOfficial](https://github.com/BattlesnakeOfficial)).

This project is being built as part of a [100 Days Of Code](https://github.com/devillexio/100-days-of-code/blob/master/log.md) challenge, and also to learn more about Go.

### Technologies

This Battlesnake uses [Go 1.14](https://golang.org/) and [Heroku](https://heroku.com).

### Strategy

"_I am a humble snake that mostly keeps to myself... and walls._" - Sawtooth

Sawtooth makes its way to the nearest wall and then hugs the wall in a clockwise direction. It favours left or right over top or bottom when the distances are equal.

### Why Sawtooth?

The name _Sawtooth_ is a homage to one of my favorite games of all time, Diablo III.

> _Sawtooth, Subterranean Monstrosity, is a Unique Rockworm (of the Demonic Serpent type) found in the Arreat Crater Level 2 in Act III of Diablo III._ - [Diablo Wiki](https://diablo.fandom.com/wiki/Sawtooth)