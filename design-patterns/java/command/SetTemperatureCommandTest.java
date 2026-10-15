public final class SetTemperatureCommandTest {
    public static void main(String[] args) {
        Thermostat thermostat = new Thermostat();
        SmartButton button = new SmartButton();
        button.setCommand(new SetTemperatureCommand(thermostat, 25));

        button.press();
        button.press();
        button.undoLast();
        if (thermostat.getCurrentTemperature() != 25) {
            throw new AssertionError("first undo should restore the prior 25-degree setting");
        }
        button.undoLast();
        if (thermostat.getCurrentTemperature() != 20) {
            throw new AssertionError("second undo should restore the initial 20-degree setting");
        }
    }
}
