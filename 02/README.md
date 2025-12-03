# 02

~~I think I can be pretty efficient about computing the number of repeated
sequences in a given number of digits. I also think that can be extended to
allow for an upper and/or lower bound within the band.~~

~~Assuming the ranges are non-intersecting, which follows from the problem
statement of ranges being leftover, this would be approximately linear runtime
over the number of ranges.~~

I have realized that I misread the objective. I don't want to know how many
there are, I need to sum them. This will be less efficient, but I didn't get
very far into it yet and a lot of the pieces can still be used. I can still be
pretty efficient within each number of digits, but will need a loop that is
something like O(sqrt(max_value))

