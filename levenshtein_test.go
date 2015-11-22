
package levenshtein

import (
  "testing"
)

func AssertDistance(t *testing.T, x, y string, expected int) {
  if cost, _ := EditDistance(x, y, DefaultOpts); cost != expected {
    t.Errorf("Distance between \"%s\" and \"%s\" was %d, expected to be %d", x, y, cost, expected)
  }
}

func AssertDistanceCal(t *testing.T, x, y string, subWeight int, expected int) {
  if cost, _ := EditDistance(x, y, Opts{ 1, 1, subWeight, false }); cost != expected {
    t.Errorf("Distance between \"%s\" and \"%s\" with substitution weight of %d was %d, expected to be %d", x, y, subWeight, cost, expected)
  }
}

func TestEditDistance(t *testing.T) {

  AssertDistance(t, "intention", "execution", 5);
  AssertDistance(t, "kitten", "sitting", 3);
  AssertDistance(t, "Saturday", "Sunday", 3);
  AssertDistance(t, "bathroom", "bedroom", 3);
  AssertDistance(t, "sometime", "everytime", 5);
  AssertDistance(t, "I tink were going work Monay", "I think we're going to work on Monday", 9);
  AssertDistance(t, "Somtims ur wrng", "Sometimes you're wrong", 7);
  AssertDistance(t, "Hw meny hurs til lunch?", "How many hours until lunch?", 5);

}

func TestEditDistanceWeight(t *testing.T) {

  AssertDistanceCal(t, "intention", "execution", 2, 8);
  AssertDistanceCal(t, "kitten", "sitting", 4, 5);
  AssertDistanceCal(t, "Saturday", "Sunday", 3, 4);
  AssertDistanceCal(t, "bathroom", "bedroom", 6, 5);
  AssertDistanceCal(t, "sometime", "everytime", 4, 7);
  AssertDistanceCal(t, "I tink were going work Monay", "I think we're going to work on Monday", 2, 9);
  AssertDistanceCal(t, "Somtims ur wrng", "Sometimes you're wrong", 3, 7);
  AssertDistanceCal(t, "Hw meny hurs til lunch?", "How many hours until lunch?", 3, 6);

}

func TestEditDistanceReturnsZeroForEqualStrings(t *testing.T) {

  AssertDistance(t, "intention", "intention", 0);
  AssertDistance(t, "kitten", "kitten", 0);
  AssertDistance(t, "Saturday", "Saturday", 0);

}

func TestEditDistanceReturnsXLenForEmptyY(t *testing.T) {

  AssertDistance(t, "intention", "", 9);
  AssertDistance(t, "kitten", "", 6);
  AssertDistance(t, "Saturday", "", 8);

}

func BenchmarkEditDistanceWord(b *testing.B) {
  for i := 0; i < b.N; i++ {
    EditDistance("intention", "execution", DefaultOpts)
  }
}

func BenchmarkEditDistanceSentence(b *testing.B) {
  for i := 0; i < b.N; i++ {
    EditDistance("Hw meny hurs til lunch?", "How many hours until lunch?", DefaultOpts)
  }
}

func BenchmarkEditDistanceLongProteinSequence(b *testing.B) {
  for i := 0; i < b.N; i++ {
    EditDistance("ADDCTACACCTGACCTCCAGGCCGATACCCCADDCCTTCACACATGAGTTTCTCAA", "ADDCTACGGGTAGGCCTTCCCCCCADDCCTTCACACATGAGTTTCTC", DefaultOpts)
  }
}
