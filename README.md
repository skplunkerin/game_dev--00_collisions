# Objectives:

This is a simple game dev project designed for learning beginner concepts using
the [Ebiten](https://github.com/hajimehoshi/ebiten) 2D Golang game library. Much
of the code will be inspired by or copied from other examples, and then modified
to suit the below beginner objectives:

- [x] get Ebiten working locally (the examples can run/work)

  - These are more complicated than thought for a 1st Ebiten dive:

    _(found out about tilemaps (see [`unused_code/tilemaps/README.md`](/unused_code/tilemaps/README.md)_
    _for my exploration and discovering of this); I'm not sure yet on how one_
    _should create tilemap JSON objects, there's likely 3rd-party tools for_
    _this, which I want to avoid using while focusing on beginner learning.)_

    - ~~[ ] begin codebase for a simple level builder:~~
    - ~~[ ] left-click to add a wall~~
    - ~~[ ] right-click to remove a wall~~
    - ~~[ ] extras:~~
      - ~~[ ] shift + click to add a dot/square (or some sort of player avatar)~~
      - ~~[ ] spacebar to start the game using the level~~
      - ~~[ ] arrow keys to navigate dot around~~

  - These are more beginner friendly:

    - [x] render a character (tile/image) on top of the tiles example map
    - [x] move the character around with arrow keys
      - [x] bonus: include `w`,`a`,`s`,`d` keys
    - [ ] add collision detection, prevent running through:
      - [x] map edges
      - [ ] house walls
      - [ ] front flowers

  - Next project will explore:
    - side-scrolling larger map by controlling the camera screen position, while
      while properly handling map edge collision detection
    - map screen transitioning (i.e. going into a door will load house interior)

---

# GAME IDEA: Italian Tomato Thrower (Vampire?)

**Desire:** create a multiplayer family friendly, horror "jump scare" game to
play with nieces and nephews that live far away.

- tomatoes:
  - randomly generated position, or always "in the kitchen"?
  - maybe vampires can "move" the tomatoes somewhere else in the kitchen to
    "hide" it?
- vampire players:
  - eat X tomatoes to "turn"
  - they "turn" into the "baddy" (chef? vampire? vampire chef?)
  - then they start singing some random opera-like italian song
    > ~ "mooooommma miiiiia, santaaaaaa mariiiiiiaaaaaahhhhh!"
    - innocents can hear it, relative to where the vampire is... making it spooky
      when you start hearing singing
  - they can throw tomatoes at non-chef's, trying to turn them
  - looks like blood, acts like vampires "biting" prey and turning them
  - innocents that are "turned" become zombie-like?
    - what should this mean?
      it would be cool to have this battle be kinda expanding territory, similar
      to Splatoon, Tony Hawk 2, etc
      where the winner at the end is the team with most territory?
      but... that wouldn't work well... that would take the "spooky, I need to
      hide from vampires" out of it
    - zombies will slowly return (walk) home, returning to normal if they reach
      the pure water
      - if the zombie sees an innocent in X radius, it'll:
        - auto track the innocent, following them (faster than a normal walk)
        - if touched, will turn that player zombie
        - players can try to "lead" the zombies back to base, to speed up the
          "respawn", but might git bit in the process
- "innocent" players:
  - maybe a few different "roles"?
    (all can shoot or splash water, in different ways/efficiencies)
    - Tomato Tracker: has a device (radar) that helps navigate where the tomato
      is hidden
    - Annihilator: water gun (picks up 1 type: sniper, machine gun, dual
      pistols), X water ballon grenades, and X trap land mines
    - Healer: special "pure juice box" to turn zombies back without them
      returning home
  - trying to hide from vampire chefs
  - collect enough "blessed water"
  - trying to "purify" zombies/chefs by spraying water on them
  - purified zombies turn back to players, returning to start
  - purified chefs "lose their song", get upset, and mope for a bit (disabled)
  - trying to stay alive until
    - timer runs out?
    - or maybe trying to achieve something before all being turned to chefs?
      - yeah! trying to find and destroy the hive tomatoes
