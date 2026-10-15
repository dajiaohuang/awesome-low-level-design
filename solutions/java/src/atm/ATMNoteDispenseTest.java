package atm;

import atm.chainofresponsibility.NoteDispenser100;
import atm.chainofresponsibility.NoteDispenser20;
import atm.chainofresponsibility.NoteDispenser50;
import atm.entities.CashDispenser;
import java.io.ByteArrayOutputStream;
import java.io.PrintStream;
import java.nio.charset.StandardCharsets;

public final class ATMNoteDispenseTest {
    public static void main(String[] args) {
        NoteDispenser100 hundreds = new NoteDispenser100(1);
        NoteDispenser50 fifties = new NoteDispenser50(1);
        NoteDispenser20 twenties = new NoteDispenser20(4);
        hundreds.setNextChain(fifties);
        fifties.setNextChain(twenties);
        CashDispenser dispenser = new CashDispenser(hundreds);

        if (!dispenser.canDispenseCash(80)) {
            throw new AssertionError("four available $20 notes should satisfy $80");
        }
        ByteArrayOutputStream output = new ByteArrayOutputStream();
        PrintStream original = System.out;
        try (PrintStream captured = new PrintStream(output, true, StandardCharsets.UTF_8)) {
            System.setOut(captured);
            dispenser.dispenseCash(80);
        } finally {
            System.setOut(original);
        }
        String payout = output.toString(StandardCharsets.UTF_8);
        if (!payout.contains("Dispensing 4 x $20") || payout.contains("$50")) {
            throw new AssertionError("payout did not use the valid four-$20 combination: " + payout);
        }
        if (dispenser.canDispenseCash(80)) {
            throw new AssertionError("dispensed notes must no longer be available");
        }
    }
}
