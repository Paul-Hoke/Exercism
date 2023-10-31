public class Lasagna {
  
  private int expectedMinutesInOven = 40;
  
  // TODO: define the 'expectedMinutesInOven()' method
  public int expectedMinutesInOven(int expectedMinutesInOven) {
    this.expectedMinutesInOven = expectedMinutesInOven;
    return this.expectedMinutesInOven;
  }
  
  public int expectedMinutesInOven() {
    return this.expectedMinutesInOven;
  }

  // TODO: define the 'remainingMinutesInOven()' method
  public int remainingMinutesInOven(int minutesSpentInOven) {
    return this.expectedMinutesInOven - minutesSpentInOven;
  }

  // TODO: define the 'preparationTimeInMinutes()' method
  public int preparationTimeInMinutes(int numLayers) {
    return numLayers * 2;
  }

  // TODO: define the 'totalTimeInMinutes()' method
  public int totalTimeInMinutes(int numLayers, int minutesSpentInOven) {
    return preparationTimeInMinutes(numLayers) + minutesSpentInOven;
  }
}
