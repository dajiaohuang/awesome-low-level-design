import java.util.Stack;

public class SetTemperatureCommand implements Command {
    private final Thermostat thermostat;
    private final int newTemperature;
    private final Stack<Integer> previousTemperatures = new Stack<>();

    public SetTemperatureCommand(Thermostat thermostat, int temperature) {
        this.thermostat = thermostat;
        this.newTemperature = temperature;
    }

    @Override
    public void execute() {
        previousTemperatures.push(thermostat.getCurrentTemperature());
        thermostat.setTemperature(newTemperature);
    }

    @Override
    public void undo() {
        thermostat.setTemperature(previousTemperatures.pop());
    }
}