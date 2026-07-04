# Grep's Design

## Flag Parsing

### Categories

#### Context lines

-A -B and -C all print lines surrounding matching pattern. This would happen in the printing layer, but it also means that grep must know at print time the location of a line. So it would be useful to have more data associated with a line than just the matched text. Really, not setting this options means above and below context is 0.

#### Counting

-c is for counting how many matches have occurred. must be >= 0. this might be as simple as having an array of matches to begin with, and then getting a count. I'm not sure if an array or slice would be the best data structure for actual storage, so I'll have to think more.

#### color

what does `never` mean?

what does `auto` mean?
what does `always` mean?


