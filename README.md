## First pass, more work to come

Inputs: 10 single digit ints - is this a string of ints or single 10 digit int?

Ouputs: bool (valid / not), or void / throw

Valid numbers: `5990128088`, `1275988113`, `4536026665`
Invalid numbers: `5990128087`, `4536016660`

Weighting Factor: input length -> 2

### Domain Model

**Option 1: bool output, string input**
Function | Input | Scenario | Output
-------- | ----- | -------- | ------
nhsnum.IsValid | string | NHS Number is valid | true
_|_ | NHS Number is invalid | false

**Option 2: bool output, int input**
Function | Input | Scenario | Output
-------- | ----- | -------- | ------
nhsnum.IsValid | int | NHS Number is valid | true
_|_ | NHS Number is invalid | false

**Option 3: error, string input**
Function | Input | Scenario | Output
-------- | ----- | -------- | ------
nhsnum.Validate | string | NHS Number is valid | void
_|_ | NHS Number is invalid | panic

**Option 4: error, int input**
Function | Input | Scenario | Output
-------- | ----- | -------- | ------
nhsnum.Validate | int | NHS Number is valid | void
_|_ | NHS Number is invalid | panic

I'll go with boolean outputs and implement both string & int versions. In other languages I would prefer to void / throw an error so any existing error handler could take care of the fail case, not sure how to accomplish that in Go.

I'm also unfamiliar with the convention of MixedCaps, please excuse the package naming if it's crude!

## String input breakdown

Index | Factor
-- | --
0 | 10
1 | 9
2 | 8
3 | 7
4 | 6
5 | 5
6 | 4
7 | 3
8 | 2

Pseudo logic
--
`maxFactor = 10`

`currentFactor = maxFactor - index`

`total += (input[index] * currentFactor)`

`remainder = total % 11`

`providedCheckDigit = input[input.length - 1]`

`calculatedCheckDigit = 11 - remainder`

`if providedCheckDigit == calculatedCheckDigit: valid`

`if providedCheckDigit == 0 && calculatedCheckDigit == 11: valid`

`otherwise invalid`

1. Split string
2. Look at each digit excluding the last, calculate running total