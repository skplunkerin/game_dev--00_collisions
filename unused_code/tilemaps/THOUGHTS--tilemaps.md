# Tilemaps

## Research

**What:** [what this json object is](#tilemap-example)

I think this is a "tilemap", possibly created from a 3rd-party tool?
"I was hoping to avoid a third party thing with tilemaps and json files"
https://discord.com/channels/842049801528016967/842056438632939570/933509688476176444

Creating your own tilemap engine tutorial; at first glance this appears to be
doing what this code looks like it's doing; hoping that this explains the math
or algorithm behind how the numbers in the JSON are created:

- https://gamedevelopment.tutsplus.com/tutorials/an-introduction-to-creating-a-tile-map-engine--gamedev-10900

## Answer:

Ok, figured it out; this is how this is working:
the image tilemap is like a puzzle, where each "square" (tile/piece) has
a number on it, starting with the first top/left tile being number 0,
then 1, 2, 3,... and so on until we get to tile 243, which is the grass.

I'm used to using the method "get X,Y position on the tilemap, grabbing
the width/height amount of pixels to the right/below X,Y position; this
will be the tile to display".

- **Simple X,Y position method:**

  - **Pros:**
    - easier to understand where on the image is being targeted
    - easier to quickly find where you would want by finding the x,y
      position by hovering the cursor over the image in Photoshop, etc
    - don't need a tool (other than an image processor that gives you x,y
      coordinates for your cursor) to build things
  - **Cons:**
    - for every tile, you will need 2 integers to identify the tile: x,y

- **Tilemap method?** (is this what it's actually called? Is a "spritesheet" the
  source image, and tilemap the json output of what you build?)

  - **Pros:**
    - only need a single integer for a tile position
  - **Cons:**
    - not simple to look at a spritesheet and know what "number" (tile ID) you
      need

## Tilemap example:

```jsonc
[
  {
    243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
    243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
    243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
    243, 218, 243, 243, 243, 243, 243, 243, 243, 243, 243, 218, 243, 244, 243,
    243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,

    243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
    243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
    243, 243, 244, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
    243, 243, 243, 243, 243, 243, 243, 243, 243, 219, 243, 243, 243, 219, 243,
    243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,

    243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
    243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
    243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
    243, 218, 243, 243, 243, 243, 243, 243, 243, 243, 243, 244, 243, 243, 243,
    243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
  },
  {
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 26, 27, 28, 29, 30, 31, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 51, 52, 53, 54, 55, 56, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 76, 77, 78, 79, 80, 81, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 101, 102, 103, 104, 105, 106, 0, 0, 0, 0,

    0, 0, 0, 0, 0, 126, 127, 128, 129, 130, 131, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 303, 303, 245, 242, 303, 303, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0,

    0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0,
  }
]
```
