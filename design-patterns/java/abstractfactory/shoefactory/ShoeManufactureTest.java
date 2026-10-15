import java.io.ByteArrayOutputStream;
import java.io.PrintStream;
import java.nio.charset.StandardCharsets;

public final class ShoeManufactureTest {
    public static void main(String[] args) {
        String dynamicFormal = capture(new String(new char[] {'F', 'o', 'r', 'm', 'a', 'l'}));
        String literalFormal = capture("Formal");
        if (!dynamicFormal.equals(literalFormal)) {
            throw new AssertionError("equal supported shoe names must select the same factory");
        }
    }

    private static String capture(String shoeType) {
        Shoe shoe = ShoeManufacture.produceShoe(shoeType);
        ByteArrayOutputStream output = new ByteArrayOutputStream();
        PrintStream original = System.out;
        try (PrintStream captured = new PrintStream(output, true, StandardCharsets.UTF_8)) {
            System.setOut(captured);
            shoe.displayBuildShoe();
        } finally {
            System.setOut(original);
        }
        return output.toString(StandardCharsets.UTF_8);
    }
}
