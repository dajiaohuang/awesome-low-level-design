import { DirectionEnum } from "./Direction";
import GreenState from "./SignalState/GreenState";
import RedState from "./SignalState/RedState";
import SignalState from "./SignalState/SignalState";
import TrafficLight from "./TrafficLight";

export default class TrafficSignalController {
    private signals: Map<DirectionEnum, TrafficLight>;
    private signalDurations: Map<DirectionEnum, Map<String, number>>;
    private pendingTimer: ReturnType<typeof setTimeout> | undefined;
    private generation = 0;

    constructor(
        signals: Map<DirectionEnum, TrafficLight>,
        signalDurations: Map<DirectionEnum, Map<String, number>>,
    ) {
        this.signals = signals;
        this.signalDurations = signalDurations;
    }

    start(startDirection: DirectionEnum) {
        const light = this.signals.get(startDirection);
        if (!light) return;
        this.generation++;
        if (this.pendingTimer !== undefined) clearTimeout(this.pendingTimer);
        this.pendingTimer = undefined;
        for (const signal of this.signals.values()) signal.setState(new RedState());
        light.setState(new GreenState());
        light.handle(this);
    }

    scheduleStateChange(
        light: TrafficLight,
        direction: DirectionEnum,
        nextState: SignalState,
        delaySeconds: number,
    ) {
        if (this.pendingTimer !== undefined) clearTimeout(this.pendingTimer);
        const generation = this.generation;
        this.pendingTimer = setTimeout(() => {
            if (generation !== this.generation) return;
            this.pendingTimer = undefined;
            light.setState(nextState);
            light.handle(this); // Handle the light after state change
        }, delaySeconds * 1000);
    }

    getSignalDuration(direction: DirectionEnum, state: SignalState) {
        return this.signalDurations.get(direction)?.get(state.getName()) || 0;
    }

    getNextDirection(current: DirectionEnum): DirectionEnum {
        const directions: DirectionEnum[] = Object.values(DirectionEnum);
        const next = (directions.indexOf(current) + 1) % directions.length;
        return directions[next];
    }

    getTrafficLight(direction: DirectionEnum) {
        const signal = this.signals.get(direction);
        if (!signal) throw new Error("Signal Not Found.");
        return signal;
    }

    manualOverride(direction: DirectionEnum) {
        // Replace the owned timer chain and red-light conflicting directions.
        this.start(direction);
    }
}
