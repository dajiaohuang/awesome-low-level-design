package atm.chainofresponsibility;

abstract class NoteDispenser implements DispenseChain {
    private DispenseChain nextChain;
    private final int noteValue;
    private int numNotes;

    public NoteDispenser(int noteValue, int numNotes) {
        this.noteValue = noteValue;
        this.numNotes = numNotes;
    }

    @Override
    public void setNextChain(DispenseChain nextChain) {
        this.nextChain = nextChain;
    }

    @Override
    public synchronized void dispense(int amount) {
        int maxNotes = Math.min(amount / noteValue, this.numNotes);
        for (int numToDispense = maxNotes; numToDispense >= 0; numToDispense--) {
            int remainingAmount = amount - (numToDispense * noteValue);
            boolean remainderDispensable = remainingAmount == 0 ||
                    (this.nextChain != null && this.nextChain.canDispense(remainingAmount));
            if (!remainderDispensable) continue;

            if (numToDispense > 0) {
                System.out.println("Dispensing " + numToDispense + " x $" + noteValue + " note(s)");
                this.numNotes -= numToDispense;
            }
            if (remainingAmount > 0) {
                this.nextChain.dispense(remainingAmount);
            }
            return;
        }
    }

    @Override
    public synchronized boolean canDispense(int amount) {
        if (amount < 0) return false;
        if (amount == 0) return true;

        int maxNotes = Math.min(amount / noteValue, this.numNotes);
        for (int numToUse = maxNotes; numToUse >= 0; numToUse--) {
            int remainingAmount = amount - (numToUse * noteValue);
            if (remainingAmount == 0 ||
                    (this.nextChain != null && this.nextChain.canDispense(remainingAmount))) {
                return true;
            }
        }
        return false;
    }
}
