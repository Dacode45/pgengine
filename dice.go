package pgengine

type Die struct {
  Rolls Int
  Sides Int
  Modifier Int
}

type Dice struct {
  mDice []Die
}

func NewDice(diceStr string) {
  dice := Dice{
    mDice: []Die{}
  }
  dice.Parse(diceStr)
}

func (dice *Dice) Parse(diceStr string) {
  size := len(diceStr)
  index := 0

  for index <= size {
    die, index := dice.ParseDie(diceStr, index)
    append(dice.mDice, die)
    index = index + 1 // Skip space
  }
}

func (dice *Dice) ParseDie(diceStr, i) {
  rolls, i := dice.ParseNumber(diceStr, i)

  i = i + 1 // Move past the D

  sides, i := dice.ParseNumber(diceStr, i)

  if i == len(diceStr) || diceStr[i] == " " {
    return Die{ rolls, sides, 0 }, i
  }

  if diceStr[i] == "+" {
    i = i + 1 // move past the +
    plus, i := dice.ParseNumber(diceStr, i)
    return Die{ Rolls: rolls, Sides: sides, Modifier: plus }, i
  }
}

func (dice *Dice) ParseNumber(str string, index Int) Int, Int {
  size := len(str)
  subStr := ""

  for i := index; i < size; i++ {
    char := str[i]

    _, err := strconv.Atoi(char)
    if err != nil {
      num, _ := strconv.Atoi(subStr)
      return num, i
    }
    subStr += char
  }
  num, _ := strconv.Atoi(subStr)
  return num, size - 1
}

func RollDie(rolls, faces, modifier Int) Int {
  var total Int

  for i := 0; i < rolls; i++ {
    total = total + rand.Intn(faces - 1) + 1
  }
  return total + modifier
}

func (dice *Dice) Roll() {
  var total Int

  for _, die := range dice.mDice {
    total = total + RollDie(die.Rolls, die.Sides, die.Modifier)
  }

  return total
}
