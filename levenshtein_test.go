
package levelshtein

import (
  "testing"
)

func AssertDistance(t *testing.T, x, y string, expected int) {
  if cost := EditDistance(x, y); cost != expected {
    t.Errorf("Expected distance between \"%s\" and \"%s\" was %d, expected to be %d", x, y, cost, expected)
  }
}

func TestEditDistance(t *testing.T) {

  // Simple word transformations
  AssertDistance(t, "intention", "execution", 5);
  AssertDistance(t, "kitten", "sitting", 3);
  AssertDistance(t, "Saturday", "Sunday", 3);
  AssertDistance(t, "bathroom", "bedroom", 3);
  AssertDistance(t, "sometime", "everytime", 5);

  // Sentence transformations
  AssertDistance(t, "I tink were going work Monay", "I think we're going to work on Monday", 9);
  AssertDistance(t, "Somtims ur wrng", "Sometimes you're wrong", 7);
  AssertDistance(t, "Hw meny hurs til lunch?", "How many hours until lunch?", 5);

}
