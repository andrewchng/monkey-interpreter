# expresision

## post fix operator

-1
!true
!
Return

## in fix operator, binary expressions - operator having 2 operands

5+5
5-5
5/5
5#5

foo == bar
foo != bar
foo < bar
foo > bar

## need to handled order of operations (operator precedence)

5 *(5 + 5)
((5+5)* 5)*5

add(2,4)
add(add(2,3), add(2,2))
max(5, add(2, (5*5)))

foo * bar / foobar
add(foo , bar)

let add = fn(x, y) { return x + y };nopo

fn(x, y) { return x + y }(5, 5)
(fn(x) { return x }(5) + 10 ) * 10

let result = if (10 > 5) { true } else { false };
