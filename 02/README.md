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

## Optimization

Even with the misread, it turns out you can avoid the loop because you can can
convert the sum into two constant time computations.

Consider the first few invalid IDs with 4 digits:
- `1010`
- `1111`
- `1212`
- `1313`

This can be rewritten as:
- `1010 + 0 * 101`
- `1010 + 1 * 101`
- `1010 + 2 * 101`
- `1010 + 3 * 101`

If you want to compute the sum up to some term, you compute these independently
and sum them together.

The left side is clearly just `N * 1010`. The right side is also a familiar
pattern, just the sum from `0` to `N - 1`. You may be familiar with the more
common sum from `1` to `N`, which is `N(N + 1)/2`.

You could convert this to the sum from `0` to `N - 1` as `(N - 1)N/2`, but it is
actually more convenient to compute the zero based index `i` and you can just
plug that directly into `i(i + 1)/2`. However, if you do then you need to
consider the left hand side as `(i + 1) * 1010`.

We also need to establish how to compute `1010`, `101`, and `i`. These are all
defined with regards to a certain number of digits `D`.

Lets start with `101`, which I have been calling the diff `d_D`. What we need is
a `1` followed by half as many zeros as digits in the full number, plus `1`.
This is easy to compute, first `h_D = 10^(D/2)+1`, where `h_D` is what I have
been calling the half magnitude. Then `d_D = h_D + 1`.

`1010` is `h_D`, but moved over by one less than half the number of digits. I
have been calling this amount the sub magnitude `s_D`, and `s_D = 10^(d_D/2)`.
`1010` is what I have been calling the basis `b_D`, and `b_D = h_D * s_D`.

`i` is also easy to compute. If you have a half bound `B`, you can compute it as
`i = B - s_D`. A half bound is just one of the repeated halves of the number,
e.g. `19` in `1919`. You can compute it from the actual lower or upper bound
divided by `h_D`. However, you still need to be careful with the bound as seen
in the old implementation.

With all this, we can define a function `f_D(i)` that computes the sum of the
invalid IDs from the 0th index up to `i` for a number of digits `D`: `f_D(i) =
h_D * (i + 1) + d_D * i(i + 1)/2`. Note that this function is well defined for
`i = -1`, and will return `0`. This means we can make it non-inclusive by
passing in a value one less than a bound. `f_D(i)` is only defined well for `0
<= i < s`

You can then use this to compute the sum between a lower bound `l` and upper
bound `u` for some number of digits `D`: `S_D(l, u) = f_D(u) - f_D(l - 1)`.

