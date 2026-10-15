abstract class NoteDispenser : IDispenseChain
{
    private IDispenseChain nextChain;
    private readonly int noteValue;
    private int numNotes;
    private readonly object dispenserLock = new object();

    public NoteDispenser(int noteValue, int numNotes)
    {
        this.noteValue = noteValue;
        this.numNotes = numNotes;
    }

    public void SetNextChain(IDispenseChain nextChain)
    {
        this.nextChain = nextChain;
    }

    public void Dispense(int amount)
    {
        lock (dispenserLock)
        {
            int maxNotes = Math.Min(amount / noteValue, numNotes);
            for (int numToDispense = maxNotes; numToDispense >= 0; numToDispense--)
            {
                int remainingAmount = amount - (numToDispense * noteValue);
                if (remainingAmount == 0)
                {
                    CommitNotes(numToDispense);
                    return;
                }
                if (nextChain != null && nextChain.CanDispense(remainingAmount))
                {
                    CommitNotes(numToDispense);
                    nextChain.Dispense(remainingAmount);
                    return;
                }
            }
        }
    }

    private void CommitNotes(int numToDispense)
    {
        if (numToDispense > 0)
        {
            Console.WriteLine($"Dispensing {numToDispense} x ${noteValue} note(s)");
            numNotes -= numToDispense;
        }
    }

    public bool CanDispense(int amount)
    {
        lock (dispenserLock)
        {
            if (amount < 0) return false;
            if (amount == 0) return true;

            int maxNotes = Math.Min(amount / noteValue, numNotes);
            for (int numToUse = maxNotes; numToUse >= 0; numToUse--)
            {
                int remainingAmount = amount - (numToUse * noteValue);
                if (remainingAmount == 0 ||
                    (nextChain != null && nextChain.CanDispense(remainingAmount)))
                {
                    return true;
                }
            }
            return false;
        }
    }
}